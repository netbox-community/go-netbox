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

import (
	"errors"
	"net/http"
	"net/url"
)

// RelatedConnection represents components that have a releated peer-device and
// peer-interface
type RelatedConnection struct {
	Device       *Device                    `json:"device"`
	ConsolePorts []*RCConsolePortIdentifier `json:"console-ports"`
	Interfaces   []*RCInterfaceIdentifier   `json:"interfaces"`
	PowerPorts   []*RCPowerPortIdentifier   `json:"power-ports"`
}

// RCConsolePortIdentifier represents a reduced version of a console port
type RCConsolePortIdentifier struct {
	ConsoleServer string `json:"console-server"`
	Name          string `json:"name"`
	Port          string `json:"port"`
}

// RCInterfaceIdentifier represents a reduced version of a device interface
type RCInterfaceIdentifier struct {
	Device    string `json:"device"`
	Interface string `json:"interface"`
	Name      string `json:"name"`
}

// RCPowerPortIdentifier represents a reduced version of a single power port
type RCPowerPortIdentifier struct {
	PDU    string `json:"pdu"`
	Name   string `json:"name"`
	Outlet string `json:"outlet"`
}

// GetRelatedConnections retreives a RelatedConnections objects from Netbox
func (s *DCIMService) GetRelatedConnections(
	options *RelatedConnectionsOptions,
) (*RelatedConnection, error) {
	req, err := s.c.newRequest(
		http.MethodGet,
		"/api/dcim/related-connections/",
		options,
	)
	if err != nil {
		return nil, err
	}

	rc := new(RelatedConnection)
	err = s.c.do(req, rc)
	if err != nil {
		return nil, err
	}
	return rc, nil
}

// RelatedConnectionsOptions is used as an argument for
// Client.DCIM.GetRelatedConnections
type RelatedConnectionsOptions struct {
	PeerDevice    string
	PeerInterface string
}

func (o *RelatedConnectionsOptions) values() (url.Values, error) {
	var err = errors.New(
		"must provide non-zero values for both peer-device and peer-interface",
	)
	if o == nil {
		return nil, err
	}

	if o.PeerDevice == "" || o.PeerInterface == "" {
		return nil, err
	}
	v := url.Values{}

	v.Set("peer-device", o.PeerDevice)

	v.Set("peer-interface", o.PeerInterface)

	return v, nil
}
