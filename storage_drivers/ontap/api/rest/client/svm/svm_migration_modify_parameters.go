// Code generated by go-swagger; DO NOT EDIT.

package svm

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
	"github.com/go-openapi/swag"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewSvmMigrationModifyParams creates a new SvmMigrationModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSvmMigrationModifyParams() *SvmMigrationModifyParams {
	return &SvmMigrationModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSvmMigrationModifyParamsWithTimeout creates a new SvmMigrationModifyParams object
// with the ability to set a timeout on a request.
func NewSvmMigrationModifyParamsWithTimeout(timeout time.Duration) *SvmMigrationModifyParams {
	return &SvmMigrationModifyParams{
		timeout: timeout,
	}
}

// NewSvmMigrationModifyParamsWithContext creates a new SvmMigrationModifyParams object
// with the ability to set a context for a request.
func NewSvmMigrationModifyParamsWithContext(ctx context.Context) *SvmMigrationModifyParams {
	return &SvmMigrationModifyParams{
		Context: ctx,
	}
}

// NewSvmMigrationModifyParamsWithHTTPClient creates a new SvmMigrationModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewSvmMigrationModifyParamsWithHTTPClient(client *http.Client) *SvmMigrationModifyParams {
	return &SvmMigrationModifyParams{
		HTTPClient: client,
	}
}

/*
SvmMigrationModifyParams contains all the parameters to send to the API endpoint

	for the svm migration modify operation.

	Typically these are written to a http.Request.
*/
type SvmMigrationModifyParams struct {

	/* Action.

	     The pause action pauses the SVM migration. This action stops data transfer and configuration replication. This operation must be performed on the destination cluster.
	The resume action resumes an SVM migration from a paused or failed state. This operation must be performed on the destination cluster.
	The cutover action triggers the cutover of an SVM from the source cluster to the destination cluster.
	The source_clean up action performs a clean up of the SVM on the source cluster.

	*/
	ActionQueryParameter *string

	/* AutoCutover.

	   Optional property that when set to true automatically performs cutover when the migration state reaches "ready for cutover".

	   Default: true
	*/
	AutoCutoverQueryParameter *bool

	/* AutoSourceCleanup.

	   Optional property that when set to true automatically cleans up the SVM on the source cluster after the migration cutover.

	   Default: true
	*/
	AutoSourceCleanupQueryParameter *bool

	/* Info.

	   Info specification
	*/
	Info *models.SvmMigration

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeoutQueryParameter *int64

	/* UUID.

	   SVM migration UUID
	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the svm migration modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmMigrationModifyParams) WithDefaults() *SvmMigrationModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the svm migration modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmMigrationModifyParams) SetDefaults() {
	var (
		autoCutoverQueryParameterDefault = bool(true)

		autoSourceCleanupQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(0)
	)

	val := SvmMigrationModifyParams{
		AutoCutoverQueryParameter:       &autoCutoverQueryParameterDefault,
		AutoSourceCleanupQueryParameter: &autoSourceCleanupQueryParameterDefault,
		ReturnTimeoutQueryParameter:     &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the svm migration modify params
func (o *SvmMigrationModifyParams) WithTimeout(timeout time.Duration) *SvmMigrationModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the svm migration modify params
func (o *SvmMigrationModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the svm migration modify params
func (o *SvmMigrationModifyParams) WithContext(ctx context.Context) *SvmMigrationModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the svm migration modify params
func (o *SvmMigrationModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the svm migration modify params
func (o *SvmMigrationModifyParams) WithHTTPClient(client *http.Client) *SvmMigrationModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the svm migration modify params
func (o *SvmMigrationModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionQueryParameter adds the action to the svm migration modify params
func (o *SvmMigrationModifyParams) WithActionQueryParameter(action *string) *SvmMigrationModifyParams {
	o.SetActionQueryParameter(action)
	return o
}

// SetActionQueryParameter adds the action to the svm migration modify params
func (o *SvmMigrationModifyParams) SetActionQueryParameter(action *string) {
	o.ActionQueryParameter = action
}

// WithAutoCutoverQueryParameter adds the autoCutover to the svm migration modify params
func (o *SvmMigrationModifyParams) WithAutoCutoverQueryParameter(autoCutover *bool) *SvmMigrationModifyParams {
	o.SetAutoCutoverQueryParameter(autoCutover)
	return o
}

// SetAutoCutoverQueryParameter adds the autoCutover to the svm migration modify params
func (o *SvmMigrationModifyParams) SetAutoCutoverQueryParameter(autoCutover *bool) {
	o.AutoCutoverQueryParameter = autoCutover
}

// WithAutoSourceCleanupQueryParameter adds the autoSourceCleanup to the svm migration modify params
func (o *SvmMigrationModifyParams) WithAutoSourceCleanupQueryParameter(autoSourceCleanup *bool) *SvmMigrationModifyParams {
	o.SetAutoSourceCleanupQueryParameter(autoSourceCleanup)
	return o
}

// SetAutoSourceCleanupQueryParameter adds the autoSourceCleanup to the svm migration modify params
func (o *SvmMigrationModifyParams) SetAutoSourceCleanupQueryParameter(autoSourceCleanup *bool) {
	o.AutoSourceCleanupQueryParameter = autoSourceCleanup
}

// WithInfo adds the info to the svm migration modify params
func (o *SvmMigrationModifyParams) WithInfo(info *models.SvmMigration) *SvmMigrationModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the svm migration modify params
func (o *SvmMigrationModifyParams) SetInfo(info *models.SvmMigration) {
	o.Info = info
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the svm migration modify params
func (o *SvmMigrationModifyParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *SvmMigrationModifyParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the svm migration modify params
func (o *SvmMigrationModifyParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithUUIDPathParameter adds the uuid to the svm migration modify params
func (o *SvmMigrationModifyParams) WithUUIDPathParameter(uuid string) *SvmMigrationModifyParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the svm migration modify params
func (o *SvmMigrationModifyParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *SvmMigrationModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ActionQueryParameter != nil {

		// query param action
		var qrAction string

		if o.ActionQueryParameter != nil {
			qrAction = *o.ActionQueryParameter
		}
		qAction := qrAction
		if qAction != "" {

			if err := r.SetQueryParam("action", qAction); err != nil {
				return err
			}
		}
	}

	if o.AutoCutoverQueryParameter != nil {

		// query param auto_cutover
		var qrAutoCutover bool

		if o.AutoCutoverQueryParameter != nil {
			qrAutoCutover = *o.AutoCutoverQueryParameter
		}
		qAutoCutover := swag.FormatBool(qrAutoCutover)
		if qAutoCutover != "" {

			if err := r.SetQueryParam("auto_cutover", qAutoCutover); err != nil {
				return err
			}
		}
	}

	if o.AutoSourceCleanupQueryParameter != nil {

		// query param auto_source_cleanup
		var qrAutoSourceCleanup bool

		if o.AutoSourceCleanupQueryParameter != nil {
			qrAutoSourceCleanup = *o.AutoSourceCleanupQueryParameter
		}
		qAutoSourceCleanup := swag.FormatBool(qrAutoSourceCleanup)
		if qAutoSourceCleanup != "" {

			if err := r.SetQueryParam("auto_source_cleanup", qAutoSourceCleanup); err != nil {
				return err
			}
		}
	}
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if o.ReturnTimeoutQueryParameter != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeoutQueryParameter != nil {
			qrReturnTimeout = *o.ReturnTimeoutQueryParameter
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
