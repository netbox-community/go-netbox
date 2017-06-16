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
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFamilyValid(t *testing.T) {
	var tests = []struct {
		desc string
		f    Family
		ok   bool
	}{
		{
			desc: "Test math.MinInt64",
			f:    math.MinInt64,
		},
		{
			desc: "Test math.MaxInt64",
			f:    math.MaxInt64,
		},
		{
			desc: "Test FamilyIPv4",
			f:    FamilyIPv4,
			ok:   true,
		},
		{
			desc: "Test FamilyIPv6",
			f:    FamilyIPv6,
			ok:   true,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("[%d] %s", i, tt.desc), func(t *testing.T) {
			if want, got := tt.ok, tt.f.Valid(); want != got {
				t.Fatalf("unexpected Family(%d).Valid():\n- want: %v\n-  got: %v",
					tt.f, want, got)
			}
		})
	}
}

type testSimple struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// ExampleNewClient demonstrates usage of the Client type.
func ExampleNewClient() {
	// Sets up a minimal, mocked NetBox server
	addr, done := exampleServer()
	defer done()

	// Creates a client configured to use the test server
	c, err := NewClient(addr, nil, nil)
	if err != nil {
		panic(fmt.Sprintf("failed to create netbox.Client: %v", err))
	}

	res := testSimple{}
	req, err := c.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		panic(err)
	}

	err = c.Do(req, &res)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", res)
}

// exampleServer creates a test HTTP server which returns its address and
// can be closed using the returned closure.
func exampleServer() (string, func()) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		simple := testSimple{
			ID:   1,
			Name: "Test 1",
			Slug: "test-1",
		}

		_ = json.NewEncoder(w).Encode(simple)
	}))

	return s.URL, func() { s.Close() }
}
