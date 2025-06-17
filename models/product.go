package models

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Quantity    int     `json:"quantity"`
	Description string  `json:"description"`
}

type PaginatedProducts struct {
	Page    int       `json:"page"`
	PerPage int       `json:"per_page"`
	Items   []Product `json:"items"`
}
