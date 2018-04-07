package main

import (
	"html/template"
	"net/http"

	models "github.com/aeckard87/iBelong/models"
	"github.com/julienschmidt/httprouter"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

type Items []struct {
	Descriptions []struct {
		Descriptor string `json:"descriptor"`
		DetailID   int    `json:"detailId"`
		ID         int    `json:"id"`
	} `json:"descriptions"`
	ID            int    `json:"id"`
	Name          string `json:"name"`
	SubCategoryID int    `json:"subCategoryID"`
	UserID        int    `json:"userID"`
}

type Users []struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

// type Category struct {
// 	Category string `gorm:"not null; unique" json:"category,omitempty"`
// 	ID       int64  `gorm:"primary_key;AUTO_INCEMENT;not null" json:"id,omitempty"`
// }

type SubCategory struct {
	ID          int64
	Subcategory string
	CategoryID  int64
}

func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/users", models.GetListUsers)
	router.GET("/items", models.GetListItems)
	router.GET("/users/:userID/items", models.ItemsByOwner)
	router.GET("/categories", models.ListCategories)
	router.GET("/categories/create", models.GetCreateCategory)
	router.POST("/categories/create", models.PostCreateCategory)

	http.ListenAndServe(":8081", router)

}

func index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tmpl := template.Must(template.ParseFiles("layout.html"))
	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	// b, _ := json.Marshal(data)
	// fmt.Println(string(b))
	tmpl.Execute(w, data)
}
