package main

import (
	"log"
	"net/http"

	"minha-api-go/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterProductRoutes(r)

	log.Println("Servidor rodando em http://localhost:3030")
	log.Fatal(http.ListenAndServe(":3030", r))
}
