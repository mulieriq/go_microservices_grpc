package data

import (
	"encoding/json"
	"fmt"
<<<<<<< HEAD
	"github.com/go-playground/validator"
	"io"
	"regexp"
=======
	"io"
>>>>>>> 1b73e57b469413ff33493d4ddab4933c271fceb2
	"time"
)

type Product struct {
	ID          int     `json:"id"`
<<<<<<< HEAD
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
=======
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"-"`
>>>>>>> 1b73e57b469413ff33493d4ddab4933c271fceb2
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
}
type Products []*Product

<<<<<<< HEAD
func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}
func validateSKU(fl validator.FieldLevel) bool {
	reg := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := reg.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}
	return true
}

=======
>>>>>>> 1b73e57b469413ff33493d4ddab4933c271fceb2
func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id) ///searched product and position kwa list
	if err != nil {
		fmt.Println("error man")
		return err
	}
	p.ID = id            //re assing pid
	productList[pos] = p //products postion replaced with updated product

	return nil
}
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

var ErrProductNotFound = fmt.Errorf("Product Not Found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {

		if p.ID == id {
			fmt.Println("data", p)
			return p, i, nil
		}

	}
	return nil, -1, ErrProductNotFound

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
