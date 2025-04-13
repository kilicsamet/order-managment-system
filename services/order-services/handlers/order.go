package handlers

import (
	"order-services/types"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (s *HandlerService) AddOrder(ctx *fiber.Ctx) error {
	var (
		order types.Order
	)

	if err := ctx.BodyParser(&order); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	resp, err := s.applicationService.CreateOrder(order)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(types.Order{
		ID:         resp.ID,
		Status:     resp.Status,
		OrderItems: resp.OrderItems,
	})
}

func (s *HandlerService) GetOrders(ctx *fiber.Ctx) error {
	resp, err := s.applicationService.GetAllOrders()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	var orders []types.Orders
	for _, order := range resp.Orders {
		orders = append(orders, types.Orders{
			ID:     order.ID,
			Status: order.Status,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(types.ListOrdersResponse{
		Orders: orders,
	})
}

func (s *HandlerService) GetOrder(ctx *fiber.Ctx) error {
	orderId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	order, err := s.applicationService.GetOrder(types.Order{ID: uint(orderId)})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(types.Order{
		ID:     order.ID,
		Status: order.Status,
	})
}

func (s *HandlerService) DeleteOrder(ctx *fiber.Ctx) error {
	orderId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	err = s.applicationService.DeleteOrder(types.Order{ID: uint(orderId)})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(types.SuccessResponse{
		Message: "Order deleted successfully",
	})
}
