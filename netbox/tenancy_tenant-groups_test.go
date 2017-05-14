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
	"reflect"
	"testing"
)

func TestTenantGroupUnmarshalJSON(t *testing.T) {
	var tests = []struct {
		desc string
		data []byte
		want *TenantGroup
	}{
		{
			desc: "full",
			data: []byte(`{ "id": 1, "name": "Tenant Group 1", "slug": "tenant-group-1", "custom_fields": {} }`),
			want: testTenantGroup(1),
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("[%d] %s", i, tt.desc), func(t *testing.T) {
			result := new(TenantGroup)
			err := json.Unmarshal(tt.data, result)
			if err != nil {
				t.Fatalf("unexpected error from TenantGroup.UnmarshalJSON: %v", err)
			}

			if want, got := tt.want, result; !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected TenantGroup:\n- want: %v\n-  got: %v", want, got)
			}
		})
	}
}

func TestTenantGroupMarshalJSON(t *testing.T) {
	var tests = []struct {
		desc string
		data *TenantGroup
		want []byte
	}{
		{
			desc: "With TenantGroup.ID",
			data: testTenantGroup(1),
			want: []byte(`{"id":1,"name":"Tenant Group 1","slug":"tenant-group-1"}`),
		},
		{
			desc: "No TenantGroup.ID",
			data: testTenantGroup(0),
			want: []byte(`{"name":"Tenant Group 0","slug":"tenant-group-0"}`),
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("[%d] %s", i, tt.desc), func(t *testing.T) {
			result, err := json.Marshal(tt.data)
			if err != nil {
				t.Fatalf("unexpected error from TenantGroup.MarshalJSON: %v", err)
			}

			if want, got := tt.want, result; bytes.Compare(want, got) != 0 {
				t.Fatalf("unexpected TenantGroup:\n- want: %s\n-  got: %s", want, got)
			}
		})
	}
}
