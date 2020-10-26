package models

import (
	"fmt"

	"loja/db"
)

//Product structure
type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

//GetAllProducts from db
func GetAllProducts() []Product {
	db := db.ConnectWithPostgres()
	defer db.Close()

	selectAll, err := db.Query("SELECT * FROM products ORDER BY id asc")

	if err != nil {
		fmt.Println("index")
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectAll.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectAll.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}
		p.ID = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}

	return products
}

//CreateNewProduct in db
func CreateNewProduct(name, description string, price float64, quantity int) {
	db := db.ConnectWithPostgres()
	defer db.Close()

	insertProduct, err := db.Prepare("INSERT INTO products(name, description, price, quantity) VALUES($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertProduct.Exec(name, description, price, quantity)
}

//DeleteProduct in db
func DeleteProduct(id string) {
	db := db.ConnectWithPostgres()
	defer db.Close()
	deleteProduct, err := db.Prepare("DELETE FROM products WHERE id=$1")

	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)
}

//EditProduct a product
func EditProduct(id string) Product {
	db := db.ConnectWithPostgres()
	defer db.Close()
	productFromDb, err := db.Query("SELECT * FROM products where id = $1", id)

	if err != nil {
		panic(err.Error())
	}

	productToUpdate := Product{}

	for productFromDb.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = productFromDb.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		productToUpdate.ID = id
		productToUpdate.Name = name
		productToUpdate.Description = description
		productToUpdate.Price = price
		productToUpdate.Quantity = quantity
	}
	return productToUpdate
}

//UpdateProduct a product in db
func UpdateProduct(id int, name, description string, price float64, quantity int) {
	db := db.ConnectWithPostgres()
	defer db.Close()

	UpdateProduct, err := db.Prepare("UPDATE products SET name = $1, description = $2, price = $3, quantity = $4 WHERE id = $5")
	if err != nil {
		panic(err.Error())
	}

	UpdateProduct.Exec(name, description, price, quantity, id)
}
