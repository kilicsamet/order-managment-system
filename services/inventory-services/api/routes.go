package api

func (s *APIService) routes() {
	s.fiberApp.Post("/api/inventory_product", s.handlerService.AddInventoryProduct)
	s.fiberApp.Put("/api/inventory_product/:id", s.handlerService.UpdateInventoryProductQuantity)
	s.fiberApp.Get("/api/inventory_products", s.handlerService.GetInventoryProducts)
	s.fiberApp.Get("/api/inventory_product/:id", s.handlerService.GetInventoryProduct)
	s.fiberApp.Delete("/api/inventory_product/:id", s.handlerService.DeleteInventoryProduct)

}
