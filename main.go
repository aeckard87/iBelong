package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	models "github.com/aeckard87/iBelong/models"
	"github.com/gorilla/mux"
)

var URI string

func main() {
	origin := os.Getenv("WORN_OUT_ORIGIN")
	wornoutPort := os.Getenv("WORN_OUT_PORT")
	path := os.Getenv("WORN_OUT_PATH")

	URI = fmt.Sprintf("%s:%s%s", origin, wornoutPort, path)
	fmt.Println("WORN_OUT_URI; " + URI)

	r := mux.NewRouter()
	// r.HandleFunc("/search/{searchTerm}", Search)
	// r.HandleFunc("/load/{dataId}", Load)
	r.HandleFunc("/", index)

	r.HandleFunc("/settings", settings)

	//Categories
	r.HandleFunc("/categories", models.ListCategories).Methods("GET")
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
	r.HandleFunc("/details", models.ListDetails).Methods("GET")
	//Create
	r.HandleFunc("/details/create", models.PostCreateDetail).Methods("POST")
	r.HandleFunc("/details/create", models.GetCreateDetail).Methods("GET")
	//Update
	r.HandleFunc("/details/update", models.PostUpdateDetail).Methods("POST")
	r.HandleFunc("/details/update", models.GetUpdateDetail).Methods("GET")
	//Delete
	r.HandleFunc("/details/delete", models.PostDeleteDetail).Methods("POST")
	r.HandleFunc("/details/delete", models.GetDeleteDetail).Methods("GET")
	//

	//Descriptors
	r.HandleFunc("/descriptors", models.ListDescriptors).Methods("GET")
	//Create
	r.HandleFunc("/descriptors/create", models.PostCreateDescriptor).Methods("POST")
	r.HandleFunc("/descriptors/create", models.GetCreateDescriptor).Methods("GET")
	//Update
	r.HandleFunc("/descriptors/update", models.PostUpdateDescriptor).Methods("POST")
	r.HandleFunc("/descriptors/update", models.GetUpdateDescriptor).Methods("GET")
	//Delete
	r.HandleFunc("/descriptors/delete", models.PostDeleteDescriptor).Methods("POST")
	r.HandleFunc("/descriptors/delete", models.GetDeleteDescriptor).Methods("GET")
	//

	//Users
	r.HandleFunc("/users", models.GetListUsers)
	r.HandleFunc("/users/{ID:[0-9]+}/items", models.ItemsByOwner).Methods("GET")
	r.HandleFunc("/users/{ID:[0-9]+}", models.GetUser).Methods("GET")
	//Create
	r.HandleFunc("/users/create", models.PostCreateUser).Methods("POST")
	r.HandleFunc("/users/create", models.GetCreateUser).Methods("GET")
	//Update
	r.HandleFunc("/users/update", models.PostUpdateUser).Methods("POST")
	r.HandleFunc("/users/update", models.GetUpdateUser).Methods("GET")
	//Delete
	r.HandleFunc("/users/delete", models.PostDeleteUser).Methods("POST")
	r.HandleFunc("/users/delete", models.GetDeleteUser).Methods("GET")

	//Items
	r.HandleFunc("/items", models.GetListItems).Methods("GET")
	r.HandleFunc("/items/{ID:[0-9]+}", models.GetListItem).Methods("GET")

	//Static file location
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	// http.Handle("/tmp/", http.StripPrefix("/tmp/", http.FileServer(http.Dir("./static/"))))

	http.Handle("/", r)
	// http.ListenAndServe(":8100", nil)
	port := os.Getenv("IBELONG_PORT")
	fmt.Println("IBELONG PORT: " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

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
func settings(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/settings/index.html"))
	tmpl.Execute(w, nil)
}
