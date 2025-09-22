package domain

type Warehouse struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	ShopID int64  `json:"shop_id"`
	Active bool   `json:"active"`
}
