package handlers

import "product-services/application"

type HandlerService struct {
	applicationService *application.ApplicationService
}

func New(appService *application.ApplicationService) *HandlerService {
	return &HandlerService{
		applicationService: appService,
	}
}
