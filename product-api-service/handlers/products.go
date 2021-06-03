package handlers

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"product-api-service/data"
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
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Println("product ", prod)
	err := data.UpdateProduct(id, &prod)
	if err!= nil {
		p.l.Println("Error data", err)
		http.Error(w, "Error", http.StatusMethodNotAllowed)
		return
	}
	if err == data.ErrProductNotFound {
		http.Error(w, "Error", http.StatusBadRequest)
		return
	}

}
func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handle POST")
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to parse data", http.StatusInternalServerError)
	}
}

type KeyProduct struct{}

func (p *Products) MiddleWareProductsValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := data.Product{}
		fmt.Println("data", r.Body)
		err := prod.FromJSON(r.Body)
		if err != nil {
			fmt.Println("Error is", err)
			http.Error(w, "Bad Request in parsing", http.StatusBadRequest)
			return
		}
		///Product validation
		err = prod.Validate()
		if err != nil {
			fmt.Println("Error is", err)
			http.Error(w, "Error validating", http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
