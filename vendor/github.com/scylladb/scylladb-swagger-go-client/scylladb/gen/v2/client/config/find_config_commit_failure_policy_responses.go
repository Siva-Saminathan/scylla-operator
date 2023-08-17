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

// FindConfigCommitFailurePolicyReader is a Reader for the FindConfigCommitFailurePolicy structure.
type FindConfigCommitFailurePolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigCommitFailurePolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigCommitFailurePolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigCommitFailurePolicyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigCommitFailurePolicyOK creates a FindConfigCommitFailurePolicyOK with default headers values
func NewFindConfigCommitFailurePolicyOK() *FindConfigCommitFailurePolicyOK {
	return &FindConfigCommitFailurePolicyOK{}
}

/*
FindConfigCommitFailurePolicyOK handles this case with default header values.

Config value
*/
type FindConfigCommitFailurePolicyOK struct {
	Payload string
}

func (o *FindConfigCommitFailurePolicyOK) GetPayload() string {
	return o.Payload
}

func (o *FindConfigCommitFailurePolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigCommitFailurePolicyDefault creates a FindConfigCommitFailurePolicyDefault with default headers values
func NewFindConfigCommitFailurePolicyDefault(code int) *FindConfigCommitFailurePolicyDefault {
	return &FindConfigCommitFailurePolicyDefault{
		_statusCode: code,
	}
}

/*
FindConfigCommitFailurePolicyDefault handles this case with default header values.

unexpected error
*/
type FindConfigCommitFailurePolicyDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config commit failure policy default response
func (o *FindConfigCommitFailurePolicyDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigCommitFailurePolicyDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigCommitFailurePolicyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigCommitFailurePolicyDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
