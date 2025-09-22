package http

import (
	"edot/internal/domain"
	"edot/internal/presenters"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	// Add usecase dependency here

}

func NewUserController() *UserController {
	return &UserController{}

}

func (uc *UserController) Create(c *fiber.Ctx) error {
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(presenters.Response{
			StatusCode: 400,
			Message:    "invalid request payload",
			Success:    false,
		})
	}
	// TODO: call usecase to create user
	return c.Status(201).JSON(presenters.Response{
		StatusCode: 201,
		Message:    "user created",
		Success:    true,
		Data:       user,
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
	// TODO: call usecase to get user by id
	user := domain.User{ID: id, Name: "Sample User"}
	return c.Status(200).JSON(presenters.Response{
		StatusCode: 200,
		Message:    "success",
		Success:    true,
		Data:       user,
	})
}

func (uc *UserController) List(c *fiber.Ctx) error {
	// TODO: call usecase to list users
	users := []domain.User{{ID: 1, Name: "User A"}, {ID: 2, Name: "User B"}}
	return c.Status(200).JSON(presenters.Response{
		StatusCode: 200,
		Message:    "success",
		Success:    true,
		Data:       users,
	})

}

// Login endpoint: authenticate user by phone or email
func (uc *UserController) Login(c *fiber.Ctx) error {
	// TODO: implement authentication logic
	return c.Status(200).JSON(presenters.Response{
		StatusCode: 200,
		Message:    "login successful (stub)",
		Success:    true,
	})
}
