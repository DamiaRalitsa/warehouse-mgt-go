# Warehouse Management API

A technical test project, built with Go. This API manages users, products, orders, shops, and warehouses, using PostgreSQL and Fiber.

## Features
- User authentication (login by phone/email)
- Product listing with stock availability
- Order checkout with stock reservation and release (background job)
- Shop management with warehouse association
- Warehouse stock management
- Product transfer between warehouses
- Activate/Deactivate warehouses (inactive stock excluded)
- RESTful endpoints
- Dockerized for easy deployment

## Project Structure
```
go.mod
go.sum
Dockerfile
Makefile
cmd/app/main.go
deployments/docker-compose.yml
internal/
  delivery/http/
    user_controller.go
    product_controller.go
    order_controller.go
    shop_controller.go
    warehouse_controller.go
    routes/route.go
  domain/
    user.go
    product.go
    order.go
    shop.go
    warehouse.go
  repositories/
    user/
    product/
    order/
    shop/
    warehouse/
  usecases/
    user/
    product/
    order/
    shop/
    warehouse/
pkg/postgres/connection.go
```

## Getting Started

### Prerequisites
- Go 1.20+
- Docker
- PostgreSQL

### Setup
1. Clone the repository:
   ```bash
   git clone https://github.com/DamiaRalitsa/warehouse-mgt-go.git
   cd warehouse-mgt-go
   ```
2. Configure your database connection in `config.json`.
3. Build and run with Docker:
   ```bash
   docker build -t warehouse-mgt-app .
   docker run -it --rm -p 8334:8334 -v $(pwd)/config.json:/config.json warehouse-mgt-app

   or simply run "make run" on terminal
   ```
4. Or run locally:
   ```bash
   go run cmd/app/main.go
   ```

### API Endpoints
#### User
- `POST /api/user/register` - Register user
- `POST /api/user/login` - Login with phone/email
- `GET /api/user/:id` - Get user by ID
- `GET /api/user/list` - List users

#### Product
- `GET /api/product/list` - List products with stock
- `GET /api/product/:id` - Get product by ID
- `POST /api/product/create` - Create product

#### Order
- `POST /api/order/checkout` - Place order, reserve stock
- `POST /api/order/release-stock` - Release reserved stock
- `GET /api/order/:id` - Get order by ID
- `GET /api/order/list` - List orders

#### Shop
- `POST /api/shop/create` - Create shop
- `GET /api/shop/:id` - Get shop by ID
- `GET /api/shop/list` - List shops
- `GET /api/shop/:id/warehouses` - List warehouses for shop

#### Warehouse
- `POST /api/warehouse/create` - Create warehouse
- `GET /api/warehouse/:id` - Get warehouse by ID
- `GET /api/warehouse/list` - List warehouses
- `POST /api/warehouse/transfer` - Transfer products between warehouses
- `POST /api/warehouse/:id/activate` - Activate warehouse
- `POST /api/warehouse/:id/deactivate` - Deactivate warehouse

## Contributing
Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

## License
MIT
