// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/scylla/models"
)

// ColumnFamilyMetricsKeyCacheHitRateByNameGetReader is a Reader for the ColumnFamilyMetricsKeyCacheHitRateByNameGet structure.
type ColumnFamilyMetricsKeyCacheHitRateByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsKeyCacheHitRateByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsKeyCacheHitRateByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsKeyCacheHitRateByNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsKeyCacheHitRateByNameGetOK creates a ColumnFamilyMetricsKeyCacheHitRateByNameGetOK with default headers values
func NewColumnFamilyMetricsKeyCacheHitRateByNameGetOK() *ColumnFamilyMetricsKeyCacheHitRateByNameGetOK {
	return &ColumnFamilyMetricsKeyCacheHitRateByNameGetOK{}
}

/*
ColumnFamilyMetricsKeyCacheHitRateByNameGetOK handles this case with default header values.

ColumnFamilyMetricsKeyCacheHitRateByNameGetOK column family metrics key cache hit rate by name get o k
*/
type ColumnFamilyMetricsKeyCacheHitRateByNameGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsKeyCacheHitRateByNameGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyMetricsKeyCacheHitRateByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsKeyCacheHitRateByNameGetDefault creates a ColumnFamilyMetricsKeyCacheHitRateByNameGetDefault with default headers values
func NewColumnFamilyMetricsKeyCacheHitRateByNameGetDefault(code int) *ColumnFamilyMetricsKeyCacheHitRateByNameGetDefault {
	return &ColumnFamilyMetricsKeyCacheHitRateByNameGetDefault{
		_statusCode: code,
	}
}

/*
ColumnFamilyMetricsKeyCacheHitRateByNameGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsKeyCacheHitRateByNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics key cache hit rate by name get default response
func (o *ColumnFamilyMetricsKeyCacheHitRateByNameGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsKeyCacheHitRateByNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsKeyCacheHitRateByNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsKeyCacheHitRateByNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
