package handlers

import (
	"bytes"
	"encoding/json"
	"minha-api-go/database"
	"minha-api-go/models"
	"minha-api-go/storage"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
)

func TestMain(m *testing.M) {
	os.Setenv("DATABASE_URL", "postgres://savio.ferreira:SENHA@localhost:5432/db_market")

	database.Connect()
	code := m.Run()
	os.Exit(code)
}

func TestGetProducts(t *testing.T) {
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

	if len(response) == 0 {
		t.Errorf("esperado pelo menos 1 produto, obtido %d", len(response))
	}
}

func TestCreateProducts(t *testing.T) {
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
}

func TestUpdateProduct(t *testing.T) {
	p := &models.Product{
		Name:        "Produto Teste",
		Price:       9.99,
		Quantity:    5,
		Description: "Produto de teste",
	}
	err := storage.CreateProduct(p)
	if err != nil {
		t.Fatalf("Erro ao criar produto no setup: %v", err)
	}

	updatedProduct := models.Product{
		ID:          p.ID,
		Name:        "Produto Atualizado",
		Price:       19.99,
		Quantity:    10,
		Description: "Produto atualizado",
	}

	body, err := json.Marshal(updatedProduct)
	if err != nil {
		t.Fatalf("Erro ao serializar produto: %v", err)
	}

	req := httptest.NewRequest(http.MethodPut, "/products/"+strconv.Itoa(p.ID), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(p.ID)})

	rr := httptest.NewRecorder()
	UpdateProduct(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("esperado status 200, obtido %d", rr.Code)
	}

}

func TestDeleteProduct(t *testing.T) {
	p := &models.Product{
		Name:        "Produto Teste",
		Price:       9.99,
		Quantity:    5,
		Description: "Produto de teste",
	}
	err := storage.CreateProduct(p)
	if err != nil {
		t.Fatalf("Erro ao criar produto no setup: %v", err)
	}

	req := httptest.NewRequest(http.MethodDelete, "/products/"+strconv.Itoa(p.ID), nil)
	req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(p.ID)})
	rr := httptest.NewRecorder()

	DeleteProduct(rr, req)

	if rr.Code != http.StatusNoContent {
		t.Errorf("esperado status 204, obtido %d", rr.Code)
	}

	_, err = storage.GetProductByID(p.ID)
	if err == nil {
		t.Errorf("Produto não foi deletado do banco")
	}
}

func TestGetProductId(t *testing.T) {
	p := &models.Product{
		Name:        "Produto Teste",
		Price:       9.99,
		Quantity:    5,
		Description: "Produto de teste",
	}
	err := storage.CreateProduct(p)
	if err != nil {
		t.Fatalf("Erro ao criar produto no setup: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/products/"+strconv.Itoa(p.ID), nil)
	req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(p.ID)})
	rr := httptest.NewRecorder()

	GetProductByID(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("esperado status 200, obtido %d", rr.Code)
	}

	var response models.Product
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Erro ao decodificar resposta JSON: %v", err)
	}

	if response.ID != p.ID || response.Name != p.Name {
		t.Errorf("Resposta inesperada: %+v", response)
	}
}
