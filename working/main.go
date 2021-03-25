package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("Hello world")
	})

	http.HandleFunc("/goodbye", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Goodbye")
	})

	http.ListenAndServe(":9090", nil)
}
