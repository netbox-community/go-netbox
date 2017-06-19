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

type nestedBase struct {
	ID   int    `json:"id"`
	URL  string `json:"url"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type simpleValueLabel struct {
	Value int    `json:"value"`
	Label string `json:"label"`
}

// ShortDeviceBay represent the short form of a Device Bay
// when a Device Bay appears in a parent device of
// another object.
type ShortDeviceBay struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ShortDevice corresponds to the simple form of
// a Device, when a Device appears as a
// parent device of another object
type ShortDevice struct {
	DeviceBay *ShortDeviceBay `json:"device_bay"`
	ID        int             `json:"id"`
	Name      string          `json:"name"`
}

// FaceType represent the face of a Device in a Rack (Front/Rear)
type FaceType simpleValueLabel

// StatusType represent status of a Device in Netbox API.
type StatusType simpleValueLabel

// A NestedDeviceRole corresponds to the Netbox API's
// nested serializer, when a DeviceRole appears as a
// parent of another object
type NestedDeviceRole nestedBase

// A NestedPlatform corresponds to the Netbox API's
// nested serializer, when a Platform appears as a
// parent of another object
type NestedPlatform nestedBase

// A NestedTenant corresponds to the Netbox API's
// nested serializer, when a Tenant appears as a
// parent of another object
type NestedTenant nestedBase

// A NestedSite corresponds to the Netbox API's
// nested serializer, when a Site appears as a
// parent of another object
type NestedSite nestedBase

// A NestedDeviceType corresponds to the Netbox API's
// nested serializer, when a DeviceType appears as a
// parent of another object
type NestedDeviceType struct {
	ID           int                 `json:"id"`
	URL          string              `json:"url"`
	Manufacturer *NestedManufacturer `json:"manufacturer"`
	Model        string              `json:"model"`
	Slug         string              `json:"slug"`
}

// A NestedRack correspondes to the Netbox API's
// nested serializer, when a Rack appears as a
// parent of another object
type NestedRack struct {
	ID          int    `json:"id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}

// A NestedIP will be used as nested serializer for
// IPv4, IPv6 ip address
type NestedIP struct {
	ID      int    `json:"id"`
	URL     string `json:"url"`
	Family  int    `json:"family"`
	Address string `json:"address"`
}

// A Device corresponds to the Netbox API's
// base serializer for Device
type Device struct {
	ID          int               `json:"id,omitempty"`
	Name        string            `json:"name"`
	DisplayName string            `json:"display_name"`
	DeviceType  *NestedDeviceType `json:"device_type"`
	DeviceRole  *NestedDeviceRole `json:"device_role"`
	Tenant      *NestedTenant     `json:"tenant,omitempty"`
	Platform    *NestedPlatform   `json:"platform,omitempty"`
	Serial      string            `json:"serial,omitempty"`
	AssetTag    string            `json:"asset_tag,omitempty"`
	Site        *NestedSite       `json:"site"`
	Rack        *NestedRack       `json:"rack,omitempty"`
	Position    int               `json:"position,omitempty"`
	Face        *FaceType         `json:"face,omitempty"`
	Parent      *ShortDevice      `json:"parent_device,omitempty"`
	Status      *StatusType       `json:"status"`
	PrimaryIP   *NestedIP         `json:"primary_ip,omitempty"`
	PrimaryIPv4 *NestedIP         `json:"primary_ip4,omitempty"`
	PrimaryIPv6 *NestedIP         `json:"primary_ip6,omitempty"`
	Comments    string            `json:"comments,omitempty"`
}

// A writableDevice corresponds to the Netbox API's
// writable serializer for a Device. It is used transparently
// when Devices are serialize into JSON.
type writableDevice struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	DeviceType  int    `json:"device_type"`
	DeviceRole  int    `json:"device_role"`
	Tenant      int    `json:"tenant,omitempty"`
	Platform    int    `json:"platform,omitempty"`
	Serial      string `json:"serial,omitempty"`
	AssetTag    string `json:"asset_tag,omitempty"`
	Site        int    `json:"site"`
	Rack        int    `json:"rack,omitempty"`
	Position    int    `json:"position,omitempty"`
	Face        string `json:"face,omitempty"`
	Parent      int    `json:"parent_device,omitempty"`
	Status      string `json:"status,omitempty"`
	PrimaryIP   int    `json:"primary_ip,omitempty"`
	PrimaryIPv4 int    `json:"primary_ip4,omitempty"`
	PrimaryIPv6 int    `json:"primary_ip6,omitempty"`
	Comments    string `json:"comments,omitempty"`
}

// MarshalJSON marshals a Device into JSON bytes,
// and is used by the standard json package.
func (d *Device) MarshalJSON() ([]byte, error) {
	var typeID, roleID, tenantID, platformID, siteID, rackID, parentID, primaryIPID, primaryIPv4ID, primaryIPv6ID int
	var status, face string

	if d.DeviceType != nil {
		typeID = d.DeviceType.ID
	}

	if d.DeviceRole != nil {
		roleID = d.DeviceRole.ID
	}

	if d.Tenant != nil {
		tenantID = d.Tenant.ID
	}

	if d.Platform != nil {
		platformID = d.Platform.ID
	}

	if d.Site != nil {
		siteID = d.Site.ID
	}

	if d.Rack != nil {
		rackID = d.Rack.ID
	}

	if d.PrimaryIP != nil {
		primaryIPID = d.PrimaryIP.ID
	}

	if d.PrimaryIPv4 != nil {
		primaryIPv4ID = d.PrimaryIPv4.ID
	}

	if d.PrimaryIPv6 != nil {
		primaryIPv6ID = d.PrimaryIPv6.ID
	}

	if d.Face != nil {
		face = d.Face.Label
	}

	if d.Parent != nil {
		parentID = d.Parent.ID
	}

	if d.Status != nil {
		status = d.Status.Label
	}

	return json.Marshal(writableDevice{
		ID:          d.ID,
		Name:        d.Name,
		DisplayName: d.DisplayName,
		DeviceType:  typeID,
		DeviceRole:  roleID,
		Tenant:      tenantID,
		Platform:    platformID,
		Serial:      d.Serial,
		AssetTag:    d.AssetTag,
		Site:        siteID,
		Rack:        rackID,
		Position:    d.Position,
		Face:        face,
		Parent:      parentID,
		Status:      status,
		PrimaryIP:   primaryIPID,
		PrimaryIPv4: primaryIPv4ID,
		PrimaryIPv6: primaryIPv6ID,
		Comments:    d.Comments,
	})

}

// ListDeviceOptions is used as an argument for Client.DCIM.Devices.List.
// Integer fileds with an *ID suffix are preferred over their string
// counterparts, and if both are set, only the *ID filed will be used.
type ListDeviceOptions struct {
	Name            string
	Serial          string
	AssetTag        string
	MacAddress      string
	IDIn            int
	SiteID          int
	Site            string
	RackGroupID     int
	RackID          int
	RoleID          int
	Role            string
	TenantID        int
	Tenant          string
	DeviceTypeID    int
	ManufacturerID  int
	Manufacturer    string
	DeviceModel     string
	PlatformID      int
	Platform        string
	Status          string
	IsConsoleServer *bool
	IsPDU           *bool
	IsNetworkDevice *bool
	HasPrimaryIP    *bool

	Query string
}

// Values generates a url.Values map from the data in ListDeviceOptions.
func (o *ListDeviceOptions) Values() (url.Values, error) {
	if o == nil {
		return nil, nil
	}

	v := url.Values{}

	status := map[string]int{
		"Offline":   0,
		"Active":    1,
		"Planned":   2,
		"Staged":    3,
		"Failed":    4,
		"Inventory": 5,
	}

	switch {
	case o.SiteID != 0:
		v.Set("site_id", strconv.Itoa(o.SiteID))
	case o.Site != "":
		v.Set("site", o.Site)
	}

	switch {
	case o.TenantID != 0:
		v.Set("tenant_id", strconv.Itoa(o.TenantID))
	case o.Tenant != "":
		v.Set("tenant", o.Tenant)
	}

	switch {
	case o.ManufacturerID != 0:
		v.Set("manufacturer_id", strconv.Itoa(o.ManufacturerID))
	case o.Manufacturer != "":
		v.Set("manufacturer", o.Manufacturer)
	}

	switch {
	case o.PlatformID != 0:
		v.Set("platform_id", strconv.Itoa(o.PlatformID))
	case o.Platform != "":
		v.Set("platform", o.Platform)
	}

	switch {
	case o.RoleID != 0:
		v.Set("role_id", strconv.Itoa(o.RoleID))
	case o.Role != "":
		v.Set("role", o.Role)
	}

	if o.Name != "" {
		v.Set("name", o.Name)
	}

	if o.Serial != "" {
		v.Set("serial", o.Serial)
	}

	if o.AssetTag != "" {
		v.Set("asset_tag", o.AssetTag)
	}

	if o.MacAddress != "" {
		v.Set("mac_address", o.MacAddress)
	}

	if o.IDIn != 0 {
		v.Set("id__in", strconv.Itoa(o.IDIn))
	}

	if o.RackGroupID != 0 {
		v.Set("rack_group_id", strconv.Itoa(o.RackGroupID))
	}

	if o.RackID != 0 {
		v.Set("rack_id", strconv.Itoa(o.RackID))
	}

	if o.DeviceTypeID != 0 {
		v.Set("device_type_id", strconv.Itoa(o.DeviceTypeID))
	}

	if o.DeviceModel != "" {
		v.Set("model", o.DeviceModel)
	}

	if o.Status != "" {
		v.Set("status", strconv.Itoa(status[o.Status]))
	}

	if o.IsConsoleServer != nil {
		v.Set("is_console_server", strconv.FormatBool(*o.IsConsoleServer))
	}

	if o.IsPDU != nil {
		v.Set("is_pdu", strconv.FormatBool(*o.IsPDU))
	}

	if o.IsNetworkDevice != nil {
		v.Set("is_network_device", strconv.FormatBool(*o.IsNetworkDevice))
	}

	if o.HasPrimaryIP != nil {
		v.Set("has_primary_ip", strconv.FormatBool(*o.HasPrimaryIP))
	}

	if o.Query != "" {
		v.Set("q", o.Query)
	}

	return v, nil
}

//go:generate go run generate_functions.go -endpoint dcim -service devices -service-name DevicesService -type-name Device -update-type-name writableDevice
//go:generate go run generate_basic_tests.go -client-endpoint DCIM -client-service Devices -endpoint dcim -service devices -service-name DevicesService -type-name Device
