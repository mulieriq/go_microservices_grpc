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
	customLog := log.New(os.Stdout, "product-api", log.LstdFlags)

	pd := handlers.NewProducts(customLog)
	serveMux := http.NewServeMux()
	serveMux.Handle("/", pd)
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
			customLog.Fatal(err)
		}
	}()
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)
	sig := <-signalChannel
	customLog.Printf("Received terminate ,graceful  shutdown %s", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(tc)

}
