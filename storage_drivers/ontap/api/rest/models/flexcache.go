// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Flexcache Defines the cache endpoint of FlexCache.
//
// swagger:model flexcache
type Flexcache struct {

	// links
	Links *FlexcacheLinks `json:"_links,omitempty"`

	// aggregates
	Aggregates []*FlexcacheAggregatesItems0 `json:"aggregates,omitempty"`

	// Number of FlexCache constituents per aggregate when the 'aggregates' field is mentioned.
	ConstituentsPerAggregate *int64 `json:"constituents_per_aggregate,omitempty"`

	// If set to true, a DR cache is created.
	DrCache *bool `json:"dr_cache,omitempty"`

	// Specifies whether or not a FlexCache volume has global file locking mode enabled. Global file locking mode is a mode where protocol read locking semantics are enforced across all FlexCaches and origins of a FlexCache volume. When global file locking mode is enabled, the "is_disconnected_mode_off_for_locks" flag is always set to "true".
	GlobalFileLockingEnabled *bool `json:"global_file_locking_enabled,omitempty"`

	// guarantee
	Guarantee *FlexcacheGuarantee `json:"guarantee,omitempty"`

	// FlexCache name
	// Example: vol1
	// Max Length: 203
	// Min Length: 1
	Name string `json:"name,omitempty"`

	// origins
	Origins []*FlexcacheRelationship `json:"origins,omitempty"`

	// The fully-qualified path in the owning SVM's namespace at which the FlexCache is mounted. The path is case insensitive and must be unique within a SVM's namespace. Path must begin with '/' and must not end with '/'. Only one FlexCache be mounted at any given junction path.
	// Example: /user/my_fc
	Path string `json:"path,omitempty"`

	// prepopulate
	Prepopulate *FlexcachePrepopulateType `json:"prepopulate,omitempty"`

	// Physical size of the FlexCache. The recommended size for a FlexCache is 10% of the origin volume. The minimum FlexCache constituent size is 1GB.
	Size int64 `json:"size,omitempty"`

	// svm
	Svm *FlexcacheSvm `json:"svm,omitempty"`

	// Specifies whether or not a Fabricpool-enabled aggregate can be used in FlexCache creation. The use_tiered_aggregate is only used when auto-provisioning a FlexCache volume.
	UseTieredAggregate *bool `json:"use_tiered_aggregate,omitempty"`

	// FlexCache UUID. Unique identifier for the FlexCache.
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	// Read Only: true
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this flexcache
func (m *Flexcache) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAggregates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGuarantee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrigins(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrepopulate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Flexcache) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *Flexcache) validateAggregates(formats strfmt.Registry) error {
	if swag.IsZero(m.Aggregates) { // not required
		return nil
	}

	for i := 0; i < len(m.Aggregates); i++ {
		if swag.IsZero(m.Aggregates[i]) { // not required
			continue
		}

		if m.Aggregates[i] != nil {
			if err := m.Aggregates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("aggregates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Flexcache) validateGuarantee(formats strfmt.Registry) error {
	if swag.IsZero(m.Guarantee) { // not required
		return nil
	}

	if m.Guarantee != nil {
		if err := m.Guarantee.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("guarantee")
			}
			return err
		}
	}

	return nil
}

func (m *Flexcache) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", m.Name, 203); err != nil {
		return err
	}

	return nil
}

func (m *Flexcache) validateOrigins(formats strfmt.Registry) error {
	if swag.IsZero(m.Origins) { // not required
		return nil
	}

	for i := 0; i < len(m.Origins); i++ {
		if swag.IsZero(m.Origins[i]) { // not required
			continue
		}

		if m.Origins[i] != nil {
			if err := m.Origins[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("origins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Flexcache) validatePrepopulate(formats strfmt.Registry) error {
	if swag.IsZero(m.Prepopulate) { // not required
		return nil
	}

	if m.Prepopulate != nil {
		if err := m.Prepopulate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("prepopulate")
			}
			return err
		}
	}

	return nil
}

func (m *Flexcache) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(m.Svm) { // not required
		return nil
	}

	if m.Svm != nil {
		if err := m.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this flexcache based on the context it is used
func (m *Flexcache) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAggregates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGuarantee(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOrigins(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrepopulate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Flexcache) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *Flexcache) contextValidateAggregates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Aggregates); i++ {

		if m.Aggregates[i] != nil {
			if err := m.Aggregates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("aggregates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Flexcache) contextValidateGuarantee(ctx context.Context, formats strfmt.Registry) error {

	if m.Guarantee != nil {
		if err := m.Guarantee.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("guarantee")
			}
			return err
		}
	}

	return nil
}

func (m *Flexcache) contextValidateOrigins(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Origins); i++ {

		if m.Origins[i] != nil {
			if err := m.Origins[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("origins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Flexcache) contextValidatePrepopulate(ctx context.Context, formats strfmt.Registry) error {

	if m.Prepopulate != nil {
		if err := m.Prepopulate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("prepopulate")
			}
			return err
		}
	}

	return nil
}

func (m *Flexcache) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if m.Svm != nil {
		if err := m.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

func (m *Flexcache) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", string(m.UUID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Flexcache) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Flexcache) UnmarshalBinary(b []byte) error {
	var res Flexcache
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FlexcacheAggregatesItems0 Aggregate
//
// swagger:model FlexcacheAggregatesItems0
type FlexcacheAggregatesItems0 struct {

	// links
	Links *FlexcacheAggregatesItems0Links `json:"_links,omitempty"`

	// name
	// Example: aggr1
	Name string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this flexcache aggregates items0
func (m *FlexcacheAggregatesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FlexcacheAggregatesItems0) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this flexcache aggregates items0 based on the context it is used
func (m *FlexcacheAggregatesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FlexcacheAggregatesItems0) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FlexcacheAggregatesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlexcacheAggregatesItems0) UnmarshalBinary(b []byte) error {
	var res FlexcacheAggregatesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FlexcacheAggregatesItems0Links flexcache aggregates items0 links
//
// swagger:model FlexcacheAggregatesItems0Links
type FlexcacheAggregatesItems0Links struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this flexcache aggregates items0 links
func (m *FlexcacheAggregatesItems0Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FlexcacheAggregatesItems0Links) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this flexcache aggregates items0 links based on the context it is used
func (m *FlexcacheAggregatesItems0Links) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FlexcacheAggregatesItems0Links) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FlexcacheAggregatesItems0Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlexcacheAggregatesItems0Links) UnmarshalBinary(b []byte) error {
	var res FlexcacheAggregatesItems0Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FlexcacheGuarantee flexcache guarantee
//
// swagger:model FlexcacheGuarantee
type FlexcacheGuarantee struct {

	// The type of space guarantee of this volume in the aggregate.
	// Enum: [volume none]
	Type *string `json:"type,omitempty"`
}

// Validate validates this flexcache guarantee
func (m *FlexcacheGuarantee) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var flexcacheGuaranteeTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["volume","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		flexcacheGuaranteeTypeTypePropEnum = append(flexcacheGuaranteeTypeTypePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// FlexcacheGuarantee
	// FlexcacheGuarantee
	// type
	// Type
	// volume
	// END DEBUGGING
	// FlexcacheGuaranteeTypeVolume captures enum value "volume"
	FlexcacheGuaranteeTypeVolume string = "volume"

	// BEGIN DEBUGGING
	// FlexcacheGuarantee
	// FlexcacheGuarantee
	// type
	// Type
	// none
	// END DEBUGGING
	// FlexcacheGuaranteeTypeNone captures enum value "none"
	FlexcacheGuaranteeTypeNone string = "none"
)

// prop value enum
func (m *FlexcacheGuarantee) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, flexcacheGuaranteeTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FlexcacheGuarantee) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("guarantee"+"."+"type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this flexcache guarantee based on context it is used
func (m *FlexcacheGuarantee) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FlexcacheGuarantee) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlexcacheGuarantee) UnmarshalBinary(b []byte) error {
	var res FlexcacheGuarantee
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FlexcacheLinks flexcache links
//
// swagger:model FlexcacheLinks
type FlexcacheLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this flexcache links
func (m *FlexcacheLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FlexcacheLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this flexcache links based on the context it is used
func (m *FlexcacheLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FlexcacheLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FlexcacheLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlexcacheLinks) UnmarshalBinary(b []byte) error {
	var res FlexcacheLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FlexcachePrepopulateType FlexCache prepopulate
//
// swagger:model FlexcachePrepopulateType
type FlexcachePrepopulateType struct {

	// dir paths
	DirPaths []string `json:"dir_paths,omitempty"`

	// exclude dir paths
	ExcludeDirPaths []string `json:"exclude_dir_paths,omitempty"`

	// Specifies whether or not the prepopulate action should search through the `dir_paths` recursively. If not set, the default value _true_ is used.
	Recurse *bool `json:"recurse,omitempty"`
}

// Validate validates this flexcache prepopulate type
func (m *FlexcachePrepopulateType) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this flexcache prepopulate type based on context it is used
func (m *FlexcachePrepopulateType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FlexcachePrepopulateType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlexcachePrepopulateType) UnmarshalBinary(b []byte) error {
	var res FlexcachePrepopulateType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FlexcacheSvm FlexCache SVM
//
// swagger:model FlexcacheSvm
type FlexcacheSvm struct {

	// links
	Links *FlexcacheSvmLinks `json:"_links,omitempty"`

	// The name of the SVM.
	//
	// Example: svm1
	Name string `json:"name,omitempty"`

	// The unique identifier of the SVM.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this flexcache svm
func (m *FlexcacheSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FlexcacheSvm) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this flexcache svm based on the context it is used
func (m *FlexcacheSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FlexcacheSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FlexcacheSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlexcacheSvm) UnmarshalBinary(b []byte) error {
	var res FlexcacheSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FlexcacheSvmLinks flexcache svm links
//
// swagger:model FlexcacheSvmLinks
type FlexcacheSvmLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this flexcache svm links
func (m *FlexcacheSvmLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FlexcacheSvmLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this flexcache svm links based on the context it is used
func (m *FlexcacheSvmLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FlexcacheSvmLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FlexcacheSvmLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlexcacheSvmLinks) UnmarshalBinary(b []byte) error {
	var res FlexcacheSvmLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
