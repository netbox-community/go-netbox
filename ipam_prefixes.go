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
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strconv"
)

// A Prefix is an IPv4 or IPv6 address prefix.
type Prefix struct {
	ID          int
	Family      Family
	Prefix      *net.IPNet
	Site        *SiteIdentifier
	VRF         *VRFIdentifier
	VLAN        *VLANIdentifier
	Status      Status
	Role        *RoleIdentifier
	Description string
}

// GetPrefix retrieves a Prefix object from NetBox by its ID.
func (s *IPAMService) GetPrefix(id int) (*Prefix, error) {
	req, err := s.c.NewRequest(
		http.MethodGet,
		fmt.Sprintf("/api/ipam/prefixes/%d", id),
		nil,
	)
	if err != nil {
		return nil, err
	}

	p := new(Prefix)
	err = s.c.Do(req, p)
	return p, err
}

// ListPrefixes retrives a list of Prefix objects from NetBox, filtered
// according to the parameters specified in options.
//
// If options is nil, all Prefixes will be retrieved.
func (s *IPAMService) ListPrefixes(options *ListPrefixesOptions) ([]*Prefix, error) {
	req, err := s.c.NewRequest(http.MethodGet, "/api/ipam/prefixes/", options)
	if err != nil {
		return nil, err
	}

	var ps []*Prefix
	err = s.c.Do(req, &ps)
	return ps, err
}

// ListPrefixesOptions is used as an argument for Client.IPAM.ListPrefixes.
// Integer fields with an *ID suffix are preferred over their string
// counterparts, and if both are set, only the *ID field will be used.
// In addition, VLANID is preferred over the site-specific VLANVID.
type ListPrefixesOptions struct {
	Family  Family
	SiteID  []int
	Site    []string
	VRFID   int
	VRF     string
	VLANID  []int
	VLANVID VLANID
	Status  int
	RoleID  []int
	Role    []string
	Parent  *net.IPNet

	// Query is a special option which enables free-form search.
	// For example, Query could be an IP address such as "8.8.8.8".
	Query string
}

// Values generates a url.Values map from the data in ListPrefixesOptions.
func (o *ListPrefixesOptions) Values() (url.Values, error) {
	if o == nil {
		return nil, nil
	}

	v := url.Values{}

	if o.Family != 0 {
		v.Set("family", strconv.Itoa(int(o.Family)))
	}

	// IDs should always be preferred over string names

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
	case o.VRFID != 0:
		v.Set("vrf_id", strconv.Itoa(o.VRFID))
	case o.VRF != "":
		v.Set("vrf", o.VRF)
	}

	// Also prefer VLAN's NetBox ID over its VLAN VID
	switch {
	case len(o.VLANID) > 0:
		for _, vid := range o.VLANID {
			v.Add("vlan_id", strconv.Itoa(vid))
		}
	case o.VLANVID != 0:
		v.Set("vlan_vid", strconv.Itoa(int(o.VLANVID)))
	}

	if o.Status != 0 {
		v.Set("status", strconv.Itoa(o.Status))
	}

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

	if o.Parent != nil {
		v.Set("parent", o.Parent.String())
	}

	if o.Query != "" {
		v.Set("q", o.Query)
	}

	return v, nil
}

// A prefix is the raw JSON representation of a Prefix.
type prefix struct {
	ID          int             `json:"id"`
	Family      Family          `json:"family"`
	Prefix      string          `json:"prefix"`
	Site        *SiteIdentifier `json:"site"`
	VRF         *VRFIdentifier  `json:"vrf"`
	VLAN        *VLANIdentifier `json:"vlan"`
	Status      Status          `json:"status"`
	Role        *RoleIdentifier `json:"role"`
	Description string          `json:"description"`
}

// MarshalJSON marshals an Prefix into JSON bytes.
func (p *Prefix) MarshalJSON() ([]byte, error) {
	return json.Marshal(prefix{
		ID:          p.ID,
		Family:      p.Family,
		Prefix:      p.Prefix.String(),
		Site:        p.Site,
		VRF:         p.VRF,
		VLAN:        p.VLAN,
		Status:      p.Status,
		Role:        p.Role,
		Description: p.Description,
	})
}

// UnmarshalJSON unmarshals JSON bytes into an Prefix, and verifies that
// the contained IP address is valid.
func (p *Prefix) UnmarshalJSON(b []byte) error {
	var raw prefix
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	_, prefix, err := net.ParseCIDR(raw.Prefix)
	if err != nil {
		return err
	}

	*p = Prefix{
		ID:          raw.ID,
		Family:      raw.Family,
		Prefix:      prefix,
		Site:        raw.Site,
		VRF:         raw.VRF,
		VLAN:        raw.VLAN,
		Status:      raw.Status,
		Role:        raw.Role,
		Description: raw.Description,
	}
	return nil
}
