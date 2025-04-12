package application

import "inventory-service/models"

type ApplicationService struct {
	modelService    *models.ModelService
	rabbitMQService *RabbitMQService
}

func New(dbType, dsn, rabbitMQURL string) (*ApplicationService, error) {
	modelService, err := models.New(dbType, dsn)
	if err != nil {
		return nil, err
	}

	rabbitMQService, err := NewRabbitMQService(rabbitMQURL)
	if err != nil {
		return nil, err
	}

	return &ApplicationService{
		modelService:    modelService,
		rabbitMQService: rabbitMQService,
	}, nil
}

func (s *ApplicationService) Migrate() {
	s.modelService.Migrate()
}

func (s *ApplicationService) GetRabbitMQService() *RabbitMQService {
	return s.rabbitMQService
}
