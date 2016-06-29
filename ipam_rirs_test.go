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
	"reflect"
	"testing"
)

func TestClientIPAMGetRIR(t *testing.T) {
	wantRIR := testRIR(1)

	c, done := testClient(t, testHandler(t, http.MethodGet, "/api/ipam/rirs/1", wantRIR))
	defer done()

	gotRIR, err := c.IPAM.GetRIR(wantRIR.ID)
	if err != nil {
		t.Fatalf("unexpected error from Client.IPAM.GetRIR: %v", err)
	}

	if want, got := *wantRIR, *gotRIR; !reflect.DeepEqual(want, got) {
		t.Fatalf("unexpected RIR:\n- want: %v\n-  got: %v", want, got)
	}
}

func TestClientIPAMListRIRs(t *testing.T) {
	wantRIRs := []*RIR{
		testRIR(1),
		testRIR(2),
		testRIR(3),
	}

	c, done := testClient(t, testHandler(t, http.MethodGet, "/api/ipam/rirs/", wantRIRs))
	defer done()

	gotRIRs, err := c.IPAM.ListRIRs()
	if err != nil {
		t.Fatalf("unexpected error from Client.IPAM.ListRIRs: %v", err)
	}

	if want, got := derefRIRs(wantRIRs), derefRIRs(gotRIRs); !reflect.DeepEqual(want, got) {
		t.Fatalf("unexpected RIRs:\n- want: %v\n-  got: %v", want, got)
	}
}

// derefRIRs is used to print values of RIRs in slice, instead of memory addresses.
func derefRIRs(rirs []*RIR) []RIR {
	r := make([]RIR, len(rirs))
	for i := range rirs {
		r[i] = *rirs[i]
	}

	return r
}
