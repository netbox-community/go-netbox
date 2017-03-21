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

// An IPAddress is an IPv4 or IPv6 address.
type IPAddress struct {
	ID          int
	Family      Family
	Address     *net.IPNet
	VRF         *VRFIdentifier
	Interface   *InterfaceIdentifier
	Description string
	NATInside   *IPAddressIdentifier
	NATOutside  *IPAddressIdentifier
}

// An IPAddressIdentifier is a reduced version of a IPAddress, returned as
// nested object in some top-level objects.  It contains information which can
// be used in subsequent API calls to identify and retrieve a full IPAddress.
type IPAddressIdentifier struct {
	ID      int
	Family  Family
	Address *net.IPNet
}

// GetIPAddress retrieves an IPAddress object from NetBox by its ID.
func (s *IPAMService) GetIPAddress(id int) (*IPAddress, error) {
	req, err := s.c.NewRequest(
		http.MethodGet,
		fmt.Sprintf("/api/ipam/ip-addresses/%d", id),
		nil,
	)
	if err != nil {
		return nil, err
	}

	ip := new(IPAddress)
	err = s.c.Do(req, ip)
	return ip, err
}

// ListIPAddresses retrives a list of IPAddress objects from NetBox, filtered according
// to the parameters specified in options.
//
// If options is nil, all IPAddresses will be retrieved.
func (s *IPAMService) ListIPAddresses(options *ListIPAddressesOptions) ([]*IPAddress, error) {
	req, err := s.c.NewRequest(http.MethodGet, "/api/ipam/ip-addresses/", options)
	if err != nil {
		return nil, err
	}

	var ips []*IPAddress
	err = s.c.Do(req, &ips)
	return ips, err
}

// An ipAddress is the raw JSON representation of an IPAddress.
type ipAddress struct {
	ID          int                  `json:"id"`
	Family      Family               `json:"family"`
	Address     string               `json:"address"`
	VRF         *VRFIdentifier       `json:"vrf"`
	Interface   *InterfaceIdentifier `json:"interface"`
	Description string               `json:"description"`
	NATInside   *IPAddressIdentifier `json:"nat_inside"`
	NATOutside  *IPAddressIdentifier `json:"nat_outside"`
}

// MarshalJSON marshals an IPAddress into JSON bytes.
func (ip *IPAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(ipAddress{
		ID:          ip.ID,
		Family:      ip.Family,
		Address:     ip.Address.String(),
		VRF:         ip.VRF,
		Interface:   ip.Interface,
		Description: ip.Description,
		NATInside:   ip.NATInside,
		NATOutside:  ip.NATOutside,
	})
}

// UnmarshalJSON unmarshals JSON bytes into an IPAddress, and verifies that
// the contained IP address is valid.
func (ip *IPAddress) UnmarshalJSON(b []byte) error {
	var raw ipAddress
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	_, ipNet, err := net.ParseCIDR(raw.Address)
	if err != nil {
		return err
	}

	*ip = IPAddress{
		ID:          raw.ID,
		Family:      raw.Family,
		Address:     ipNet,
		VRF:         raw.VRF,
		Interface:   raw.Interface,
		Description: raw.Description,
		NATInside:   raw.NATInside,
		NATOutside:  raw.NATOutside,
	}
	return nil
}

// An ipAddressIdentifier is the raw JSON representation of an IPAddressIdentifier.
type ipAddressIdentifier struct {
	ID      int    `json:"id"`
	Family  Family `json:"family"`
	Address string `json:"address"`
}

// MarshalJSON marshals an IPAddressIdentifier into JSON bytes.
func (ip *IPAddressIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(ipAddressIdentifier{
		ID:      ip.ID,
		Family:  ip.Family,
		Address: ip.Address.String(),
	})
}

// UnmarshalJSON unmarshals JSON bytes into an IPAddressIdentifier, and verifies that
// the contained IP address is valid.
func (ip *IPAddressIdentifier) UnmarshalJSON(b []byte) error {
	var raw ipAddressIdentifier
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	_, ipNet, err := net.ParseCIDR(raw.Address)
	if err != nil {
		return err
	}

	*ip = IPAddressIdentifier{
		ID:      raw.ID,
		Family:  raw.Family,
		Address: ipNet,
	}
	return nil
}

// ListIPAddressesOptions is used as an argument for Client.IPAM.ListIPAddresses.
// Integer fields with an *ID suffix are preferred over their string
// counterparts, and if both are set, only the *ID field will be used.
type ListIPAddressesOptions struct {
	Family      Family
	VRFID       []int
	VRF         string
	InterfaceID []int
	DeviceID    []int
	Device      []string

	// Query is a special option which enables free-form search.
	// For example, Query could be an IP address such as "8.8.8.8".
	Query string
}

// values generates a url.Values map from the data in ListIPAddressesOptions.
func (o *ListIPAddressesOptions) values() (url.Values, error) {
	if o == nil {
		return nil, nil
	}

	v := url.Values{}

	if o.Family != 0 {
		v.Set("family", strconv.Itoa(int(o.Family)))
	}

	for _, i := range o.InterfaceID {
		v.Add("interface_id", strconv.Itoa(i))
	}

	// IDs should always be preferred over string names

	switch {
	case len(o.VRFID) > 0:
		for _, vid := range o.VRFID {
			v.Add("vrf_id", strconv.Itoa(vid))
		}
	case o.VRF != "":
		v.Set("vrf", o.VRF)
	}

	switch {
	case len(o.DeviceID) > 0:
		for _, d := range o.DeviceID {
			v.Add("device_id", strconv.Itoa(d))
		}
	case len(o.Device) > 0:
		for _, d := range o.Device {
			v.Add("device", d)
		}
	}

	if o.Query != "" {
		v.Set("q", o.Query)
	}

	return v, nil
}
