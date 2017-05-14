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

func TestTenantGet(t *testing.T) {
	var tests = []struct {
		desc      string
		want      *Tenant
		wantGroup *TenantGroup
	}{
		{
			desc:      "Without TenantGroup",
			want:      testTenantWithGroup(1, nil),
			wantGroup: nil,
		},
		{
			desc:      "With TenantGroup",
			want:      testTenantWithGroup(1, testTenantGroup(1)),
			wantGroup: testTenantGroup(1),
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("[%d] %s", i, tt.desc), func(t *testing.T) {
			serverData := serverDataTenant(*tt.want)

			c, done := testClient(t, testHandler(t, http.MethodGet, "/api/tenancy/tenants/1/", &serverData))
			defer done()

			res, err := c.Tenancy.Tenants.Get(1)
			if err != nil {
				t.Fatalf("unexpected error from Client.Tenancy.Tenants.Get: %v", err)
			}

			if want, got := tt.want, res; !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected Tenant:\n- want: %v\n-  got: %v", want, got)
			}
			if want, got := tt.wantGroup, res.Group; !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected TenantGroup:\n- want: %v\n-  got: %v", want, got)
			}
		})
	}
}

func TestTenantUnmarshalJSON(t *testing.T) {
	var tests = []struct {
		desc string
		data []byte
		want *Tenant
	}{
		{
			desc: "Nil Group",
			data: []byte(`{ "id": 1, "name": "Tenant 1", "slug": "tenant-1", "group": null, "description": "Tenant 1 Description", "comments": "Tenant 1 Comments", "custom_fields": {} }`),
			want: testTenantWithGroup(1, nil),
		},
		{
			desc: "With Group",
			data: []byte(`{ "id": 1, "name": "Tenant 1", "slug": "tenant-1", "group": {"id": 1, "name": "Tenant Group 1", "slug": "tenant-group-1"}, "description": "Tenant 1 Description", "comments": "Tenant 1 Comments", "custom_fields": {} }`),
			want: testTenant(1),
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("[%d] %s", i, tt.desc), func(t *testing.T) {
			result := new(Tenant)
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

func TestTenantMarshalJSON(t *testing.T) {
	var tests = []struct {
		desc string
		data *Tenant
		want []byte
	}{
		{
			desc: "Nil Group",
			data: testTenantWithGroup(1, nil),
			want: []byte(`{"id":1,"name":"Tenant 1","slug":"tenant-1","description":"Tenant 1 Description","comments":"Tenant 1 Comments"}`),
		},
		{
			desc: "With Group",
			data: testTenant(1),
			want: []byte(`{"id":1,"name":"Tenant 1","slug":"tenant-1","description":"Tenant 1 Description","comments":"Tenant 1 Comments","group":1}`),
		},
		{
			desc: "No Tenant.ID",
			data: testTenantWithGroup(0, nil),
			want: []byte(`{"name":"Tenant 0","slug":"tenant-0","description":"Tenant 0 Description","comments":"Tenant 0 Comments"}`),
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("[%d] %s", i, tt.desc), func(t *testing.T) {
			result, err := json.Marshal(tt.data)
			if err != nil {
				t.Fatalf("unexpected error from Tenant.MarshalJSON: %v", err)
			}

			if want, got := tt.want, result; bytes.Compare(want, got) != 0 {
				t.Fatalf("unexpected Tenant:\n- want: %s\n-  got: %s", want, got)
			}
		})
	}
}

func TestListTenantOptions(t *testing.T) {
	var tests = []struct {
		desc string
		o    *ListTenantOptions
		v    url.Values
	}{
		{
			desc: "empty options",
		},
		{
			desc: "full options",
			o: &ListTenantOptions{
				Name:    "Hello",
				IDIn:    "1,2,3",
				GroupID: 1,
				Query:   "World",
			},
			v: url.Values{
				"name":     []string{"Hello"},
				"id__in":   []string{"1,2,3"},
				"group_id": []string{"1"},
				"q":        []string{"World"},
			},
		},
		{
			desc: "group vs group_id",
			o: &ListTenantOptions{
				GroupID: 1,
				Group:   "Group1",
			},
			v: url.Values{
				"group_id": []string{"1"},
			},
		},
		{
			desc: "group name",
			o: &ListTenantOptions{
				Group: "Group 1",
			},
			v: url.Values{
				"group": []string{"Group 1"},
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
