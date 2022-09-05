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

// NewCompactionManagerForceUserDefinedCompactionPostParams creates a new CompactionManagerForceUserDefinedCompactionPostParams object
// with the default values initialized.
func NewCompactionManagerForceUserDefinedCompactionPostParams() *CompactionManagerForceUserDefinedCompactionPostParams {
	var ()
	return &CompactionManagerForceUserDefinedCompactionPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCompactionManagerForceUserDefinedCompactionPostParamsWithTimeout creates a new CompactionManagerForceUserDefinedCompactionPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCompactionManagerForceUserDefinedCompactionPostParamsWithTimeout(timeout time.Duration) *CompactionManagerForceUserDefinedCompactionPostParams {
	var ()
	return &CompactionManagerForceUserDefinedCompactionPostParams{

		timeout: timeout,
	}
}

// NewCompactionManagerForceUserDefinedCompactionPostParamsWithContext creates a new CompactionManagerForceUserDefinedCompactionPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewCompactionManagerForceUserDefinedCompactionPostParamsWithContext(ctx context.Context) *CompactionManagerForceUserDefinedCompactionPostParams {
	var ()
	return &CompactionManagerForceUserDefinedCompactionPostParams{

		Context: ctx,
	}
}

// NewCompactionManagerForceUserDefinedCompactionPostParamsWithHTTPClient creates a new CompactionManagerForceUserDefinedCompactionPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCompactionManagerForceUserDefinedCompactionPostParamsWithHTTPClient(client *http.Client) *CompactionManagerForceUserDefinedCompactionPostParams {
	var ()
	return &CompactionManagerForceUserDefinedCompactionPostParams{
		HTTPClient: client,
	}
}

/*
CompactionManagerForceUserDefinedCompactionPostParams contains all the parameters to send to the API endpoint
for the compaction manager force user defined compaction post operation typically these are written to a http.Request
*/
type CompactionManagerForceUserDefinedCompactionPostParams struct {

	/*DataFiles
	  a comma separated list of sstable file to compact. must contain keyspace and columnfamily name in path(for 2.1+) or file name itself

	*/
	DataFiles string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the compaction manager force user defined compaction post params
func (o *CompactionManagerForceUserDefinedCompactionPostParams) WithTimeout(timeout time.Duration) *CompactionManagerForceUserDefinedCompactionPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the compaction manager force user defined compaction post params
func (o *CompactionManagerForceUserDefinedCompactionPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the compaction manager force user defined compaction post params
func (o *CompactionManagerForceUserDefinedCompactionPostParams) WithContext(ctx context.Context) *CompactionManagerForceUserDefinedCompactionPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the compaction manager force user defined compaction post params
func (o *CompactionManagerForceUserDefinedCompactionPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the compaction manager force user defined compaction post params
func (o *CompactionManagerForceUserDefinedCompactionPostParams) WithHTTPClient(client *http.Client) *CompactionManagerForceUserDefinedCompactionPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the compaction manager force user defined compaction post params
func (o *CompactionManagerForceUserDefinedCompactionPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDataFiles adds the dataFiles to the compaction manager force user defined compaction post params
func (o *CompactionManagerForceUserDefinedCompactionPostParams) WithDataFiles(dataFiles string) *CompactionManagerForceUserDefinedCompactionPostParams {
	o.SetDataFiles(dataFiles)
	return o
}

// SetDataFiles adds the dataFiles to the compaction manager force user defined compaction post params
func (o *CompactionManagerForceUserDefinedCompactionPostParams) SetDataFiles(dataFiles string) {
	o.DataFiles = dataFiles
}

// WriteToRequest writes these params to a swagger request
func (o *CompactionManagerForceUserDefinedCompactionPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param data_files
	qrDataFiles := o.DataFiles
	qDataFiles := qrDataFiles
	if qDataFiles != "" {
		if err := r.SetQueryParam("data_files", qDataFiles); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
