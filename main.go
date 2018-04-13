package main

import (
	"html/template"
	"net/http"

	models "github.com/aeckard87/iBelong/models"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	// r.HandleFunc("/search/{searchTerm}", Search)
	// r.HandleFunc("/load/{dataId}", Load)
	r.HandleFunc("/", index)

	//Categories
	r.HandleFunc("/categories", models.ListCategories).Methods("GET")
	// r.HandleFunc("/categories/{ID}", models.ListCategory).Methods("GET")
	//Create
	r.HandleFunc("/categories/create", models.PostCreateCategory).Methods("POST")
	r.HandleFunc("/categories/create", models.GetCreateCategory).Methods("GET")
	//Update
	r.HandleFunc("/categories/update", models.PostUpdateCategory).Methods("POST")
	r.HandleFunc("/categories/update", models.GetUpdateCategory).Methods("GET")
	//Delete
	r.HandleFunc("/categories/delete", models.PostDeleteCategory).Methods("POST")
	r.HandleFunc("/categories/delete", models.GetDeleteCategory).Methods("GET")

	//SubCategories
	r.HandleFunc("/subcategories", models.ListSubCategories).Methods("GET")
	//Create
	r.HandleFunc("/subcategories/create", models.PostCreateSubCategory).Methods("POST")
	r.HandleFunc("/subcategories/create", models.GetCreateSubCategory).Methods("GET")
	//Update
	r.HandleFunc("/subcategories/update", models.PostUpdateSubCategory).Methods("POST")
	r.HandleFunc("/subcategories/update", models.GetUpdateSubCategory).Methods("GET")
	//Delete
	r.HandleFunc("/subcategories/delete", models.PostDeleteSubCategory).Methods("POST")
	r.HandleFunc("/subcategories/delete", models.GetDeleteSubCategory).Methods("GET")
	//
	//Details
	//
	//Descriptors
	//
	//Users
	r.HandleFunc("/users", models.GetListUsers)
	r.HandleFunc("/users/{ID}/items", models.ItemsByOwner).Methods("GET")

	//Items
	r.HandleFunc("/items", models.GetListItems).Methods("GET")

	//Static file location
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	// http.Handle("/tmp/", http.StripPrefix("/tmp/", http.FileServer(http.Dir("./static/"))))

	http.Handle("/", r)
	http.ListenAndServe(":8100", nil)

	//----------------//

	// router := httprouter.New()
	// http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./static"))))

	// router.GET("/", index)

	// //Categories
	// router.GET("/categories", models.ListCategories)
	// //Create
	// router.GET("/categories/create", models.GetCreateCategory)
	// router.POST("/categories/create", models.PostCreateCategory)
	// //delete
	// router.GET("/categories/delete", models.GetDeleteCategory)
	// router.POST("/categories/delete", models.PostDeleteCategory)

	// //SubCategories
	// //Details
	// //Descriptors
	// //Users
	// router.GET("/users", models.GetListUsers)
	// router.GET("/users/:userID/items", models.ItemsByOwner)
	// //Items
	// router.GET("/items", models.GetListItems)

	// http.ListenAndServe(":8081", router)

}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}
