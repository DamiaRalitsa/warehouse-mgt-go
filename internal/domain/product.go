package domain

type Product struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Stock    int    `json:"stock"`
	Reserved int    `json:"reserved"`
}
