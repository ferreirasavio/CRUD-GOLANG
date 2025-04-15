package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"minha-api-go/models"
	"minha-api-go/storage"

	"github.com/gorilla/mux"
)

func TestGetProducts(t *testing.T) {
	storage.Products = []models.Product{
		{ID: 1, Name: "Produto Teste", Price: 9.99, Quantity: 5, Description: "Produto de teste"},
	}

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	rr := httptest.NewRecorder()

	GetProducts(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("esperado status 200, obtido %d", rr.Code)
	}

	var response []models.Product
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Erro ao decodificar resposta JSON: %v", err)
	}

	if len(response) != 1 || response[0].Name != "Produto Teste" {
		t.Errorf("Resposta inesperada: %+v", response)
	}
}

func TestCreateProducts(t *testing.T) {
	storage.Products = []models.Product{}

	newProduct := models.Product{
		Name:        "Produto Teste",
		Price:       9.99,
		Quantity:    5,
		Description: "Produto de teste",
	}

	body, err := json.Marshal(newProduct)
	if err != nil {
		t.Fatalf("Erro ao serializar produto: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	CreateProduct(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("esperado status 201, obtido %d", rr.Code)
	}

	if len(storage.Products) != 1 {
		t.Errorf("esperado 1 produto, obtido %d", len(storage.Products))
	}
	if storage.Products[0].Name != newProduct.Name {
		t.Errorf("esperado nome %s, obtido %s", newProduct.Name, storage.Products[0].Name)
	}
	if storage.Products[0].Price != newProduct.Price {
		t.Errorf("esperado preço %f, obtido %f", newProduct.Price, storage.Products[0].Price)
	}
	if storage.Products[0].Quantity != newProduct.Quantity {
		t.Errorf("esperado quantidade %d, obtido %d", newProduct.Quantity, storage.Products[0].Quantity)
	}
	if storage.Products[0].Description != newProduct.Description {
		t.Errorf("esperado descrição %s, obtido %s", newProduct.Description, storage.Products[0].Description)
	}
	if storage.Products[0].ID != 1 {
		t.Errorf("esperado ID 1, obtido %d", storage.Products[0].ID)
	}
}

func TestUpdateProduct(t *testing.T) {
	storage.Products = []models.Product{
		{ID: 1, Name: "Produto Teste", Price: 9.99, Quantity: 5, Description: "Produto de teste"},
	}

	updatedProduct := models.Product{
		Name:        "Produto Atualizado",
		Price:       19.99,
		Quantity:    10,
		Description: "Produto atualizado",
	}

	body, err := json.Marshal(updatedProduct)
	if err != nil {
		t.Fatalf("Erro ao serializar produto: %v", err)
	}

	req := httptest.NewRequest(http.MethodPut, "/products/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	rr := httptest.NewRecorder()
	UpdateProduct(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("esperado status 200, obtido %d", rr.Code)
	}

	updated := storage.Products[0]

	if updated.Name != updatedProduct.Name {
		t.Errorf("esperado nome %s, obtido %s", updatedProduct.Name, updated.Name)
	}
	if updated.Price != updatedProduct.Price {
		t.Errorf("esperado preço %f, obtido %f", updatedProduct.Price, updated.Price)
	}
	if updated.Quantity != updatedProduct.Quantity {
		t.Errorf("esperado quantidade %d, obtido %d", updatedProduct.Quantity, updated.Quantity)
	}
	if updated.Description != updatedProduct.Description {
		t.Errorf("esperado descrição %s, obtido %s", updatedProduct.Description, updated.Description)
	}
}

func TestDeleteProduct(t *testing.T) {
	storage.Products = []models.Product{
		{ID: 1, Name: "Produto Teste", Price: 9.99, Quantity: 5, Description: "Produto de teste"},
	}

	req := httptest.NewRequest(http.MethodDelete, "/products/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	rr := httptest.NewRecorder()

	DeleteProduct(rr, req)

	if rr.Code != http.StatusNoContent {
		t.Errorf("esperado status 200, obtido %d", rr.Code)
	}

	if len(storage.Products) != 0 {
		t.Errorf("esperado 0 produtos, obtido %d", len(storage.Products))
	}
}

func TestGetProductId(t *testing.T) {
	storage.Products = []models.Product{
		{ID: 1, Name: "Produto Teste", Price: 9.99, Quantity: 5, Description: "Produto de teste"},
	}

	req := httptest.NewRequest(http.MethodGet, "/products/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	rr := httptest.NewRecorder()

	GetProductByID(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("esperado status 200, obtido %d", rr.Code)
	}

	var response models.Product
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Erro ao decodificar resposta JSON: %v", err)
	}

	if response.Name != "Produto Teste" {
		t.Errorf("Resposta inesperada: %+v", response)
	}
}
