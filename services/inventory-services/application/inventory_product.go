package application

import (
	"encoding/json"
	"fmt"
	"inventory-service/models"
	"inventory-service/types"
	"log"
	"strings"
	"time"

	"github.com/streadway/amqp"
)

type InventoryMessage struct {
	OrderID string `json:"order_id"`
	Items   []struct {
		ProductID int `json:"product_id"`
		Quantity  int `json:"quantity"`
	} `json:"items"`
}

func (s *ApplicationService) CreateInventoryProduct(req types.InventoryProduct) (types.InventoryProduct, error) {

	product := models.InventoryProduct{
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	createdProduct, err := s.modelService.CreateInventoryProduct(product)
	if err != nil {
		return types.InventoryProduct{}, err
	}

	return types.InventoryProduct{
		ID:        createdProduct.ID,
		ProductID: createdProduct.ProductID,
		Quantity:  createdProduct.Quantity,
	}, nil
}

func (s *ApplicationService) ListenToOrderQueue() {
	for {
		msgs, err := s.rabbitMQService.ConsumeMessages("order_queue")
		if err != nil {
			log.Printf("Error consuming messages: %s. Retrying in 2 seconds...", err)
			time.Sleep(2 * time.Second)
			continue
		}

		for msg := range msgs {
			if err := s.processOrder(msg); err != nil {
				log.Printf("Error processing order: %s", err)
			}
			msg.Ack(false)
		}
	}
}

func (s *ApplicationService) processOrder(msg amqp.Delivery) error {
	var orderMsg InventoryMessage
	if err := json.Unmarshal(msg.Body, &orderMsg); err != nil {
		log.Printf("Error unmarshalling message: %s", err)
		return err
	}

	failedItems, err := s.checkStockForOrder(orderMsg)
	if err != nil {
		log.Printf("Error checking stock: %s", err)
		return err
	}

	if len(failedItems) > 0 {
		return s.sendFailMessage(orderMsg, failedItems)
	}

	return s.processSuccessfulOrder(orderMsg)
}

func (s *ApplicationService) checkStockForOrder(orderMsg InventoryMessage) ([]string, error) {
	var failedItems []string
	for _, product := range orderMsg.Items {
		stock, err := s.CheckProductStock(product.ProductID)
		if err != nil {
			log.Printf("Error checking stock for product %d: %s", product.ProductID, err)
			continue
		}

		if stock < product.Quantity {
			failedItems = append(failedItems, fmt.Sprintf("Product %d has insufficient stock.", product.ProductID))
		}
	}
	return failedItems, nil
}

func (s *ApplicationService) sendFailMessage(orderMsg InventoryMessage, failedItems []string) error {
	failMessage := types.SuccessMessage{
		OrderID: orderMsg.OrderID,
		Status:  "failed",
		Message: fmt.Sprintf("The following items have insufficient stock: %s", strings.Join(failedItems, ", ")),
	}

	failMessageBody, err := json.Marshal(failMessage)
	if err != nil {
		log.Printf("Error marshalling fail message: %s", err)
		return err
	}

	err = s.rabbitMQService.PublishMessage("success_queue", failMessageBody)
	if err != nil {
		log.Printf("Error publishing fail message to success queue: %s", err)
		return err
	}

	return nil
}

func (s *ApplicationService) processSuccessfulOrder(orderMsg InventoryMessage) error {
	for _, product := range orderMsg.Items {
		if err := s.UpdateInventoryProductQuantity(uint(product.ProductID), -product.Quantity); err != nil {
			log.Printf("Error updating inventory for product %d: %s", product.ProductID, err)
			return err
		}
	}

	successMessage := types.SuccessMessage{
		OrderID: orderMsg.OrderID,
		Status:  "success",
		Message: "Order successfully processed.",
	}

	successMessageBody, err := json.Marshal(successMessage)
	if err != nil {
		log.Printf("Error marshalling success message: %s", err)
		return err
	}

	if err := s.rabbitMQService.PublishMessage("success_queue", successMessageBody); err != nil {
		log.Printf("Error publishing success message to success queue: %s", err)
		return err
	}

	return nil
}

func (s *ApplicationService) UpdateInventoryProductQuantity(productID uint, quantityChange int) error {
	inventoryProduct, err := s.modelService.GetInventoryProductByProductID(productID)
	if err != nil {
		return fmt.Errorf("failed to get product from inventory: %v", err)
	}

	newQuantity := inventoryProduct.Quantity + quantityChange

	err = s.modelService.UpdateInventoryProductQuantity(productID, newQuantity)
	if err != nil {
		return fmt.Errorf("failed to update inventory quantity: %v", err)
	}

	log.Printf("Inventory for Product ID %d updated to %d", productID, newQuantity)
	return nil
}

func (s *ApplicationService) CheckProductStock(productID int) (int, error) {
	inventoryProduct, err := s.modelService.GetInventoryProductByProductID(uint(productID))
	if err != nil {
		return 0, fmt.Errorf("failed to get product from inventory: %v", err)
	}

	return inventoryProduct.Quantity, nil
}

func (s *ApplicationService) GetInventoryProduct(req types.InventoryProduct) (types.InventoryProduct, error) {
	connResp, err := s.modelService.GetInventoryProduct(int(req.ID))
	if err != nil {
		return types.InventoryProduct{}, err
	}

	return types.InventoryProduct{
		ID:        connResp.ID,
		ProductID: connResp.ProductID,
		Quantity:  connResp.Quantity,
	}, nil
}

func (s *ApplicationService) DeleteInventoryProduct(req types.InventoryProduct) error {
	err := s.modelService.DeleteInventoryProduct(req.ID)
	if err != nil {
		return err
	}

	return fmt.Errorf("Product deleted successfully")
}

func (s *ApplicationService) GetAllInventoryProducts() (types.ListInventoryProductsResponse, error) {
	var products types.ListInventoryProductsResponse
	connResp, err := s.modelService.GetAllInventoryProducts()
	if err != nil {
		return types.ListInventoryProductsResponse{}, err
	}

	for _, product := range connResp {
		response := types.InventoryProduct{
			ID:        product.ID,
			ProductID: product.ProductID,
			Quantity:  product.Quantity,
		}

		products.InventoryProducts = append(products.InventoryProducts, response)
	}

	return products, nil
}
