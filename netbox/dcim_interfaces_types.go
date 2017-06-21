// Copyright 2017 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package netbox

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// FormFactor represent network interface type (1000BASE-T/SFP)
type FormFactor simpleValueLabel

// NestedInterfaceConnection corresponds to the Netbox API's
// nested serializer, when an interface connection appears as a
// parent of another object.
type NestedInterfaceConnection struct {
	ID     int    `json:"id"`
	URL    string `json:"url"`
	Status bool   `json:"connection_status"`
}

// ShortNestedInterface corresponds to the Netbox API's
// nested serializer, when an interface appears as a
// parent of another interface (LAG Interface)
type ShortNestedInterface struct {
	ID   int    `json:"id"`
	URL  string `json:"url"`
	Name string `json:"name"`
}

// NestedInterface corresponds to the Netbox API's
// nested serializer, when an interface appears as a
// parent of another object.
type NestedInterface struct {
	ID          int          `json:"id"`
	URL         string       `json:"URL"`
	Device      NestedDevice `json:"device"`
	Name        string       `json:"name"`
	FormFactor  `json:"form_factor"`
	LAG         *ShortNestedInterface `json:"lag,omitempty"`
	MAC         string                `json:"mac_address,omitempty"`
	MGMT        bool                  `json:"mgmt_only"`
	Description string                `json:"description,omitempty"`
}

// Interface corresponde to the Netbox API's
// base serializer for Interfaces.
type Interface struct {
	ID                 int          `json:"id,omitempty"`
	Device             NestedDevice `json:"device"`
	Name               string       `json:"name"`
	FormFactor         `json:"form_factor"`
	LAG                *ShortNestedInterface      `json:"lag,omitempty"`
	MAC                string                     `json:"mac_address,omitempty"`
	MGMT               bool                       `json:"mgmt_only"`
	Description        string                     `json:"description,omitempty"`
	Connection         *NestedInterfaceConnection `json:"connection,omitempty"`
	ConnectedInterface *NestedInterface           `json:"connected_interface,omitempty"`
}

// A writableInterface corresponds to the Netbox API's
// writable serializer for an Interface. It is used transparently
// when Interfaces are serialized in to JSON.
type writableInterface struct {
	ID                 int    `json:"id,omitempty"`
	Device             int    `json:"device"`
	Name               string `json:"name"`
	FormFactor         int    `json:"form_factor"`
	LAG                *int   `json:"lag,omitempty"`
	MAC                string `json:"mac_address,omitempty"`
	MGMT               bool   `json:"mgmt_only"`
	Description        string `json:"description,omitempty"`
	Connection         *int   `json:"connection,omitempty"`
	ConnectedInterface *int   `json:"connected_interface,omitempty"`
}

// MarshalJSON marshals an Interface into JSON bytes,
// and is used by the standard json package.
func (i *Interface) MarshalJSON() ([]byte, error) {
	var lagID, connectionID, connectedInterfaceID *int

	if i.LAG != nil {
		lagID = &i.LAG.ID
	}

	if i.Connection != nil {
		connectionID = &i.Connection.ID
	}

	if i.ConnectedInterface != nil {
		connectedInterfaceID = &i.ConnectedInterface.ID
	}

	return json.Marshal(writableInterface{
		ID:                 i.ID,
		Device:             i.Device.ID,
		Name:               i.Name,
		FormFactor:         i.FormFactor.Value,
		LAG:                lagID,
		MAC:                i.MAC,
		MGMT:               i.MGMT,
		Description:        i.Description,
		Connection:         connectionID,
		ConnectedInterface: connectedInterfaceID,
	})

}

// ListInterfaceOptions is used as an argument for Client.DCUM.Interfaces.List.
// Integer fileds with *ID suffix are preferred over their string
// counterparts, and if both are set, only the *ID filed will be used
type ListInterfaceOptions struct {
	Name          string
	FormFactor    *int
	DeviceID      int
	Device        string
	InterfaceType string
	LagID         int
	MacAddress    string
}

// Values generates a url.Values map from the data in ListInterfaceOptions.
func (o *ListInterfaceOptions) Values() (url.Values, error) {
	if o == nil {
		return nil, nil
	}

	v := url.Values{}

	if o.Name != "" {
		v.Set("name", o.Name)
	}

	switch {
	case o.DeviceID != 0:
		v.Set("device_id", strconv.Itoa(o.DeviceID))
	case o.Device != "":
		v.Set("device", o.Device)
	}

	if o.FormFactor != nil {
		v.Set("form_factor", strconv.Itoa(*o.FormFactor))
	}

	if o.InterfaceType != "" {
		v.Set("type", o.InterfaceType)
	}

	if o.LagID != 0 {
		v.Set("lag_id", strconv.Itoa(o.LagID))
	}

	if o.MacAddress != "" {
		v.Set("mac_address", o.MacAddress)
	}

	return v, nil
}

//go:generate go run generate_functions.go -endpoint dcim -service interfaces -service-name InterfacesService -type-name Interface -update-type-name writableInterface
//go:generate go run generate_basic_tests.go -client-endpoint DCIM -client-service Interfaces -endpoint dcim -service interfaces -service-name InterfacesService -type-name Interface
