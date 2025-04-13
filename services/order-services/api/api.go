package api

import (
	"order-services/application"
	"order-services/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type APIService struct {
	listenAddr     string
	fiberApp       *fiber.App
	handlerService *handlers.HandlerService
}

func New(listenAddr string, appService *application.ApplicationService) *APIService {
	if !fiber.IsChild() {
		appService.Migrate()
	}

	handService := handlers.New(appService)
	service := &APIService{
		listenAddr: listenAddr,
		fiberApp: fiber.New(fiber.Config{
			DisableStartupMessage: false,
			Prefork:               false,
		}),
		handlerService: handService,
	}
	service.fiberApp.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))
	service.routes()
	return service
}

func (s *APIService) Stop() {
	s.fiberApp.Shutdown()
}

func (s *APIService) Start() {
	go func() {
		s.fiberApp.Listen(s.listenAddr)
	}()
}
