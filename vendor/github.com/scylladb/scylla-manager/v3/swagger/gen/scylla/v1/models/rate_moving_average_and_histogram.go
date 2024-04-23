// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RateMovingAverageAndHistogram rate_moving_average_and_histogram
//
// # A timer metric which aggregates timing durations and provides duration statistics, plus throughput statistics
//
// swagger:model rate_moving_average_and_histogram
type RateMovingAverageAndHistogram struct {

	// hist
	Hist *Histogram `json:"hist,omitempty"`

	// meter
	Meter *RateMovingAverage `json:"meter,omitempty"`
}

// Validate validates this rate moving average and histogram
func (m *RateMovingAverageAndHistogram) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHist(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMeter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RateMovingAverageAndHistogram) validateHist(formats strfmt.Registry) error {

	if swag.IsZero(m.Hist) { // not required
		return nil
	}

	if m.Hist != nil {
		if err := m.Hist.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hist")
			}
			return err
		}
	}

	return nil
}

func (m *RateMovingAverageAndHistogram) validateMeter(formats strfmt.Registry) error {

	if swag.IsZero(m.Meter) { // not required
		return nil
	}

	if m.Meter != nil {
		if err := m.Meter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RateMovingAverageAndHistogram) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RateMovingAverageAndHistogram) UnmarshalBinary(b []byte) error {
	var res RateMovingAverageAndHistogram
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}