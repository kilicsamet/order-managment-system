package api

func (s *APIService) routes() {
	s.fiberApp.Post("/api/order", s.handlerService.AddOrder)
	s.fiberApp.Get("/api/orders", s.handlerService.GetOrders)
	s.fiberApp.Get("/api/order/:id", s.handlerService.GetOrder)
	s.fiberApp.Delete("/api/order/:id", s.handlerService.DeleteOrder)
	s.fiberApp.Get("/api/order_items", s.handlerService.GetOrderItems)
	s.fiberApp.Get("/api/order_item/:id", s.handlerService.GetOrderItem)

}
