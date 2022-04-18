package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rest_api/handler"
	"time"
)

func main() {

	logger := log.New(os.Stdin, "product-api", log.LstdFlags)
	productHandler := handler.NewProducts(logger)

	serveMux := http.NewServeMux()

	serveMux.Handle("/products", productHandler)

	server := http.Server{
		Addr:         ":9090",
		Handler:      serveMux,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Printf("Error starting server %s ", err)
		}
	}()

	signalChanel := make(chan os.Signal)
	signal.Notify(signalChanel, os.Interrupt)
	signal.Notify(signalChanel, os.Kill)

	channelData := <-signalChanel

	logger.Printf("System shutdown gracefully %s", channelData)

	timeOutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := server.Shutdown(timeOutContext)
	if err != nil {
		logger.Println(err)
	}
}
