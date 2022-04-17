package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	g.l.Println("Goodbye")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Iza", http.StatusBadGateway)
		return
	}

	bytes, err := fmt.Fprintf(w, "Hey, here is your data %s", data)
	if err != nil {
		http.Error(w, "bad request man", http.StatusBadGateway)
		log.Printf("Bytes sent %d", bytes)

	}
}
