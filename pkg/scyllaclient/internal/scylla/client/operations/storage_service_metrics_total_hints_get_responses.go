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

// StorageServiceMetricsTotalHintsGetReader is a Reader for the StorageServiceMetricsTotalHintsGet structure.
type StorageServiceMetricsTotalHintsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceMetricsTotalHintsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceMetricsTotalHintsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceMetricsTotalHintsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceMetricsTotalHintsGetOK creates a StorageServiceMetricsTotalHintsGetOK with default headers values
func NewStorageServiceMetricsTotalHintsGetOK() *StorageServiceMetricsTotalHintsGetOK {
	return &StorageServiceMetricsTotalHintsGetOK{}
}

/*
StorageServiceMetricsTotalHintsGetOK handles this case with default header values.

StorageServiceMetricsTotalHintsGetOK storage service metrics total hints get o k
*/
type StorageServiceMetricsTotalHintsGetOK struct {
	Payload int32
}

func (o *StorageServiceMetricsTotalHintsGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *StorageServiceMetricsTotalHintsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageServiceMetricsTotalHintsGetDefault creates a StorageServiceMetricsTotalHintsGetDefault with default headers values
func NewStorageServiceMetricsTotalHintsGetDefault(code int) *StorageServiceMetricsTotalHintsGetDefault {
	return &StorageServiceMetricsTotalHintsGetDefault{
		_statusCode: code,
	}
}

/*
StorageServiceMetricsTotalHintsGetDefault handles this case with default header values.

internal server error
*/
type StorageServiceMetricsTotalHintsGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service metrics total hints get default response
func (o *StorageServiceMetricsTotalHintsGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceMetricsTotalHintsGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceMetricsTotalHintsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceMetricsTotalHintsGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
