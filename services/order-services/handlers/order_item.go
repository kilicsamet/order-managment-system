package handlers

import (
	"order-services/types"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (s *HandlerService) AddOrderItem(ctx *fiber.Ctx) error {
	var (
		orderItem types.OrderItem
	)

	if err := ctx.BodyParser(&orderItem); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	resp, err := s.applicationService.CreateOrderItem(orderItem)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(types.OrderItem{
		ID:        resp.ID,
		OrderID:   resp.OrderID,
		ProductID: resp.ProductID,
		Quantity:  resp.Quantity,
	})
}

func (s *HandlerService) GetOrderItems(ctx *fiber.Ctx) error {
	resp, err := s.applicationService.GetAllOrderItems()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	var orderItems []types.OrderItem
	for _, orderItem := range resp.OrderItems {
		orderItems = append(orderItems, types.OrderItem{
			ID:        orderItem.ID,
			OrderID:   orderItem.OrderID,
			ProductID: orderItem.ProductID,
			Quantity:  orderItem.Quantity,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(types.ListOrderItemsResponse{
		OrderItems: orderItems,
	})
}

func (s *HandlerService) GetOrderItem(ctx *fiber.Ctx) error {
	orderItemId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	orderItem, err := s.applicationService.GetOrderItem(types.OrderItem{ID: uint(orderItemId)})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(types.OrderItem{
		ID:        orderItem.ID,
		OrderID:   orderItem.OrderID,
		ProductID: orderItem.ProductID,
		Quantity:  orderItem.Quantity,
	})
}
