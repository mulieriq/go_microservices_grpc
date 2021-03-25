package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
}

func (h *Hello) ServeHTTp(writter http.ResponseWriter, request http.Request) {

	log.Println("Goodbye")
	d, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err)
		http.Error(writter, "OOps", http.StatusBadRequest)
		return
	}
	log.Printf("Data %s", d)
	fmt.Fprintf(writter, "Hello  User : %s", d)
}
