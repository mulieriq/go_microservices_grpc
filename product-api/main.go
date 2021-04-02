package main

import (
	"context"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
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
	serveMux := mux.NewRouter()


	getRouter := serveMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", pd.GetProducts)

	putRouter := serveMux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", pd.UpdateProduct)
	putRouter.Use(pd.MiddleWareProductsValidation)

	postRouter := serveMux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", pd.AddProduct)
	postRouter.Use(pd.MiddleWareProductsValidation)

	 ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"})) //allows all origins
	//serveMux.Handle("/products", pd)
	server := &http.Server{
		Addr:         ":9090",
		Handler:      ch(serveMux),
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
