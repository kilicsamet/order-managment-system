package application

import (
	"order-services/models"
	"order-services/types"
)

func (s *ApplicationService) CreateOrderItem(req types.OrderItem) (types.OrderItem, error) {

	orderItem := models.OrderItem{
		OrderID:   req.OrderID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	createdOrderItem, err := s.modelService.CreateOrderItem(orderItem)
	if err != nil {
		return types.OrderItem{}, err
	}

	return types.OrderItem{
		ID:        createdOrderItem.ID,
		OrderID:   createdOrderItem.OrderID,
		ProductID: createdOrderItem.ProductID,
		Quantity:  createdOrderItem.Quantity,
	}, nil
}

func (s *ApplicationService) GetOrderItem(req types.OrderItem) (types.OrderItem, error) {
	connResp, err := s.modelService.GetOrderItem((req.ID))
	if err != nil {
		return types.OrderItem{}, err
	}

	return types.OrderItem{
		ID:        connResp.ID,
		OrderID:   connResp.OrderID,
		ProductID: connResp.ProductID,
		Quantity:  connResp.Quantity,
	}, nil
}

func (s *ApplicationService) GetAllOrderItems() (types.ListOrderItemsResponse, error) {
	var orderItems types.ListOrderItemsResponse
	connResp, err := s.modelService.GetAllOrderItems()
	if err != nil {
		return types.ListOrderItemsResponse{}, err
	}

	for _, orderItem := range connResp {
		response := types.OrderItem{
			ID:        orderItem.ID,
			OrderID:   orderItem.OrderID,
			ProductID: orderItem.ProductID,
			Quantity:  orderItem.Quantity,
		}

		orderItems.OrderItems = append(orderItems.OrderItems, response)
	}

	return orderItems, nil
}
