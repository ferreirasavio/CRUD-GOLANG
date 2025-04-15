package routes

import (
	"minha-api-go/handlers"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(r *mux.Router) {
	r.HandleFunc("/", handlers.HandleRoot).Methods("GET")
	r.HandleFunc("/products", handlers.GetProducts).Methods("GET")
	r.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id}", handlers.GetProductByID).Methods("GET")
	r.HandleFunc("/products/{id}", handlers.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", handlers.DeleteProduct).Methods("DELETE")
}
