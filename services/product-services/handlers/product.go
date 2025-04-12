package handlers

import (
	"product-services/types"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (s *HandlerService) AddProduct(ctx *fiber.Ctx) error {
	var (
		product types.Product
	)

	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	resp, err := s.applicationService.CreateProduct(product)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(types.Product{
		ID:          resp.ID,
		Name:        resp.Name,
		ImageURL:    resp.ImageURL,
		Description: resp.Description,
		Price:       resp.Price,
	})
}

func (s *HandlerService) GetProducts(ctx *fiber.Ctx) error {
	resp, err := s.applicationService.GetAllProducts()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	var products []types.Product
	for _, product := range resp.Products {
		products = append(products, types.Product{
			ID:          product.ID,
			Name:        product.Name,
			ImageURL:    product.ImageURL,
			Description: product.Description,
			Price:       product.Price,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(types.ListProductsResponse{
		Products: products,
	})
}

func (s *HandlerService) GetProduct(ctx *fiber.Ctx) error {
	productId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	product, err := s.applicationService.GetProduct(types.Product{ID: uint(productId)})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	// Ürünü döndür
	return ctx.Status(fiber.StatusOK).JSON(types.Product{
		ID:          product.ID,
		Name:        product.Name,
		ImageURL:    product.ImageURL,
		Description: product.Description,
		Price:       product.Price,
	})
}

func (s *HandlerService) UpdateProduct(ctx *fiber.Ctx) error {
	productId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	var product types.Product
	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	if product.ID != uint(productId) {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Error: "Product ID in URL and body must match",
		})
	}

	// UpdateProduct fonksiyonunu çağır
	updatedProduct, err := s.applicationService.UpdateProduct(product)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	// Güncellenmiş ürünü döndür
	return ctx.Status(fiber.StatusOK).JSON(types.Product{
		ID:          updatedProduct.ID,
		Name:        updatedProduct.Name,
		ImageURL:    updatedProduct.ImageURL,
		Description: updatedProduct.Description,
		Price:       updatedProduct.Price,
	})
}

func (s *HandlerService) DeleteProduct(ctx *fiber.Ctx) error {
	productId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	err = s.applicationService.DeleteProduct(types.Product{ID: uint(productId)})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	// Silme başarılı olursa mesaj döndürüyoruz
	return ctx.Status(fiber.StatusOK).JSON(types.SuccessResponse{
		Message: "Product deleted successfully",
	})
}
