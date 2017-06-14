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
)

// A Tenant is a representation of netbox tenants
type Tenant struct {
	ID          int          `json:"id,omitempty"`
	Name        string       `json:"name"`
	Slug        string       `json:"slug"`
	Description string       `json:"description"`
	Comments    string       `json:"comments"`
	Group       *TenantGroup `json:"Group"`
}

// A updateTenant is the internal representation of tenants
// needed to POST/PUT/PATCH to Netbox API
type updateTenant struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Comments    string `json:"comments"`
	Group       int    `json:"group,omitempty"`
}

// MarshalJSON marshals an Tenant into JSON bytes.
func (t *Tenant) MarshalJSON() ([]byte, error) {
	group := 0
	if t.Group != nil {
		group = t.Group.ID
	}
	return json.Marshal(updateTenant{
		ID:          t.ID,
		Name:        t.Name,
		Slug:        t.Slug,
		Description: t.Description,
		Comments:    t.Comments,
		Group:       group,
	})
}

// ListTenantOptions is used as an argument for Client.Tenancy.Tenants.List.
// Integer fields with an *ID suffix are preferred over their string
// counterparts, and if both are set, only the *ID field will be used.
type ListTenantOptions struct {
	Name    string
	IDIn    string
	GroupID int
	Group   string

	Query string
}

// Values generates a url.Values map from the data in ListTenantOptions.
func (o *ListTenantOptions) Values() (url.Values, error) {
	if o == nil {
		return nil, nil
	}

	v := url.Values{}

	switch {
	case o.GroupID != 0:
		v.Set("group_id", strconv.Itoa(o.GroupID))
	case o.Group != "":
		v.Set("group", o.Group)
	}

	if o.Name != "" {
		v.Set("name", o.Name)
	}

	if o.IDIn != "" {
		v.Set("id__in", o.IDIn)
	}

	if o.Query != "" {
		v.Set("q", o.Query)
	}

	return v, nil
}

//go:generate go run generate_functions.go -type-name Tenant -update-type-name updateTenant -service-name TenantsService -endpoint tenancy -service tenants
//go:generate go run generate_basic_tests.go -type-name Tenant -service-name TenantsService -endpoint tenancy -service tenants -client-endpoint Tenancy -client-service Tenants
