package model

type Product struct {
	Id    int64    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Quantity int `json:"quantity"`
	Category string `json:"category"`
	Description string `json:"description"`
}

type Products []Product