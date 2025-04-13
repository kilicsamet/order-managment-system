package types

type OrderItem struct {
	ID        uint `json:"id"`
	OrderID   uint `json:"order_id"`
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

type CreateOrderItemRequest struct {
	OrderID   uint `json:"order_id"`
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

type GetOrderItemResponse struct {
	OrderItem `json:",inline"`
}

type ListOrderItemsResponse struct {
	OrderItems []OrderItem `json:"order_items"`
}
