package http

import (
	"edot/internal/domain"
	"edot/internal/presenters"
	"edot/internal/usecases/warehouse"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type WarehouseController struct {
	Usecase *warehouse.WarehouseUsecase
}

func NewWarehouseController() *WarehouseController {
	return &WarehouseController{
		Usecase: warehouse.NewWarehouseUsecase(),
	}
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

// TransferProduct endpoint: transfer products between warehouses
func (wc *WarehouseController) TransferProduct(c *fiber.Ctx) error {
	var req struct {
		ProductID       int64 `json:"product_id"`
		FromWarehouseID int64 `json:"from_warehouse_id"`
		ToWarehouseID   int64 `json:"to_warehouse_id"`
		Quantity        int   `json:"quantity"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(presenters.Response{
			StatusCode: 400,
			Message:    "invalid request payload",
			Success:    false,
		})
	}
	err := wc.Usecase.TransferProduct(req.ProductID, req.FromWarehouseID, req.ToWarehouseID, req.Quantity)
	if err != nil {
		return c.Status(500).JSON(presenters.Response{
			StatusCode: 500,
			Message:    "failed to transfer product",
			Success:    false,
			Error:      err.Error(),
		})
	}
	return c.Status(200).JSON(presenters.Response{
		StatusCode: 200,
		Message:    "product transferred",
		Success:    true,
	})
}

// Activate endpoint: activate a warehouse
func (wc *WarehouseController) Activate(c *fiber.Ctx) error {
	warehouseID, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(presenters.Response{
			StatusCode: 400,
			Message:    "invalid warehouse id",
			Success:    false,
		})
	}
	err = wc.Usecase.SetActiveStatus(warehouseID, true)
	if err != nil {
		return c.Status(500).JSON(presenters.Response{
			StatusCode: 500,
			Message:    "failed to activate warehouse",
			Success:    false,
			Error:      err.Error(),
		})
	}
	return c.Status(200).JSON(presenters.Response{
		StatusCode: 200,
		Message:    "warehouse activated",
		Success:    true,
	})
}

// Deactivate endpoint: deactivate a warehouse
func (wc *WarehouseController) Deactivate(c *fiber.Ctx) error {
	warehouseID, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(presenters.Response{
			StatusCode: 400,
			Message:    "invalid warehouse id",
			Success:    false,
		})
	}
	err = wc.Usecase.SetActiveStatus(warehouseID, false)
	if err != nil {
		return c.Status(500).JSON(presenters.Response{
			StatusCode: 500,
			Message:    "failed to deactivate warehouse",
			Success:    false,
			Error:      err.Error(),
		})
	}
	return c.Status(200).JSON(presenters.Response{
		StatusCode: 200,
		Message:    "warehouse deactivated",
		Success:    true,
	})
}
