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

// CommitlogMetricsCompletedTasksGetReader is a Reader for the CommitlogMetricsCompletedTasksGet structure.
type CommitlogMetricsCompletedTasksGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CommitlogMetricsCompletedTasksGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCommitlogMetricsCompletedTasksGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCommitlogMetricsCompletedTasksGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCommitlogMetricsCompletedTasksGetOK creates a CommitlogMetricsCompletedTasksGetOK with default headers values
func NewCommitlogMetricsCompletedTasksGetOK() *CommitlogMetricsCompletedTasksGetOK {
	return &CommitlogMetricsCompletedTasksGetOK{}
}

/*
CommitlogMetricsCompletedTasksGetOK handles this case with default header values.

CommitlogMetricsCompletedTasksGetOK commitlog metrics completed tasks get o k
*/
type CommitlogMetricsCompletedTasksGetOK struct {
	Payload interface{}
}

func (o *CommitlogMetricsCompletedTasksGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CommitlogMetricsCompletedTasksGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCommitlogMetricsCompletedTasksGetDefault creates a CommitlogMetricsCompletedTasksGetDefault with default headers values
func NewCommitlogMetricsCompletedTasksGetDefault(code int) *CommitlogMetricsCompletedTasksGetDefault {
	return &CommitlogMetricsCompletedTasksGetDefault{
		_statusCode: code,
	}
}

/*
CommitlogMetricsCompletedTasksGetDefault handles this case with default header values.

internal server error
*/
type CommitlogMetricsCompletedTasksGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the commitlog metrics completed tasks get default response
func (o *CommitlogMetricsCompletedTasksGetDefault) Code() int {
	return o._statusCode
}

func (o *CommitlogMetricsCompletedTasksGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CommitlogMetricsCompletedTasksGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *CommitlogMetricsCompletedTasksGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
