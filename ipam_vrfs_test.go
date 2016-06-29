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

func TestClientIPAMGetVRF(t *testing.T) {
	wantVRF := testVRF(1)

	c, done := testClient(t, testHandler(t, http.MethodGet, "/api/ipam/vrfs/1", wantVRF))
	defer done()

	gotVRF, err := c.IPAM.GetVRF(wantVRF.ID)
	if err != nil {
		t.Fatalf("unexpected error from Client.IPAM.GetVRF: %v", err)
	}

	if want, got := *wantVRF, *gotVRF; !reflect.DeepEqual(want, got) {
		t.Fatalf("unexpected VRF:\n- want: %v\n-  got: %v", want, got)
	}
}

func TestClientIPAMListVRFs(t *testing.T) {
	wantVRFs := []*VRF{
		testVRF(1),
		testVRF(2),
		testVRF(3),
	}

	c, done := testClient(t, testHandler(t, http.MethodGet, "/api/ipam/vrfs/", wantVRFs))
	defer done()

	gotVRFs, err := c.IPAM.ListVRFs()
	if err != nil {
		t.Fatalf("unexpected error from Client.IPAM.ListVRFs: %v", err)
	}

	if want, got := derefVRFs(wantVRFs), derefVRFs(gotVRFs); !reflect.DeepEqual(want, got) {
		t.Fatalf("unexpected VRFs:\n- want: %v\n-  got: %v", want, got)
	}
}

// derefVRFs is used to print values of VRFs in slice, instead of memory addresses.
func derefVRFs(vrfs []*VRF) []VRF {
	v := make([]VRF, len(vrfs))
	for i := range vrfs {
		v[i] = *vrfs[i]
	}

	return v
}
