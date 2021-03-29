package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"product-api/product-api/handlers"
	"time"
)

func main() {
	customeLog := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(customeLog)
	gh := handlers.NewGoodbye(customeLog)
	pd := handlers.NewProducts(customeLog)
	serveMux := http.NewServeMux()
	serveMux.Handle("/", hh) //testing api
	serveMux.Handle("/goodbye", gh) //Testing api
	serveMux.Handle("/products", pd)  ///Production code
	server := &http.Server{
		Addr:         ":9090",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			customeLog.Fatal(err)
		}
	}()
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)
	sig := <-signalChannel
	customeLog.Printf("Received terminate ,graceful  shutdown %s", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(tc)
	//http.ListenAndServe(":9090", serveMux)
}
