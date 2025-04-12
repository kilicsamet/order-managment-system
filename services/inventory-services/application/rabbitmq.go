package application

import (
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
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQService{
		connection: conn,
		channel:    channel,
	}, nil
}

func (s *RabbitMQService) ConsumeMessages(queueName string) (<-chan amqp.Delivery, error) {
	msgs, err := s.channel.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}

func (s *RabbitMQService) PublishMessage(queueName string, body []byte) error {
	// Kuyruğu kontrol et ve oluştur (Eğer yoksa)
	_, err := s.channel.QueueDeclare(
		queueName, // Kuyruğun adı
		true,      // Kuyrukta mesajlar kalıcı olsun mu
		false,     // Kuyruk otomatik olarak silinsin mi
		false,     // Kuyruğu yalnızca tek bir tüketiciye mi teslim et
		false,     // Kuyruğun var olup olmadığını kontrol etmeden oluşturma
		nil,       // Ek parametreler
	)
	if err != nil {
		log.Printf("Error declaring queue %s: %s", queueName, err)
		return err
	}

	// Kuyruğa mesaj gönder
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
		log.Printf("Failed to publish message to queue %s: %s", queueName, err)
		return err
	}

	return nil
}

func (s *RabbitMQService) Close() {
	if s.channel != nil {
		s.channel.Close()
	}
	if s.connection != nil {
		s.connection.Close()
	}
}
