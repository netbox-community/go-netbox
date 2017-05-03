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

// RelatedConnection represents components that have a related peer-device and
// peer-interface.
type RelatedConnection struct {
	Device       *Device        `json:"device"`
	ConsolePorts []*ConsolePort `json:"console-ports"`
	Interfaces   []*Interface   `json:"interfaces"`
	PowerPorts   []*PowerPort   `json:"power-ports"`
}

// ConsolePortIdentifier represents a reduced version of a console port.
type ConsolePortIdentifier struct {
	Device string `json:"device"`
	Name   string `json:"name"`
	Port   string `json:"port"`
}

// PowerPortIdentifier represents a reduced version of a single power port.
type PowerPortIdentifier struct {
	Device string `json:"device"`
	Name   string `json:"name"`
	Outlet string `json:"outlet"`
}

// GetRelatedConnections retrieves a RelatedConnection object from NetBox.
func (s *DCIMService) GetRelatedConnections(
	peerDevice, peerInterface string,
) (*RelatedConnection, error) {
	req, err := s.c.NewRequest(
		http.MethodGet,
		"/api/dcim/related-connections/",
		&relatedConnectionsOptions{
			peerDevice,
			peerInterface,
		},
	)
	if err != nil {
		return nil, err
	}

	rc := new(RelatedConnection)
	err = s.c.Do(req, rc)
	if err != nil {
		return nil, err
	}
	return rc, nil
}

// relatedConnectionsOptions is used as an argument for
// Client.DCIM.GetRelatedConnections.
type relatedConnectionsOptions struct {
	PeerDevice    string
	PeerInterface string
}

func (o *relatedConnectionsOptions) Values() (url.Values, error) {
	err := errors.New(
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
