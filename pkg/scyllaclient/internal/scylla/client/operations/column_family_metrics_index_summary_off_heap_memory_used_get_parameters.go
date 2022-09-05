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

// NewColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedGetParams creates a new ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedGetParams() *ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedGetParams {

	return &ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedGetParamsWithTimeout creates a new ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedGetParams {

	return &ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedGetParamsWithContext creates a new ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedGetParams {

	return &ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedGetParamsWithHTTPClient creates a new ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedGetParams {

	return &ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedGetParams{
		HTTPClient: client,
	}
}

/*
ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedGetParams contains all the parameters to send to the API endpoint
for the column family metrics index summary off heap memory used get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics index summary off heap memory used get params
func (o *ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics index summary off heap memory used get params
func (o *ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics index summary off heap memory used get params
func (o *ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics index summary off heap memory used get params
func (o *ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics index summary off heap memory used get params
func (o *ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics index summary off heap memory used get params
func (o *ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
