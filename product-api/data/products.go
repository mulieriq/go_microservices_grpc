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
	SKU         string  `json:"-"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
}
type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
func (p *Product) FromJSON(at io.Reader) error {
	e := json.NewDecoder(at)
	return e.Decode(p)
}

func GetProducts() Products {
	return productList
}
func AddProduct(p *Product) {
	p.ID = getNextId()
	productList = append(productList, p)

}
func updateProduct(id int, p *Product) error {
    _,pos,err :=findProduct(id) ///searched product and position kwa list
	if err !=nil{
		return err

	}
	p.ID=id //re assing pid
	productList[pos] = p //products postion replaced with updated product
}

var ErrProductNotFound = fmt.Errorf("Product Not Found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {

		if p.ID == id {
			return p, i, nil
		}

	}
	return nil, 0, ErrProductNotFound

}

func getNextId() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Milky Coffee",
		Price:       2.45,
		SKU:         "fajkfja",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "String Coffee",
		Price:       1.99,
		SKU:         "e4ggaf",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
