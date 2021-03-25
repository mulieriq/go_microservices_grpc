package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func (g *Goodbye) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	g.l.Println("Root Route") //logg
	fmt.Printf("Data %d",request.Body)
	writer.Write([]byte("Hello Goodbye"))
}

func (g *Goodbye) ServerHTTP(writer http.ResponseWriter,request *http.Request) {

}


func NewGoodbye(l*log.Logger) *Goodbye  {
	return &Goodbye{l}
}

