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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewStorageServiceUpdateSnitchPostParams creates a new StorageServiceUpdateSnitchPostParams object
// with the default values initialized.
func NewStorageServiceUpdateSnitchPostParams() *StorageServiceUpdateSnitchPostParams {
	var ()
	return &StorageServiceUpdateSnitchPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageServiceUpdateSnitchPostParamsWithTimeout creates a new StorageServiceUpdateSnitchPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageServiceUpdateSnitchPostParamsWithTimeout(timeout time.Duration) *StorageServiceUpdateSnitchPostParams {
	var ()
	return &StorageServiceUpdateSnitchPostParams{

		timeout: timeout,
	}
}

// NewStorageServiceUpdateSnitchPostParamsWithContext creates a new StorageServiceUpdateSnitchPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageServiceUpdateSnitchPostParamsWithContext(ctx context.Context) *StorageServiceUpdateSnitchPostParams {
	var ()
	return &StorageServiceUpdateSnitchPostParams{

		Context: ctx,
	}
}

// NewStorageServiceUpdateSnitchPostParamsWithHTTPClient creates a new StorageServiceUpdateSnitchPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageServiceUpdateSnitchPostParamsWithHTTPClient(client *http.Client) *StorageServiceUpdateSnitchPostParams {
	var ()
	return &StorageServiceUpdateSnitchPostParams{
		HTTPClient: client,
	}
}

/*
StorageServiceUpdateSnitchPostParams contains all the parameters to send to the API endpoint
for the storage service update snitch post operation typically these are written to a http.Request
*/
type StorageServiceUpdateSnitchPostParams struct {

	/*Dynamic
	  When true dynamicsnitch is used

	*/
	Dynamic bool
	/*DynamicBadnessThreshold
	  Dynamic badness threshold, (default 0.0)

	*/
	DynamicBadnessThreshold *string
	/*DynamicResetInterval
	  integer, in ms (default 600,000)

	*/
	DynamicResetInterval *int32
	/*DynamicUpdateInterval
	  integer, in ms (default 100)

	*/
	DynamicUpdateInterval *int32
	/*EpSnitchClassName
	  The canonical path name for a class implementing IEndpointSnitch

	*/
	EpSnitchClassName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage service update snitch post params
func (o *StorageServiceUpdateSnitchPostParams) WithTimeout(timeout time.Duration) *StorageServiceUpdateSnitchPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service update snitch post params
func (o *StorageServiceUpdateSnitchPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage service update snitch post params
func (o *StorageServiceUpdateSnitchPostParams) WithContext(ctx context.Context) *StorageServiceUpdateSnitchPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service update snitch post params
func (o *StorageServiceUpdateSnitchPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service update snitch post params
func (o *StorageServiceUpdateSnitchPostParams) WithHTTPClient(client *http.Client) *StorageServiceUpdateSnitchPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service update snitch post params
func (o *StorageServiceUpdateSnitchPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDynamic adds the dynamic to the storage service update snitch post params
func (o *StorageServiceUpdateSnitchPostParams) WithDynamic(dynamic bool) *StorageServiceUpdateSnitchPostParams {
	o.SetDynamic(dynamic)
	return o
}

// SetDynamic adds the dynamic to the storage service update snitch post params
func (o *StorageServiceUpdateSnitchPostParams) SetDynamic(dynamic bool) {
	o.Dynamic = dynamic
}

// WithDynamicBadnessThreshold adds the dynamicBadnessThreshold to the storage service update snitch post params
func (o *StorageServiceUpdateSnitchPostParams) WithDynamicBadnessThreshold(dynamicBadnessThreshold *string) *StorageServiceUpdateSnitchPostParams {
	o.SetDynamicBadnessThreshold(dynamicBadnessThreshold)
	return o
}

// SetDynamicBadnessThreshold adds the dynamicBadnessThreshold to the storage service update snitch post params
func (o *StorageServiceUpdateSnitchPostParams) SetDynamicBadnessThreshold(dynamicBadnessThreshold *string) {
	o.DynamicBadnessThreshold = dynamicBadnessThreshold
}

// WithDynamicResetInterval adds the dynamicResetInterval to the storage service update snitch post params
func (o *StorageServiceUpdateSnitchPostParams) WithDynamicResetInterval(dynamicResetInterval *int32) *StorageServiceUpdateSnitchPostParams {
	o.SetDynamicResetInterval(dynamicResetInterval)
	return o
}

// SetDynamicResetInterval adds the dynamicResetInterval to the storage service update snitch post params
func (o *StorageServiceUpdateSnitchPostParams) SetDynamicResetInterval(dynamicResetInterval *int32) {
	o.DynamicResetInterval = dynamicResetInterval
}

// WithDynamicUpdateInterval adds the dynamicUpdateInterval to the storage service update snitch post params
func (o *StorageServiceUpdateSnitchPostParams) WithDynamicUpdateInterval(dynamicUpdateInterval *int32) *StorageServiceUpdateSnitchPostParams {
	o.SetDynamicUpdateInterval(dynamicUpdateInterval)
	return o
}

// SetDynamicUpdateInterval adds the dynamicUpdateInterval to the storage service update snitch post params
func (o *StorageServiceUpdateSnitchPostParams) SetDynamicUpdateInterval(dynamicUpdateInterval *int32) {
	o.DynamicUpdateInterval = dynamicUpdateInterval
}

// WithEpSnitchClassName adds the epSnitchClassName to the storage service update snitch post params
func (o *StorageServiceUpdateSnitchPostParams) WithEpSnitchClassName(epSnitchClassName string) *StorageServiceUpdateSnitchPostParams {
	o.SetEpSnitchClassName(epSnitchClassName)
	return o
}

// SetEpSnitchClassName adds the epSnitchClassName to the storage service update snitch post params
func (o *StorageServiceUpdateSnitchPostParams) SetEpSnitchClassName(epSnitchClassName string) {
	o.EpSnitchClassName = epSnitchClassName
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServiceUpdateSnitchPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param dynamic
	qrDynamic := o.Dynamic
	qDynamic := swag.FormatBool(qrDynamic)
	if qDynamic != "" {
		if err := r.SetQueryParam("dynamic", qDynamic); err != nil {
			return err
		}
	}

	if o.DynamicBadnessThreshold != nil {

		// query param dynamic_badness_threshold
		var qrDynamicBadnessThreshold string
		if o.DynamicBadnessThreshold != nil {
			qrDynamicBadnessThreshold = *o.DynamicBadnessThreshold
		}
		qDynamicBadnessThreshold := qrDynamicBadnessThreshold
		if qDynamicBadnessThreshold != "" {
			if err := r.SetQueryParam("dynamic_badness_threshold", qDynamicBadnessThreshold); err != nil {
				return err
			}
		}

	}

	if o.DynamicResetInterval != nil {

		// query param dynamic_reset_interval
		var qrDynamicResetInterval int32
		if o.DynamicResetInterval != nil {
			qrDynamicResetInterval = *o.DynamicResetInterval
		}
		qDynamicResetInterval := swag.FormatInt32(qrDynamicResetInterval)
		if qDynamicResetInterval != "" {
			if err := r.SetQueryParam("dynamic_reset_interval", qDynamicResetInterval); err != nil {
				return err
			}
		}

	}

	if o.DynamicUpdateInterval != nil {

		// query param dynamic_update_interval
		var qrDynamicUpdateInterval int32
		if o.DynamicUpdateInterval != nil {
			qrDynamicUpdateInterval = *o.DynamicUpdateInterval
		}
		qDynamicUpdateInterval := swag.FormatInt32(qrDynamicUpdateInterval)
		if qDynamicUpdateInterval != "" {
			if err := r.SetQueryParam("dynamic_update_interval", qDynamicUpdateInterval); err != nil {
				return err
			}
		}

	}

	// query param ep_snitch_class_name
	qrEpSnitchClassName := o.EpSnitchClassName
	qEpSnitchClassName := qrEpSnitchClassName
	if qEpSnitchClassName != "" {
		if err := r.SetQueryParam("ep_snitch_class_name", qEpSnitchClassName); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
