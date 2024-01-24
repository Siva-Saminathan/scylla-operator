// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TableProgress table progress
//
// swagger:model TableProgress
type TableProgress struct {

	// completed at
	// Format: date-time
	CompletedAt *strfmt.DateTime `json:"completed_at,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// failed
	Failed int64 `json:"failed,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`

	// skipped
	Skipped int64 `json:"skipped,omitempty"`

	// started at
	// Format: date-time
	StartedAt *strfmt.DateTime `json:"started_at,omitempty"`

	// table
	Table string `json:"table,omitempty"`

	// uploaded
	Uploaded int64 `json:"uploaded,omitempty"`
}

// Validate validates this table progress
func (m *TableProgress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompletedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TableProgress) validateCompletedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CompletedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("completed_at", "body", "date-time", m.CompletedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TableProgress) validateStartedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.StartedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("started_at", "body", "date-time", m.StartedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TableProgress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TableProgress) UnmarshalBinary(b []byte) error {
	var res TableProgress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
