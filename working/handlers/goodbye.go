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

func (g *Goodbye) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	g.l.Println("Root Route") //logg
	d,_:=ioutil.ReadAll(request.Body)
	fmt.Fprintf(writer,"Data %s",d)
	writer.Write([]byte("Hello Goodbye"))
}

func (g *Goodbye) ServerHTTP(writer http.ResponseWriter,request *http.Request) {

}


func NewGoodbye(l*log.Logger) *Goodbye  {
	return &Goodbye{l}
}

