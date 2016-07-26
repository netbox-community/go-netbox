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
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestGetRelatedConnection(t *testing.T) {
	want := testRelatedConnection(1)

	c, done := testClient(
		t,
		testHandler(t, http.MethodGet, "/api/dcim/related-connections/", want),
	)
	defer done()

	got, err := c.DCIM.GetRelatedConnections(
		&RelatedConnectionsOptions{
			PeerDevice:    "device1",
			PeerInterface: "interface1",
		},
	)
	if err != nil {
		t.Fatalf("unexpected error from Client.DCIM.GetRelatedConnections: %v", err)
	}

	if !rCEqual(want, got) {
		t.Fatalf(
			"unexpected related-connections payload:\n- want: %v\n- got: %v",
			*want,
			*got,
		)
	}
}

func TestGetRelatedConnectionWithOptions(t *testing.T) {
	var happyPathTests = []struct {
		desc string
		o    *RelatedConnectionsOptions
		want url.Values
	}{
		{
			desc: "normal string length",
			o: &RelatedConnectionsOptions{
				PeerDevice:    "test-abc_def-0123456789",
				PeerInterface: "test-zyx-wvu-987654321",
			},
			want: url.Values{
				"peer-device":    []string{"test-abc_def-0123456789"},
				"peer-interface": []string{"test-zyx-wvu-987654321"},
			},
		},
		{
			desc: "super long string",
			o: &RelatedConnectionsOptions{
				PeerDevice:    "test-abc_def-01234567890000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
				PeerInterface: "test-zyx-wvu-98765432100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
			},
			want: url.Values{
				"peer-device":    []string{"test-abc_def-01234567890000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"},
				"peer-interface": []string{"test-zyx-wvu-98765432100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"},
			},
		},
		{
			desc: "super short string",
			o: &RelatedConnectionsOptions{
				PeerDevice:    "a",
				PeerInterface: "z",
			},
			want: url.Values{
				"peer-device":    []string{"a"},
				"peer-interface": []string{"z"},
			},
		},
	}
	for i, test := range happyPathTests {
		t.Logf("[%02d] happy path test %q", i, test.desc)

		got, err := test.o.values()
		if err != nil {
			t.Fatalf("unexpected Values error: %v", err)
		}

		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf(
				"unexpected url.Values map:\n- want: %v\n-  got: %v",
				test.want,
				got,
			)
		}
	}

	var negativeTests = []struct {
		desc string
		o    *RelatedConnectionsOptions
	}{
		{
			desc: "only peer-device",
			o: &RelatedConnectionsOptions{
				PeerDevice: "test-abc_def-0123456789",
			},
		},
		{
			desc: "only peer-interface",
			o: &RelatedConnectionsOptions{
				PeerInterface: "test-zyx-wvu-987654321",
			},
		},
		{
			desc: "no peer-device or peer-interface",
			o:    &RelatedConnectionsOptions{},
		},
	}
	for i, test := range negativeTests {
		t.Logf("[%02d] happy path test %q", i, test.desc)

		if _, err := test.o.values(); err == nil {
			t.Fatal("expected Error but got nil")
		}
	}
}

func rCEqual(i, j *RelatedConnection) bool {
	if ok := deviceEqual(i.Device, j.Device); !ok {
		return false
	}

	if len(i.ConsolePorts) != len(j.ConsolePorts) {
		return false
	}
	for n := range i.ConsolePorts {
		if i.ConsolePorts[n].Name != j.ConsolePorts[n].Name {
			return false
		}
		if i.ConsolePorts[n].ConsoleServer != j.ConsolePorts[n].ConsoleServer {
			return false
		}
		if i.ConsolePorts[n].Port != j.ConsolePorts[n].Port {
			return false
		}
	}

	if len(i.Interfaces) != len(j.Interfaces) {
		return false
	}
	for n := range i.Interfaces {
		if i.Interfaces[n].Name != j.Interfaces[n].Name {
			return false
		}
		if i.Interfaces[n].Device != j.Interfaces[n].Device {
			return false
		}
		if i.Interfaces[n].Interface != j.Interfaces[n].Interface {
			return false
		}
	}

	if len(i.PowerPorts) != len(j.PowerPorts) {
		return false
	}
	for n := range i.PowerPorts {
		if i.PowerPorts[n].Name != j.PowerPorts[n].Name {
			return false
		}
		if i.PowerPorts[n].PDU != j.PowerPorts[n].PDU {
			return false
		}
		if i.PowerPorts[n].Outlet != j.PowerPorts[n].Outlet {
			return false
		}
	}
	return true
}
