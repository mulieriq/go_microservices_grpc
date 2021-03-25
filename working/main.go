package main

import (
	"log"
	"net/http"
	"os"
	"working/working/handlers"
)

func main() {
	log := log.New(os.Stdout,"product-api",log.LstdFlags)
	helloHandlerLogger := handlers.NewHello(log)
	http.ListenAndServe(":9090", nil)
}
