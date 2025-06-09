package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func createProduct(db *sql.DB, p Product) error {
	query := "INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id,name,price;"
	_, err := db.Exec(query, p.Name, p.Price)
	if err != nil {
		return err
	}
	return nil

}

func getProduct(db *sql.DB, id int) (Product, error) {
	query := "SELECT id, name, price FROM products WHERE id = $1;"
	row := db.QueryRow(query, id)
	var p Product
	err := row.Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}

func updateProduct(db *sql.DB, p Product) (Product, error) {
	query := "UPDATE products SET name = $1, price = $2 WHERE id = $3 RETURNING id, name, price;"
	row := db.QueryRow(query, p.Name, p.Price, p.ID)
	var updatedProduct Product
	err := row.Scan(&updatedProduct.ID, &updatedProduct.Name, &updatedProduct.Price)
	if err != nil {
		return Product{}, err
	}
	fmt.Printf("Product updated with ID: %d\n", updatedProduct.ID)
	return updatedProduct, nil
}

func deleteProduct(db *sql.DB, id int) error {
	query := "DELETE FROM products WHERE id = $1;"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	fmt.Printf("Product with ID %d deleted successfully.\n", id)
	return nil
}

func getProducts(db *sql.DB) ([]Product, error) {
	query := "SELECT id, name, price FROM products"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
