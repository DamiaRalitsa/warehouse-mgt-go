package http

import (
	"edot/internal/domain"
	"edot/internal/presenters"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type WarehouseController struct{}

func NewWarehouseController() *WarehouseController {
	return &WarehouseController{}
}

func (wc *WarehouseController) Create(c *fiber.Ctx) error {
	var warehouse domain.Warehouse
	if err := c.BodyParser(&warehouse); err != nil {
		return c.Status(400).JSON(presenters.Response{
			StatusCode: 400,
			Message:    "invalid request payload",
			Success:    false,
		})
	}
	return c.Status(201).JSON(presenters.Response{
		StatusCode: 201,
		Message:    "warehouse created",
		Success:    true,
		Data:       warehouse,
	})
}

func (wc *WarehouseController) GetByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(presenters.Response{
			StatusCode: 400,
			Message:    "invalid warehouse id",
			Success:    false,
		})
	}
	warehouse := domain.Warehouse{ID: id, Name: "Sample Warehouse"}
	return c.Status(200).JSON(presenters.Response{
		StatusCode: 200,
		Message:    "success",
		Success:    true,
		Data:       warehouse,
	})
}

func (wc *WarehouseController) List(c *fiber.Ctx) error {
	warehouses := []domain.Warehouse{{ID: 1, Name: "Warehouse 1"}, {ID: 2, Name: "Warehouse 2"}}
	return c.Status(200).JSON(presenters.Response{
		StatusCode: 200,
		Message:    "success",
		Success:    true,
		Data:       warehouses,
	})
}
