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

// StorageServiceKeyspaceCleanupByKeyspacePostReader is a Reader for the StorageServiceKeyspaceCleanupByKeyspacePost structure.
type StorageServiceKeyspaceCleanupByKeyspacePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceKeyspaceCleanupByKeyspacePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceKeyspaceCleanupByKeyspacePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceKeyspaceCleanupByKeyspacePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceKeyspaceCleanupByKeyspacePostOK creates a StorageServiceKeyspaceCleanupByKeyspacePostOK with default headers values
func NewStorageServiceKeyspaceCleanupByKeyspacePostOK() *StorageServiceKeyspaceCleanupByKeyspacePostOK {
	return &StorageServiceKeyspaceCleanupByKeyspacePostOK{}
}

/*
StorageServiceKeyspaceCleanupByKeyspacePostOK handles this case with default header values.

StorageServiceKeyspaceCleanupByKeyspacePostOK storage service keyspace cleanup by keyspace post o k
*/
type StorageServiceKeyspaceCleanupByKeyspacePostOK struct {
	Payload int32
}

func (o *StorageServiceKeyspaceCleanupByKeyspacePostOK) GetPayload() int32 {
	return o.Payload
}

func (o *StorageServiceKeyspaceCleanupByKeyspacePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageServiceKeyspaceCleanupByKeyspacePostDefault creates a StorageServiceKeyspaceCleanupByKeyspacePostDefault with default headers values
func NewStorageServiceKeyspaceCleanupByKeyspacePostDefault(code int) *StorageServiceKeyspaceCleanupByKeyspacePostDefault {
	return &StorageServiceKeyspaceCleanupByKeyspacePostDefault{
		_statusCode: code,
	}
}

/*
StorageServiceKeyspaceCleanupByKeyspacePostDefault handles this case with default header values.

internal server error
*/
type StorageServiceKeyspaceCleanupByKeyspacePostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service keyspace cleanup by keyspace post default response
func (o *StorageServiceKeyspaceCleanupByKeyspacePostDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceKeyspaceCleanupByKeyspacePostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceKeyspaceCleanupByKeyspacePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceKeyspaceCleanupByKeyspacePostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
