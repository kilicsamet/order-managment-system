package api

func (s *APIService) routes() {
	s.fiberApp.Post("/api/product", s.handlerService.AddProduct)
	s.fiberApp.Get("/api/products", s.handlerService.GetProducts)
	s.fiberApp.Get("/api/product/:id", s.handlerService.GetProduct)
	s.fiberApp.Put("/api/product/:id", s.handlerService.UpdateProduct)
	s.fiberApp.Delete("/api/product/:id", s.handlerService.DeleteProduct)

}
