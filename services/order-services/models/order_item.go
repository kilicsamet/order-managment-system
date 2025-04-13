package models

import (
	"fmt"

	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	OrderID   uint `gorm:"not null"`
	ProductID uint `gorm:"not null"`
	Quantity  int  `gorm:"type:int;not null"`
}

func (s *ModelService) CreateOrderItem(orderItem OrderItem) (OrderItem, error) {
	err := s.db.Create(&orderItem).Error
	if err != nil {
		return OrderItem{}, fmt.Errorf("failed to create order item: %w", err)
	}
	return orderItem, nil
}

func (s *ModelService) GetOrderItem(id uint) (*OrderItem, error) {
	orderItem := &OrderItem{}
	tx := s.db.First(orderItem, id)
	if tx.Error != nil {
		return nil, fmt.Errorf("order item with id %d not found: %w", id, tx.Error)
	}
	return orderItem, nil
}

func (s *ModelService) GetAllOrderItems() ([]*OrderItem, error) {
	orderItems := []*OrderItem{}
	tx := s.db.Find(&orderItems)
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to retrieve order items: %w", tx.Error)
	}
	return orderItems, nil
}

func (s *ModelService) GetOrderItemsByOrderID(orderID uint) ([]*OrderItem, error) {
	orderItems := []*OrderItem{}
	tx := s.db.Where("order_id = ?", orderID).Find(&orderItems)
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to retrieve order items for order_id %d: %w", orderID, tx.Error)
	}
	return orderItems, nil
}

func (s *ModelService) DeleteOrderItem(id uint) error {
	if err := s.db.Delete(&OrderItem{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete order item: %w", err)
	}
	return nil
}
