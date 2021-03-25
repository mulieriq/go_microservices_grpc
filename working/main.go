package main

import (
	"log"
	"net/http"
	"os"
	"working/working/handlers"
)

func main() {
	log := log.New(os.Stdout,"product-api",log.LstdFlags)
	hh := handlers.NewHello(log)
	servemux:=http.NewServeMux()
	servemux.Handle("/",hh)
	http.ListenAndServe(":9090", servemux)
}
