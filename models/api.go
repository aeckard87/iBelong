package models

import (
	"fmt"
	"os"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// API type
type API struct {
	// schema
	Schema string

	// Origin
	Origin string

	// PORT
	Port string

	//Path
	Path string

	//URI
	URI string

	//Obj
	// Obj string `json:"obj"`
}

/* polymorph Category category false */

/* polymorph Category id false */

// Validate validates this category
func (m *API) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *API) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *API) UnmarshalBinary(b []byte) error {
	var res API
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GetAPIPath interface
func (m *API) GetAPIPath() {
	schema := os.Getenv("WORN_OUT_SCHEMA")
	origin := os.Getenv("WORN_OUT_ORIGIN")
	port := os.Getenv("WORN_OUT_PORT")
	path := os.Getenv("WORN_OUT_PATH")

	uri := fmt.Sprintf("%s%s:%s%s", schema, origin, port, path)
	fmt.Println("WORN_OUT_URI; " + uri)
	m.Schema = schema
	m.Origin = origin
	m.Port = port
	m.Path = path
	m.URI = uri
}
