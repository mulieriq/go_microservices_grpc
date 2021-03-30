package handlers

import (
	"log"
	"net/http"
	"product-api/product-api/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}
	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}
	if r.Method == http.MethodPut {
	  path:=r.URL.Path

	}
	//catch all
	w.WriteHeader(http.StatusMethodNotAllowed)

}
func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handle POST")
	prod := &data.Product{}
	p.l.Printf("data %#v", prod)
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	p.l.Printf("Prod:  %#v", prod)
	data.AddProduct(prod)
}
func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to parse data", http.StatusInternalServerError)
	}

}
