package http

import (
	"edot/internal/domain"
	"edot/internal/presenters"
	"edot/internal/usecases/warehouse"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ShopController struct {
	WarehouseUsecase *warehouse.WarehouseUsecase
}

func NewShopController() *ShopController {
	return &ShopController{
		WarehouseUsecase: warehouse.NewWarehouseUsecase(),
	}
}

func (sc *ShopController) Create(c *fiber.Ctx) error {
	var shop domain.Shop
	if err := c.BodyParser(&shop); err != nil {
		return c.Status(400).JSON(presenters.Response{
			StatusCode: 400,
			Message:    "invalid request payload",
			Success:    false,
		})
	}
	return c.Status(201).JSON(presenters.Response{
		StatusCode: 201,
		Message:    "shop created",
		Success:    true,
		Data:       shop,
	})
}

func (sc *ShopController) GetByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(presenters.Response{
			StatusCode: 400,
			Message:    "invalid shop id",
			Success:    false,
		})
	}
	shop := domain.Shop{ID: id, Name: "Sample Shop"}
	return c.Status(200).JSON(presenters.Response{
		StatusCode: 200,
		Message:    "success",
		Success:    true,
		Data:       shop,
	})
}

func (sc *ShopController) List(c *fiber.Ctx) error {
	shops := []domain.Shop{{ID: 1, Name: "Shop"}}
	return c.Status(200).JSON(presenters.Response{
		StatusCode: 200,
		Message:    "success",
		Success:    true,
		Data:       shops,
	})

}

// ListWarehouses endpoint: list warehouses for a shop
func (sc *ShopController) ListWarehouses(c *fiber.Ctx) error {
	shopID, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(presenters.Response{
			StatusCode: 400,
			Message:    "invalid shop id",
			Success:    false,
		})
	}
	warehouses, err := sc.WarehouseUsecase.ListByShopID(shopID)
	if err != nil {
		return c.Status(500).JSON(presenters.Response{
			StatusCode: 500,
			Message:    "failed to fetch warehouses",
			Success:    false,
			Error:      err.Error(),
		})
	}
	return c.Status(200).JSON(presenters.Response{
		StatusCode: 200,
		Message:    "success",
		Success:    true,
		Data:       warehouses,
	})
}
