package application

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQService struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

func NewRabbitMQService(rabbitMQURL string) (*RabbitMQService, error) {
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	return &RabbitMQService{
		connection: conn,
		channel:    channel,
	}, nil
}

func (s *RabbitMQService) ensureChannelOpen() error {
	// Kanalda bir işlem yaparak hatayı kontrol et
	if s.channel == nil {
		// Eğer kanal yoksa, yeni bir kanal oluştur
		if err := s.reconnectChannel(); err != nil {
			return fmt.Errorf("failed to reconnect channel: %w", err)
		}
	}

	// Kanalda bir işlem yapmayı deneyelim
	err := s.channel.Publish(
		"",
		"health_check_queue", // Sağlık kontrolü için geçici bir kuyruk adı
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte("health check"),
		},
	)
	if err != nil {
		// Eğer kanal kapalıysa ya da bir hata varsa, kanal yeniden açılmalıdır
		log.Printf("Channel is closed or an error occurred: %v", err)
		if err := s.reconnectChannel(); err != nil {
			return fmt.Errorf("failed to reconnect channel: %w", err)
		}
	}

	return nil
}

func (s *RabbitMQService) reconnectChannel() error {
	if s.connection == nil || s.connection.IsClosed() {
		conn, err := amqp.Dial("your-rabbitmq-url")
		if err != nil {
			return fmt.Errorf("failed to connect to RabbitMQ: %w", err)
		}
		s.connection = conn
	}

	// Yeni kanal oluştur
	channel, err := s.connection.Channel()
	if err != nil {
		return fmt.Errorf("failed to create a new channel: %w", err)
	}
	s.channel = channel
	return nil
}

func (s *RabbitMQService) PublishMessage(queueName string, body []byte) error {
	// Kanalın açık olup olmadığını kontrol et
	if err := s.ensureChannelOpen(); err != nil {
		return err
	}

	_, err := s.channel.QueueDeclare(
		queueName,
		true,  // Durable: Kuyruk sunucuda kalıcı olacak
		false, // AutoDelete: Kuyruk boşaldığında otomatik silinmez
		false, // Exclusive: Diğer bağlantılardan erişim engellenmez
		false, // NoWait: Bekleme yok
		nil,   // Args: Opsiyonel parametreler
	)
	if err != nil {
		return fmt.Errorf("error declaring queue %s: %w", queueName, err)
	}

	err = s.channel.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message to queue %s: %w", queueName, err)
	}

	return nil
}

func (s *RabbitMQService) ConsumeMessages(queueName string) (<-chan amqp.Delivery, error) {
	// Consume işlemi, kanal ve kuyruk varsa başlayacaktır
	msgs, err := s.channel.Consume(
		queueName,
		"",
		true,  // AutoAck: Mesajın otomatik olarak onaylanması
		false, // Exclusive: Sadece bu kanal kullanabilir
		false, // NoLocal: Mesajın sadece bu kanal tarafından alınmasını sağlar
		false, // NoWait: Cevap beklemeden işlemi başlat
		nil,   // Args: Opsiyonel parametreler
	)
	if err != nil {
		return nil, fmt.Errorf("failed to consume messages from queue %s: %w", queueName, err)
	}

	return msgs, nil
}

func (s *RabbitMQService) Close() {
	if s.channel != nil {
		err := s.channel.Close()
		if err != nil {
			log.Printf("Error closing channel: %v", err)
		}
	}
	if s.connection != nil {
		err := s.connection.Close()
		if err != nil {
			log.Printf("Error closing connection: %v", err)
		}
	}
}
