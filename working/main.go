package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("Hello world")
	})

	http.HandleFunc("/goodbye", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Goodbye")
		writer.Write([]byte("Hi"))
		d, err := ioutil.ReadAll(request.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Data %s",d)
		writer.Write(d)

	})

	http.ListenAndServe(":9090", nil)
}
