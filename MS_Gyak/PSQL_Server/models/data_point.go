// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DataPoint data point
// swagger:model DataPoint
type DataPoint struct {

	// data point description
	DataPointDescription interface{} `json:"DataPointDescription,omitempty"`

	// ID
	ID int64 `json:"ID,omitempty"`

	// module
	Module interface{} `json:"Module,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp strfmt.DateTime `json:"Timestamp,omitempty"`

	// value
	Value float32 `json:"Value,omitempty"`
}

// Validate validates this data point
func (m *DataPoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataPoint) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("Timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataPoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataPoint) UnmarshalBinary(b []byte) error {
	var res DataPoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
