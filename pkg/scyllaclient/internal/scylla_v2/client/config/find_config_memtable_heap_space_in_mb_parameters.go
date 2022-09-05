// Code generated by go-swagger; DO NOT EDIT.

package config

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

// NewFindConfigMemtableHeapSpaceInMbParams creates a new FindConfigMemtableHeapSpaceInMbParams object
// with the default values initialized.
func NewFindConfigMemtableHeapSpaceInMbParams() *FindConfigMemtableHeapSpaceInMbParams {

	return &FindConfigMemtableHeapSpaceInMbParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigMemtableHeapSpaceInMbParamsWithTimeout creates a new FindConfigMemtableHeapSpaceInMbParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigMemtableHeapSpaceInMbParamsWithTimeout(timeout time.Duration) *FindConfigMemtableHeapSpaceInMbParams {

	return &FindConfigMemtableHeapSpaceInMbParams{

		timeout: timeout,
	}
}

// NewFindConfigMemtableHeapSpaceInMbParamsWithContext creates a new FindConfigMemtableHeapSpaceInMbParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigMemtableHeapSpaceInMbParamsWithContext(ctx context.Context) *FindConfigMemtableHeapSpaceInMbParams {

	return &FindConfigMemtableHeapSpaceInMbParams{

		Context: ctx,
	}
}

// NewFindConfigMemtableHeapSpaceInMbParamsWithHTTPClient creates a new FindConfigMemtableHeapSpaceInMbParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigMemtableHeapSpaceInMbParamsWithHTTPClient(client *http.Client) *FindConfigMemtableHeapSpaceInMbParams {

	return &FindConfigMemtableHeapSpaceInMbParams{
		HTTPClient: client,
	}
}

/*
FindConfigMemtableHeapSpaceInMbParams contains all the parameters to send to the API endpoint
for the find config memtable heap space in mb operation typically these are written to a http.Request
*/
type FindConfigMemtableHeapSpaceInMbParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config memtable heap space in mb params
func (o *FindConfigMemtableHeapSpaceInMbParams) WithTimeout(timeout time.Duration) *FindConfigMemtableHeapSpaceInMbParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config memtable heap space in mb params
func (o *FindConfigMemtableHeapSpaceInMbParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config memtable heap space in mb params
func (o *FindConfigMemtableHeapSpaceInMbParams) WithContext(ctx context.Context) *FindConfigMemtableHeapSpaceInMbParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config memtable heap space in mb params
func (o *FindConfigMemtableHeapSpaceInMbParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config memtable heap space in mb params
func (o *FindConfigMemtableHeapSpaceInMbParams) WithHTTPClient(client *http.Client) *FindConfigMemtableHeapSpaceInMbParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config memtable heap space in mb params
func (o *FindConfigMemtableHeapSpaceInMbParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigMemtableHeapSpaceInMbParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
