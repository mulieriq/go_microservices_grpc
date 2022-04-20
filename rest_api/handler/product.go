package handler

import (
	"log"
	"net/http"
	"rest_api/data"
	"time"
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
	w.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {

	time.Sleep(10 * time.Second)
	listOfProducts := data.GetProducts()
	err := listOfProducts.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json data", http.StatusInternalServerError)
	}
}
