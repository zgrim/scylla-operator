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

// NewFindConfigMurmur3PartitionerIgnoreMsbBitsParams creates a new FindConfigMurmur3PartitionerIgnoreMsbBitsParams object
// with the default values initialized.
func NewFindConfigMurmur3PartitionerIgnoreMsbBitsParams() *FindConfigMurmur3PartitionerIgnoreMsbBitsParams {

	return &FindConfigMurmur3PartitionerIgnoreMsbBitsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigMurmur3PartitionerIgnoreMsbBitsParamsWithTimeout creates a new FindConfigMurmur3PartitionerIgnoreMsbBitsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigMurmur3PartitionerIgnoreMsbBitsParamsWithTimeout(timeout time.Duration) *FindConfigMurmur3PartitionerIgnoreMsbBitsParams {

	return &FindConfigMurmur3PartitionerIgnoreMsbBitsParams{

		timeout: timeout,
	}
}

// NewFindConfigMurmur3PartitionerIgnoreMsbBitsParamsWithContext creates a new FindConfigMurmur3PartitionerIgnoreMsbBitsParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigMurmur3PartitionerIgnoreMsbBitsParamsWithContext(ctx context.Context) *FindConfigMurmur3PartitionerIgnoreMsbBitsParams {

	return &FindConfigMurmur3PartitionerIgnoreMsbBitsParams{

		Context: ctx,
	}
}

// NewFindConfigMurmur3PartitionerIgnoreMsbBitsParamsWithHTTPClient creates a new FindConfigMurmur3PartitionerIgnoreMsbBitsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigMurmur3PartitionerIgnoreMsbBitsParamsWithHTTPClient(client *http.Client) *FindConfigMurmur3PartitionerIgnoreMsbBitsParams {

	return &FindConfigMurmur3PartitionerIgnoreMsbBitsParams{
		HTTPClient: client,
	}
}

/*
FindConfigMurmur3PartitionerIgnoreMsbBitsParams contains all the parameters to send to the API endpoint
for the find config murmur3 partitioner ignore msb bits operation typically these are written to a http.Request
*/
type FindConfigMurmur3PartitionerIgnoreMsbBitsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config murmur3 partitioner ignore msb bits params
func (o *FindConfigMurmur3PartitionerIgnoreMsbBitsParams) WithTimeout(timeout time.Duration) *FindConfigMurmur3PartitionerIgnoreMsbBitsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config murmur3 partitioner ignore msb bits params
func (o *FindConfigMurmur3PartitionerIgnoreMsbBitsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config murmur3 partitioner ignore msb bits params
func (o *FindConfigMurmur3PartitionerIgnoreMsbBitsParams) WithContext(ctx context.Context) *FindConfigMurmur3PartitionerIgnoreMsbBitsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config murmur3 partitioner ignore msb bits params
func (o *FindConfigMurmur3PartitionerIgnoreMsbBitsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config murmur3 partitioner ignore msb bits params
func (o *FindConfigMurmur3PartitionerIgnoreMsbBitsParams) WithHTTPClient(client *http.Client) *FindConfigMurmur3PartitionerIgnoreMsbBitsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config murmur3 partitioner ignore msb bits params
func (o *FindConfigMurmur3PartitionerIgnoreMsbBitsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigMurmur3PartitionerIgnoreMsbBitsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
