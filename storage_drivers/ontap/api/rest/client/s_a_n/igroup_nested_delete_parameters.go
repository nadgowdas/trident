// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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
)

// NewIgroupNestedDeleteParams creates a new IgroupNestedDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIgroupNestedDeleteParams() *IgroupNestedDeleteParams {
	return &IgroupNestedDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIgroupNestedDeleteParamsWithTimeout creates a new IgroupNestedDeleteParams object
// with the ability to set a timeout on a request.
func NewIgroupNestedDeleteParamsWithTimeout(timeout time.Duration) *IgroupNestedDeleteParams {
	return &IgroupNestedDeleteParams{
		timeout: timeout,
	}
}

// NewIgroupNestedDeleteParamsWithContext creates a new IgroupNestedDeleteParams object
// with the ability to set a context for a request.
func NewIgroupNestedDeleteParamsWithContext(ctx context.Context) *IgroupNestedDeleteParams {
	return &IgroupNestedDeleteParams{
		Context: ctx,
	}
}

// NewIgroupNestedDeleteParamsWithHTTPClient creates a new IgroupNestedDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewIgroupNestedDeleteParamsWithHTTPClient(client *http.Client) *IgroupNestedDeleteParams {
	return &IgroupNestedDeleteParams{
		HTTPClient: client,
	}
}

/* IgroupNestedDeleteParams contains all the parameters to send to the API endpoint
   for the igroup nested delete operation.

   Typically these are written to a http.Request.
*/
type IgroupNestedDeleteParams struct {

	/* AllowDeleteWhileMapped.

	     Allows the deletion of a nested initiator group from of a mapped initiator group.<br/>
	Deleting a nested initiator group from a mapped initiator group means
	that the LUNs, to which the initiator group is mapped, are no longer
	available to the initiators nested below the initiator group being
	removed.  This might cause a disruption in the availability of
	data.<br/>
	<b>This parameter should be used with caution.</b>

	*/
	AllowDeleteWhileMappedQueryParameter *bool

	/* IgroupUUID.

	   The unique identifier of the parent initiator group.

	*/
	IgroupUUIDPathParameter string

	/* UUID.

	   The unique identifier of the nested initiator group.

	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the igroup nested delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IgroupNestedDeleteParams) WithDefaults() *IgroupNestedDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the igroup nested delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IgroupNestedDeleteParams) SetDefaults() {
	var (
		allowDeleteWhileMappedQueryParameterDefault = bool(false)
	)

	val := IgroupNestedDeleteParams{
		AllowDeleteWhileMappedQueryParameter: &allowDeleteWhileMappedQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the igroup nested delete params
func (o *IgroupNestedDeleteParams) WithTimeout(timeout time.Duration) *IgroupNestedDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the igroup nested delete params
func (o *IgroupNestedDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the igroup nested delete params
func (o *IgroupNestedDeleteParams) WithContext(ctx context.Context) *IgroupNestedDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the igroup nested delete params
func (o *IgroupNestedDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the igroup nested delete params
func (o *IgroupNestedDeleteParams) WithHTTPClient(client *http.Client) *IgroupNestedDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the igroup nested delete params
func (o *IgroupNestedDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllowDeleteWhileMappedQueryParameter adds the allowDeleteWhileMapped to the igroup nested delete params
func (o *IgroupNestedDeleteParams) WithAllowDeleteWhileMappedQueryParameter(allowDeleteWhileMapped *bool) *IgroupNestedDeleteParams {
	o.SetAllowDeleteWhileMappedQueryParameter(allowDeleteWhileMapped)
	return o
}

// SetAllowDeleteWhileMappedQueryParameter adds the allowDeleteWhileMapped to the igroup nested delete params
func (o *IgroupNestedDeleteParams) SetAllowDeleteWhileMappedQueryParameter(allowDeleteWhileMapped *bool) {
	o.AllowDeleteWhileMappedQueryParameter = allowDeleteWhileMapped
}

// WithIgroupUUIDPathParameter adds the igroupUUID to the igroup nested delete params
func (o *IgroupNestedDeleteParams) WithIgroupUUIDPathParameter(igroupUUID string) *IgroupNestedDeleteParams {
	o.SetIgroupUUIDPathParameter(igroupUUID)
	return o
}

// SetIgroupUUIDPathParameter adds the igroupUuid to the igroup nested delete params
func (o *IgroupNestedDeleteParams) SetIgroupUUIDPathParameter(igroupUUID string) {
	o.IgroupUUIDPathParameter = igroupUUID
}

// WithUUIDPathParameter adds the uuid to the igroup nested delete params
func (o *IgroupNestedDeleteParams) WithUUIDPathParameter(uuid string) *IgroupNestedDeleteParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the igroup nested delete params
func (o *IgroupNestedDeleteParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *IgroupNestedDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AllowDeleteWhileMappedQueryParameter != nil {

		// query param allow_delete_while_mapped
		var qrAllowDeleteWhileMapped bool

		if o.AllowDeleteWhileMappedQueryParameter != nil {
			qrAllowDeleteWhileMapped = *o.AllowDeleteWhileMappedQueryParameter
		}
		qAllowDeleteWhileMapped := swag.FormatBool(qrAllowDeleteWhileMapped)
		if qAllowDeleteWhileMapped != "" {

			if err := r.SetQueryParam("allow_delete_while_mapped", qAllowDeleteWhileMapped); err != nil {
				return err
			}
		}
	}

	// path param igroup.uuid
	if err := r.SetPathParam("igroup.uuid", o.IgroupUUIDPathParameter); err != nil {
		return err
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