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
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/swagger/gen/agent/models"
)

// NewJobProgressParams creates a new JobProgressParams object
// with the default values initialized.
func NewJobProgressParams() *JobProgressParams {
	var ()
	return &JobProgressParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJobProgressParamsWithTimeout creates a new JobProgressParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJobProgressParamsWithTimeout(timeout time.Duration) *JobProgressParams {
	var ()
	return &JobProgressParams{

		timeout: timeout,
	}
}

// NewJobProgressParamsWithContext creates a new JobProgressParams object
// with the default values initialized, and the ability to set a context for a request
func NewJobProgressParamsWithContext(ctx context.Context) *JobProgressParams {
	var ()
	return &JobProgressParams{

		Context: ctx,
	}
}

// NewJobProgressParamsWithHTTPClient creates a new JobProgressParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewJobProgressParamsWithHTTPClient(client *http.Client) *JobProgressParams {
	var ()
	return &JobProgressParams{
		HTTPClient: client,
	}
}

/*JobProgressParams contains all the parameters to send to the API endpoint
for the job progress operation typically these are written to a http.Request
*/
type JobProgressParams struct {

	/*Jobinfo
	  Job info params with id and long polling

	*/
	Jobinfo *models.JobInfoParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the job progress params
func (o *JobProgressParams) WithTimeout(timeout time.Duration) *JobProgressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the job progress params
func (o *JobProgressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the job progress params
func (o *JobProgressParams) WithContext(ctx context.Context) *JobProgressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the job progress params
func (o *JobProgressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the job progress params
func (o *JobProgressParams) WithHTTPClient(client *http.Client) *JobProgressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the job progress params
func (o *JobProgressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJobinfo adds the jobinfo to the job progress params
func (o *JobProgressParams) WithJobinfo(jobinfo *models.JobInfoParams) *JobProgressParams {
	o.SetJobinfo(jobinfo)
	return o
}

// SetJobinfo adds the jobinfo to the job progress params
func (o *JobProgressParams) SetJobinfo(jobinfo *models.JobInfoParams) {
	o.Jobinfo = jobinfo
}

// WriteToRequest writes these params to a swagger request
func (o *JobProgressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Jobinfo != nil {
		if err := r.SetBodyParam(o.Jobinfo); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
