package http

import (
	"edot/internal/domain"
	"edot/internal/presenters"
	"edot/internal/usecases/user"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	Usecase *user.UserUsecase
}

func NewUserController() *UserController {
	return &UserController{
		Usecase: user.NewUserUsecase(),
	}
}

func (uc *UserController) Create(c *fiber.Ctx) error {
	var u domain.User
	if err := c.BodyParser(&u); err != nil {
		return c.Status(400).JSON(presenters.Response{
			StatusCode: 400,
			Message:    "invalid request payload",
			Success:    false,
		})
	}
	if err := uc.Usecase.Create(&u); err != nil {
		return c.Status(500).JSON(presenters.Response{
			StatusCode: 500,
			Message:    "failed to create user",
			Success:    false,
			Error:      err.Error(),
		})
	}
	return c.Status(201).JSON(presenters.Response{
		StatusCode: 201,
		Message:    "user created",
		Success:    true,
		Data:       u,
	})
}

func (uc *UserController) GetByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(presenters.Response{
			StatusCode: 400,
			Message:    "invalid user id",
			Success:    false,
		})
	}
	user, err := uc.Usecase.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(presenters.Response{
			StatusCode: 404,
			Message:    "user not found",
			Success:    false,
			Error:      err.Error(),
		})
	}
	return c.Status(200).JSON(presenters.Response{
		StatusCode: 200,
		Message:    "success",
		Success:    true,
		Data:       user,
	})
}

func (uc *UserController) List(c *fiber.Ctx) error {
	users, err := uc.Usecase.GetAll()
	if err != nil {
		return c.Status(500).JSON(presenters.Response{
			StatusCode: 500,
			Message:    "failed to fetch users",
			Success:    false,
			Error:      err.Error(),
		})
	}
	return c.Status(200).JSON(presenters.Response{
		StatusCode: 200,
		Message:    "success",
		Success:    true,
		Data:       users,
	})
}

// Login endpoint: authenticate user by phone or email
func (uc *UserController) Login(c *fiber.Ctx) error {
	var req struct {
		Identifier string `json:"identifier"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(presenters.Response{
			StatusCode: 400,
			Message:    "invalid request payload",
			Success:    false,
		})
	}
	user, err := uc.Usecase.Login(req.Identifier)
	if err != nil {
		return c.Status(401).JSON(presenters.Response{
			StatusCode: 401,
			Message:    "authentication failed",
			Success:    false,
			Error:      err.Error(),
		})
	}
	return c.Status(200).JSON(presenters.Response{
		StatusCode: 200,
		Message:    "login successful",
		Success:    true,
		Data:       user,
	})
}
