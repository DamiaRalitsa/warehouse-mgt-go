package routes

import (
	"encoding/json"
	"os"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rs/zerolog"

	"edot/internal/delivery/http"
)

type RouteConfig struct {
	App                 *fiber.App
	UserController      *http.UserController
	ProductController   *http.ProductController
	OrderController     *http.OrderController
	ShopController      *http.ShopController
	WarehouseController *http.WarehouseController
}

func NewRouteConfig() *RouteConfig {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	app := fiber.New(fiber.Config{
		Prefork:     false,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		BodyLimit:   100 * 1024 * 1024,
	})

	app.Use(func(c *fiber.Ctx) error {
		if c.Path() == "/metrics" {
			return c.Next()
		}
		return fiberzerolog.New(fiberzerolog.Config{
			Logger: &logger,
		})(c)
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	userController := http.NewUserController()
	productController := http.NewProductController()
	orderController := http.NewOrderController()
	shopController := http.NewShopController()
	warehouseController := http.NewWarehouseController()

	routeConfig := RouteConfig{
		App:                 app,
		UserController:      userController,
		ProductController:   productController,
		OrderController:     orderController,
		ShopController:      shopController,
		WarehouseController: warehouseController,
	}

	routeConfig.SetupRoute()
	return &routeConfig
}

func (rc *RouteConfig) SetupRoute() {
	userGroup := rc.App.Group("/api/user")
	userGroup.Post("/create", rc.UserController.Create)
	userGroup.Get("/:id", rc.UserController.GetByID)
	userGroup.Get("/list", rc.UserController.List)

	productGroup := rc.App.Group("/api/product")
	productGroup.Post("/create", rc.ProductController.Create)
	productGroup.Get("/:id", rc.ProductController.GetByID)
	productGroup.Get("/list", rc.ProductController.List)

	orderGroup := rc.App.Group("/api/order")
	orderGroup.Post("/create", rc.OrderController.Create)
	orderGroup.Get("/:id", rc.OrderController.GetByID)
	orderGroup.Get("/list", rc.OrderController.List)

	shopGroup := rc.App.Group("/api/shop")
	shopGroup.Post("/create", rc.ShopController.Create)
	shopGroup.Get("/:id", rc.ShopController.GetByID)
	shopGroup.Get("/list", rc.ShopController.List)

	warehouseGroup := rc.App.Group("/api/warehouse")
	warehouseGroup.Post("/create", rc.WarehouseController.Create)
	warehouseGroup.Get("/:id", rc.WarehouseController.GetByID)
	warehouseGroup.Get("/list", rc.WarehouseController.List)
}

func (rc *RouteConfig) Listen(address string) {
	rc.App.Listen(address)
}
