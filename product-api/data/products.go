package data

import (
	"time"
)

type Product struct {
	ID          int `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       float32 `json:"price"`
	SKU         string `json:"-"`
	CreatedOn   string `json:"-"`
	UpdatedOn   string `json:"-"`
}

func GetProducts() []*Product {
	return productList
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Milky Coffee",
		Price:       2.45,
		SKU: "fajkfja",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "String Coffee",
		Price:       1.99,
		SKU: "e4ggaf",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
