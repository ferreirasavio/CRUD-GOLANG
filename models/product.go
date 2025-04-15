package models

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Quantity    int     `json:"quantity"`
	Description string  `json:"description"`
}
