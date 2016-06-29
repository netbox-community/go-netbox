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

func TestClientIPAMGetIPAddress(t *testing.T) {
	wantIP := testIPAddress(FamilyIPv4, 1)

	c, done := testClient(t, testHandler(t, http.MethodGet, "/api/ipam/ip-addresses/1", wantIP))
	defer done()

	gotIP, err := c.IPAM.GetIPAddress(wantIP.ID)
	if err != nil {
		t.Fatalf("unexpected error from Client.IPAM.GetIPAddress: %v", err)
	}

	if want, got := *wantIP, *gotIP; !ipAddressesEqual(want, got) {
		t.Fatalf("unexpected IPAddress:\n- want: %v\n-  got: %v", want, got)
	}
}

func TestClientIPAMListIPAddresses(t *testing.T) {
	wantIPs := []*IPAddress{
		testIPAddress(FamilyIPv4, 1),
		testIPAddress(FamilyIPv6, 2),
	}

	c, done := testClient(t, testHandler(t, http.MethodGet, "/api/ipam/ip-addresses/", wantIPs))
	defer done()

	gotIPs, err := c.IPAM.ListIPAddresses(nil)
	if err != nil {
		t.Fatalf("unexpected error from Client.IPAM.ListIPAddresses: %v", err)
	}

	want := derefIPAddresses(wantIPs)
	got := derefIPAddresses(gotIPs)
	if !ipAddressesSlicesEqual(want, got) {
		t.Fatalf("unexpected IPs:\n- want: %v\n-  got: %v", want, got)
	}
}

func TestListIPAddressesOptionsValues(t *testing.T) {
	var tests = []struct {
		desc string
		o    *ListIPAddressesOptions
		v    url.Values
	}{
		{
			desc: "empty options",
		},
		{
			desc: "family only",
			o: &ListIPAddressesOptions{
				Family: FamilyIPv4,
			},
			v: url.Values{
				"family": []string{strconv.Itoa(int(FamilyIPv4))},
			},
		},
		{
			desc: "1 vrf_id only",
			o: &ListIPAddressesOptions{
				VRFID: []int{1},
			},
			v: url.Values{
				"vrf_id": []string{"1"},
			},
		},
		{
			desc: "3 vrf_ids only",
			o: &ListIPAddressesOptions{
				VRFID: []int{1, 2, 3},
			},
			v: url.Values{
				"vrf_id": []string{"1", "2", "3"},
			},
		},
		{
			desc: "vrf only",
			o: &ListIPAddressesOptions{
				VRF: "vrf",
			},
			v: url.Values{
				"vrf": []string{"vrf"},
			},
		},
		{
			desc: "vrf and vrf_id, vrf_id preferred",
			o: &ListIPAddressesOptions{
				VRF:   "vrf",
				VRFID: []int{1},
			},
			v: url.Values{
				"vrf_id": []string{"1"},
			},
		},
		{
			desc: "1 interface_id only",
			o: &ListIPAddressesOptions{
				InterfaceID: []int{2},
			},
			v: url.Values{
				"interface_id": []string{"2"},
			},
		},
		{
			desc: "3 interface_ids only",
			o: &ListIPAddressesOptions{
				InterfaceID: []int{2, 3, 4},
			},
			v: url.Values{
				"interface_id": []string{"2", "3", "4"},
			},
		},
		{
			desc: "1 device_id only",
			o: &ListIPAddressesOptions{
				DeviceID: []int{3},
			},
			v: url.Values{
				"device_id": []string{"3"},
			},
		},
		{
			desc: "3 device_ids only",
			o: &ListIPAddressesOptions{
				DeviceID: []int{3, 4, 5},
			},
			v: url.Values{
				"device_id": []string{"3", "4", "5"},
			},
		},
		{
			desc: "1 device only",
			o: &ListIPAddressesOptions{
				Device: []string{"device"},
			},
			v: url.Values{
				"device": []string{"device"},
			},
		},
		{
			desc: "3 devices only",
			o: &ListIPAddressesOptions{
				Device: []string{"a", "b", "c"},
			},
			v: url.Values{
				"device": []string{"a", "b", "c"},
			},
		},
		{
			desc: "device and device_id, device_id preferred",
			o: &ListIPAddressesOptions{
				Device:   []string{"device"},
				DeviceID: []int{3},
			},
			v: url.Values{
				"device_id": []string{"3"},
			},
		},
		{
			desc: "q only",
			o: &ListIPAddressesOptions{
				Query: "query",
			},
			v: url.Values{
				"q": []string{"query"},
			},
		},
		{
			desc: "all options",
			o: &ListIPAddressesOptions{
				Family:      FamilyIPv4,
				VRFID:       []int{1},
				VRF:         "vrf",
				InterfaceID: []int{2},
				DeviceID:    []int{3},
				Device:      []string{"device"},
				Query:       "query",
			},
			v: url.Values{
				"family":       []string{strconv.Itoa(int(FamilyIPv4))},
				"vrf_id":       []string{"1"},
				"interface_id": []string{"2"},
				"device_id":    []string{"3"},
				"q":            []string{"query"},
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

func TestIPAddressMarshalJSON(t *testing.T) {
	var tests = []struct {
		desc string
		ip   *IPAddress
		b    []byte
	}{
		{
			desc: "IPv4 address",
			ip:   testIPAddress(FamilyIPv4, 1),
			b:    []byte(`{"id":1,"family":4,"address":"8.8.8.0/24","vrf":{"id":1,"name":"VRFIdentifier 1","rd":"rd 1"},"interface":{"id":1,"device":{"id":1,"name":"DeviceIdentifier 1"},"name":"InterfaceIdentifier 1"},"description":"description 1","nat_inside":{"id":1,"family":4,"address":"8.8.8.0/24"},"nat_outside":{"id":1,"family":4,"address":"8.8.8.0/24"}}`),
		},
		{
			desc: "IPv6 address",
			ip:   testIPAddress(FamilyIPv6, 2),
			b:    []byte(`{"id":2,"family":6,"address":"2001:4860:4860::/48","vrf":{"id":2,"name":"VRFIdentifier 2","rd":"rd 2"},"interface":{"id":2,"device":{"id":2,"name":"DeviceIdentifier 2"},"name":"InterfaceIdentifier 2"},"description":"description 2","nat_inside":{"id":2,"family":6,"address":"2001:4860:4860::/48"},"nat_outside":{"id":2,"family":6,"address":"2001:4860:4860::/48"}}`),
		},
	}

	for i, tt := range tests {
		t.Logf("[%02d] test %q", i, tt.desc)

		b, err := json.Marshal(tt.ip)
		if err != nil {
			t.Fatalf("unexpected JSON marshal error: %v", err)
		}

		if want, got := tt.b, b; !bytes.Equal(want, got) {
			t.Fatalf("unexpected JSON bytes:\n- want: %v\n-  got: %v",
				string(want), string(got))
		}
	}
}

func TestIPAddressUnmarshalJSON(t *testing.T) {
	var tests = []struct {
		desc string
		b    []byte
		ip   *IPAddress
		err  error
	}{
		{
			desc: "invalid IP address",
			b:    []byte(`{"address":"foo"}`),
			err: &net.ParseError{
				Type: "CIDR address",
				Text: "foo",
			},
		},
		{
			desc: "IPv4 address",
			b:    []byte(`{"id":1,"family":4,"address":"8.8.8.8/24","vrf":{"id":1,"name":"VRFIdentifier 1","rd":"rd 1"},"interface":{"id":1,"device":{"id":1,"name":"DeviceIdentifier 1"},"name":"InterfaceIdentifier 1"}}`),
			ip:   testIPAddress(FamilyIPv4, 1),
		},
		{
			desc: "IPv6 address",
			b:    []byte(`{"id":2,"family":6,"address":"2001:4860:4860::8888/48","vrf":{"id":2,"name":"VRFIdentifier 2","rd":"rd 2"},"interface":{"id":2,"device":{"id":2,"name":"DeviceIdentifier 2"},"name":"InterfaceIdentifier 2"}}`),
			ip:   testIPAddress(FamilyIPv6, 2),
		},
	}

	for i, tt := range tests {
		t.Logf("[%02d] test %q", i, tt.desc)

		ip := new(IPAddress)
		err := json.Unmarshal(tt.b, ip)

		if want, got := tt.err, err; !reflect.DeepEqual(want, got) {
			t.Fatalf("unexpected error:\n- want: %v\n-  got: %v",
				want, got)
		}
		if err != nil {
			continue
		}

		if want, got := *tt.ip, *ip; !ipAddressesEqual(want, got) {
			t.Fatalf("unexpected IPAddress:\n- want: %v\n-  got: %v",
				want, got)
		}
	}
}

func TestIPAddressIdentifierMarshalJSON(t *testing.T) {
	var tests = []struct {
		desc string
		ip   *IPAddressIdentifier
		b    []byte
	}{
		{
			desc: "IPv4 address",
			ip:   testIPAddressIdentifier(FamilyIPv4, 1),
			b:    []byte(`{"id":1,"family":4,"address":"8.8.8.0/24"}`),
		},
		{
			desc: "IPv6 address",
			ip:   testIPAddressIdentifier(FamilyIPv6, 2),
			b:    []byte(`{"id":2,"family":6,"address":"2001:4860:4860::/48"}`),
		},
	}

	for i, tt := range tests {
		t.Logf("[%02d] test %q", i, tt.desc)

		b, err := json.Marshal(tt.ip)
		if err != nil {
			t.Fatalf("unexpected JSON marshal error: %v", err)
		}

		if want, got := tt.b, b; !bytes.Equal(want, got) {
			t.Fatalf("unexpected JSON bytes:\n- want: %v\n-  got: %v",
				string(want), string(got))
		}
	}
}

func TestIPAddressIdentifierUnmarshalJSON(t *testing.T) {
	var tests = []struct {
		desc string
		b    []byte
		ip   *IPAddressIdentifier
		err  error
	}{
		{
			desc: "invalid IP address",
			b:    []byte(`{"address":"foo"}`),
			err: &net.ParseError{
				Type: "CIDR address",
				Text: "foo",
			},
		},
		{
			desc: "IPv4 address",
			b:    []byte(`{"id":1,"family":4,"address":"8.8.8.8/24"}`),
			ip:   testIPAddressIdentifier(FamilyIPv4, 1),
		},
		{
			desc: "IPv6 address",
			b:    []byte(`{"id":2,"family":6,"address":"2001:4860:4860::8888/48"}`),
			ip:   testIPAddressIdentifier(FamilyIPv6, 2),
		},
	}

	for i, tt := range tests {
		t.Logf("[%02d] test %q", i, tt.desc)

		ip := new(IPAddressIdentifier)
		err := json.Unmarshal(tt.b, ip)

		if want, got := tt.err, err; !reflect.DeepEqual(want, got) {
			t.Fatalf("unexpected error:\n- want: %v\n-  got: %v",
				want, got)
		}
		if err != nil {
			continue
		}

		if want, got := *tt.ip, *ip; !ipAddressIdentifiersEqual(want, got) {
			t.Fatalf("unexpected IPAddressIdentifier:\n- want: %v\n-  got: %v",
				want, got)
		}
	}
}

func ipAddressesSlicesEqual(a, b []IPAddress) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !ipAddressesEqual(a[i], b[i]) {
			return false
		}
	}

	return true
}

func ipAddressesEqual(a, b IPAddress) bool {
	if a.ID != b.ID {
		return false
	}

	if a.Family != b.Family {
		return false
	}

	if !a.Address.IP.Equal(b.Address.IP) {
		return false
	}

	if a.Address.String() != b.Address.String() {
		return false
	}

	if !reflect.DeepEqual(a.VRF, b.VRF) {
		return false
	}

	if !reflect.DeepEqual(a.Interface, b.Interface) {
		return false
	}

	return true
}

func ipAddressIdentifiersEqual(a, b IPAddressIdentifier) bool {
	if a.ID != b.ID {
		return false
	}

	if a.Family != b.Family {
		return false
	}

	if !a.Address.IP.Equal(b.Address.IP) {
		return false
	}

	return true
}

// Used to print values of IPAddresses in slice, instead of memory addresses.
func derefIPAddresses(ips []*IPAddress) []IPAddress {
	ip := make([]IPAddress, len(ips))
	for i := range ips {
		ip[i] = *ips[i]
	}

	return ip
}
