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

// A Site is a physical location where devices may reside.
type Site struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	Facility        string `json:"facility"`
	ASN             ASN    `json:"asn"`
	PhysicalAddress string `json:"physical_address"`
	ShippingAddress string `json:"shipping_address"`
	Comments        string `json:"comments"`
	CountPrefixes   int    `json:"count_prefixes"`
	CountVLANs      int    `json:"count_vlans"`
	CountRacks      int    `json:"count_racks"`
	CountDevices    int    `json:"count_devices"`
	CountCircuits   int    `json:"count_circuits"`
}

// A SiteIdentifier is a reduced version of a Site, returned as a nested
// object in some top-level objects.  It contains information which can
// be used in subsequent API calls to identify and retrieve a full Site.
type SiteIdentifier struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// GetSite retrieves a Site object from NetBox by its ID.
func (s *DCIMService) GetSite(id int) (*Site, error) {
	req, err := s.c.NewRequest(
		http.MethodGet,
		fmt.Sprintf("/api/dcim/sites/%d", id),
		nil,
	)
	if err != nil {
		return nil, err
	}

	st := new(Site)
	err = s.c.Do(req, st)
	return st, err
}

// ListSites retrives a list of Site objects from NetBox.
func (s *DCIMService) ListSites() ([]*Site, error) {
	req, err := s.c.NewRequest(http.MethodGet, "/api/dcim/sites/", nil)
	if err != nil {
		return nil, err
	}

	var sts []*Site
	err = s.c.Do(req, &sts)
	return sts, err
}
