package storage

import (
	"context"
	"errors"
	"minha-api-go/database"
	"minha-api-go/models"
)

func GetProducts() ([]models.Product, error) {
	rows, err := database.DB.Query(context.Background(), "SELECT id, name, description, price, quantity FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Quantity); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func GetProductByID(id int) (*models.Product, error) {
	var p models.Product
	query := `SELECT id, name, description, price, quantity FROM products WHERE id=$1`
	err := database.DB.QueryRow(context.Background(), query, id).Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Quantity)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func CreateProduct(p *models.Product) error {
	query := `INSERT INTO products (name, description, price, quantity) VALUES ($1, $2, $3, $4) RETURNING id`
	return database.DB.QueryRow(context.Background(), query, p.Name, p.Description, p.Price, p.Quantity).Scan(&p.ID)
}

func UpdateProduct(p *models.Product) error {
	query := `UPDATE products SET name=$1, description=$2, price=$3, quantity=$4 WHERE id=$5`
	cmdTag, err := database.DB.Exec(context.Background(), query, p.Name, p.Description, p.Price, p.Quantity, p.ID)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return errors.New("produto não encontrado")
	}
	return nil
}

func DeleteProduct(id int) error {
	cmdTag, err := database.DB.Exec(context.Background(), "DELETE FROM products WHERE id=$1", id)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return errors.New("produto não encontrado")
	}
	return nil
}
