package data

import (
	"encoding/json"
	"fmt"
	"io"
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

func (p *Product) FromJSON(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(p)

}

var ErrProductNotFound = fmt.Errorf("product not found")

func UpdateProduct(id int, p *Product) (plist Products, ind int, err error) {
	for index, prod := range productList {
		if prod.ID == id {
			p.ID = id
			productList[index] = p

		}
		return productList, index, nil
	}
	return nil, -1, ErrProductNotFound

}
func AddProduct(p *Product) {
	p.ID = getProductId()
	productList = append(productList, p)

}
func getProductId() int {
	id := productList[len(productList)-1]
	nextID := id.ID + 1
	return nextID
}
func (p *Products) ToJSON(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
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
