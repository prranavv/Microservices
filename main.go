package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"project/handlers"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "productapi", log.LstdFlags)
	ph := handlers.NewProducts(l)

	sm := mux.NewRouter()
	getrouter := sm.Methods(http.MethodGet).Subrouter()
	getrouter.HandleFunc("/", ph.GetProducts)

	putrouter := sm.Methods(http.MethodPut).Subrouter()
	putrouter.HandleFunc("/{id:[0-9]+}", ph.Updateproducts)
	putrouter.Use(ph.MiddlewareProductValidation)
	//sm.Handle("/products", ph)
	postrouter := sm.Methods(http.MethodPost).Subrouter()
	postrouter.HandleFunc("/", ph.Addproduct)
	postrouter.Use(ph.MiddlewareProductValidation)

	deleterouter := sm.Methods(http.MethodDelete).Subrouter()
	deleterouter.HandleFunc("/", ph.DeleteProducts)
	deleterouter.Use(ph.MiddlewareProductValidation)
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		l.Println("Listening on port 9090")
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)
	signal.Notify(sigchan, os.Kill)
	sig := <-sigchan
	l.Println("Recieved terminate,gracefull shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
