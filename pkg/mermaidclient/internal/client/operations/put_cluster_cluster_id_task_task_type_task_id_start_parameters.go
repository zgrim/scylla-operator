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

// NewPutClusterClusterIDTaskTaskTypeTaskIDStartParams creates a new PutClusterClusterIDTaskTaskTypeTaskIDStartParams object
// with the default values initialized.
func NewPutClusterClusterIDTaskTaskTypeTaskIDStartParams() *PutClusterClusterIDTaskTaskTypeTaskIDStartParams {
	var ()
	return &PutClusterClusterIDTaskTaskTypeTaskIDStartParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutClusterClusterIDTaskTaskTypeTaskIDStartParamsWithTimeout creates a new PutClusterClusterIDTaskTaskTypeTaskIDStartParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutClusterClusterIDTaskTaskTypeTaskIDStartParamsWithTimeout(timeout time.Duration) *PutClusterClusterIDTaskTaskTypeTaskIDStartParams {
	var ()
	return &PutClusterClusterIDTaskTaskTypeTaskIDStartParams{

		timeout: timeout,
	}
}

// NewPutClusterClusterIDTaskTaskTypeTaskIDStartParamsWithContext creates a new PutClusterClusterIDTaskTaskTypeTaskIDStartParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutClusterClusterIDTaskTaskTypeTaskIDStartParamsWithContext(ctx context.Context) *PutClusterClusterIDTaskTaskTypeTaskIDStartParams {
	var ()
	return &PutClusterClusterIDTaskTaskTypeTaskIDStartParams{

		Context: ctx,
	}
}

// NewPutClusterClusterIDTaskTaskTypeTaskIDStartParamsWithHTTPClient creates a new PutClusterClusterIDTaskTaskTypeTaskIDStartParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutClusterClusterIDTaskTaskTypeTaskIDStartParamsWithHTTPClient(client *http.Client) *PutClusterClusterIDTaskTaskTypeTaskIDStartParams {
	var ()
	return &PutClusterClusterIDTaskTaskTypeTaskIDStartParams{
		HTTPClient: client,
	}
}

/*
PutClusterClusterIDTaskTaskTypeTaskIDStartParams contains all the parameters to send to the API endpoint
for the put cluster cluster ID task task type task ID start operation typically these are written to a http.Request
*/
type PutClusterClusterIDTaskTaskTypeTaskIDStartParams struct {

	/*ClusterID*/
	ClusterID string
	/*Continue*/
	Continue bool
	/*TaskID*/
	TaskID string
	/*TaskType*/
	TaskType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put cluster cluster ID task task type task ID start params
func (o *PutClusterClusterIDTaskTaskTypeTaskIDStartParams) WithTimeout(timeout time.Duration) *PutClusterClusterIDTaskTaskTypeTaskIDStartParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put cluster cluster ID task task type task ID start params
func (o *PutClusterClusterIDTaskTaskTypeTaskIDStartParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put cluster cluster ID task task type task ID start params
func (o *PutClusterClusterIDTaskTaskTypeTaskIDStartParams) WithContext(ctx context.Context) *PutClusterClusterIDTaskTaskTypeTaskIDStartParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put cluster cluster ID task task type task ID start params
func (o *PutClusterClusterIDTaskTaskTypeTaskIDStartParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put cluster cluster ID task task type task ID start params
func (o *PutClusterClusterIDTaskTaskTypeTaskIDStartParams) WithHTTPClient(client *http.Client) *PutClusterClusterIDTaskTaskTypeTaskIDStartParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put cluster cluster ID task task type task ID start params
func (o *PutClusterClusterIDTaskTaskTypeTaskIDStartParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the put cluster cluster ID task task type task ID start params
func (o *PutClusterClusterIDTaskTaskTypeTaskIDStartParams) WithClusterID(clusterID string) *PutClusterClusterIDTaskTaskTypeTaskIDStartParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the put cluster cluster ID task task type task ID start params
func (o *PutClusterClusterIDTaskTaskTypeTaskIDStartParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithContinue adds the continueVar to the put cluster cluster ID task task type task ID start params
func (o *PutClusterClusterIDTaskTaskTypeTaskIDStartParams) WithContinue(continueVar bool) *PutClusterClusterIDTaskTaskTypeTaskIDStartParams {
	o.SetContinue(continueVar)
	return o
}

// SetContinue adds the continue to the put cluster cluster ID task task type task ID start params
func (o *PutClusterClusterIDTaskTaskTypeTaskIDStartParams) SetContinue(continueVar bool) {
	o.Continue = continueVar
}

// WithTaskID adds the taskID to the put cluster cluster ID task task type task ID start params
func (o *PutClusterClusterIDTaskTaskTypeTaskIDStartParams) WithTaskID(taskID string) *PutClusterClusterIDTaskTaskTypeTaskIDStartParams {
	o.SetTaskID(taskID)
	return o
}

// SetTaskID adds the taskId to the put cluster cluster ID task task type task ID start params
func (o *PutClusterClusterIDTaskTaskTypeTaskIDStartParams) SetTaskID(taskID string) {
	o.TaskID = taskID
}

// WithTaskType adds the taskType to the put cluster cluster ID task task type task ID start params
func (o *PutClusterClusterIDTaskTaskTypeTaskIDStartParams) WithTaskType(taskType string) *PutClusterClusterIDTaskTaskTypeTaskIDStartParams {
	o.SetTaskType(taskType)
	return o
}

// SetTaskType adds the taskType to the put cluster cluster ID task task type task ID start params
func (o *PutClusterClusterIDTaskTaskTypeTaskIDStartParams) SetTaskType(taskType string) {
	o.TaskType = taskType
}

// WriteToRequest writes these params to a swagger request
func (o *PutClusterClusterIDTaskTaskTypeTaskIDStartParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// query param continue
	qrContinue := o.Continue
	qContinue := swag.FormatBool(qrContinue)
	if qContinue != "" {
		if err := r.SetQueryParam("continue", qContinue); err != nil {
			return err
		}
	}

	// path param task_id
	if err := r.SetPathParam("task_id", o.TaskID); err != nil {
		return err
	}

	// path param task_type
	if err := r.SetPathParam("task_type", o.TaskType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
