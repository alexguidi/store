package controllers

import (
	"fmt"
	"html/template"
	"log"
	"loja/models"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

//Index with all products saved in db
func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.GetAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}

//New product page
func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

//Insert a new product in db
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceConverted, err := strconv.ParseFloat(price, 64)

		if err != nil {
			log.Println("Error to convert price: ", err)
		}

		quantityConverted, err := strconv.Atoi(quantity)

		if err != nil {
			log.Println("Error to convert quantity: ", err)
		}

		models.CreateNewProduct(name, description, priceConverted, quantityConverted)
	}

	http.Redirect(w, r, "/", 301)
}

//Delete a product
func Delete(w http.ResponseWriter, r *http.Request) {
	prodID := r.URL.Query().Get("id")
	models.DeleteProduct(prodID)
	http.Redirect(w, r, "/", 301)
}

//Edit a product
func Edit(w http.ResponseWriter, r *http.Request) {
	prodID := r.URL.Query().Get("id")

	product := models.EditProduct(prodID)
	temp.ExecuteTemplate(w, "Edit", product)
}

//Update a product
func Update(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Update controller")

	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		idConverted, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Error during conversion of id to int: ", err)
		}

		priceConverted, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error during conversion of price to float64: ", err)
		}

		quantityConverted, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error during conversion of quantity to int: ", err)
		}

		fmt.Println(idConverted, name, description, priceConverted, quantityConverted)

		models.UpdateProduct(idConverted, name, description, priceConverted, quantityConverted)
	}

	http.Redirect(w, r, "/", 301)
}
