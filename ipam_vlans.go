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
	"net/url"
	"strconv"
)

// A VLANID is a 12 bit integer VLAN ID.  Use its Valid method to determine
// if the contained value is a valid VLAN ID.
type VLANID int

// Valid determines if a VLANID contains a valid VLAN ID.
func (v VLANID) Valid() bool {
	// Cannot be less than 0, cannot be greater than 4096.
	return v >= 0 && v <= 4096
}

// A VLAN is a Virtual LAN object which can be assigned to a Site.
type VLAN struct {
	ID          int             `json:"id"`
	Site        *SiteIdentifier `json:"site"`
	VID         VLANID          `json:"vid"`
	Name        string          `json:"name"`
	Status      Status          `json:"status"`
	Role        *RoleIdentifier `json:"role"`
	DisplayName string          `json:"display_name"`
}

// A VLANIdentifier is a reduced version of a VLAN, returned as a nested
// object in some top-level objects.  It contains information which can
// be used in subsequent API calls to identify and retrieve a full VLAN.
type VLANIdentifier struct {
	ID          int    `json:"id"`
	VID         VLANID `json:"vid"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}

// GetVLAN retrieves a VLAN object from NetBox by its ID.
func (s *IPAMService) GetVLAN(id int) (*VLAN, error) {
	req, err := s.c.NewRequest(
		http.MethodGet,
		fmt.Sprintf("/api/ipam/vlans/%d", id),
		nil,
	)
	if err != nil {
		return nil, err
	}

	vlan := new(VLAN)
	err = s.c.Do(req, vlan)
	return vlan, err
}

// ListVLANs retrives a list of VLAN objects from NetBox, filtered according
// to the parameters specified in options.
//
// If options is nil, all VLANs will be retrieved.
func (s *IPAMService) ListVLANs(options *ListVLANsOptions) ([]*VLAN, error) {
	req, err := s.c.NewRequest(http.MethodGet, "/api/ipam/vlans/", options)
	if err != nil {
		return nil, err
	}

	var vlans []*VLAN
	err = s.c.Do(req, &vlans)
	return vlans, err
}

// ListVLANsOptions is used as an argument for Client.IPAM.ListVLANs.
// Integer fields with an *ID suffix are preferred over their string
// counterparts, and if both are set, only the *ID field will be used.
type ListVLANsOptions struct {
	Name     string
	Role     []string
	RoleID   []int
	Site     []string
	SiteID   []int
	Status   []string
	StatusID int
	VID      VLANID
}

// Values generates a url.Values map from the data in ListVLANsOptions.
func (o *ListVLANsOptions) Values() (url.Values, error) {
	if o == nil {
		return nil, nil
	}

	v := url.Values{}

	if o.Name != "" {
		v.Set("name", o.Name)
	}

	// IDs should always be preferred over string names

	switch {
	case len(o.RoleID) > 0:
		for _, r := range o.RoleID {
			v.Add("role_id", strconv.Itoa(r))
		}
	case len(o.Role) > 0:
		for _, r := range o.Role {
			v.Add("role", r)
		}
	}

	switch {
	case len(o.SiteID) > 0:
		for _, s := range o.SiteID {
			v.Add("site_id", strconv.Itoa(s))
		}
	case len(o.Site) > 0:
		for _, s := range o.Site {
			v.Add("site", s)
		}
	}

	switch {
	case o.StatusID != 0:
		v.Set("status_id", strconv.Itoa(o.StatusID))
	case len(o.Status) > 0:
		for _, s := range o.Status {
			v.Add("status", s)
		}
	}

	if o.VID != 0 {
		v.Set("vid", strconv.Itoa(int(o.VID)))
	}

	return v, nil
}
