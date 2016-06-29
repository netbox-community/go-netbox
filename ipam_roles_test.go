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

func TestClientIPAMGetRole(t *testing.T) {
	wantRole := testRole(1)

	c, done := testClient(t, testHandler(t, http.MethodGet, "/api/ipam/roles/1", wantRole))
	defer done()

	gotRole, err := c.IPAM.GetRole(wantRole.ID)
	if err != nil {
		t.Fatalf("unexpected error from Client.IPAM.GetRole: %v", err)
	}

	if want, got := *wantRole, *gotRole; !reflect.DeepEqual(want, got) {
		t.Fatalf("unexpected Role:\n- want: %v\n-  got: %v", want, got)
	}
}

func TestClientIPAMListRoles(t *testing.T) {
	wantRoles := []*Role{
		testRole(1),
		testRole(2),
		testRole(3),
	}

	c, done := testClient(t, testHandler(t, http.MethodGet, "/api/ipam/roles/", wantRoles))
	defer done()

	gotRoles, err := c.IPAM.ListRoles()
	if err != nil {
		t.Fatalf("unexpected error from Client.IPAM.ListRoles: %v", err)
	}

	if want, got := derefRoles(wantRoles), derefRoles(gotRoles); !reflect.DeepEqual(want, got) {
		t.Fatalf("unexpected Roles:\n- want: %v\n-  got: %v", want, got)
	}
}

// derefRoles is used to print values of Roles in slice, instead of memory addresses.
func derefRoles(roles []*Role) []Role {
	r := make([]Role, len(roles))
	for i := range roles {
		r[i] = *roles[i]
	}

	return r
}
