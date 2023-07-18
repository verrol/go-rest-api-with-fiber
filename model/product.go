package model

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Quantity int `json:"quantity"`
	Category string `json:"category"`
	Description string `json:"description"`
}

type Products []Product