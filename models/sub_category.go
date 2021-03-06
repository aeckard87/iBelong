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

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SubCategory sub category
// swagger:model SubCategory

type SubCategory struct {

	// id
	ID int64 `gorm:"primary_key;auto_incement;not null" json:"id,omitempty"`

	// subcategory
	Subcategory string `gorm:"unique; not null" json:"subcategory,omitempty"`

	//Category Foreign Key
	CategoryID int64 `gorm:"foreignkey:CategoryID;not null" json:"category_id,omitempty"`
}

/* polymorph SubCategory id false */

/* polymorph SubCategory subcategory false */

// Validate validates this sub category
func (m *SubCategory) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SubCategory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubCategory) UnmarshalBinary(b []byte) error {
	var res SubCategory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

func subcategoriesByCategory(id int64) []SubCategory {
	// fmt.Println("ID", id)
	url := fmt.Sprintf("http://127.0.0.1:9000/v1/categories/%v/subcategories", id)
	// fmt.Println(url)

	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	var subcategories []SubCategory
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			fmt.Println(err2)
		}
		// bodyString := string(bodyBytes)
		// fmt.Println(bodyString)
		json.Unmarshal(bodyBytes, &subcategories)
		return subcategories
	} else {
		return subcategories
	}
}

func ListSubCategories(w http.ResponseWriter, r *http.Request) { //, ps httprouter.Params) {
	tmpl := template.Must(template.ParseFiles("templates/subcategories/listsubcategoriesbycategory.html"))
	url := "http://localhost:9000/v1/categories"
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

func GetCreateSubCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET SubCategory/Create")
	categories := GetCategories()
	// tmpl := template.Must(template.ParseFiles("templates/subcategories/createSubCategory.html"))
	tmpl, err := template.New("createSubCategory.html").Funcs(template.FuncMap{
		"subcatByCat": descriptorsByDetail,
	}).ParseFiles("templates/subcategories/createSubCategory.html")
	if err != nil {
		fmt.Println(err)
	}
	tmpl.Execute(w, categories)
}

func PostCreateSubCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST SubCategory/Create")
	var subcategory SubCategory
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	subcategory.Subcategory = r.PostForm.Get("name")
	b, err := json.Marshal(subcategory)
	if err != nil {
		fmt.Println(err)
		return
	}

	request_url := fmt.Sprintf("http://localhost:9000/v1/categories/%s/subcategories", r.PostForm.Get("categoryID"))
	fmt.Println(request_url)
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

	ListSubCategories(w, r)

}

func GetDeleteSubCategory(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/subcategories/deleteSubcategory.html"))
	url := "http://localhost:9000/v1/categories"
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

func PostDeleteSubCategory(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	request_url := "http://localhost:9000/v1/subcategories/" + r.PostForm.Get("subcategoryID")
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

	http.Redirect(w, r, "http://localhost:8100/subcategories", 301)

}

func GetUpdateSubCategory(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/subcategories/updateSubcategory.html"))
	url := "http://localhost:9000/v1/categories"
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

func PostUpdateSubCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST Category/Update")
	var subcategory SubCategory
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	subcategory.Subcategory = r.PostForm.Get("name")
	fmt.Println(subcategory.Subcategory)
	fmt.Println(r.PostForm.Get("subcategoryID"))
	b, err := json.Marshal(subcategory)
	if err != nil {
		fmt.Println(err)
		return
	}

	request_url := "http://localhost:9000/v1/subcategories/" + r.PostForm.Get("subcategoryID")
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

	http.Redirect(w, r, "http://localhost:8100/subcategories", 301)

}
