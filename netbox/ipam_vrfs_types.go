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
	"strings"
)

// A VRF is a representation of netbox vrf
type VRF struct {
	ID            int     `json:"id,omitempty"`
	Name          string  `json:"name"`
	RD            string  `json:"rd"`
	EnforceUnique bool    `json:"enforce_unique"`
	Description   string  `json:"description"`
	Tenant        *Tenant `json:"tenant,omitempty"`
}

// A writableVRF corresponds to the Netbox API's
// writable serializer for an VRF. It is used transparently
// when VRF are serialized in to JSON.
type writableVRF struct {
	ID            int    `json:"id,omitempty"`
	Name          string `json:"name"`
	RD            string `json:"rd"`
	EnforceUnique bool   `json:"enforce_unique"`
	Description   string `json:"description"`
	Tenant        *int   `json:"tenant,omitempty"`
}

// MarshalJSON marshals an VRF into JSON bytes,
// and is used by the standard json package.
func (v *VRF) MarshalJSON() ([]byte, error) {
	var tenantID *int
	if v.Tenant != nil {
		tenantID = &v.Tenant.ID
	}
	return json.Marshal(writableVRF{
		ID:            v.ID,
		Name:          v.Name,
		RD:            v.RD,
		EnforceUnique: v.EnforceUnique,
		Description:   v.Description,
		Tenant:        tenantID,
	})
}

// ListVRFOptions is used as an argument for Client.IPAM.VRFs.List.
type ListVRFOptions struct {
	Name          string
	RD            string
	EnforceUnique *bool
	IDIn          []uint64
	TenantID      int
	Tenant        string
	Query         string
}

// Values generates a url.Values map from the data in ListVRFOptions.
func (o *ListVRFOptions) Values() (url.Values, error) {
	if o == nil {
		return nil, nil
	}

	v := url.Values{}

	if o.Name != "" {
		v.Set("name", o.Name)
	}

	if o.RD != "" {
		v.Set("rd", o.RD)
	}

	if o.EnforceUnique != nil {
		if *o.EnforceUnique {
			v.Set("enforce_unique", "True")
		} else {
			v.Set("enforce_unique", "False")
		}
	}

	if len(o.IDIn) > 0 {
		vals := make([]string, len(o.IDIn))
		for i, k := range o.IDIn {
			vals[i] = strconv.FormatUint(k, 10)
		}
		v.Set("id__in", strings.Join(vals, ","))
	}

	switch {
	case o.TenantID != 0:
		v.Set("tenant_id", strconv.Itoa(o.TenantID))
	case o.Tenant != "":
		v.Set("tenant", o.Tenant)
	}

	if o.Query != "" {
		v.Set("q", o.Query)
	}

	return v, nil
}

//go:generate go run generate_functions.go -type-name VRF -update-type-name writableVRF -service-name VRFsService -endpoint ipam -service vrfs
//go:generate go run generate_basic_tests.go -type-name VRF -service-name IpamVRFService -endpoint ipam -service vrfs -client-endpoint IPAM -client-service VRFs
