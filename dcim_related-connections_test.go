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

	got, err := c.DCIM.GetRelatedConnections("device1", "interface1")
	if err != nil {
		t.Fatalf("unexpected error from Client.DCIM.GetRelatedConnections: %v", err)
	}

	if !rcEqual(want, got) {
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
		o    *relatedConnectionsOptions
		want url.Values
	}{
		{
			desc: "normal string length",
			o: &relatedConnectionsOptions{
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
			o: &relatedConnectionsOptions{
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
			o: &relatedConnectionsOptions{
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

		got, err := test.o.Values()
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
		o    *relatedConnectionsOptions
	}{
		{
			desc: "only peer-device",
			o: &relatedConnectionsOptions{
				PeerDevice: "test-abc_def-0123456789",
			},
		},
		{
			desc: "only peer-interface",
			o: &relatedConnectionsOptions{
				PeerInterface: "test-zyx-wvu-987654321",
			},
		},
		{
			desc: "no peer-device or peer-interface",
			o:    &relatedConnectionsOptions{},
		},
	}
	for i, test := range negativeTests {
		t.Logf("[%02d] negative path test %q", i, test.desc)

		if _, err := test.o.Values(); err == nil {
			t.Fatal("expected Error but got nil")
		}
	}
}

func rcEqual(a, b *RelatedConnection) bool {
	if ok := deviceEqual(a.Device, b.Device); !ok {
		return false
	}

	if !reflect.DeepEqual(a.ConsolePorts, b.ConsolePorts) {
		return false
	}

	if !reflect.DeepEqual(a.Interfaces, b.Interfaces) {
		return false
	}

	if !reflect.DeepEqual(a.PowerPorts, b.PowerPorts) {
		return false
	}

	return true
}
