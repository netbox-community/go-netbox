// Copyright 2016 The go-netbox Authors.
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

// Interface represents an interface object.
type Interface struct {
	ID                 int              `json:"id"`
	Name               string           `json:"name"`
	FormFactor         string           `json:"form_factor"`
	MacAddress         string           `json:"mac_address"`
	MgmtOnly           bool             `json:"mgmt_only"`
	Description        string           `json:"description"`
	IsConnected        bool             `json:"is_connected"`
	ConnectedInterface *InterfaceDetail `json:"connected_interface"`
}

// InterfaceDetail represents an interface-detail object.
type InterfaceDetail struct {
	ID          int               `json:"id"`
	Device      *DeviceIdentifier `json:"device"`
	Name        string            `json:"name"`
	FormFactor  string            `json:"form_factor"`
	MacAddress  string            `json:"mac_address"`
	MgmtOnly    bool              `json:"mgmt_only"`
	Description string            `json:"description"`
	IsConnected bool              `json:"is_connected"`
}

// An InterfaceIdentifier is a reduced version of an Interface, returned as a
// nested object in some top-level objects.  It contains information which can
// be used in subsequent API calls to identify and retrieve a full Interface.
type InterfaceIdentifier struct {
	ID     int               `json:"id"`
	Device *DeviceIdentifier `json:"device"`
	Name   string            `json:"name"`
}
