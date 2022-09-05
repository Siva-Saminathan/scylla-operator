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

// NewStorageServiceRPCServerGetParams creates a new StorageServiceRPCServerGetParams object
// with the default values initialized.
func NewStorageServiceRPCServerGetParams() *StorageServiceRPCServerGetParams {

	return &StorageServiceRPCServerGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageServiceRPCServerGetParamsWithTimeout creates a new StorageServiceRPCServerGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageServiceRPCServerGetParamsWithTimeout(timeout time.Duration) *StorageServiceRPCServerGetParams {

	return &StorageServiceRPCServerGetParams{

		timeout: timeout,
	}
}

// NewStorageServiceRPCServerGetParamsWithContext creates a new StorageServiceRPCServerGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageServiceRPCServerGetParamsWithContext(ctx context.Context) *StorageServiceRPCServerGetParams {

	return &StorageServiceRPCServerGetParams{

		Context: ctx,
	}
}

// NewStorageServiceRPCServerGetParamsWithHTTPClient creates a new StorageServiceRPCServerGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageServiceRPCServerGetParamsWithHTTPClient(client *http.Client) *StorageServiceRPCServerGetParams {

	return &StorageServiceRPCServerGetParams{
		HTTPClient: client,
	}
}

/*
StorageServiceRPCServerGetParams contains all the parameters to send to the API endpoint
for the storage service Rpc server get operation typically these are written to a http.Request
*/
type StorageServiceRPCServerGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage service Rpc server get params
func (o *StorageServiceRPCServerGetParams) WithTimeout(timeout time.Duration) *StorageServiceRPCServerGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service Rpc server get params
func (o *StorageServiceRPCServerGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage service Rpc server get params
func (o *StorageServiceRPCServerGetParams) WithContext(ctx context.Context) *StorageServiceRPCServerGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service Rpc server get params
func (o *StorageServiceRPCServerGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service Rpc server get params
func (o *StorageServiceRPCServerGetParams) WithHTTPClient(client *http.Client) *StorageServiceRPCServerGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service Rpc server get params
func (o *StorageServiceRPCServerGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServiceRPCServerGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
