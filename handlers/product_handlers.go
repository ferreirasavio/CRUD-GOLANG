package handlers

import (
	"encoding/json"
	"log"
	"minha-api-go/models"
	"minha-api-go/storage"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Bem-vindo à API de produtos!"))
}

func GetProducts(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products, err := storage.GetProducts()
	if err != nil {
		http.Error(w, "Erro ao buscar produtos", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	product, err := storage.GetProductByID(id)
	if err != nil {
		http.Error(w, "Produto não encontrado", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct models.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		log.Printf("Erro ao decodificar: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if newProduct.Name == "" {
		http.Error(w, "Nome do produto não foi informado.", http.StatusBadRequest)
		return
	}

	if newProduct.Price <= 0 {
		http.Error(w, "Preço do produto não foi informado.", http.StatusBadRequest)
		return
	}

	if newProduct.Quantity <= 0 {
		http.Error(w, "Quantidade do produto não foi informada.", http.StatusBadRequest)
		return
	}

	if err := storage.CreateProduct(&newProduct); err != nil {
		http.Error(w, "Erro ao criar produto", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newProduct)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	existingProduct, err := storage.GetProductByID(id)
	if err != nil {
		http.Error(w, "Produto não encontrado", http.StatusNotFound)
		return
	}

	var updatedProduct models.Product
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	if updatedProduct.Name != "" {
		existingProduct.Name = updatedProduct.Name
	}
	if updatedProduct.Description != "" {
		existingProduct.Description = updatedProduct.Description
	}
	if updatedProduct.Price > 0 {
		existingProduct.Price = updatedProduct.Price
	}
	if updatedProduct.Quantity > 0 {
		existingProduct.Quantity = updatedProduct.Quantity
	}

	err = storage.UpdateProduct(existingProduct)
	if err != nil {
		http.Error(w, "Erro ao atualizar produto", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existingProduct)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = storage.DeleteProduct(id)
	if err != nil {
		if err.Error() == "produto não encontrado" {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Erro ao deletar produto", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
