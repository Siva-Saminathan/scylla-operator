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

// CacheServiceRowCacheCapacityPostReader is a Reader for the CacheServiceRowCacheCapacityPost structure.
type CacheServiceRowCacheCapacityPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CacheServiceRowCacheCapacityPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCacheServiceRowCacheCapacityPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCacheServiceRowCacheCapacityPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCacheServiceRowCacheCapacityPostOK creates a CacheServiceRowCacheCapacityPostOK with default headers values
func NewCacheServiceRowCacheCapacityPostOK() *CacheServiceRowCacheCapacityPostOK {
	return &CacheServiceRowCacheCapacityPostOK{}
}

/*
CacheServiceRowCacheCapacityPostOK handles this case with default header values.

CacheServiceRowCacheCapacityPostOK cache service row cache capacity post o k
*/
type CacheServiceRowCacheCapacityPostOK struct {
}

func (o *CacheServiceRowCacheCapacityPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCacheServiceRowCacheCapacityPostDefault creates a CacheServiceRowCacheCapacityPostDefault with default headers values
func NewCacheServiceRowCacheCapacityPostDefault(code int) *CacheServiceRowCacheCapacityPostDefault {
	return &CacheServiceRowCacheCapacityPostDefault{
		_statusCode: code,
	}
}

/*
CacheServiceRowCacheCapacityPostDefault handles this case with default header values.

internal server error
*/
type CacheServiceRowCacheCapacityPostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the cache service row cache capacity post default response
func (o *CacheServiceRowCacheCapacityPostDefault) Code() int {
	return o._statusCode
}

func (o *CacheServiceRowCacheCapacityPostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CacheServiceRowCacheCapacityPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *CacheServiceRowCacheCapacityPostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
