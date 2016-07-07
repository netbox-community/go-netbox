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

	// Retrieve an IPAddress with ID 1
	ip, err := c.IPAM.GetIPAddress(1)
	if err != nil {
		panic(fmt.Sprintf("failed to retrieve IP address: %v", err))
	}

	fmt.Printf("IP #%03d: %s (%s)\n", ip.ID, ip.Address.String(), ip.Family)

	// Output:
	// IP #001: 192.168.1.1/32 (IPv4)
}

// exampleServer creates a test HTTP server which returns its address and
// can be closed using the returned closure.
func exampleServer() (string, func()) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := struct {
			ID      int    `json:"id"`
			Family  Family `json:"family"`
			Address string `json:"address"`
		}{
			ID:      1,
			Family:  FamilyIPv4,
			Address: "192.168.1.1/32",
		}

		_ = json.NewEncoder(w).Encode(ip)
	}))

	return s.URL, func() { s.Close() }
}
