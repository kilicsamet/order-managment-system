package application

import (
	"product-services/models"
)

type ApplicationService struct {
	modelService *models.ModelService
}

func New(dbType, dsn string) (*ApplicationService, error) {
	modelService, err := models.New(dbType, dsn)
	if err != nil {
		return nil, err
	}
	return &ApplicationService{
		modelService: modelService,
	}, nil
}

func (s *ApplicationService) Migrate() {
	s.modelService.Migrate()
}
