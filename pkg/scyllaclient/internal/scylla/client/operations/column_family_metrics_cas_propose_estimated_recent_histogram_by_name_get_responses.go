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

// ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetReader is a Reader for the ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGet structure.
type ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetOK creates a ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetOK with default headers values
func NewColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetOK() *ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetOK {
	return &ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetOK{}
}

/*
ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetOK handles this case with default header values.

ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetOK column family metrics cas propose estimated recent histogram by name get o k
*/
type ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetOK struct {
}

func (o *ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetDefault creates a ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetDefault with default headers values
func NewColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetDefault(code int) *ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetDefault {
	return &ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetDefault{
		_statusCode: code,
	}
}

/*
ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics cas propose estimated recent histogram by name get default response
func (o *ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsCasProposeEstimatedRecentHistogramByNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
