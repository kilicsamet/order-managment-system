package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `gorm:"type:varchar(100);not null"`
	ImageURL    string  `gorm:"type:varchar(255);not null"`
	Description string  `gorm:"type:text"`
	Price       float64 `gorm:"type:decimal(10,2);not null"`
}

func (s *ModelService) CreateProduct(product Product) (Product, error) {
	err := s.db.Create(&product).Error
	return product, err
}

func (s *ModelService) GetProduct(id int) (*Product, error) {
	product := &Product{}
	tx := s.db.First(product, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return product, nil
}

func (s *ModelService) GetAllProducts() ([]*Product, error) {
	products := []*Product{}
	tx := s.db.Find(&products)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return products, nil
}

func (s *ModelService) UpdateProduct(product Product) error {
	return s.db.Save(&product).Error
}

func (s *ModelService) DeleteProduct(id int) error {
	return s.db.Delete(&Product{}, id).Error
}
