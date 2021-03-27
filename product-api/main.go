package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"product-api/product-api/handlers"
)

func main() {
	log := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(log)
	gh := handlers.NewGoodbye(log)
	 pd:= handlers.NewProducts(log)
	servemux := http.NewServeMux()
	servemux.Handle("/", hh)
	servemux.Handle("/goodbye", gh)
	servemux.Handle("/products",pd)
	server := &http.Server{
		Addr:         ":9090",
		Handler:      servemux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)
	sig := <-signalChannel
	log.Printf("Received terminate ,graceful  shutdown %s", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(tc)
	//http.ListenAndServe(":9090", servemux)
}
