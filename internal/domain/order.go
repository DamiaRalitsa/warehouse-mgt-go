package domain

type Order struct {
	ID         int64 `json:"id"`
	UserID     int64 `json:"user_id"`
	ProductID  int64 `json:"product_id"`
	ShopID     int64 `json:"shop_id"`
	ReservedAt int64 `json:"reserved_at"` // unix timestamp
	Paid       bool  `json:"paid"`
}
