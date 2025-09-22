package http

import (
	"edot/internal/domain"
	"edot/internal/presenters"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ShopController struct{}

func NewShopController() *ShopController {
	return &ShopController{}
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
