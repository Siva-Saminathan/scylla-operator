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

// StorageServiceSlowQueryPostReader is a Reader for the StorageServiceSlowQueryPost structure.
type StorageServiceSlowQueryPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceSlowQueryPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceSlowQueryPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceSlowQueryPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceSlowQueryPostOK creates a StorageServiceSlowQueryPostOK with default headers values
func NewStorageServiceSlowQueryPostOK() *StorageServiceSlowQueryPostOK {
	return &StorageServiceSlowQueryPostOK{}
}

/*
StorageServiceSlowQueryPostOK handles this case with default header values.

StorageServiceSlowQueryPostOK storage service slow query post o k
*/
type StorageServiceSlowQueryPostOK struct {
}

func (o *StorageServiceSlowQueryPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageServiceSlowQueryPostDefault creates a StorageServiceSlowQueryPostDefault with default headers values
func NewStorageServiceSlowQueryPostDefault(code int) *StorageServiceSlowQueryPostDefault {
	return &StorageServiceSlowQueryPostDefault{
		_statusCode: code,
	}
}

/*
StorageServiceSlowQueryPostDefault handles this case with default header values.

internal server error
*/
type StorageServiceSlowQueryPostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service slow query post default response
func (o *StorageServiceSlowQueryPostDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceSlowQueryPostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceSlowQueryPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceSlowQueryPostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
