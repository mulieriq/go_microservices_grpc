package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"product-api/product-api/data"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //extracting id from mux
	id, _ := strconv.Atoi(vars["id"])
	p.l.Println("Handle put", id)

	p.l.Println("product ", prod)
	errorp := data.UpdateProduct(id, prod)
	if errorp != nil {
		p.l.Println("eroor data", errorp)
		http.Error(w, "Erro", http.StatusMethodNotAllowed)
		return
	}
	if errorp == data.ErrProductNotFound {
		http.Error(w, "Erro", http.StatusBadRequest)
		return
	}

}
func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
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
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to parse data", http.StatusInternalServerError)
	}
}

type KeyProduct struct

	func(p *Products) MiddleWareProductsValidation(w http.ResponseWriter, r *http.Request) http.Handler {
	return http.HandlerFunc(w
	http.ResponseWriter, r * http.Request){

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
	http.Error(w, "Bad Request in parsing", http.StatusBadRequest)
	return
	}
	ctx := r.Context().Value(KeyProduct{}, prop)
	req := r.WithContext(ctx)
	next.ServeHTTP(w, r)

	}
	}
