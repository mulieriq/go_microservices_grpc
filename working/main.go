package main

import (
	"log"
	"net/http"
)

func main()  {
	http.HandleFunc("/",func(http.ResponseWriter , *http.Request){
		log.Println("Hello world")
	})
	
	http.HandleFunc("/goodbye", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Goodbye")
		writer.Write([]byte("Hi"))
		log.Println(request.Body)
	})

	http.ListenAndServe(":9090",nil)
}