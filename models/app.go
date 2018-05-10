package models

import (
	"fmt"
	"os"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// App type
type App struct {
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
}

// Validate validates this category
func (m *App) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *App) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *App) UnmarshalBinary(b []byte) error {
	var res App
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GetAppPath interface
func (m *App) GetAppPath() {
	schema := os.Getenv("WORN_OUT_SCHEMA")
	origin := os.Getenv("IBELONG_ORIGIN")
	port := os.Getenv("IBELONG_PORT")
	path := os.Getenv("IBELONG_PATH")

	uri := fmt.Sprintf("%s%s:%s%s", schema, origin, port, path)
	fmt.Println("WORN_OUT_URI; " + uri)
	m.Schema = schema
	m.Origin = origin
	m.Port = port
	m.Path = path
	m.URI = uri
}
