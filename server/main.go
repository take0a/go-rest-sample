package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/take0a/go-rest-sample/module/customers"
	"github.com/take0a/go-rest-sample/module/orders"
	"github.com/take0a/go-rest-sample/module/products"
)

func main() {
	dsn := os.Getenv("DSN")
	db, err := NewDB(dsn)
	if err != nil {
		log.Printf("NewDB error %s", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Mount("/customers", customers.NewRouter(db))
	r.Mount("/orders", orders.NewRouter(db))
	r.Mount("/products", products.NewRouter(db))

	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	err = http.ListenAndServe(addr, r)
	if err != nil {
		log.Printf("ListenAndServe error %s", err)
	}
}
