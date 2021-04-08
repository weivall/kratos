// Code generated by go-swagger; DO NOT EDIT.

package public

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

// NewGetSelfServiceRegistrationFlowParams creates a new GetSelfServiceRegistrationFlowParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSelfServiceRegistrationFlowParams() *GetSelfServiceRegistrationFlowParams {
	return &GetSelfServiceRegistrationFlowParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSelfServiceRegistrationFlowParamsWithTimeout creates a new GetSelfServiceRegistrationFlowParams object
// with the ability to set a timeout on a request.
func NewGetSelfServiceRegistrationFlowParamsWithTimeout(timeout time.Duration) *GetSelfServiceRegistrationFlowParams {
	return &GetSelfServiceRegistrationFlowParams{
		timeout: timeout,
	}
}

// NewGetSelfServiceRegistrationFlowParamsWithContext creates a new GetSelfServiceRegistrationFlowParams object
// with the ability to set a context for a request.
func NewGetSelfServiceRegistrationFlowParamsWithContext(ctx context.Context) *GetSelfServiceRegistrationFlowParams {
	return &GetSelfServiceRegistrationFlowParams{
		Context: ctx,
	}
}

// NewGetSelfServiceRegistrationFlowParamsWithHTTPClient creates a new GetSelfServiceRegistrationFlowParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSelfServiceRegistrationFlowParamsWithHTTPClient(client *http.Client) *GetSelfServiceRegistrationFlowParams {
	return &GetSelfServiceRegistrationFlowParams{
		HTTPClient: client,
	}
}

/* GetSelfServiceRegistrationFlowParams contains all the parameters to send to the API endpoint
   for the get self service registration flow operation.

   Typically these are written to a http.Request.
*/
type GetSelfServiceRegistrationFlowParams struct {

	/* ID.

	     The Registration Flow ID

	The value for this parameter comes from `flow` URL Query parameter sent to your
	application (e.g. `/registration?flow=abcde`).
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get self service registration flow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSelfServiceRegistrationFlowParams) WithDefaults() *GetSelfServiceRegistrationFlowParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get self service registration flow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSelfServiceRegistrationFlowParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get self service registration flow params
func (o *GetSelfServiceRegistrationFlowParams) WithTimeout(timeout time.Duration) *GetSelfServiceRegistrationFlowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get self service registration flow params
func (o *GetSelfServiceRegistrationFlowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get self service registration flow params
func (o *GetSelfServiceRegistrationFlowParams) WithContext(ctx context.Context) *GetSelfServiceRegistrationFlowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get self service registration flow params
func (o *GetSelfServiceRegistrationFlowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get self service registration flow params
func (o *GetSelfServiceRegistrationFlowParams) WithHTTPClient(client *http.Client) *GetSelfServiceRegistrationFlowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get self service registration flow params
func (o *GetSelfServiceRegistrationFlowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get self service registration flow params
func (o *GetSelfServiceRegistrationFlowParams) WithID(id string) *GetSelfServiceRegistrationFlowParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get self service registration flow params
func (o *GetSelfServiceRegistrationFlowParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetSelfServiceRegistrationFlowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param id
	qrID := o.ID
	qID := qrID
	if qID != "" {

		if err := r.SetQueryParam("id", qID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}