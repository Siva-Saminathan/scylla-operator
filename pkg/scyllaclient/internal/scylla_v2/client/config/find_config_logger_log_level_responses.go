// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/scylla_v2/models"
)

// FindConfigLoggerLogLevelReader is a Reader for the FindConfigLoggerLogLevel structure.
type FindConfigLoggerLogLevelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigLoggerLogLevelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigLoggerLogLevelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigLoggerLogLevelDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigLoggerLogLevelOK creates a FindConfigLoggerLogLevelOK with default headers values
func NewFindConfigLoggerLogLevelOK() *FindConfigLoggerLogLevelOK {
	return &FindConfigLoggerLogLevelOK{}
}

/*
FindConfigLoggerLogLevelOK handles this case with default header values.

Config value
*/
type FindConfigLoggerLogLevelOK struct {
	Payload []string
}

func (o *FindConfigLoggerLogLevelOK) GetPayload() []string {
	return o.Payload
}

func (o *FindConfigLoggerLogLevelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigLoggerLogLevelDefault creates a FindConfigLoggerLogLevelDefault with default headers values
func NewFindConfigLoggerLogLevelDefault(code int) *FindConfigLoggerLogLevelDefault {
	return &FindConfigLoggerLogLevelDefault{
		_statusCode: code,
	}
}

/*
FindConfigLoggerLogLevelDefault handles this case with default header values.

unexpected error
*/
type FindConfigLoggerLogLevelDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config logger log level default response
func (o *FindConfigLoggerLogLevelDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigLoggerLogLevelDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigLoggerLogLevelDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigLoggerLogLevelDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
