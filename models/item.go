// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/gorilla/mux"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Item item
// swagger:model Item

type Item struct {

	// descriptions
	Descriptions []Descriptor `json:"descriptions"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// sub category ID
	SubCategoryID int64 `json:"subCategoryID,omitempty"`

	// user ID
	UserID int64 `json:"userID,omitempty"`
}

type Owner struct {
	FirstName string
	LastName  string
	Username  string
	ID        int
	Email     string
	Items     []Item
}

/* polymorph Item descriptions false */

/* polymorph Item id false */

/* polymorph Item name false */

/* polymorph Item subCategoryID false */

/* polymorph Item userID false */

// Validate validates this item
func (m *Item) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Item) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Item) UnmarshalBinary(b []byte) error {
	var res Item
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

func GetListItems(w http.ResponseWriter, r *http.Request) {
	itemTmpl := template.Must(template.ParseFiles("templates/items/items.html"))
	url := "http://localhost:9000/v1/items"
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		// err
	}
	defer resp.Body.Close()
	var items []Item
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

func ItemsByOwner(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET ItemsByOwner")
	vars := mux.Vars(r)
	id := vars["ID"]

	tmpl := template.Must(template.ParseFiles("./templates/items/itemsByOwner.html"))
	itemUrl := "http://localhost:9000/v1/users/" + id + "/items"
	userUrl := "http://localhost:9000/v1/users/" + id

	var client http.Client
	itemResp, err := client.Get(itemUrl)
	userResp, err := client.Get(userUrl)
	if err != nil {
		// err
	}
	defer itemResp.Body.Close()
	defer userResp.Body.Close()

	var items []Item
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
