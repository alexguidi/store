package controllers

import (
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
