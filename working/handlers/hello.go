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
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	h.l.Println("Goodbye")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Iza", http.StatusBadGateway)
		return
	}

	bytes, err := fmt.Fprintf(rw, "Hey, here is your data %s", data)
	if err != nil {
		http.Error(rw, "bad request man", http.StatusBadGateway)
		log.Printf("Bytes sent %d", bytes)

	}
}
