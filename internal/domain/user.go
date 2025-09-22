package domain

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone,omitempty"`
	Email string `json:"email,omitempty"`
}
