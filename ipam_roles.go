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

// A Role is a role tag which can be applied to an object such as a
// Prefix or VLAN.  Role values are typically used to indicate the role
// of an object, such as infrastructure, management, customer, etc.
type Role struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	Weight int    `json:"weight"`
}

// A RoleIdentifier is a reduced version of a Role, returned as a nested
// object in some top-level objects.  It contains information which can
// be used in subsequent API calls to identify and retrieve a full Role.
type RoleIdentifier struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// GetRole retrieves a Role object from NetBox by its ID.
func (s *IPAMService) GetRole(id int) (*Role, error) {
	req, err := s.c.NewRequest(
		http.MethodGet,
		fmt.Sprintf("/api/ipam/roles/%d", id),
		nil,
	)
	if err != nil {
		return nil, err
	}

	r := new(Role)
	err = s.c.Do(req, r)
	return r, err
}

// ListRoles retrieves a list of Role objects from NetBox.
func (s *IPAMService) ListRoles() ([]*Role, error) {
	req, err := s.c.NewRequest(http.MethodGet, "/api/ipam/roles/", nil)
	if err != nil {
		return nil, err
	}

	var rs []*Role
	err = s.c.Do(req, &rs)
	return rs, err
}
