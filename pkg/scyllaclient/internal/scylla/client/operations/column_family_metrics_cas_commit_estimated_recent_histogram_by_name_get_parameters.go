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

// NewColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams creates a new ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams() *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParamsWithTimeout creates a new ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParamsWithContext creates a new ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParamsWithHTTPClient creates a new ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams{
		HTTPClient: client,
	}
}

/*
ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams contains all the parameters to send to the API endpoint
for the column family metrics cas commit estimated recent histogram by name get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams struct {

	/*Name
	  The column family name in keyspace:name format

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics cas commit estimated recent histogram by name get params
func (o *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics cas commit estimated recent histogram by name get params
func (o *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics cas commit estimated recent histogram by name get params
func (o *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics cas commit estimated recent histogram by name get params
func (o *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics cas commit estimated recent histogram by name get params
func (o *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics cas commit estimated recent histogram by name get params
func (o *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the column family metrics cas commit estimated recent histogram by name get params
func (o *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams) WithName(name string) *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the column family metrics cas commit estimated recent histogram by name get params
func (o *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
