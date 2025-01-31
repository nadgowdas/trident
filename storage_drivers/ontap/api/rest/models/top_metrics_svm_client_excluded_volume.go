// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TopMetricsSvmClientExcludedVolume List of volumes and their details as to why they are not included in the SVM activity tracking REST API.
//
// swagger:model top_metrics_svm_client_excluded_volume
type TopMetricsSvmClientExcludedVolume struct {

	// reason
	Reason *TopMetricsSvmClientExcludedVolumeReason `json:"reason,omitempty"`

	// volume
	Volume *TopMetricsSvmClientExcludedVolumeVolume `json:"volume,omitempty"`
}

// Validate validates this top metrics svm client excluded volume
func (m *TopMetricsSvmClientExcludedVolume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolume(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmClientExcludedVolume) validateReason(formats strfmt.Registry) error {
	if swag.IsZero(m.Reason) { // not required
		return nil
	}

	if m.Reason != nil {
		if err := m.Reason.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reason")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsSvmClientExcludedVolume) validateVolume(formats strfmt.Registry) error {
	if swag.IsZero(m.Volume) { // not required
		return nil
	}

	if m.Volume != nil {
		if err := m.Volume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this top metrics svm client excluded volume based on the context it is used
func (m *TopMetricsSvmClientExcludedVolume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReason(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolume(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmClientExcludedVolume) contextValidateReason(ctx context.Context, formats strfmt.Registry) error {

	if m.Reason != nil {
		if err := m.Reason.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reason")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsSvmClientExcludedVolume) contextValidateVolume(ctx context.Context, formats strfmt.Registry) error {

	if m.Volume != nil {
		if err := m.Volume.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopMetricsSvmClientExcludedVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsSvmClientExcludedVolume) UnmarshalBinary(b []byte) error {
	var res TopMetricsSvmClientExcludedVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsSvmClientExcludedVolumeReason top metrics svm client excluded volume reason
//
// swagger:model TopMetricsSvmClientExcludedVolumeReason
type TopMetricsSvmClientExcludedVolumeReason struct {

	// Warning code indicating why the volume is not included in the SVM activity tracking REST API.
	// Example: 111411207
	// Read Only: true
	Code string `json:"code,omitempty"`

	// Details why the volume is not included in the SVM activity tracking REST API.
	// Example: The volume is offline.
	// Read Only: true
	Message string `json:"message,omitempty"`
}

// Validate validates this top metrics svm client excluded volume reason
func (m *TopMetricsSvmClientExcludedVolumeReason) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this top metrics svm client excluded volume reason based on the context it is used
func (m *TopMetricsSvmClientExcludedVolumeReason) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmClientExcludedVolumeReason) contextValidateCode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "reason"+"."+"code", "body", string(m.Code)); err != nil {
		return err
	}

	return nil
}

func (m *TopMetricsSvmClientExcludedVolumeReason) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "reason"+"."+"message", "body", string(m.Message)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopMetricsSvmClientExcludedVolumeReason) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsSvmClientExcludedVolumeReason) UnmarshalBinary(b []byte) error {
	var res TopMetricsSvmClientExcludedVolumeReason
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsSvmClientExcludedVolumeVolume top metrics svm client excluded volume volume
//
// swagger:model TopMetricsSvmClientExcludedVolumeVolume
type TopMetricsSvmClientExcludedVolumeVolume struct {

	// links
	Links *TopMetricsSvmClientExcludedVolumeVolumeLinks `json:"_links,omitempty"`

	// The name of the volume.
	// Example: volume1
	Name string `json:"name,omitempty"`

	// Unique identifier for the volume. This corresponds to the instance-uuid that is exposed in the CLI and ONTAPI. It does not change due to a volume move.
	// Example: 028baa66-41bd-11e9-81d5-00a0986138f7
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this top metrics svm client excluded volume volume
func (m *TopMetricsSvmClientExcludedVolumeVolume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmClientExcludedVolumeVolume) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this top metrics svm client excluded volume volume based on the context it is used
func (m *TopMetricsSvmClientExcludedVolumeVolume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmClientExcludedVolumeVolume) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopMetricsSvmClientExcludedVolumeVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsSvmClientExcludedVolumeVolume) UnmarshalBinary(b []byte) error {
	var res TopMetricsSvmClientExcludedVolumeVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsSvmClientExcludedVolumeVolumeLinks top metrics svm client excluded volume volume links
//
// swagger:model TopMetricsSvmClientExcludedVolumeVolumeLinks
type TopMetricsSvmClientExcludedVolumeVolumeLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this top metrics svm client excluded volume volume links
func (m *TopMetricsSvmClientExcludedVolumeVolumeLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmClientExcludedVolumeVolumeLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this top metrics svm client excluded volume volume links based on the context it is used
func (m *TopMetricsSvmClientExcludedVolumeVolumeLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmClientExcludedVolumeVolumeLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopMetricsSvmClientExcludedVolumeVolumeLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsSvmClientExcludedVolumeVolumeLinks) UnmarshalBinary(b []byte) error {
	var res TopMetricsSvmClientExcludedVolumeVolumeLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
