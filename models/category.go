// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/aeckard87/WornOut/models"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/gorilla/mux"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Category category
// swagger:model Category
// Utilize GetCategories function and also update redirects!

type Category struct {

	// category
	Category string `gorm:"not null; unique" json:"category,omitempty"`

	// id
	ID int64 `gorm:"primary_key;AUTO_INCEMENT;not null" json:"id,omitempty"`
}

/* polymorph Category category false */

/* polymorph Category id false */

// Validate validates this category
func (m *Category) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Category) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Category) UnmarshalBinary(b []byte) error {
	var res Category
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

func GetCategories() []Category {
	url := "http://10.0.0.13:8081/aeckard87/wornOut/v1/categories/"
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		// err
	}
	defer resp.Body.Close()
	var categories []Category
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			fmt.Println(err2)
		}
		// bodyString := string(bodyBytes)
		// fmt.Println(bodyString)
		json.Unmarshal(bodyBytes, &categories)
		return categories
	} else {
		return categories
	}
}

func GetCreateCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET Category/Create")
	var category models.Category
	tmpl := template.Must(template.ParseFiles("templates/categories/createCategory.html"))
	tmpl.Execute(w, category)
}

func PostCreateCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST Category/Create")
	var category Category
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	category.Category = r.PostForm.Get("name")
	b, err := json.Marshal(category)
	if err != nil {
		fmt.Println(err)
		return
	}

	request_url := "http://10.0.0.13:8081/aeckard87/wornOut/v1/categories"
	req, err := http.NewRequest("POST", request_url, bytes.NewBuffer(b))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	// ListCategories(w, r, ps)
	ListCategories(w, r)

}

func GetUpdateCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET Category/Update")
	tmpl := template.Must(template.ParseFiles("templates/categories/updateCategory.html"))
	tmpl.Execute(w, GetCategories())
}

func PostUpdateCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST Category/Create")
	var category Category
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	category.Category = r.PostForm.Get("name")
	// category.ID, err := strconv.Atoi(r.PostForm.Get("categoryID"))
	b, err := json.Marshal(category)
	if err != nil {
		fmt.Println(err)
		return
	}

	request_url := "http://10.0.0.13:8081/aeckard87/wornOut/v1/categories/" + r.PostForm.Get("categoryID")
	req, err := http.NewRequest("PUT", request_url, bytes.NewBuffer(b))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	http.Redirect(w, r, "http://10.0.0.13:8100/categories", 301)

}

//Not working as expected
func ListCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]

	tmpl := template.Must(template.ParseFiles("templates/categories/listCategory.html"))
	url := fmt.Sprintf("http://10.0.0.13:8081/aeckard87/wornOut/v1/categories/%v", id)
	fmt.Println("URL", url)

	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		// err
	}
	defer resp.Body.Close()

	var categories Category
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			fmt.Println(err2)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
		json.Unmarshal(bodyBytes, &categories)
		fmt.Println("Unmarshaled", categories)
		tmpl.Execute(w, categories)
	} else {
		tmpl.Execute(w, categories)
	}
}

func ListCategories(w http.ResponseWriter, r *http.Request) { //, ps httprouter.Params) {
	tmpl := template.Must(template.ParseFiles("templates/categories/listCategories.html"))
	url := "http://10.0.0.13:8081/aeckard87/wornOut/v1/categories"
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		// err
	}
	defer resp.Body.Close()
	var categories []Category
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			fmt.Println(err2)
		}
		// bodyString := string(bodyBytes)
		// fmt.Println(bodyString)
		json.Unmarshal(bodyBytes, &categories)
		tmpl.Execute(w, categories)
	} else {
		tmpl.Execute(w, categories)
	}
}

func GetDeleteCategory(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/categories/deleteCategory.html"))
	url := "http://10.0.0.13:8081/aeckard87/wornOut/v1/categories"
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		// err
	}
	defer resp.Body.Close()
	var categories []Category
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			fmt.Println(err2)
		}
		// bodyString := string(bodyBytes)
		// fmt.Println(bodyString)
		json.Unmarshal(bodyBytes, &categories)
		tmpl.Execute(w, categories)
	} else {
		tmpl.Execute(w, categories)
	}
}

func PostDeleteCategory(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	for key, value := range r.PostForm {
		fmt.Printf("Key: %s\tValue: %s", key, value)
	}

	request_url := "http://10.0.0.13:8081/aeckard87/wornOut/v1/categories/" + r.PostForm.Get("categoryID")
	fmt.Println(request_url)
	req, err := http.NewRequest("DELETE", request_url, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	ListCategories(w, r)

}
