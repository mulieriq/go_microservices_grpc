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
	gh:=handlers.NewGoodbye(log)
	servemux:=http.NewServeMux()
	servemux.Handle("/",hh)
	servemux.Handle("/goodbye",gh)
	http.ListenAndServe(":9090", servemux)
}
