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
	"testing"
)

func TestFamilyValid(t *testing.T) {
	var tests = []struct {
		f  Family
		ok bool
	}{
		{
			f: math.MinInt64,
		},
		{
			f: math.MaxInt64,
		},
		{
			f:  FamilyIPv4,
			ok: true,
		},
		{
			f:  FamilyIPv6,
			ok: true,
		},
	}

	for _, tt := range tests {
		if want, got := tt.ok, tt.f.Valid(); want != got {
			t.Fatalf("unexpected Family(%d).Valid():\n- want: %v\n-  got: %v",
				tt.f, want, got)
		}
	}
}
