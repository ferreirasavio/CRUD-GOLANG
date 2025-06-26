package main

import (
	"log"
	"net/http"
	"os"

	"minha-api-go/database"
	"minha-api-go/routes"

	"github.com/gorilla/mux"
)

func main() {
	database.Connect()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3030"
	}

	r := mux.NewRouter()
	routes.RegisterProductRoutes(r)

	log.Printf("Servidor rodando na porta %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
