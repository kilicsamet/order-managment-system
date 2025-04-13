package types

type Orders struct {
	ID     uint   `json:"id"`
	Status string `json:"status"`
}
type Order struct {
	ID         uint               `json:"id"`
	Status     string             `json:"status"`
	OrderItems []OrderItemMessage `json:"order_items"`
}

type CreateOrderRequest struct {
	Status string `json:"status"`
}

type GetOrderResponse struct {
	Order `json:",inline"`
}

type ListOrdersResponse struct {
	Orders []Orders `json:"orders"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}
type OrderMessage struct {
	OrderID string             `json:"order_id"`
	Items   []OrderItemMessage `json:"items"`
}

type OrderItemMessage struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
type SuccessMessage struct {
	OrderID string `json:"order_id"`
	Status  string `json:"status"`
	Message string `json:"message"`
}
