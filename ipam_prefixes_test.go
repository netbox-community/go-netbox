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
	"bytes"
	"encoding/json"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"testing"
)

func TestClientIPAMGetPrefix(t *testing.T) {
	wantPrefix := testPrefix(FamilyIPv4, 1)

	c, done := testClient(t, testHandler(t, http.MethodGet, "/api/ipam/prefixes/1", wantPrefix))
	defer done()

	gotPrefix, err := c.IPAM.GetPrefix(wantPrefix.ID)
	if err != nil {
		t.Fatalf("unexpected error from Client.IPAM.GetPrefix: %v", err)
	}

	if want, got := *wantPrefix, *gotPrefix; !prefixesEqual(want, got) {
		t.Fatalf("unexpected Prefix:\n- want: %v\n-  got: %v", want, got)
	}
}

func TestClientIPAMListPrefixes(t *testing.T) {
	wantPrefixes := []*Prefix{
		testPrefix(FamilyIPv4, 1),
		testPrefix(FamilyIPv6, 2),
	}

	c, done := testClient(t, testHandler(t, http.MethodGet, "/api/ipam/prefixes/", wantPrefixes))
	defer done()

	gotPrefixes, err := c.IPAM.ListPrefixes(nil)
	if err != nil {
		t.Fatalf("unexpected error from Client.IPAM.ListPrefixes: %v", err)
	}

	want := derefPrefixes(wantPrefixes)
	got := derefPrefixes(gotPrefixes)
	if !prefixSlicesEqual(want, got) {
		t.Fatalf("unexpected Prefixes:\n- want: %v\n-  got: %v", want, got)
	}
}

func TestListPrefixesOptionsValues(t *testing.T) {
	var tests = []struct {
		desc string
		o    *ListPrefixesOptions
		v    url.Values
	}{
		{
			desc: "empty options",
		},
		{
			desc: "family only",
			o: &ListPrefixesOptions{
				Family: FamilyIPv4,
			},
			v: url.Values{
				"family": []string{strconv.Itoa(int(FamilyIPv4))},
			},
		},
		{
			desc: "1 site_id only",
			o: &ListPrefixesOptions{
				SiteID: []int{1},
			},
			v: url.Values{
				"site_id": []string{"1"},
			},
		},
		{
			desc: "3 site_ids only",
			o: &ListPrefixesOptions{
				SiteID: []int{1, 2, 3},
			},
			v: url.Values{
				"site_id": []string{"1", "2", "3"},
			},
		},
		{
			desc: "1 site only",
			o: &ListPrefixesOptions{
				Site: []string{"site"},
			},
			v: url.Values{
				"site": []string{"site"},
			},
		},
		{
			desc: "3 sites only",
			o: &ListPrefixesOptions{
				Site: []string{"sitefoo", "sitebar", "sitebaz"},
			},
			v: url.Values{
				"site": []string{"sitefoo", "sitebar", "sitebaz"},
			},
		},
		{
			desc: "site and site_id, site_id preferred",
			o: &ListPrefixesOptions{
				Site:   []string{"site"},
				SiteID: []int{1},
			},
			v: url.Values{
				"site_id": []string{"1"},
			},
		},
		{
			desc: "vrf_id only",
			o: &ListPrefixesOptions{
				VRFID: 2,
			},
			v: url.Values{
				"vrf_id": []string{"2"},
			},
		},
		{
			desc: "vrf only",
			o: &ListPrefixesOptions{
				VRF: "vrf",
			},
			v: url.Values{
				"vrf": []string{"vrf"},
			},
		},
		{
			desc: "vrf and vrf_id, vrf_id preferred",
			o: &ListPrefixesOptions{
				VRF:   "vrf",
				VRFID: 2,
			},
			v: url.Values{
				"vrf_id": []string{"2"},
			},
		},
		{
			desc: "1 vlan_id only",
			o: &ListPrefixesOptions{
				VLANID: []int{3},
			},
			v: url.Values{
				"vlan_id": []string{"3"},
			},
		},
		{
			desc: "3 vlan_ids only",
			o: &ListPrefixesOptions{
				VLANID: []int{3, 4, 5},
			},
			v: url.Values{
				"vlan_id": []string{"3", "4", "5"},
			},
		},
		{
			desc: "vlan_vid only",
			o: &ListPrefixesOptions{
				VLANVID: 4094,
			},
			v: url.Values{
				"vlan_vid": []string{"4094"},
			},
		},
		{
			desc: "vlan and vlan_id, vlan_id preferred",
			o: &ListPrefixesOptions{
				VLANVID: 4094,
				VLANID:  []int{3},
			},
			v: url.Values{
				"vlan_id": []string{"3"},
			},
		},
		{
			desc: "status only",
			o: &ListPrefixesOptions{
				Status: 4,
			},
			v: url.Values{
				"status": []string{"4"},
			},
		},
		{
			desc: "1 role_id only",
			o: &ListPrefixesOptions{
				RoleID: []int{5},
			},
			v: url.Values{
				"role_id": []string{"5"},
			},
		},
		{
			desc: "3 role_ids only",
			o: &ListPrefixesOptions{
				RoleID: []int{5, 6, 7},
			},
			v: url.Values{
				"role_id": []string{"5", "6", "7"},
			},
		},
		{
			desc: "1 role only",
			o: &ListPrefixesOptions{
				Role: []string{"role"},
			},
			v: url.Values{
				"role": []string{"role"},
			},
		},
		{
			desc: "3 roles only",
			o: &ListPrefixesOptions{
				Role: []string{"rolefoo", "rolebar", "rolebaz"},
			},
			v: url.Values{
				"role": []string{"rolefoo", "rolebar", "rolebaz"},
			},
		},
		{
			desc: "role and role_id, role_id preferred",
			o: &ListPrefixesOptions{
				Role:   []string{"role"},
				RoleID: []int{5},
			},
			v: url.Values{
				"role_id": []string{"5"},
			},
		},
		{
			desc: "parent only",
			o: &ListPrefixesOptions{
				Parent: &net.IPNet{
					IP:   net.ParseIP("::1"),
					Mask: net.CIDRMask(128, 128),
				},
			},
			v: url.Values{
				"parent": []string{"::1/128"},
			},
		},
		{
			desc: "q only",
			o: &ListPrefixesOptions{
				Query: "query",
			},
			v: url.Values{
				"q": []string{"query"},
			},
		},
		{
			desc: "all options",
			o: &ListPrefixesOptions{
				Family:  FamilyIPv4,
				SiteID:  []int{1},
				Site:    []string{"site"},
				VRFID:   2,
				VRF:     "vrf",
				VLANID:  []int{3},
				VLANVID: 4094,
				Status:  4,
				RoleID:  []int{5},
				Role:    []string{"role"},
				Parent: &net.IPNet{
					IP:   net.ParseIP("::1"),
					Mask: net.CIDRMask(128, 128),
				},
				Query: "query",
			},
			v: url.Values{
				"family":  []string{strconv.Itoa(int(FamilyIPv4))},
				"site_id": []string{"1"},
				"vrf_id":  []string{"2"},
				"vlan_id": []string{"3"},
				"status":  []string{"4"},
				"role_id": []string{"5"},
				"parent":  []string{"::1/128"},
				"q":       []string{"query"},
			},
		},
	}

	for i, tt := range tests {
		t.Logf("[%02d] test %q", i, tt.desc)

		v, err := tt.o.values()
		if err != nil {
			t.Fatalf("unexpected Values error: %v", err)
		}

		if want, got := tt.v, v; !reflect.DeepEqual(want, got) {
			t.Fatalf("unexpected url.Values map:\n- want: %v\n-  got: %v",
				want, got)
		}
	}
}

func TestPrefixMarshalJSON(t *testing.T) {
	var tests = []struct {
		desc string
		p    *Prefix
		b    []byte
	}{
		{
			desc: "IPv4 prefix",
			p:    testPrefix(FamilyIPv4, 1),
			b:    []byte(`{"id":1,"family":4,"prefix":"8.8.0.0/16","site":{"id":1,"name":"SiteIdentifier 1","slug":"siteidentifier1"},"vrf":{"id":1,"name":"VRFIdentifier 1","rd":"rd 1"},"vlan":{"id":1,"vid":1,"name":"vlanidentifier 1","display_name":"VLANIdentifier 1"},"status":1,"role":{"id":1,"name":"RoleIdentifier 1","slug":"roleidentifier1"},"description":"description 1"}`),
		},
		{
			desc: "IPv6 prefix",
			p:    testPrefix(FamilyIPv6, 2),
			b:    []byte(`{"id":2,"family":6,"prefix":"2001:4860::/32","site":{"id":2,"name":"SiteIdentifier 2","slug":"siteidentifier2"},"vrf":{"id":2,"name":"VRFIdentifier 2","rd":"rd 2"},"vlan":{"id":2,"vid":2,"name":"vlanidentifier 2","display_name":"VLANIdentifier 2"},"status":1,"role":{"id":2,"name":"RoleIdentifier 2","slug":"roleidentifier2"},"description":"description 2"}`),
		},
	}

	for i, tt := range tests {
		t.Logf("[%02d] test %q", i, tt.desc)

		b, err := json.Marshal(tt.p)
		if err != nil {
			t.Fatalf("unexpected JSON marshal error: %v", err)
		}

		if want, got := tt.b, b; !bytes.Equal(want, got) {
			t.Fatalf("unexpected JSON bytes:\n- want: %v\n-  got: %v",
				string(want), string(got))
		}
	}
}

func TestPrefixUnmarshalJSON(t *testing.T) {
	var tests = []struct {
		desc string
		b    []byte
		p    *Prefix
		err  error
	}{
		{
			desc: "invalid prefix",
			b:    []byte(`{"prefix":"foo"}`),
			err: &net.ParseError{
				Type: "CIDR address",
				Text: "foo",
			},
		},
		{
			desc: "IPv4 prefix",
			b:    []byte(`{"id":1,"family":4,"prefix":"8.8.0.0/16","site":{"id":1,"name":"SiteIdentifier 1","slug":"siteidentifier1"},"vrf":{"id":1,"name":"VRFIdentifier 1","rd":"rd 1"},"vlan":{"id":1,"vid":1,"name":"vlanidentifier 1","display_name":"VLANIdentifier 1"},"status":1,"role":{"id":1,"name":"RoleIdentifier 1","slug":"roleidentifier1"},"description":"description 1"}`),
			p:    testPrefix(FamilyIPv4, 1),
		},
		{
			desc: "IPv6 prefix",
			b:    []byte(`{"id":2,"family":6,"prefix":"2001:4860::/32","site":{"id":2,"name":"SiteIdentifier 2","slug":"siteidentifier2"},"vrf":{"id":2,"name":"VRFIdentifier 2","rd":"rd 2"},"vlan":{"id":2,"vid":2,"name":"vlanidentifier 2","display_name":"VLANIdentifier 2"},"status":1,"role":{"id":2,"name":"RoleIdentifier 2","slug":"roleidentifier2"},"description":"description 2"}`),
			p:    testPrefix(FamilyIPv6, 2),
		},
	}

	for i, tt := range tests {
		t.Logf("[%02d] test %q", i, tt.desc)

		p := new(Prefix)
		err := json.Unmarshal(tt.b, p)

		if want, got := tt.err, err; !reflect.DeepEqual(want, got) {
			t.Fatalf("unexpected error:\n- want: %v\n-  got: %v",
				want, got)
		}
		if err != nil {
			continue
		}

		if want, got := *tt.p, *p; !prefixesEqual(want, got) {
			t.Fatalf("unexpected Prefix:\n- want: %v\n-  got: %v",
				want, got)
		}
	}
}

func prefixSlicesEqual(a, b []Prefix) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !prefixesEqual(a[i], b[i]) {
			return false
		}
	}

	return true
}

func prefixesEqual(a, b Prefix) bool {
	if a.ID != b.ID {
		return false
	}

	if a.Family != b.Family {
		return false
	}

	if a.Prefix.String() != b.Prefix.String() {
		return false
	}

	if !reflect.DeepEqual(a.Site, b.Site) {
		return false
	}

	if !reflect.DeepEqual(a.VRF, b.VRF) {
		return false
	}

	if !reflect.DeepEqual(a.VLAN, b.VLAN) {
		return false
	}

	if a.Status != b.Status {
		return false
	}

	if !reflect.DeepEqual(a.Role, b.Role) {
		return false
	}

	if a.Description != b.Description {
		return false
	}

	return true
}

// Used to print values of Prefixs in slice, instead of memory addresses.
func derefPrefixes(prefixes []*Prefix) []Prefix {
	p := make([]Prefix, len(prefixes))
	for i := range prefixes {
		p[i] = *prefixes[i]
	}

	return p
}
