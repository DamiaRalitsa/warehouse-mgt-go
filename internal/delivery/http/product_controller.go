package http

import (
	"edot/internal/domain"
	"edot/internal/presenters"
	"edot/internal/usecases/product"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	Usecase *product.ProductUsecase
}

func NewProductController() *ProductController {
	return &ProductController{
		Usecase: product.NewProductUsecase(),
	}
}

func (pc *ProductController) Create(c *fiber.Ctx) error {
	var product domain.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(presenters.Response{
			StatusCode: 400,
			Message:    "invalid request payload",
			Success:    false,
		})
	}
	return c.Status(201).JSON(presenters.Response{
		StatusCode: 201,
		Message:    "product created",
		Success:    true,
		Data:       product,
	})
}

func (pc *ProductController) GetByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(presenters.Response{
			StatusCode: 400,
			Message:    "invalid product id",
			Success:    false,
		})
	}
	product := domain.Product{ID: id, Name: "Sample Product"}
	return c.Status(200).JSON(presenters.Response{
		StatusCode: 200,
		Message:    "success",
		Success:    true,
		Data:       product,
	})
}

func (pc *ProductController) List(c *fiber.Ctx) error {
	products, err := pc.Usecase.GetAll()
	if err != nil {
		return c.Status(500).JSON(presenters.Response{
			StatusCode: 500,
			Message:    "failed to fetch products",
			Success:    false,
			Error:      err.Error(),
		})
	}
	return c.Status(200).JSON(presenters.Response{
		StatusCode: 200,
		Message:    "success",
		Success:    true,
		Data:       products,
	})
}
