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

// ColumnFamilyMetricsLiveDiskSpaceUsedByNameGetReader is a Reader for the ColumnFamilyMetricsLiveDiskSpaceUsedByNameGet structure.
type ColumnFamilyMetricsLiveDiskSpaceUsedByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsLiveDiskSpaceUsedByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsLiveDiskSpaceUsedByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsLiveDiskSpaceUsedByNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsLiveDiskSpaceUsedByNameGetOK creates a ColumnFamilyMetricsLiveDiskSpaceUsedByNameGetOK with default headers values
func NewColumnFamilyMetricsLiveDiskSpaceUsedByNameGetOK() *ColumnFamilyMetricsLiveDiskSpaceUsedByNameGetOK {
	return &ColumnFamilyMetricsLiveDiskSpaceUsedByNameGetOK{}
}

/*
ColumnFamilyMetricsLiveDiskSpaceUsedByNameGetOK handles this case with default header values.

ColumnFamilyMetricsLiveDiskSpaceUsedByNameGetOK column family metrics live disk space used by name get o k
*/
type ColumnFamilyMetricsLiveDiskSpaceUsedByNameGetOK struct {
	Payload int32
}

func (o *ColumnFamilyMetricsLiveDiskSpaceUsedByNameGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *ColumnFamilyMetricsLiveDiskSpaceUsedByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsLiveDiskSpaceUsedByNameGetDefault creates a ColumnFamilyMetricsLiveDiskSpaceUsedByNameGetDefault with default headers values
func NewColumnFamilyMetricsLiveDiskSpaceUsedByNameGetDefault(code int) *ColumnFamilyMetricsLiveDiskSpaceUsedByNameGetDefault {
	return &ColumnFamilyMetricsLiveDiskSpaceUsedByNameGetDefault{
		_statusCode: code,
	}
}

/*
ColumnFamilyMetricsLiveDiskSpaceUsedByNameGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsLiveDiskSpaceUsedByNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics live disk space used by name get default response
func (o *ColumnFamilyMetricsLiveDiskSpaceUsedByNameGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsLiveDiskSpaceUsedByNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsLiveDiskSpaceUsedByNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsLiveDiskSpaceUsedByNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
