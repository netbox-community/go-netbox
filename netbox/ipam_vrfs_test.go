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

func TestVRFGet(t *testing.T) {
	var tests = []struct {
		desc string
		want *VRF
	}{
		{
			desc: "Without Tenant",
			want: testVRFWithTenant(1, nil),
		},
		{
			desc: "With Tenant",
			want: testVRFWithTenant(1, testTenantWithGroup(1, nil)),
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("[%d] %s", i, tt.desc), func(t *testing.T) {
			serverData := serverDataVRF(*tt.want)

			c, done := testClient(t, testHandler(t, http.MethodGet, "/api/ipam/vrfs/1/", &serverData))
			defer done()

			res, err := c.IPAM.VRFs.Get(1)
			if err != nil {
				t.Fatalf("unexpected error from Client.IPAM.VRFs.Get: %v", err)
			}

			if want, got := tt.want, res; !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected Tenant:\n- want: %v\n-  got: %v", want, got)
			}
		})
	}
}

func TestVRFUnmarshalJSON(t *testing.T) {
	var tests = []struct {
		desc string
		data []byte
		want *VRF
	}{
		{
			desc: "Nil Group",
			data: []byte(`{ "id": 1, "name": "VRF 1", "rd": "vrf-1", "tenant": null, "enforce_unique": true, "description": "VRF 1 Description", "custom_fields": {} }`),
			want: testVRFWithTenant(1, nil),
		},
		{
			desc: "With Group",
			data: []byte(`{ "id": 1, "name": "VRF 1", "rd": "vrf-1", "tenant": { "id": 1, "name": "Tenant 1", "slug": "tenant-1" }, "enforce_unique": true, "description": "VRF 1 Description", "custom_fields": {} }`),
			want: testVRFWithTenant(1, testTenantIdentifier(1)),
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("[%d] %s", i, tt.desc), func(t *testing.T) {
			result := new(VRF)
			err := json.Unmarshal(tt.data, result)
			if err != nil {
				t.Fatalf("unexpected error from Tenant.UnmarshalJSON: %v", err)
			}

			if want, got := tt.want, result; !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected Tenant:\n- want: %v\n-  got: %v", want, got)
			}
		})
	}
}

func TestVRFMarshalJSON(t *testing.T) {
	var tests = []struct {
		desc string
		data *VRF
		want []byte
	}{
		{
			desc: "Nil Group",
			data: testVRFWithTenant(1, nil),
			want: []byte(`{"id":1,"name":"VRF 1","rd":"vrf-1","enforce_unique":true,"description":"VRF 1 Description"}`),
		},
		{
			desc: "With Group",
			data: testVRFWithTenant(1, testTenantIdentifier(1)),
			want: []byte(`{"id":1,"name":"VRF 1","rd":"vrf-1","enforce_unique":true,"description":"VRF 1 Description","tenant":1}`),
		},
		{
			desc: "No VRF.ID",
			data: testVRFWithTenant(0, nil),
			want: []byte(`{"name":"VRF 0","rd":"vrf-0","enforce_unique":true,"description":"VRF 0 Description"}`),
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("[%d] %s", i, tt.desc), func(t *testing.T) {
			result, err := json.Marshal(tt.data)
			if err != nil {
				t.Fatalf("unexpected error from VRF.MarshalJSON: %v", err)
			}

			if want, got := tt.want, result; bytes.Compare(want, got) != 0 {
				t.Fatalf("unexpected VRF:\n- want: %s\n-  got: %s", want, got)
			}
		})
	}
}

func TestListVRFOptions(t *testing.T) {
	enforceUniqueTrue := true
	enforceUniqueFalse := false
	var tests = []struct {
		desc string
		o    *ListVRFOptions
		v    url.Values
	}{
		{
			desc: "empty options",
		},
		{
			desc: "full options",
			o: &ListVRFOptions{
				Name:          "Hello",
				RD:            "hello",
				EnforceUnique: &enforceUniqueTrue,
				IDIn:          []uint64{1, 2, 3},
				TenantID:      1,
				Query:         "World",
			},
			v: url.Values{
				"name":           []string{"Hello"},
				"rd":             []string{"hello"},
				"id__in":         []string{"1,2,3"},
				"enforce_unique": []string{"True"},
				"tenant_id":      []string{"1"},
				"q":              []string{"World"},
			},
		},
		{
			desc: "tenant vs tenant_id",
			o: &ListVRFOptions{
				TenantID: 1,
				Tenant:   "Tenant 1",
			},
			v: url.Values{
				"tenant_id": []string{"1"},
			},
		},
		{
			desc: "tenant name",
			o: &ListVRFOptions{
				Tenant: "Tenant 1",
			},
			v: url.Values{
				"tenant": []string{"Tenant 1"},
			},
		},
		{
			desc: "enforce_unique default value",
			o:    &ListVRFOptions{},
			v:    url.Values{},
		},
		{
			desc: "enforce_unique true",
			o: &ListVRFOptions{
				EnforceUnique: &enforceUniqueTrue,
			},
			v: url.Values{
				"enforce_unique": []string{"True"},
			},
		},
		{
			desc: "enforce_unique false",
			o: &ListVRFOptions{
				EnforceUnique: &enforceUniqueFalse,
			},
			v: url.Values{
				"enforce_unique": []string{"False"},
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
