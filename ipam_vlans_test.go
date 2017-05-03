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
	"math"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestVLANIDValid(t *testing.T) {
	var tests = []struct {
		id VLANID
		ok bool
	}{
		{
			id: math.MinInt64,
		},
		{
			id: -1,
		},
		{
			id: 4097,
		},
		{
			id: math.MaxInt64,
		},
		{
			id: 0,
			ok: true,
		},
		{
			id: 4096,
			ok: true,
		},
	}

	for _, tt := range tests {
		if want, got := tt.ok, tt.id.Valid(); want != got {
			t.Fatalf("unexpected VLANID(%d).Valid():\n- want: %v\n-  got: %v",
				tt.id, want, got)
		}
	}
}

func TestClientIPAMGetVLAN(t *testing.T) {
	wantVLAN := testVLAN(1)

	c, done := testClient(t, testHandler(t, http.MethodGet, "/api/ipam/vlans/1", wantVLAN))
	defer done()

	gotVLAN, err := c.IPAM.GetVLAN(wantVLAN.ID)
	if err != nil {
		t.Fatalf("unexpected error from Client.IPAM.GetVLAN: %v", err)
	}

	if want, got := *wantVLAN, *gotVLAN; !reflect.DeepEqual(want, got) {
		t.Fatalf("unexpected VLAN:\n- want: %v\n-  got: %v", want, got)
	}
}

func TestClientIPAMListVLANs(t *testing.T) {
	wantVLANs := []*VLAN{
		testVLAN(1),
		testVLAN(2),
		testVLAN(3),
	}

	c, done := testClient(t, testHandler(t, http.MethodGet, "/api/ipam/vlans/", wantVLANs))
	defer done()

	gotVLANs, err := c.IPAM.ListVLANs(nil)
	if err != nil {
		t.Fatalf("unexpected error from Client.IPAM.ListVLANs: %v", err)
	}

	if want, got := derefVLANs(wantVLANs), derefVLANs(gotVLANs); !reflect.DeepEqual(want, got) {
		t.Fatalf("unexpected VLANs:\n- want: %v\n-  got: %v", want, got)
	}
}

func TestListVLANsOptionsValues(t *testing.T) {
	var tests = []struct {
		desc string
		o    *ListVLANsOptions
		v    url.Values
	}{
		{
			desc: "empty options",
		},
		{
			desc: "name only",
			o: &ListVLANsOptions{
				Name: "name",
			},
			v: url.Values{
				"name": []string{"name"},
			},
		},
		{
			desc: "1 role_id only",
			o: &ListVLANsOptions{
				RoleID: []int{1},
			},
			v: url.Values{
				"role_id": []string{"1"},
			},
		},
		{
			desc: "3 role_ids only",
			o: &ListVLANsOptions{
				RoleID: []int{1, 2, 3},
			},
			v: url.Values{
				"role_id": []string{"1", "2", "3"},
			},
		},
		{
			desc: "1 role only",
			o: &ListVLANsOptions{
				Role: []string{"role"},
			},
			v: url.Values{
				"role": []string{"role"},
			},
		},
		{
			desc: "3 roles only",
			o: &ListVLANsOptions{
				Role: []string{"rolefoo", "rolebar", "rolebaz"},
			},
			v: url.Values{
				"role": []string{"rolefoo", "rolebar", "rolebaz"},
			},
		},
		{
			desc: "role and role_id, role_id preferred",
			o: &ListVLANsOptions{
				Role:   []string{"role"},
				RoleID: []int{1},
			},
			v: url.Values{
				"role_id": []string{"1"},
			},
		},
		{
			desc: "1 site_id only",
			o: &ListVLANsOptions{
				SiteID: []int{2},
			},
			v: url.Values{
				"site_id": []string{"2"},
			},
		},
		{
			desc: "3 site_ids only",
			o: &ListVLANsOptions{
				SiteID: []int{2, 3, 4},
			},
			v: url.Values{
				"site_id": []string{"2", "3", "4"},
			},
		},
		{
			desc: "1 site only",
			o: &ListVLANsOptions{
				Site: []string{"site"},
			},
			v: url.Values{
				"site": []string{"site"},
			},
		},
		{
			desc: "3 sites only",
			o: &ListVLANsOptions{
				Site: []string{"sitefoo", "sitebar", "sitebaz"},
			},
			v: url.Values{
				"site": []string{"sitefoo", "sitebar", "sitebaz"},
			},
		},
		{
			desc: "site and site_id, site_id preferred",
			o: &ListVLANsOptions{
				Site:   []string{"site"},
				SiteID: []int{2},
			},
			v: url.Values{
				"site_id": []string{"2"},
			},
		},
		{
			desc: "status_id only",
			o: &ListVLANsOptions{
				StatusID: 3,
			},
			v: url.Values{
				"status_id": []string{"3"},
			},
		},
		{
			desc: "1 status only",
			o: &ListVLANsOptions{
				Status: []string{"status"},
			},
			v: url.Values{
				"status": []string{"status"},
			},
		},
		{
			desc: "3 statuses only",
			o: &ListVLANsOptions{
				Status: []string{"statusfoo", "statusbar", "statusbaz"},
			},
			v: url.Values{
				"status": []string{"statusfoo", "statusbar", "statusbaz"},
			},
		},
		{
			desc: "status and status_id, status_id preferred",
			o: &ListVLANsOptions{
				Status:   []string{"status"},
				StatusID: 3,
			},
			v: url.Values{
				"status_id": []string{"3"},
			},
		},
		{
			desc: "vid only",
			o: &ListVLANsOptions{
				VID: 4,
			},
			v: url.Values{
				"vid": []string{"4"},
			},
		},
		{
			desc: "all options",
			o: &ListVLANsOptions{
				Name:     "name",
				Role:     []string{"role"},
				RoleID:   []int{1},
				Site:     []string{"site"},
				SiteID:   []int{2},
				Status:   []string{"status"},
				StatusID: 3,
				VID:      4,
			},
			v: url.Values{
				"name":      []string{"name"},
				"role_id":   []string{"1"},
				"site_id":   []string{"2"},
				"status_id": []string{"3"},
				"vid":       []string{"4"},
			},
		},
	}

	for i, tt := range tests {
		t.Logf("[%02d] test %q", i, tt.desc)

		v, err := tt.o.Values()
		if err != nil {
			t.Fatalf("unexpected Values error: %v", err)
		}

		if want, got := tt.v, v; !reflect.DeepEqual(want, got) {
			t.Fatalf("unexpected url.Values map:\n- want: %v\n-  got: %v",
				want, got)
		}
	}
}

// Used to print values of VLANs in slice, instead of memory addresses.
func derefVLANs(vlans []*VLAN) []VLAN {
	v := make([]VLAN, len(vlans))
	for i := range vlans {
		v[i] = *vlans[i]
	}

	return v
}
