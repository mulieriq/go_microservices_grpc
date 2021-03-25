package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"
	"working/working/handlers"
)

func main() {
	log := log.New(os.Stdout,"product-api",log.LstdFlags)
	hh := handlers.NewHello(log)
	gh:=handlers.NewGoodbye(log)
	servemux:=http.NewServeMux()
	servemux.Handle("/",hh)
	servemux.Handle("/goodbye",gh)
	server:=&http.Server{
		Addr: ":9090",
		Handler: servemux,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}
	go func() {
		err:=server.ListenAndServe()
		if err!=nil{
			log.Fatal(err)
		}
	}()

	tc,_ :=context.WithTimeout(context.Background(),30*time.Second)
	server.Shutdown(tc)
	//http.ListenAndServe(":9090", servemux)
}
