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

// NewStorageServiceIncrementalBackupsGetParams creates a new StorageServiceIncrementalBackupsGetParams object
// with the default values initialized.
func NewStorageServiceIncrementalBackupsGetParams() *StorageServiceIncrementalBackupsGetParams {

	return &StorageServiceIncrementalBackupsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageServiceIncrementalBackupsGetParamsWithTimeout creates a new StorageServiceIncrementalBackupsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageServiceIncrementalBackupsGetParamsWithTimeout(timeout time.Duration) *StorageServiceIncrementalBackupsGetParams {

	return &StorageServiceIncrementalBackupsGetParams{

		timeout: timeout,
	}
}

// NewStorageServiceIncrementalBackupsGetParamsWithContext creates a new StorageServiceIncrementalBackupsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageServiceIncrementalBackupsGetParamsWithContext(ctx context.Context) *StorageServiceIncrementalBackupsGetParams {

	return &StorageServiceIncrementalBackupsGetParams{

		Context: ctx,
	}
}

// NewStorageServiceIncrementalBackupsGetParamsWithHTTPClient creates a new StorageServiceIncrementalBackupsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageServiceIncrementalBackupsGetParamsWithHTTPClient(client *http.Client) *StorageServiceIncrementalBackupsGetParams {

	return &StorageServiceIncrementalBackupsGetParams{
		HTTPClient: client,
	}
}

/*
StorageServiceIncrementalBackupsGetParams contains all the parameters to send to the API endpoint
for the storage service incremental backups get operation typically these are written to a http.Request
*/
type StorageServiceIncrementalBackupsGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage service incremental backups get params
func (o *StorageServiceIncrementalBackupsGetParams) WithTimeout(timeout time.Duration) *StorageServiceIncrementalBackupsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service incremental backups get params
func (o *StorageServiceIncrementalBackupsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage service incremental backups get params
func (o *StorageServiceIncrementalBackupsGetParams) WithContext(ctx context.Context) *StorageServiceIncrementalBackupsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service incremental backups get params
func (o *StorageServiceIncrementalBackupsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service incremental backups get params
func (o *StorageServiceIncrementalBackupsGetParams) WithHTTPClient(client *http.Client) *StorageServiceIncrementalBackupsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service incremental backups get params
func (o *StorageServiceIncrementalBackupsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServiceIncrementalBackupsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
