package main

import (
	"log"
	"net/http"
)

func main()  {
	http.HandleFunc("/",func(http.ResponseWriter , *http.Request){
		log.Println("Hello world")
	})

	http.ListenAndServe(":9090",nil)
}