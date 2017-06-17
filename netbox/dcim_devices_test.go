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
	_ "bytes"
	"encoding/json"
	"fmt"
	"net/http"
	_ "net/url"
	"reflect"
	"testing"
)

func TestDeviceGet(t *testing.T) {
	var tests = []struct {
		desc string
		want *Device
	}{
		{
			desc: "Device with Platform",
			want: testDevice(1),
		},
	}

	for idx, tt := range tests {
		t.Run(fmt.Sprintf("[%d] %s", idx, tt.desc), func(t *testing.T) {
			serverData := serverDataDevice(*tt.want)

			c, done := testClient(t, testHandler(t, http.MethodGet, "/api/dcim/devices/1/", &serverData))
			defer done()

			res, err := c.DCIM.Devices.Get(1)
			if err != nil {
				t.Fatalf("unexpected error from c.DCIM.Devices.Get: %v", err)
			}

			if want, got := tt.want, res; !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected Device\n- want: %v\n- got: %v", want, got)
			}
		})
	}
}

func TestDeviceUnmarshalJSON(t *testing.T) {
	var tests = []struct {
		desc string
		data []byte
		want *Device
	}{
		{
			desc: "Maximum device",
			data: []byte(`{"id": 1, "name": "Device 1", "display_name": "Device 1", "device_type": { "id": 11, "url": "http://localhost/api/dcim/device-types/11/", "manufacturer": { "id": 21, "url": "http://localhost/api/dcim/manufacturers/21/", "name": "Manufacturer Name", "slug": "mfg-name"}, "model": "Device Type Model", "slug": "device-type-model"}, "device_role": { "id": 31, "url": "http://localhost/api/dcim/device-roles/31/", "name": "Device Role Name", "slug": "device-role-name"}, "tenant": { "id": 41, "url": "http://localhost/api/tenancy/tenants/41/", "name": "Tenant Name", "slug": "tenant-name" }, "platform": { "id": 51, "url": "http://localhost/api/dcim/platforms/51", "name": "Platform Name", "slug": "platform-name" }, "serial": "Serial", "asset_tag": "Tag", "site": { "id": 61, "url": "http://localhost/api/dcim/sites/61/", "name": "Site Name", "slug": "site-name" }, "rack": { "id": 71, "url": "http://localhost/api/dcim/racks/71/", "name": "Rack Name", "display_name": "Rack Name" }, "position": 81, "face": { "value": 0, "label": "Front" }, "status": { "value": 1, "label": "Active" } }`),
			want: &Device{
				ID:          1,
				Name:        "Device 1",
				DisplayName: "Device 1",
				DeviceType: &NestedDeviceType{
					ID:  11,
					URL: "http://localhost/api/dcim/device-types/11/",
					Manufacturer: &NestedManufacturer{
						ID:   21,
						URL:  "http://localhost/api/dcim/manufacturers/21/",
						Name: "Manufacturer Name",
						Slug: "mfg-name",
					},
					Model: "Device Type Model",
					Slug:  "device-type-model",
				},
				DeviceRole: &NestedDeviceRole{
					ID:   31,
					URL:  "http://localhost/api/dcim/device-roles/31/",
					Name: "Device Role Name",
					Slug: "device-role-name",
				},
				Tenant: &NestedTenant{
					ID:   41,
					URL:  "http://localhost/api/tenancy/tenants/41/",
					Name: "Tenant Name",
					Slug: "tenant-name",
				},
				Platform: &NestedPlatform{
					ID:   51,
					URL:  "http://localhost/api/dcim/platforms/51",
					Name: "Platform Name",
					Slug: "platform-name",
				},
				Serial:   "Serial",
				AssetTag: "Tag",
				Site: &NestedSite{
					ID:   61,
					URL:  "http://localhost/api/dcim/sites/61/",
					Name: "Site Name",
					Slug: "site-name",
				},
				Rack: &NestedRack{
					ID:          71,
					URL:         "http://localhost/api/dcim/racks/71/",
					Name:        "Rack Name",
					DisplayName: "Rack Name",
				},
				Position: 81,
				Face: &FaceType{
					Value: 0,
					Label: "Front",
				},
				Status: &StatusType{
					Value: 1,
					Label: "Active",
				},
			},
		},
	}

	for idx, tt := range tests {
		t.Run(fmt.Sprintf("[%d] %s", idx, tt.desc), func(t *testing.T) {
			result := new(Device)
			err := json.Unmarshal(tt.data, result)
			if err != nil {
				t.Fatalf("unexpected error from Device.UnmarshalJSON: %v", err)
			}

			if want, got := tt.want, result; !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected Device:\n- want: %v\n- got: %v", want, got)
			}
		})
	}
}

func testNestedDeviceRole(id int) *NestedDeviceRole {
	return &NestedDeviceRole{
		ID:   id,
		URL:  fmt.Sprintf("http://localhost/api/dcim/device-roles/%d/", id),
		Name: fmt.Sprintf("Device Role %d", id),
		Slug: fmt.Sprintf("device-role-%d", id),
	}
}

func testNestedPlatform(id int) *NestedPlatform {
	return &NestedPlatform{
		ID:   id,
		URL:  fmt.Sprintf("http://localhost/api/dcim/platforms/%d/", id),
		Name: fmt.Sprintf("Platform %d", id),
		Slug: fmt.Sprintf("platform-%d", id),
	}
}

func testNestedTenant(id int) *NestedTenant {
	return &NestedTenant{
		ID:   id,
		URL:  fmt.Sprintf("http://localhost/api/tenancy/tenants/%d/", id),
		Name: fmt.Sprintf("Tenant %d", id),
		Slug: fmt.Sprintf("tenant-%d", id),
	}
}

func testNestedSite(id int) *NestedSite {
	return &NestedSite{
		ID:   id,
		URL:  fmt.Sprintf("http://localhost/api/dcim/sites/%d/", id),
		Name: fmt.Sprintf("Site %d", id),
		Slug: fmt.Sprintf("site-%d", id),
	}
}

func testNestedRack(id int) *NestedRack {
	return &NestedRack{
		ID:          id,
		URL:         fmt.Sprintf("http://localhost/api/dcim/racks/%d/", id),
		Name:        fmt.Sprintf("Rack %d", id),
		DisplayName: fmt.Sprintf("Rack %d", id),
	}
}

func testNestedDeviceType(id int) *NestedDeviceType {
	manufacturerID := id + 1000
	return &NestedDeviceType{
		ID:           id,
		URL:          fmt.Sprintf("http://localhost/api/dcim/device-types/%d/", id),
		Manufacturer: testNestedManufacturer(manufacturerID),
		Model:        fmt.Sprintf("Device Type Model of %d", id),
		Slug:         fmt.Sprintf("test-device-type-model-%d", id),
	}
}

func testDevice(id int) *Device {
	roleID := id + 1000
	typeID := id + 2000
	tenantID := id + 3000
	platformID := id + 4000
	siteID := id + 5000
	rackID := id + 6000

	return testDeviceHelper(id, false, testNestedDeviceType(typeID), testNestedDeviceRole(roleID), testNestedTenant(tenantID), testNestedPlatform(platformID), testNestedSite(siteID), testNestedRack(rackID))
}

func testDeviceCreate(id int) *Device {
	roleID := id + 1000
	typeID := id + 2000
	tenantID := id + 3000
	platformID := id + 4000
	siteID := id + 5000
	rackID := id + 6000
	return testDeviceHelper(id, true, testNestedDeviceType(typeID), testNestedDeviceRole(roleID), testNestedTenant(tenantID), testNestedPlatform(platformID), testNestedSite(siteID), testNestedRack(rackID))
}

func testDeviceHelper(id int, create bool, devicetype *NestedDeviceType, role *NestedDeviceRole, tenant *NestedTenant, platform *NestedPlatform, site *NestedSite, rack *NestedRack) *Device {
	deviceID := id

	if create {
		deviceID = 0
	}

	return &Device{
		ID:          deviceID,
		Name:        fmt.Sprintf("Device %d", id),
		DisplayName: fmt.Sprintf("Device %d", id),
		DeviceType:  devicetype,
		DeviceRole:  role,
		Tenant:      tenant,
		Platform:    platform,
		Site:        site,
		Rack:        rack,
	}
}
