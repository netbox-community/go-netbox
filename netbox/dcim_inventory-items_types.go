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

// A NestedManufacturer corresponds to the Netbox API's
// nested serializer, when a manufacturer appears as a
// parent of another object.
//
// https://netbox.readthedocs.io/en/stable/api/overview/#serialization
type NestedManufacturer struct {
	ID   int    `json:"id"`
	URL  string `json:"url"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// A NestedDevice corresponds to the Netbox API's
// nested serializer, when a device appears as a
// parent of another object.
type NestedDevice struct {
	ID          int    `json:"id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}

// An InventoryItem corresponds to the Netbox API's
// base serializer for Inventory Items.
type InventoryItem struct {
	ID           int                 `json:"id,omitempty"`
	Device       NestedDevice        `json:"device"`
	Parent       *int                `json:"parent,omitempty"`
	Name         string              `json:"name"`
	Manufacturer *NestedManufacturer `json:"manufacturer,omitempty"`
	PartID       string              `json:"part_id,omitempty"`
	Serial       string              `json:"serial,omitempty"`
	Discovered   bool                `json:"discovered,omitempty"`
}

// A writableInventoryItem corresponds to the Netbox API's
// writable serializer for an InventoryItem. It is used transparently
// when IventoryItems are serialized in to JSON.
type writableInventoryItem struct {
	ID           int    `json:"id,omitempty"`
	Device       int    `json:"device"`
	Parent       *int   `json:"parent,omitempty"`
	Name         string `json:"name"`
	Manufacturer *int   `json:"manufacturer,omitempty"`
	PartID       string `json:"part_id,omitempty"`
	Serial       string `json:"serial,omitempty"`
	Discovered   bool   `json:"discovered,omitempty"`
}

// MarshalJSON marshals an InventoryItem into JSON bytes,
// and is used by the standard json package.
func (i *InventoryItem) MarshalJSON() ([]byte, error) {
	var manufacturerID *int
	if i.Manufacturer != nil {
		manufacturerID = &i.Manufacturer.ID
	}
	return json.Marshal(writableInventoryItem{
		ID:           i.ID,
		Device:       i.Device.ID,
		Parent:       i.Parent,
		Name:         i.Name,
		Manufacturer: manufacturerID,
		PartID:       i.PartID,
		Serial:       i.Serial,
		Discovered:   i.Discovered,
	})
}

// ListInventoryItemOptions is used as an argument for Client.DCIM.InventoryItems.List.
// Integer fields with an *ID suffix are preferred over their string
// counterparts, and if both are set, only the *ID field will be used.
type ListInventoryItemOptions struct {
	Name     string
	DeviceID int
	Device   string
}

// Values generates a url.Values map from the data in ListTenantOptions.
func (o *ListInventoryItemOptions) Values() (url.Values, error) {
	if o == nil {
		return nil, nil
	}

	v := url.Values{}

	if o.Name != "" {
		v.Set("name", o.Name)
	}

	switch {
	case o.DeviceID != 0:
		v.Set("device_id", strconv.Itoa(o.DeviceID))
	case o.Device != "":
		v.Set("device", o.Device)
	}

	return v, nil
}

//go:generate go run generate_functions.go -type-name InventoryItem -update-type-name writableInventoryItem -service-name InventoryItemsService -endpoint dcim -service inventory-items
//go:generate go run generate_basic_tests.go -type-name InventoryItem -service-name InventoryItemsService -endpoint dcim -service inventory-items -client-endpoint DCIM -client-service InventoryItems
