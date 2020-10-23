package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func connectWithPostgres() *sql.DB {
	conn := "user=postgres dbname=alura_store password=example host=localhost port=5432 sslmode=disable"

	db, err := sql.Open("postgres", conn)

	if err != nil {
		panic(err.Error())
	}

	return db
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := connectWithPostgres()
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

	temp.ExecuteTemplate(w, "Index", products)
}
