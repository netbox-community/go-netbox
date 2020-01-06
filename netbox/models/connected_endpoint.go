package models

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ConnectedEndpoint Connection endpoint
// swagger:model ConnectedEndpoint
type ConnectedEndpoint struct {
	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Url
	// Read Only: true
	URL strfmt.URI `json:"url,omitempty"`

	// device
	// Required: true
	Device *NestedDevice `json:"device"`

	// Name
	// Max Length: 64
	Name *string `json:"name,omitempty"`

	Cable int64 `json:"cable,omitempty"`

	ConnectionStatus *DeviceInterfaceConnectionStatus `json:"connection_status,omitempty"`

	Circuit *Circuit `json:"circuit,omitempty"`
}

// Validate validates this interface
func (m *ConnectedEndpoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDevice(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateConnectionStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConnectedEndpoint) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	if m.Device != nil {

		if err := m.Device.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *ConnectedEndpoint) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 64); err != nil {
		return err
	}

	return nil
}

func (m *ConnectedEndpoint) validateConnectionStatus(formats strfmt.Registry) error {

	if err := validate.Required("connection_status", "body", m.ConnectionStatus); err != nil {
		return err
	}

	if m.ConnectionStatus != nil {

		if err := m.ConnectionStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection_status")
			}
			return err
		}
	}

	return nil
}
