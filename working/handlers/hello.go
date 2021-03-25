package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}
func (h *Hello) ServeHTTP(writter http.ResponseWriter, request *http.Request) {

	h.l.Println("Root Route") //logger
	d, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err)
		http.Error(writter, "OOps", http.StatusBadRequest)
		return
	}
	log.Printf("Data %s", d)
	fmt.Fprintf(writter, "Hello  User : %s", d)
}
