package application

import (
	"fmt"
	"product-services/models"
	"product-services/types"
)

func (s *ApplicationService) CreateProduct(req types.Product) (types.Product, error) {

	product := models.Product{
		Name:        req.Name,
		ImageURL:    req.ImageURL,
		Description: req.Description,
		Price:       req.Price,
	}

	createdProduct, err := s.modelService.CreateProduct(product)
	if err != nil {
		return types.Product{}, err
	}

	return types.Product{
		ID:          createdProduct.ID,
		Name:        createdProduct.Name,
		ImageURL:    createdProduct.ImageURL,
		Description: createdProduct.Description,
		Price:       createdProduct.Price,
	}, nil
}

func (s *ApplicationService) GetProduct(req types.Product) (types.Product, error) {
	connResp, err := s.modelService.GetProduct(int(req.ID))
	if err != nil {
		return types.Product{}, err
	}

	return types.Product{
		ID:          connResp.ID,
		Name:        connResp.Name,
		ImageURL:    connResp.ImageURL,
		Description: connResp.Description,
		Price:       connResp.Price,
	}, nil
}

func (s *ApplicationService) UpdateProduct(req types.Product) (types.Product, error) {

	product, err := s.modelService.GetProduct(int(req.ID))
	if err != nil {
		return types.Product{}, err
	}

	product.ID = req.ID
	product.Name = req.Name
	product.ImageURL = req.ImageURL
	product.Description = req.Description
	product.Price = req.Price

	err = s.modelService.UpdateProduct(*product)
	if err != nil {
		return types.Product{}, err
	}

	return types.Product{
		ID:          product.ID,
		Name:        product.Name,
		ImageURL:    product.ImageURL,
		Description: product.Description,
		Price:       product.Price,
	}, nil
}

func (s *ApplicationService) DeleteProduct(req types.Product) error {
	err := s.modelService.DeleteProduct(int(req.ID))
	if err != nil {
		return err
	}

	return fmt.Errorf("Product deleted successfully")
}

func (s *ApplicationService) GetAllProducts() (types.ListProductsResponse, error) {
	var products types.ListProductsResponse
	connResp, err := s.modelService.GetAllProducts()
	if err != nil {
		return types.ListProductsResponse{}, err
	}

	for _, product := range connResp {
		response := types.Product{
			ID:          product.ID,
			Name:        product.Name,
			ImageURL:    product.ImageURL,
			Description: product.Description,
			Price:       product.Price,
		}

		products.Products = append(products.Products, response)
	}

	return products, nil
}
