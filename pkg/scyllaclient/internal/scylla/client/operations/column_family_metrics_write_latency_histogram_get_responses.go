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

// ColumnFamilyMetricsWriteLatencyHistogramGetReader is a Reader for the ColumnFamilyMetricsWriteLatencyHistogramGet structure.
type ColumnFamilyMetricsWriteLatencyHistogramGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsWriteLatencyHistogramGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsWriteLatencyHistogramGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsWriteLatencyHistogramGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsWriteLatencyHistogramGetOK creates a ColumnFamilyMetricsWriteLatencyHistogramGetOK with default headers values
func NewColumnFamilyMetricsWriteLatencyHistogramGetOK() *ColumnFamilyMetricsWriteLatencyHistogramGetOK {
	return &ColumnFamilyMetricsWriteLatencyHistogramGetOK{}
}

/*
ColumnFamilyMetricsWriteLatencyHistogramGetOK handles this case with default header values.

ColumnFamilyMetricsWriteLatencyHistogramGetOK column family metrics write latency histogram get o k
*/
type ColumnFamilyMetricsWriteLatencyHistogramGetOK struct {
	Payload []*models.Histogram
}

func (o *ColumnFamilyMetricsWriteLatencyHistogramGetOK) GetPayload() []*models.Histogram {
	return o.Payload
}

func (o *ColumnFamilyMetricsWriteLatencyHistogramGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsWriteLatencyHistogramGetDefault creates a ColumnFamilyMetricsWriteLatencyHistogramGetDefault with default headers values
func NewColumnFamilyMetricsWriteLatencyHistogramGetDefault(code int) *ColumnFamilyMetricsWriteLatencyHistogramGetDefault {
	return &ColumnFamilyMetricsWriteLatencyHistogramGetDefault{
		_statusCode: code,
	}
}

/*
ColumnFamilyMetricsWriteLatencyHistogramGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsWriteLatencyHistogramGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics write latency histogram get default response
func (o *ColumnFamilyMetricsWriteLatencyHistogramGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsWriteLatencyHistogramGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsWriteLatencyHistogramGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsWriteLatencyHistogramGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
