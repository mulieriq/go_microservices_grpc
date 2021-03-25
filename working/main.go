package main

import (
	"fmt"
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
		d, err := ioutil.ReadAll(request.Body)
		if err != nil {
			log.Fatal(err)
			http.Error(writer,"OOps",http.StatusBadRequest)
			return
		}
		log.Printf("Data %s",d)
		fmt.Fprintf(writer,"Hello  User : %s",d)
	})

	http.ListenAndServe(":9090", nil)
}
