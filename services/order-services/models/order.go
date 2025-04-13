package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Status     string      `gorm:"type:varchar(20);not null"`
	OrderItems []OrderItem `json:"order_items"`
}

func (s *ModelService) CreateOrder(order Order) (Order, error) {
	err := s.db.Create(&order).Error
	return order, err
}

func (s *ModelService) GetOrder(id int) (*Order, error) {
	order := &Order{}
	tx := s.db.First(order, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return order, nil
}

func (s *ModelService) GetAllOrders() ([]*Order, error) {
	orders := []*Order{}
	tx := s.db.Find(&orders)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return orders, nil
}

func (s *ModelService) UpdateOrder(order Order) error {
	return s.db.Save(&order).Error
}

func (s *ModelService) DeleteOrder(id int) error {
	return s.db.Delete(&Order{}, id).Error
}
