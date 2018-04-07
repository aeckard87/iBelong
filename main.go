package main

import (
	"html/template"
	"net/http"

	models "github.com/aeckard87/iBelong/models"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", index)

	//Categories
	router.GET("/categories", models.ListCategories)
	router.GET("/categories/create", models.GetCreateCategory)
	router.POST("/categories/create", models.PostCreateCategory)
	//SubCategories
	//Details
	//Descriptors
	//Users
	router.GET("/users", models.GetListUsers)
	router.GET("/users/:userID/items", models.ItemsByOwner)
	//Items
	router.GET("/items", models.GetListItems)

	http.ListenAndServe(":8081", router)

}

func index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tmpl := template.Must(template.ParseFiles("layout.html"))
	tmpl.Execute(w, nil)
}
