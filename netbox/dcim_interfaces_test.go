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
	"reflect"
	"testing"
)

func TestInterfaceGet(t *testing.T) {
	var tests = []struct {
		desc string
		want *Interface
	}{
		{
			desc: "Maximum interface",
			want: testInterface(1),
		},
	}

	for idx, tt := range tests {
		t.Run(fmt.Sprintf("[%d] %s", idx, tt.desc), func(t *testing.T) {
			serverData := serverDataInterface(*tt.want)

			c, done := testClient(t, testHandler(t, http.MethodGet, "/api/dcim/interfaces/1/", &serverData))
			defer done()

			res, err := c.DCIM.Interfaces.Get(1)
			if err != nil {
				t.Fatalf("unexpected error from Client.DCIM.Interfaces.Get: %v", err)
			}

			if want, got := tt.want, res; !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected Interface\n= want: %v\n-got: %v", want, got)
			}
		})
	}
}

func TestInterfaceUnmarshalJSON(t *testing.T) {
	var tests = []struct {
		desc string
		data []byte
		want *Interface
	}{
		{
			desc: "Maximum interface",
			data: []byte(`{"id": 1, "device": { "id": 101, "url": "http://localhost/api/dcim/devices/101/", "name": "Device 101", "display_name": "Display Device 101" }, "name": "Interface 1", "form_factor": {"value": 1000, "label": "1000BASE-T"}, "mgmt_only": false, "connection": { "id": 201, "url": "http://localhost/api/dcim/interface-connections/201/", "connection_status": true}, "connected_interface": { "id": 301, "url": "http://localhost/api/dcim/interfaces/301/", "device": {"id": 401, "url": "http://localhost/api/dcim/devices/401/", "name": "Device 401", "display_name": "Display Device 401"}, "name":"Interface 301", "form_factor": {"value": 1000, "label": "1000BASE-T"}, "lag": {"id": 501, "url": "http://localhost/api/dcim/interfaces/501/", "name":"Interface 501" }, "mgmt_only": false } }`),
			want: testInterface(1),
		},
	}

	for idx, tt := range tests {
		t.Run(fmt.Sprintf("[%d] %s", idx, tt.desc), func(t *testing.T) {
			res := new(Interface)
			err := json.Unmarshal(tt.data, res)
			if err != nil {
				t.Fatalf("unexpected error from Interface.UnmarshalJSON: %v", err)
			}

			if want, got := tt.want, res; !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected Interface:\n- want: %v\n- got: %v", want, got)
			}
		})
	}
}

func TestInterfaceMarshalJSON(t *testing.T) {
	var tests = []struct {
		desc string
		data *Interface
		want []byte
	}{
		{
			desc: "Maximum interface",
			data: testInterface(1),
			want: []byte(`{"id":1,"device":101,"name":"Interface 1","form_factor":1000,"mgmt_only":false,"connection":201,"connected_interface":301}`),
		},
	}

	for idx, tt := range tests {
		t.Run(fmt.Sprintf("[%d] %s", idx, tt.desc), func(t *testing.T) {
			res, err := json.Marshal(tt.data)
			if err != nil {
				t.Fatalf("unexpected error from Interface.MarshalJSON: %v", err)
			}

			if want, got := tt.want, res; bytes.Compare(want, got) != 0 {
				t.Fatalf("unexpected JSON:\n- want: %v\n- got: %v", string(want), string(got))
			}
		})
	}
}

func testNestedDevice(id int) NestedDevice {
	return NestedDevice{
		ID:          id,
		URL:         fmt.Sprintf("http://localhost/api/dcim/devices/%d/", id),
		Name:        fmt.Sprintf("Device %d", id),
		DisplayName: fmt.Sprintf("Display Device %d", id),
	}
}

func testShortNestedInterface(id int) *ShortNestedInterface {
	return &ShortNestedInterface{
		ID:   id,
		URL:  fmt.Sprintf("http://localhost/api/dcim/interfaces/%d/", id),
		Name: fmt.Sprintf("Interface %d", id),
	}
}

func testNestedInterfaceConnection(id int) *NestedInterfaceConnection {
	return &NestedInterfaceConnection{
		ID:     id,
		URL:    fmt.Sprintf("http://localhost/api/dcim/interface-connections/%d/", id),
		Status: true,
	}
}

func testFormFactor() FormFactor {
	return FormFactor{
		Value: 1000,
		Label: "1000BASE-T",
	}
}

func testNestedInterface(id int) *NestedInterface {
	deviceID := id + 100
	lagID := id + 200
	return &NestedInterface{
		ID:         id,
		URL:        fmt.Sprintf("http://localhost/api/dcim/interfaces/%d/", id),
		Device:     testNestedDevice(deviceID),
		Name:       fmt.Sprintf("Interface %d", id),
		FormFactor: testFormFactor(),
		LAG:        testShortNestedInterface(lagID),
		MGMT:       false,
	}
}

func testInterface(id int) *Interface {
	return testInterfaceHelper(id, false)
}

func testInterfaceCreate(id int) *Interface {
	return testInterfaceHelper(id, true)
}

func testInterfaceHelper(id int, create bool) *Interface {
	interfaceID := id
	if create {
		interfaceID = 0
	}

	nestedDeviceID := id + 100
	connectionID := id + 200
	connectedInterfaceID := id + 300
	return &Interface{
		ID:                 interfaceID,
		Device:             testNestedDevice(nestedDeviceID),
		Name:               fmt.Sprintf("Interface %d", id),
		FormFactor:         testFormFactor(),
		MGMT:               false,
		Connection:         testNestedInterfaceConnection(connectionID),
		ConnectedInterface: testNestedInterface(connectedInterfaceID),
	}
}
