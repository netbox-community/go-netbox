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
	"fmt"
	"net/http"
)

// A VRF is a Virtual Routing and Forwarding device.
type VRF struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	RD          string `json:"rd"`
	Description string `json:"description"`
}

// A VRFIdentifier is a VRF returned as a nested object in some top-level
// objects.  It contains information which can be used in subsequent API
// calls to identify and retrieve a full VRF.
type VRFIdentifier struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	RD   string `json:"rd"`
}

// GetVRF retrieves a VRF object from NetBox by its ID.
func (s *IPAMService) GetVRF(id int) (*VRF, error) {
	req, err := s.c.NewRequest(
		http.MethodGet,
		fmt.Sprintf("/api/ipam/vrfs/%d", id),
		nil,
	)
	if err != nil {
		return nil, err
	}

	v := new(VRF)
	err = s.c.Do(req, v)
	return v, err
}

// ListVRFs retrives a list of VRF objects from NetBox.
func (s *IPAMService) ListVRFs() ([]*VRF, error) {
	req, err := s.c.NewRequest(http.MethodGet, "/api/ipam/vrfs/", nil)
	if err != nil {
		return nil, err
	}

	var vs []*VRF
	err = s.c.Do(req, &vs)
	return vs, err
}
