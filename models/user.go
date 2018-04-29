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

// User user
// swagger:model User

type User struct {

	// id
	ID int64 `json:"id,omitempty"`

	// username
	Username string `json:"username,omitempty"`

	// first name
	FirstName string `json:"firstName,omitempty"`

	// last name
	LastName string `json:"lastName,omitempty"`

	// email
	Email string `json:"email,omitempty"`
}

/* polymorph User email false */

/* polymorph User firstName false */

/* polymorph User id false */

/* polymorph User lastName false */

/* polymorph User username false */

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

func GetListUsers(w http.ResponseWriter, r *http.Request) {
	usersTmpl := template.Must(template.ParseFiles("templates/users/users.html"))
	url := "http://10.0.0.13:8081/aeckard87/wornOut/v1/users"
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		// err
	}
	defer resp.Body.Close()
	var users []User
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

func GetCreateUser(w http.ResponseWriter, r *http.Request) {
	usersTmpl := template.Must(template.ParseFiles("templates/users/createUser.html"))
	url := "http://10.0.0.13:8081/aeckard87/wornOut/v1/users"
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		// err
	}
	defer resp.Body.Close()
	var users []User
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
func PostCreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST User/Create")
	var user User
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	user.FirstName = r.PostForm.Get("firstname")
	user.LastName = r.PostForm.Get("lastname")
	user.Email = r.PostForm.Get("email")
	user.Username = r.PostForm.Get("username")
	fmt.Println(user)

	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	request_url := "http://10.0.0.13:8081/aeckard87/wornOut/v1/users"
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
	GetListUsers(w, r)

}

func GetUpdateUser(w http.ResponseWriter, r *http.Request) {
	usersTmpl := template.Must(template.ParseFiles("templates/users/updateUser.html"))
	url := "http://10.0.0.13:8081/aeckard87/wornOut/v1/users"
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		// err
	}
	defer resp.Body.Close()
	var users []User
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
func PostUpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST User/Update")
	var user User
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	user.FirstName = r.PostForm.Get("firstname")
	user.LastName = r.PostForm.Get("lastname")
	user.Email = r.PostForm.Get("email")
	user.Username = r.PostForm.Get("username")
	// user.ID, err := strconv.Atoi(r.PostForm.Get("userID"))
	fmt.Println(user)
	fmt.Println(r.PostForm.Get("userID"))

	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	request_url := "http://10.0.0.13:8081/aeckard87/wornOut/v1/users/" + r.PostForm.Get("userID")
	req, err := http.NewRequest("PUT", request_url, bytes.NewBuffer(b))
	// req.Header.Set("X-Custom-Header", "myvalue")
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

	http.Redirect(w, r, "http://10.0.0.13:8100/users", 301)

}

func GetDeleteUser(w http.ResponseWriter, r *http.Request) {
	usersTmpl := template.Must(template.ParseFiles("templates/users/deleteUser.html"))
	url := "http://10.0.0.13:8081/aeckard87/wornOut/v1/users"
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		// err
	}
	defer resp.Body.Close()
	var users []User
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
func PostDeleteUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	// for key, value := range r.PostForm {
	// 	fmt.Printf("Key: %s\tValue: %s", key, value)
	// }

	request_url := "http://10.0.0.13:8081/aeckard87/wornOut/v1/users/" + r.PostForm.Get("userID")
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

	GetListUsers(w, r)

}
