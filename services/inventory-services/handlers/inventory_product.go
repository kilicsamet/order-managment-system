package handlers

import (
	"encoding/json"
	"inventory-service/types"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (s *HandlerService) AddInventoryProduct(ctx *fiber.Ctx) error {
	var (
		productInventory types.InventoryProduct
	)

	if err := ctx.BodyParser(&productInventory); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	resp, err := s.applicationService.CreateInventoryProduct(productInventory)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(types.InventoryProduct{
		ID:        resp.ID,
		ProductID: resp.ProductID,
		Quantity:  resp.Quantity,
	})
}

func (s *HandlerService) UpdateInventoryProductQuantity(ctx *fiber.Ctx) error {
	productId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	var updateRequest types.UpdateQuantityRequest
	if err := ctx.BodyParser(&updateRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	updateRequest.ProductID = uint(productId)

	messageBody, err := json.Marshal(updateRequest)
	if err != nil {
		log.Println("Failed to marshal quantity update:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}
	rabbitMQService := s.applicationService.GetRabbitMQService()
	err = rabbitMQService.PublishMessage("inventory_update_queue", messageBody)
	if err != nil {
		log.Println("Failed to publish message:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(types.SuccessResponse{
		Message: "Inventory quantity update message sent to queue successfully",
	})
}

func (s *HandlerService) GetInventoryProducts(ctx *fiber.Ctx) error {
	resp, err := s.applicationService.GetAllInventoryProducts()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	var products []types.InventoryProduct
	for _, product := range resp.InventoryProducts {
		products = append(products, types.InventoryProduct{
			ID:        product.ID,
			ProductID: product.ProductID,
			Quantity:  product.Quantity,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(types.ListInventoryProductsResponse{
		InventoryProducts: products,
	})
}

func (s *HandlerService) GetInventoryProduct(ctx *fiber.Ctx) error {
	productId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	product, err := s.applicationService.GetInventoryProduct(types.InventoryProduct{ID: uint(productId)})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	// Ürünü döndür
	return ctx.Status(fiber.StatusOK).JSON(types.InventoryProduct{
		ID:        product.ID,
		ProductID: product.ProductID,
		Quantity:  product.Quantity,
	})
}

func (s *HandlerService) DeleteInventoryProduct(ctx *fiber.Ctx) error {
	productId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	err = s.applicationService.DeleteInventoryProduct(types.InventoryProduct{ID: uint(productId)})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(types.SuccessResponse{
		Message: "Product deleted successfully",
	})
}
