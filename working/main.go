package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, request *http.Request) {
		log.Println("Hello World")
		data, err := ioutil.ReadAll(request.Body)
		if err != nil {
			http.Error(rw, "Oop", http.StatusBadRequest)
			return
		}
		bytes, err := fmt.Fprintf(rw, "Hello %s", data)
		log.Printf("Bytes Received %d", bytes)
		if err != nil {
			http.Error(rw, "bad request man", http.StatusBadGateway)
			log.Printf("Bytes sent %d", bytes)

		}
	})

	http.HandleFunc("/goodbye", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Goodbye")
		data, err := ioutil.ReadAll(request.Body)
		if err != nil {
			http.Error(writer, "Iza", http.StatusBadGateway)
			return
		}

		bytes, err := fmt.Fprintf(writer, "Hey, here is your data %s", data)
		if err != nil {
			http.Error(writer, "bad request man", http.StatusBadGateway)
			log.Printf("Bytes sent %d", bytes)

		}
	})

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		print(err)
	}
}
