package routes

import (
	"loja/controllers"
	"net/http"
)

//LoadRoutes maps urls to the controller
func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
}
