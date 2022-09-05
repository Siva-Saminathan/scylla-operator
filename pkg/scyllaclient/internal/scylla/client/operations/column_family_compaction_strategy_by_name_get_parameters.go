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

// NewColumnFamilyCompactionStrategyByNameGetParams creates a new ColumnFamilyCompactionStrategyByNameGetParams object
// with the default values initialized.
func NewColumnFamilyCompactionStrategyByNameGetParams() *ColumnFamilyCompactionStrategyByNameGetParams {
	var ()
	return &ColumnFamilyCompactionStrategyByNameGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyCompactionStrategyByNameGetParamsWithTimeout creates a new ColumnFamilyCompactionStrategyByNameGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyCompactionStrategyByNameGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyCompactionStrategyByNameGetParams {
	var ()
	return &ColumnFamilyCompactionStrategyByNameGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyCompactionStrategyByNameGetParamsWithContext creates a new ColumnFamilyCompactionStrategyByNameGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyCompactionStrategyByNameGetParamsWithContext(ctx context.Context) *ColumnFamilyCompactionStrategyByNameGetParams {
	var ()
	return &ColumnFamilyCompactionStrategyByNameGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyCompactionStrategyByNameGetParamsWithHTTPClient creates a new ColumnFamilyCompactionStrategyByNameGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyCompactionStrategyByNameGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyCompactionStrategyByNameGetParams {
	var ()
	return &ColumnFamilyCompactionStrategyByNameGetParams{
		HTTPClient: client,
	}
}

/*
ColumnFamilyCompactionStrategyByNameGetParams contains all the parameters to send to the API endpoint
for the column family compaction strategy by name get operation typically these are written to a http.Request
*/
type ColumnFamilyCompactionStrategyByNameGetParams struct {

	/*Name
	  The column family name in keyspace:name format

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family compaction strategy by name get params
func (o *ColumnFamilyCompactionStrategyByNameGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyCompactionStrategyByNameGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family compaction strategy by name get params
func (o *ColumnFamilyCompactionStrategyByNameGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family compaction strategy by name get params
func (o *ColumnFamilyCompactionStrategyByNameGetParams) WithContext(ctx context.Context) *ColumnFamilyCompactionStrategyByNameGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family compaction strategy by name get params
func (o *ColumnFamilyCompactionStrategyByNameGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family compaction strategy by name get params
func (o *ColumnFamilyCompactionStrategyByNameGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyCompactionStrategyByNameGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family compaction strategy by name get params
func (o *ColumnFamilyCompactionStrategyByNameGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the column family compaction strategy by name get params
func (o *ColumnFamilyCompactionStrategyByNameGetParams) WithName(name string) *ColumnFamilyCompactionStrategyByNameGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the column family compaction strategy by name get params
func (o *ColumnFamilyCompactionStrategyByNameGetParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyCompactionStrategyByNameGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
