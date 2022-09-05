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

// FindConfigJoinRingReader is a Reader for the FindConfigJoinRing structure.
type FindConfigJoinRingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigJoinRingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigJoinRingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigJoinRingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigJoinRingOK creates a FindConfigJoinRingOK with default headers values
func NewFindConfigJoinRingOK() *FindConfigJoinRingOK {
	return &FindConfigJoinRingOK{}
}

/*
FindConfigJoinRingOK handles this case with default header values.

Config value
*/
type FindConfigJoinRingOK struct {
	Payload bool
}

func (o *FindConfigJoinRingOK) GetPayload() bool {
	return o.Payload
}

func (o *FindConfigJoinRingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigJoinRingDefault creates a FindConfigJoinRingDefault with default headers values
func NewFindConfigJoinRingDefault(code int) *FindConfigJoinRingDefault {
	return &FindConfigJoinRingDefault{
		_statusCode: code,
	}
}

/*
FindConfigJoinRingDefault handles this case with default header values.

unexpected error
*/
type FindConfigJoinRingDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config join ring default response
func (o *FindConfigJoinRingDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigJoinRingDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigJoinRingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigJoinRingDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
