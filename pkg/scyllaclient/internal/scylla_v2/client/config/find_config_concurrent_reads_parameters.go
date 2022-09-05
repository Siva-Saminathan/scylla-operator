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

// NewFindConfigConcurrentReadsParams creates a new FindConfigConcurrentReadsParams object
// with the default values initialized.
func NewFindConfigConcurrentReadsParams() *FindConfigConcurrentReadsParams {

	return &FindConfigConcurrentReadsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigConcurrentReadsParamsWithTimeout creates a new FindConfigConcurrentReadsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigConcurrentReadsParamsWithTimeout(timeout time.Duration) *FindConfigConcurrentReadsParams {

	return &FindConfigConcurrentReadsParams{

		timeout: timeout,
	}
}

// NewFindConfigConcurrentReadsParamsWithContext creates a new FindConfigConcurrentReadsParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigConcurrentReadsParamsWithContext(ctx context.Context) *FindConfigConcurrentReadsParams {

	return &FindConfigConcurrentReadsParams{

		Context: ctx,
	}
}

// NewFindConfigConcurrentReadsParamsWithHTTPClient creates a new FindConfigConcurrentReadsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigConcurrentReadsParamsWithHTTPClient(client *http.Client) *FindConfigConcurrentReadsParams {

	return &FindConfigConcurrentReadsParams{
		HTTPClient: client,
	}
}

/*
FindConfigConcurrentReadsParams contains all the parameters to send to the API endpoint
for the find config concurrent reads operation typically these are written to a http.Request
*/
type FindConfigConcurrentReadsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config concurrent reads params
func (o *FindConfigConcurrentReadsParams) WithTimeout(timeout time.Duration) *FindConfigConcurrentReadsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config concurrent reads params
func (o *FindConfigConcurrentReadsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config concurrent reads params
func (o *FindConfigConcurrentReadsParams) WithContext(ctx context.Context) *FindConfigConcurrentReadsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config concurrent reads params
func (o *FindConfigConcurrentReadsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config concurrent reads params
func (o *FindConfigConcurrentReadsParams) WithHTTPClient(client *http.Client) *FindConfigConcurrentReadsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config concurrent reads params
func (o *FindConfigConcurrentReadsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigConcurrentReadsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
