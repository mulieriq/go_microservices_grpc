package data

import (
	"encoding/json"
	"io"
	"log"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	log.Println("here the data is ready .....")
	return encoder.Encode(p)
}

func GetProducts() Products {
	return productList

}

var productList = []*Product{

	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		SKU:         "abed323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		Price:       2.45,
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		SKU:         "fjk323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		Price:       1.99,
	},
}
