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
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/http/httptest"
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

// TODO(cglaubitz) implement ping, test this piece of code
// ExampleNewClient demonstrates usage of the Client type.
func ExampleNewClient() {
	// Sets up a minimal, mocked NetBox server
	addr, done := exampleServer()
	defer done()

	// Creates a client configured to use the test server
	c, err := NewClient(addr, nil)
	if err != nil {
		panic(fmt.Sprintf("failed to create netbox.Client: %v", err))
	}

	// Retrieve API Information
	api, err := c.Ping()
	if err != nil {
		panic(fmt.Sprintf("failed to Ping API: %v", err))
	}

	fmt.Printf("API Version   %s\n", api.Version)
	fmt.Printf("API Endpoints %v\n", api.Endpoints)

}

// exampleServer creates a test HTTP server which returns its address and
// can be closed using the returned closure.
func exampleServer() (string, func()) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := struct {
			Dcim     string `json:"dcim"`
			Circuits string `json:"circuits"`
			Ipam     string `json:"ipam"`
			Secrets  string `json:"secrets"`
			Tenancy  string `json:"tenancy"`
			Extras   string `json:"extras"`
		}{
			Dcim:     "https://netbox.example.com/api/dcim/",
			Circuits: "https://netbox.example.com/api/circuits/",
			Ipam:     "https://netbox.example.com/api/ipam/",
			Secrets:  "https://netbox.example.com/api/secrets/",
			Tenancy:  "https://netbox.example.com/api/tenancy/",
			Extras:   "https://netbox.example.com/api/extras/",
		}

		w.Header().Set("API-Version", "2.0")

		_ = json.NewEncoder(w).Encode(ip)
	}))

	return s.URL, func() { s.Close() }
}
