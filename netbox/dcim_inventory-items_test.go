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
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestInventoryItemGet(t *testing.T) {
	var tests = []struct {
		desc             string
		want             *InventoryItem
		wantManufacturer *NestedManufacturer
	}{
		{
			desc: "InventoryItem without Manufacturer",
			want: testInventoryItem(1),
		},
		{
			desc: "InventoryItem with Manufacturer",
			want: testInventoryItemWithManufacturer(1),
		},
	}

	for idx, ii := range tests {
		t.Run(fmt.Sprintf("[%d] %s", idx, ii.desc), func(t *testing.T) {
			serverData := serverDataInventoryItem(*ii.want)

			c, done := testClient(t, testHandler(t, http.MethodGet, "/api/dcim/inventory-items/1/", &serverData))
			defer done()

			res, err := c.DCIM.InventoryItems.Get(1)
			if err != nil {
				t.Fatalf("unexpected error from c.DCIM.InventoryItems.Get: %v", err)
			}

			if want, got := ii.want, res; !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected InventoryItem\n- want: %v\n- got: %v", want, got)
			}
		})
	}
}

func TestInventoryItemUnmarshalJSON(t *testing.T) {
	parentID := 4
	var tests = []struct {
		desc string
		data []byte
		want *InventoryItem
	}{
		{
			desc: "Minimal inventory item",
			data: []byte(`{ "id": 4, "device": { "id": 1, "url": "http://localhost:8000/api/dcim/devices/1/", "name": "Test Device 1", "display_name": "Test Device 1"}, "parent": null, "name": "device name", "manufacturer": null }`),
			want: &InventoryItem{
				ID: 4,
				Device: NestedDevice{
					ID:          1,
					URL:         "http://localhost:8000/api/dcim/devices/1/",
					Name:        "Test Device 1",
					DisplayName: "Test Device 1",
				},
				Name:         "device name",
				Parent:       nil,
				Manufacturer: nil,
			},
		},
		{
			desc: "Maximal inventory item",
			data: []byte(`{ "id": 2, "device": { "id": 1, "url": "http://localhost:8000/api/dcim/devices/1/", "name": "Test Device 1", "display_name": "Test Device 1" }, "parent": 4, "name": "the device name", "manufacturer": { "id": 10, "url": "http://localhost:8000/api/dcim/manufacturers/10/", "name": "manufacturer name", "slug": "mfg-name" }, "part_id": "the part ID", "serial": "the serial", "discovered": true }`),
			want: &InventoryItem{
				ID: 2,
				Device: NestedDevice{
					ID:          1,
					URL:         "http://localhost:8000/api/dcim/devices/1/",
					Name:        "Test Device 1",
					DisplayName: "Test Device 1",
				},
				Parent: &parentID,
				Name:   "the device name",
				Manufacturer: &NestedManufacturer{
					ID:   10,
					URL:  "http://localhost:8000/api/dcim/manufacturers/10/",
					Name: "manufacturer name",
					Slug: "mfg-name",
				},
				PartID:     "the part ID",
				Serial:     "the serial",
				Discovered: true,
			},
		},
	}

	for idx, ii := range tests {
		t.Run(fmt.Sprintf("[%d] %s", idx, ii.desc), func(t *testing.T) {
			result := new(InventoryItem)
			err := json.Unmarshal(ii.data, result)
			if err != nil {
				t.Fatalf("unexpected error from InventoryItem.UnmarshalJSON: %v", err)
			}

			if want, got := ii.want, result; !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected InventoryItem:\n- want: %v\n-  got: %v", want, got)
			}
		})
	}
}

func TestInventoryItemMarshalJSON(t *testing.T) {
	var tests = []struct {
		desc string
		data *InventoryItem
		want []byte
	}{
		{
			desc: "Inventory item without manufacturer",
			data: testInventoryItem(1),
			want: []byte(`{"id":1,"device":10001,"parent":null,"name":"Inventory Item 1","part_id":"Part ID 1","serial":"Serial 1","discovered":true}`),
		},
		{
			desc: "Inventory item with manufacturer",
			data: testInventoryItemWithManufacturer(2),
			want: []byte(`{"device":10002,"parent":null,"name":"Inventory Item 2","manufacturer":20002,"part_id":"Part ID 2","serial":"Serial 2","discovered":true}`),
		},
	}

	for idx, ii := range tests {
		t.Run(fmt.Sprintf("[%d] %s", idx, ii.desc), func(t *testing.T) {
			result, err := json.Marshal(ii.data)
			if err != nil {
				t.Fatalf("unexpected error from writableInventoryItem.MarshalJSON: %v", err)
			}

			if want, got := ii.want, result; bytes.Compare(want, got) != 0 {
				t.Fatalf("unexpected JSON:\n- want: %v\n-  got: %v", string(want), string(got))
			}
		})
	}
}

func TestWritableInventoryItemUnmarshalJSON(t *testing.T) {
	parentID := 4
	manufacturerID := 6
	var tests = []struct {
		desc string
		data []byte
		want *writableInventoryItem
	}{
		{
			desc: "Minimal inventory item",
			data: []byte(`{ "device": 1, "parent": null, "name": "device name" }`),
			want: &writableInventoryItem{Device: 1, Name: "device name"},
		},
		{
			desc: "Maximal inventory item",
			data: []byte(`{ "id": 2, "device": 3, "parent": 4, "name": "the device name", "manufacturer": 6, "part_id": "the part ID", "serial": "the serial", "discovered": true }`),
			want: &writableInventoryItem{
				ID:           2,
				Device:       3,
				Parent:       &parentID,
				Name:         "the device name",
				Manufacturer: &manufacturerID,
				PartID:       "the part ID",
				Serial:       "the serial",
				Discovered:   true,
			},
		},
	}

	for idx, ii := range tests {
		t.Run(fmt.Sprintf("[%d] %s", idx, ii.desc), func(t *testing.T) {
			result := new(writableInventoryItem)
			err := json.Unmarshal(ii.data, result)
			if err != nil {
				t.Fatalf("unexpected error from writableInventoryItem.UnmarshalJSON: %v", err)
			}

			if want, got := ii.want, result; !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected writableInventoryItem:\n- want: %v\n-  got: %v", want, got)
			}
		})
	}
}

func TestListInventoryItemOptions(t *testing.T) {
	var tests = []struct {
		desc string
		o    *ListInventoryItemOptions
		v    url.Values
	}{
		{
			desc: "empty options",
		},
		{
			desc: "full options",
			o: &ListInventoryItemOptions{
				Name:     "Hello",
				DeviceID: 1,
				Device:   "node_name_1",
			},
			v: url.Values{
				"name":      []string{"Hello"},
				"device_id": []string{"1"},
			},
		},
		{
			desc: "device name",
			o: &ListInventoryItemOptions{
				Name:   "Hello",
				Device: "node_name_1",
			},
			v: url.Values{
				"name":   []string{"Hello"},
				"device": []string{"node_name_1"},
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("[%d] %s", i, tt.desc), func(t *testing.T) {
			v, err := tt.o.Values()
			if err != nil {
				t.Fatalf("unexpected Values error: %v", err)
			}

			if want, got := tt.v, v; !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected url.Values map:\n- want: %v\n-  got: %v", want, got)
			}
		})
	}
}

func testInventoryItem(id int) *InventoryItem {
	return testInventoryItemHelper(id, false, nil)
}

func testInventoryItemCreate(id int) *InventoryItem {
	return testInventoryItemHelper(id, true, nil)
}

func testInventoryItemWithManufacturer(id int) *InventoryItem {
	return testInventoryItemHelper(id, true, testNestedManufacturer(id+20000))
}

func testNestedManufacturer(id int) *NestedManufacturer {
	return &NestedManufacturer{
		ID:   id,
		URL:  fmt.Sprintf("http://localhost:8000/api/dcim/manufacturers/%d/", id),
		Name: fmt.Sprintf("Test nested manufacturer %d", id),
		Slug: fmt.Sprintf("test-nested-manufacturer-%d", id),
	}
}

func testInventoryItemHelper(id int, create bool, manufacturer *NestedManufacturer) *InventoryItem {
	deviceID := id + 10000
	device := NestedDevice{
		ID:          deviceID,
		URL:         fmt.Sprintf("http://example.host/api/dcim/devices/%d/", deviceID),
		Name:        fmt.Sprintf("Device %d", deviceID),
		DisplayName: fmt.Sprintf("Display Device %d", deviceID),
	}

	itemID := id
	if create {
		itemID = 0
	}
	return &InventoryItem{
		ID:           itemID,
		Device:       device,
		Parent:       nil,
		Name:         fmt.Sprintf("Inventory Item %d", id),
		Manufacturer: manufacturer,
		PartID:       fmt.Sprintf("Part ID %d", id),
		Serial:       fmt.Sprintf("Serial %d", id),
		Discovered:   true,
	}
}
