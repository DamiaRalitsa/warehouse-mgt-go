package http

import (
	"edot/internal/domain"
	"edot/internal/presenters"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type OrderController struct{}

func NewOrderController() *OrderController {
	return &OrderController{}
}

func (oc *OrderController) Create(c *fiber.Ctx) error {
	var order domain.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(presenters.Response{
			StatusCode: 400,
			Message:    "invalid request payload",
			Success:    false,
		})
	}
	return c.Status(201).JSON(presenters.Response{
		StatusCode: 201,
		Message:    "order created",
		Success:    true,
		Data:       order,
	})
}

func (oc *OrderController) GetByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(presenters.Response{
			StatusCode: 400,
			Message:    "invalid order id",
			Success:    false,
		})
	}
	order := domain.Order{ID: id, UserID: 1, ProductID: 1, ShopID: 1}
	return c.Status(200).JSON(presenters.Response{
		StatusCode: 200,
		Message:    "success",
		Success:    true,
		Data:       order,
	})
}

func (oc *OrderController) List(c *fiber.Ctx) error {
	orders := []domain.Order{{ID: 1, UserID: 1, ProductID: 1, ShopID: 1}}
	return c.Status(200).JSON(presenters.Response{
		StatusCode: 200,
		Message:    "success",
		Success:    true,
		Data:       orders,
	})
}
