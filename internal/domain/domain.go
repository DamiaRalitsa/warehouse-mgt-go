package domain

// Common domain interfaces for all entities

type UserRepository interface {
	Create(user *User) error
	GetByID(id int64) (*User, error)
	List() ([]User, error)
}

type ProductRepository interface {
	Create(product *Product) error
	GetByID(id int64) (*Product, error)
	List() ([]Product, error)
}

type OrderRepository interface {
	Create(order *Order) error
	GetByID(id int64) (*Order, error)
	List() ([]Order, error)
}

type ShopRepository interface {
	Create(shop *Shop) error
	GetByID(id int64) (*Shop, error)
	List() ([]Shop, error)
}

type WarehouseRepository interface {
	Create(warehouse *Warehouse) error
	GetByID(id int64) (*Warehouse, error)
	List() ([]Warehouse, error)
}
