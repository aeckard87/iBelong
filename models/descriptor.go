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

// Descriptor descriptor
// swagger:model Descriptor

type Descriptor struct {

	// descriptor
	Descriptor string `json:"descriptor,omitempty"`

	// detail Id
	DetailID int64 `gorm:"foreignkey:DetailID;not null" json:"detailId,omitempty"`

	// id
	ID int64 `gorm:"primary_key;auto_incement" json:"id,omitempty"`
}

/* polymorph Descriptor descriptor false */

/* polymorph Descriptor detailId false */

/* polymorph Descriptor id false */

// Validate validates this descriptor
func (m *Descriptor) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Descriptor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Descriptor) UnmarshalBinary(b []byte) error {
	var res Descriptor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

func descriptorsByDetail(id int64) []Descriptor {
	// fmt.Println("ID", id)
	url := fmt.Sprintf("http://127.0.0.1:9000/v1/details/%v/descriptors", id)
	// fmt.Println(url)

	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	var descriptors []Descriptor
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			fmt.Println(err2)
		}
		// bodyString := string(bodyBytes)
		// fmt.Println(bodyString)
		json.Unmarshal(bodyBytes, &descriptors)
		return descriptors
	} else {
		return descriptors
	}
}

func ListDescriptors(w http.ResponseWriter, r *http.Request) { //, ps httprouter.Params) {
	tmpl := template.Must(template.ParseFiles("templates/descriptors/listdescriptorsbydetail.html"))
	url := "http://10.0.0.13:8081/aeckard87/wornOut/v1/details"
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		// err
	}
	defer resp.Body.Close()
	var details []Detail
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			fmt.Println(err2)
		}
		// bodyString := string(bodyBytes)
		// fmt.Println(bodyString)
		json.Unmarshal(bodyBytes, &details)
		tmpl.Execute(w, details)
	} else {
		tmpl.Execute(w, details)
	}
}

func GetCreateDescriptor(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET Descriptor/Create")
	details := GetDetails()
	// tmpl := template.Must(template.ParseFiles("templates/descriptors/createDescriptor.html"))
	tmpl, err := template.New("createDescriptor.html").Funcs(template.FuncMap{
		"subcatByCat": descriptorsByDetail,
	}).ParseFiles("templates/descriptors/createDescriptor.html")
	if err != nil {
		fmt.Println(err)
	}
	tmpl.Execute(w, details)
}

func PostCreateDescriptor(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST Descriptor/Create")
	var descriptor Descriptor
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	descriptor.Descriptor = r.PostForm.Get("name")
	b, err := json.Marshal(descriptor)
	if err != nil {
		fmt.Println(err)
		return
	}

	request_url := fmt.Sprintf("http://10.0.0.13:8081/aeckard87/wornOut/v1/details/%s/descriptors", r.PostForm.Get("detailID"))
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

	GetCreateDescriptor(w, r)

}

func GetDeleteDescriptor(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/descriptors/deleteDescriptor.html"))
	url := "http://10.0.0.13:8081/aeckard87/wornOut/v1/details"
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		// err
	}
	defer resp.Body.Close()
	var details []Detail
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			fmt.Println(err2)
		}
		// bodyString := string(bodyBytes)
		// fmt.Println(bodyString)
		json.Unmarshal(bodyBytes, &details)
		tmpl.Execute(w, details)
	} else {
		tmpl.Execute(w, details)
	}
}

func PostDeleteDescriptor(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	request_url := "http://10.0.0.13:8081/aeckard87/wornOut/v1/descriptors/" + r.PostForm.Get("descriptorID")
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

	http.Redirect(w, r, "http://10.0.0.13:8100/descriptors", 301)

}

func GetUpdateDescriptor(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/descriptors/updateDescriptor.html"))
	url := "http://10.0.0.13:8081/aeckard87/wornOut/v1/details"
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		// err
	}
	defer resp.Body.Close()
	var details []Detail
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			fmt.Println(err2)
		}
		// bodyString := string(bodyBytes)
		// fmt.Println(bodyString)
		json.Unmarshal(bodyBytes, &details)
		tmpl.Execute(w, details)
	} else {
		tmpl.Execute(w, details)
	}
}

func PostUpdateDescriptor(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST Detail/Update")
	var descriptor Descriptor
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	descriptor.Descriptor = r.PostForm.Get("name")
	fmt.Println(descriptor.Descriptor)
	fmt.Println(r.PostForm.Get("descriptorID"))
	b, err := json.Marshal(descriptor)
	if err != nil {
		fmt.Println(err)
		return
	}

	request_url := "http://10.0.0.13:8081/aeckard87/wornOut/v1/descriptors/" + r.PostForm.Get("descriptorID")
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

	http.Redirect(w, r, "http://10.0.0.13:8100/descriptors", 301)

}
