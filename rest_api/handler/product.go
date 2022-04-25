package handler

import (
	"log"
	"net/http"
	"regexp"
	"rest_api/data"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}

	if r.Method == http.MethodPut {
		regex := `/([0-9]+)`
		urlPath := r.URL.Path
		reg := regexp.MustCompile(regex)
		groups := reg.FindAllStringSubmatch(urlPath, -1)
		p.l.Println("groups", groups)

		if len(groups) != 1 {
			http.Error(w, "more than 1 id", http.StatusBadRequest)
			return
		}

		if len(groups[0]) != 2 {
			http.Error(w, "More than one capture group", http.StatusBadRequest)
			return
		}

		productId := groups[0][1]
		productIdS, err := strconv.Atoi(productId)

		if err != nil {
			http.Error(w, "Unable to convert id", http.StatusInternalServerError)
			return
		}
		p.updateProduct(w, r, productIdS)

		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Products) updateProduct(w http.ResponseWriter, r *http.Request, prodId int) {

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal", http.StatusBadRequest)
	}

	plist, _, err2 := data.UpdateProduct(prodId, prod)
	if err2 == data.ErrProductNotFound {
		http.Error(w, "PRODUCT NOT FOUND", http.StatusInternalServerError)
		return
	}
	err1 := plist.ToJSON(w)
	if err1 != nil {
		http.Error(w, "Unable to marshal json data", http.StatusInternalServerError)
		return
	}

}

func (p *Products) getProducts(w http.ResponseWriter) {
	listOfProducts := data.GetProducts()
	err := listOfProducts.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json data", http.StatusInternalServerError)
		return
	}
}

func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle get products")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal", http.StatusBadRequest)
	}
	p.l.Println("Data %#s", prod)
	data.AddProduct(prod)

	products := data.GetProducts()
	err1 := products.ToJSON(w)
	if err1 != nil {
		http.Error(w, "Error when returning the data", http.StatusBadRequest)
		return
	}

}
