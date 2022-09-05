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

// NewFindConfigSkipWaitForGossipToSettleParams creates a new FindConfigSkipWaitForGossipToSettleParams object
// with the default values initialized.
func NewFindConfigSkipWaitForGossipToSettleParams() *FindConfigSkipWaitForGossipToSettleParams {

	return &FindConfigSkipWaitForGossipToSettleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigSkipWaitForGossipToSettleParamsWithTimeout creates a new FindConfigSkipWaitForGossipToSettleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigSkipWaitForGossipToSettleParamsWithTimeout(timeout time.Duration) *FindConfigSkipWaitForGossipToSettleParams {

	return &FindConfigSkipWaitForGossipToSettleParams{

		timeout: timeout,
	}
}

// NewFindConfigSkipWaitForGossipToSettleParamsWithContext creates a new FindConfigSkipWaitForGossipToSettleParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigSkipWaitForGossipToSettleParamsWithContext(ctx context.Context) *FindConfigSkipWaitForGossipToSettleParams {

	return &FindConfigSkipWaitForGossipToSettleParams{

		Context: ctx,
	}
}

// NewFindConfigSkipWaitForGossipToSettleParamsWithHTTPClient creates a new FindConfigSkipWaitForGossipToSettleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigSkipWaitForGossipToSettleParamsWithHTTPClient(client *http.Client) *FindConfigSkipWaitForGossipToSettleParams {

	return &FindConfigSkipWaitForGossipToSettleParams{
		HTTPClient: client,
	}
}

/*
FindConfigSkipWaitForGossipToSettleParams contains all the parameters to send to the API endpoint
for the find config skip wait for gossip to settle operation typically these are written to a http.Request
*/
type FindConfigSkipWaitForGossipToSettleParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config skip wait for gossip to settle params
func (o *FindConfigSkipWaitForGossipToSettleParams) WithTimeout(timeout time.Duration) *FindConfigSkipWaitForGossipToSettleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config skip wait for gossip to settle params
func (o *FindConfigSkipWaitForGossipToSettleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config skip wait for gossip to settle params
func (o *FindConfigSkipWaitForGossipToSettleParams) WithContext(ctx context.Context) *FindConfigSkipWaitForGossipToSettleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config skip wait for gossip to settle params
func (o *FindConfigSkipWaitForGossipToSettleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config skip wait for gossip to settle params
func (o *FindConfigSkipWaitForGossipToSettleParams) WithHTTPClient(client *http.Client) *FindConfigSkipWaitForGossipToSettleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config skip wait for gossip to settle params
func (o *FindConfigSkipWaitForGossipToSettleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigSkipWaitForGossipToSettleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
