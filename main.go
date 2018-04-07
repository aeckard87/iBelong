package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
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

type Owner struct {
	FirstName string
	LastName  string
	Username  string
	ID        int
	Email     string
	Items     Items
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
	router.GET("/users", listUsers)
	router.GET("/items", listItems)
	router.GET("/users/:userID/items", itemsByOwner)
	router.GET("/categories", models.ListCategories)
	router.GET("/categories/create", createCategory)
	router.POST("/categories/create", models.CreateCategory)

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

func createCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var category models.Category
	tmpl := template.Must(template.ParseFiles("templates/categories/createCategory.html"))
	tmpl.Execute(w, category)
}

func listItems(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	itemTmpl := template.Must(template.ParseFiles("items.html"))
	url := "http://localhost:9000/v1/items"
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		// err
	}
	defer resp.Body.Close()
	var items Items
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			fmt.Println(err2)
		}
		// bodyString := string(bodyBytes)
		// fmt.Println(bodyString)
		json.Unmarshal(bodyBytes, &items)
		itemTmpl.Execute(w, items)
	} else {
		itemTmpl.Execute(w, items)
	}
}

func listUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	usersTmpl := template.Must(template.ParseFiles("users.html"))
	url := "http://localhost:9000/v1/users"
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		// err
	}
	defer resp.Body.Close()
	var users Users
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			fmt.Println(err2)
		}
		// bodyString := string(bodyBytes)
		// fmt.Println(bodyString)
		json.Unmarshal(bodyBytes, &users)
		usersTmpl.Execute(w, users)
	} else {
		usersTmpl.Execute(w, users)
	}
}
func itemsByOwner(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	tmpl := template.Must(template.ParseFiles("itemsByOwner.html"))
	itemUrl := "http://localhost:9000/v1/users/" + ps.ByName("userID") + "/items"
	userUrl := "http://localhost:9000/v1/users/" + ps.ByName("userID")
	var client http.Client
	itemResp, err := client.Get(itemUrl)
	userResp, err := client.Get(userUrl)
	if err != nil {
		// err
	}
	defer itemResp.Body.Close()
	defer userResp.Body.Close()
	var items Items
	var owner Owner

	if itemResp.StatusCode == http.StatusOK && userResp.StatusCode == http.StatusOK {
		itemBytes, err2 := ioutil.ReadAll(itemResp.Body)
		userBytes, err2 := ioutil.ReadAll(userResp.Body)
		if err2 != nil {
			fmt.Println(err2)
		}

		// bodyString := string(bodyBytes)
		// fmt.Println(bodyString)
		json.Unmarshal(itemBytes, &items)
		json.Unmarshal(userBytes, &owner)
		owner.Items = items

		tmpl.Execute(w, owner)
	} else {
		tmpl.Execute(w, owner)
	}
}
