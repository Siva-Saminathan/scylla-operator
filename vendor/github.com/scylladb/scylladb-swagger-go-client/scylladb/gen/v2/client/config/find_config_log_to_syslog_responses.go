// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylladb-swagger-go-client/scylladb/gen/v2/models"
)

// FindConfigLogToSyslogReader is a Reader for the FindConfigLogToSyslog structure.
type FindConfigLogToSyslogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigLogToSyslogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigLogToSyslogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigLogToSyslogDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigLogToSyslogOK creates a FindConfigLogToSyslogOK with default headers values
func NewFindConfigLogToSyslogOK() *FindConfigLogToSyslogOK {
	return &FindConfigLogToSyslogOK{}
}

/*
FindConfigLogToSyslogOK handles this case with default header values.

Config value
*/
type FindConfigLogToSyslogOK struct {
	Payload bool
}

func (o *FindConfigLogToSyslogOK) GetPayload() bool {
	return o.Payload
}

func (o *FindConfigLogToSyslogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigLogToSyslogDefault creates a FindConfigLogToSyslogDefault with default headers values
func NewFindConfigLogToSyslogDefault(code int) *FindConfigLogToSyslogDefault {
	return &FindConfigLogToSyslogDefault{
		_statusCode: code,
	}
}

/*
FindConfigLogToSyslogDefault handles this case with default header values.

unexpected error
*/
type FindConfigLogToSyslogDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config log to syslog default response
func (o *FindConfigLogToSyslogDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigLogToSyslogDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigLogToSyslogDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigLogToSyslogDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
