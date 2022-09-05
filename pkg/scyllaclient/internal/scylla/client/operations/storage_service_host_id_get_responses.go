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

// StorageServiceHostIDGetReader is a Reader for the StorageServiceHostIDGet structure.
type StorageServiceHostIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceHostIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceHostIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceHostIDGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceHostIDGetOK creates a StorageServiceHostIDGetOK with default headers values
func NewStorageServiceHostIDGetOK() *StorageServiceHostIDGetOK {
	return &StorageServiceHostIDGetOK{}
}

/*
StorageServiceHostIDGetOK handles this case with default header values.

StorageServiceHostIDGetOK storage service host Id get o k
*/
type StorageServiceHostIDGetOK struct {
	Payload []*models.Mapper
}

func (o *StorageServiceHostIDGetOK) GetPayload() []*models.Mapper {
	return o.Payload
}

func (o *StorageServiceHostIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageServiceHostIDGetDefault creates a StorageServiceHostIDGetDefault with default headers values
func NewStorageServiceHostIDGetDefault(code int) *StorageServiceHostIDGetDefault {
	return &StorageServiceHostIDGetDefault{
		_statusCode: code,
	}
}

/*
StorageServiceHostIDGetDefault handles this case with default header values.

internal server error
*/
type StorageServiceHostIDGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service host Id get default response
func (o *StorageServiceHostIDGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceHostIDGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceHostIDGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceHostIDGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
