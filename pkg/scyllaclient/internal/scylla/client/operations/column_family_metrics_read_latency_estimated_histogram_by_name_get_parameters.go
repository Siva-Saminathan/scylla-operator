// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetParams creates a new ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetParams() *ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetParamsWithTimeout creates a new ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetParamsWithContext creates a new ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetParamsWithHTTPClient creates a new ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetParams{
		HTTPClient: client,
	}
}

/*
ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetParams contains all the parameters to send to the API endpoint
for the column family metrics read latency estimated histogram by name get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetParams struct {

	/*Name
	  The column family name in keyspace:name format

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics read latency estimated histogram by name get params
func (o *ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics read latency estimated histogram by name get params
func (o *ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics read latency estimated histogram by name get params
func (o *ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics read latency estimated histogram by name get params
func (o *ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics read latency estimated histogram by name get params
func (o *ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics read latency estimated histogram by name get params
func (o *ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the column family metrics read latency estimated histogram by name get params
func (o *ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetParams) WithName(name string) *ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the column family metrics read latency estimated histogram by name get params
func (o *ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
