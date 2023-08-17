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
	"github.com/go-openapi/strfmt"
)

// NewFindConfigDefaultLogLevelParams creates a new FindConfigDefaultLogLevelParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindConfigDefaultLogLevelParams() *FindConfigDefaultLogLevelParams {
	return &FindConfigDefaultLogLevelParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigDefaultLogLevelParamsWithTimeout creates a new FindConfigDefaultLogLevelParams object
// with the ability to set a timeout on a request.
func NewFindConfigDefaultLogLevelParamsWithTimeout(timeout time.Duration) *FindConfigDefaultLogLevelParams {
	return &FindConfigDefaultLogLevelParams{
		timeout: timeout,
	}
}

// NewFindConfigDefaultLogLevelParamsWithContext creates a new FindConfigDefaultLogLevelParams object
// with the ability to set a context for a request.
func NewFindConfigDefaultLogLevelParamsWithContext(ctx context.Context) *FindConfigDefaultLogLevelParams {
	return &FindConfigDefaultLogLevelParams{
		Context: ctx,
	}
}

// NewFindConfigDefaultLogLevelParamsWithHTTPClient creates a new FindConfigDefaultLogLevelParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindConfigDefaultLogLevelParamsWithHTTPClient(client *http.Client) *FindConfigDefaultLogLevelParams {
	return &FindConfigDefaultLogLevelParams{
		HTTPClient: client,
	}
}

/*
FindConfigDefaultLogLevelParams contains all the parameters to send to the API endpoint

	for the find config default log level operation.

	Typically these are written to a http.Request.
*/
type FindConfigDefaultLogLevelParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find config default log level params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindConfigDefaultLogLevelParams) WithDefaults() *FindConfigDefaultLogLevelParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find config default log level params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindConfigDefaultLogLevelParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find config default log level params
func (o *FindConfigDefaultLogLevelParams) WithTimeout(timeout time.Duration) *FindConfigDefaultLogLevelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config default log level params
func (o *FindConfigDefaultLogLevelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config default log level params
func (o *FindConfigDefaultLogLevelParams) WithContext(ctx context.Context) *FindConfigDefaultLogLevelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config default log level params
func (o *FindConfigDefaultLogLevelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config default log level params
func (o *FindConfigDefaultLogLevelParams) WithHTTPClient(client *http.Client) *FindConfigDefaultLogLevelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config default log level params
func (o *FindConfigDefaultLogLevelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigDefaultLogLevelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
