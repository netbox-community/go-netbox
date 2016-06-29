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

// An RIR is a Regional Internet Registry which manages allocation of IP
// addresses.
type RIR struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// An RIRIdentifier is an RIR returned as a nested object in some top-level
// objects.  Though RIR and RIRIdentifier currently share the same fields,
// this may not always be the case.  It contains information which can be
// used in subsequent API calls to identify and retrieve a full RIR.
type RIRIdentifier struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// GetRIR retrieves an RIR object from NetBox by its ID.
func (s *IPAMService) GetRIR(id int) (*RIR, error) {
	req, err := s.c.newRequest(
		http.MethodGet,
		fmt.Sprintf("/api/ipam/rirs/%d", id),
		nil,
	)
	if err != nil {
		return nil, err
	}

	r := new(RIR)
	err = s.c.do(req, r)
	return r, err
}

// ListRIRs retrives a list of RIR objects from NetBox.
func (s *IPAMService) ListRIRs() ([]*RIR, error) {
	req, err := s.c.newRequest(http.MethodGet, "/api/ipam/rirs/", nil)
	if err != nil {
		return nil, err
	}

	var rs []*RIR
	err = s.c.do(req, &rs)
	return rs, err
}
