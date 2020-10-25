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

	selectAll, err := db.Query("SELECT * FROM products")

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
