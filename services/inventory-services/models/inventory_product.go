package models

import (
	"fmt"

	"gorm.io/gorm"
)

type InventoryProduct struct {
	gorm.Model
	ProductID uint `gorm:"type:int;not null"`
	Quantity  int  `gorm:"type:int;not null"`
}

func (s *ModelService) CreateInventoryProduct(inventoryProduct InventoryProduct) (InventoryProduct, error) {
	if err := s.db.Create(&inventoryProduct).Error; err != nil {
		return InventoryProduct{}, fmt.Errorf("failed to create inventory product: %w", err)
	}
	return inventoryProduct, nil
}

func (s *ModelService) GetInventoryProductByProductID(productID uint) (*InventoryProduct, error) {
	inventoryProduct := &InventoryProduct{}
	tx := s.db.Where("product_id = ?", productID).First(inventoryProduct)
	if tx.Error != nil {
		return nil, fmt.Errorf("product with id %d not found: %w", productID, tx.Error)
	}
	return inventoryProduct, nil
}

func (s *ModelService) GetInventoryProduct(id int) (*InventoryProduct, error) {
	product := &InventoryProduct{}
	tx := s.db.First(product, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return product, nil
}

func (s *ModelService) UpdateInventoryProductQuantity(productID uint, quantity int) error {
	if quantity < 0 {
		return fmt.Errorf("quantity cannot be less than zero")
	}

	inventoryProduct := &InventoryProduct{}
	tx := s.db.Where("product_id = ?", productID).First(inventoryProduct)
	if tx.Error != nil {
		return fmt.Errorf("product with id %d not found: %w", productID, tx.Error)
	}

	inventoryProduct.Quantity = quantity

	tx = s.db.Begin()
	if tx.Error != nil {
		return fmt.Errorf("failed to start transaction: %w", tx.Error)
	}

	if err := tx.Save(inventoryProduct).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update inventory product quantity: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (s *ModelService) DeleteInventoryProduct(productID uint) error {
	if err := s.db.Where("product_id = ?", productID).Delete(&InventoryProduct{}).Error; err != nil {
		return fmt.Errorf("failed to delete inventory product: %w", err)
	}
	return nil
}

func (s *ModelService) GetAllInventoryProducts() ([]*InventoryProduct, error) {
	inventoryProducts := []*InventoryProduct{}
	tx := s.db.Find(&inventoryProducts)
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to retrieve inventory products: %w", tx.Error)
	}
	return inventoryProducts, nil
}
