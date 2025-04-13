package application

import (
	"encoding/json"
	"fmt"
	"log"
	"order-services/models"
	"order-services/types"

	"github.com/google/uuid"
)

type InventoryMessage struct {
	OrderID string `json:"order_id"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (s *ApplicationService) CreateOrder(req types.Order) (types.Order, error) {
	order := models.Order{
		Status: req.Status,
	}
	var orderItems []types.OrderItemMessage

	for _, item := range req.OrderItems {
		orderItem := types.OrderItemMessage{
			ProductID: int(item.ProductID),
			Quantity:  item.Quantity,
		}

		orderItems = append(orderItems, orderItem)
	}
	orderID := uuid.New().String()

	err := s.sendOrderMessage(orderID, orderItems)
	if err != nil {
		return types.Order{}, fmt.Errorf("failed to send order message: %w", err)
	}

	success, message := s.ListenToSuccessQueue(orderID)

	if success {
		order.Status = "success"
		createdOrder, err := s.modelService.CreateOrder(order)
		if err != nil {
			return types.Order{}, err
		}

		for _, item := range req.OrderItems {
			orderItem := models.OrderItem{
				OrderID:   createdOrder.ID,
				ProductID: uint(item.ProductID),
				Quantity:  item.Quantity,
			}

			_, err := s.modelService.CreateOrderItem(orderItem)
			if err != nil {
				s.modelService.DeleteOrder(int(createdOrder.ID))
				return types.Order{}, fmt.Errorf("failed to create order item: %w", err)
			}
		}

		return types.Order{
			ID:         createdOrder.ID,
			Status:     "success",
			OrderItems: req.OrderItems,
		}, nil
	} else {
		return types.Order{}, fmt.Errorf(message)
	}

}
func (s *ApplicationService) sendOrderMessage(orderID string, orderItems []types.OrderItemMessage) error {
	failMessage := types.OrderMessage{
		OrderID: orderID,
		Items:   orderItems,
	}

	failMessageBody, err := json.Marshal(failMessage)
	if err != nil {
		log.Printf("Error marshalling fail message: %s", err)
		return err
	}

	err = s.rabbitMQService.PublishMessage("order_queue", failMessageBody)
	if err != nil {
		log.Printf("Error publishing fail message to success queue: %s", err)
		return err
	}

	return nil
}
func (s *ApplicationService) ListenToSuccessQueue(orderID string) (bool, string) {
	msgs, err := s.rabbitMQService.ConsumeMessages("success_queue")
	if err != nil {
		return false, "Failed to consume messages"
	}

	for msg := range msgs {
		var orderMsg types.SuccessMessage

		err := json.Unmarshal(msg.Body, &orderMsg)
		if err != nil {
			msg.Nack(false, false)
			continue
		}
		if orderMsg.OrderID == orderID && orderMsg.Status == "success" {
			msg.Ack(true)
			return true, ""
		}

		msg.Ack(true)
		return false, orderMsg.Message
	}

	return false, "Order processing failed or status is not success."
}

func (s *ApplicationService) GetOrder(req types.Order) (types.Order, error) {
	connResp, err := s.modelService.GetOrder(int(req.ID))
	if err != nil {
		return types.Order{}, err
	}

	return types.Order{
		ID:     connResp.ID,
		Status: connResp.Status,
	}, nil
}

func (s *ApplicationService) DeleteOrder(req types.Order) error {
	err := s.modelService.DeleteOrder(int(req.ID))
	if err != nil {
		return err
	}

	return fmt.Errorf("Order deleted successfully")
}

func (s *ApplicationService) GetAllOrders() (types.ListOrdersResponse, error) {
	var orders types.ListOrdersResponse
	connResp, err := s.modelService.GetAllOrders()
	if err != nil {
		return types.ListOrdersResponse{}, err
	}

	for _, order := range connResp {
		response := types.Orders{
			ID:     order.ID,
			Status: order.Status,
		}

		orders.Orders = append(orders.Orders, response)
	}

	return orders, nil
}
