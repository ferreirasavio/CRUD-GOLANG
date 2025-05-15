package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"minha-api-go/models"
	"minha-api-go/storage"

	"github.com/gorilla/mux"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Bem-vindo à API de produtos!"))
}

func GetProducts(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storage.Products)
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for _, product := range storage.Products {
		if product.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(product)
			return
		}
	}
	http.Error(w, "Produto não encontrado", http.StatusNotFound)
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

	newProduct.ID = len(storage.Products) + 1
	storage.Products = append(storage.Products, newProduct)
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

	var updatedProduct models.Product
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		log.Printf("Erro ao decodificar: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	found := false
	updated := false
	var existingProduct models.Product

	for i, product := range storage.Products {
		if product.ID == id {
			found = true
			existingProduct = product

			if updatedProduct.Name != "" && updatedProduct.Name != existingProduct.Name {
				storage.Products[i].Name = updatedProduct.Name
				updated = true
			}
			if updatedProduct.Description != "" && updatedProduct.Description != existingProduct.Description {
				storage.Products[i].Description = updatedProduct.Description
				updated = true
			}
			if updatedProduct.Price > 0 && updatedProduct.Price != existingProduct.Price {
				storage.Products[i].Price = updatedProduct.Price
				updated = true
			}
			if updatedProduct.Quantity > 0 && updatedProduct.Quantity != existingProduct.Quantity {
				storage.Products[i].Quantity = updatedProduct.Quantity
				updated = true
			}

			if updated {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(storage.Products[i])
				return
			}
			break
		}
	}

	if found && !updated {
		http.Error(w, "Nenhum dado foi alterado.", http.StatusOK)
		return
	}

	if !found {
		http.Error(w, "Produto não encontrado", http.StatusNotFound)
		return
	}
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for i, product := range storage.Products {
		if product.ID == id {
			storage.Products = append(storage.Products[:i], storage.Products[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Produto não encontrado", http.StatusNotFound)
}
