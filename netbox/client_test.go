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
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strconv"
	"testing"
)

func TestClientBadJSON(t *testing.T) {
	c, done := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("foo"))
	})
	defer done()

	req, err := c.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal("expected no error, but an error returned")
	}

	// Pass empty struct to trigger JSON unmarshaling path
	var v struct{}

	err = c.Do(req, &v)
	if _, ok := err.(*json.SyntaxError); !ok {
		t.Fatalf("unexpected error type: %T", err)
	}
}

func TestClientBadStatusCode(t *testing.T) {
	var tests = []struct {
		desc       string
		data       []byte
		statusCode int
		want       error
	}{
		{
			desc:       "403, but no json result",
			data:       []byte("foo"),
			statusCode: http.StatusForbidden,
			want:       errors.New("403 - foo"),
		},
		{
			desc:       "403, with json, but without detail",
			data:       []byte(`{"error_msg": "some error occurred"}`),
			statusCode: http.StatusForbidden,
			want:       errors.New(`403 - {"error_msg": "some error occurred"}`),
		},
		{
			desc:       "500, but correct json",
			data:       []byte(`{"detail": "some error occurred"}`),
			statusCode: http.StatusInternalServerError,
			want:       errors.New("500 - some error occurred"),
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("[%d] %s", i, tt.desc), func(t *testing.T) {
			c, done := testClient(t, func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.statusCode)
				w.Write(tt.data)
			})
			defer done()

			req, err := c.NewRequest(http.MethodGet, "/", nil)
			if err != nil {
				t.Fatal("expected no error, but an error returned")
			}

			var v struct{}
			err = c.Do(req, &v)
			if want, got := tt.want, err; !reflect.DeepEqual(want, got) {
				t.Fatalf("expected error:\n- want: %v\n-  got: %v", want, got)
			}
		})
	}
}

func TestNewJSONRequest(t *testing.T) {
	c, done := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("foo"))
	})
	defer done()
	wantBody := "{\"id\":1,\"name\":\"Test 1\"}\n"
	wantHeader := "application/json; charset=utf-8"

	req, err := c.NewJSONRequest(http.MethodPost, "/", nil, &struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{
		ID:   1,
		Name: "Test 1",
	})
	if err != nil {
		t.Fatal("expected no error, but an error returned")
	}

	res, err := ioutil.ReadAll(req.Body)
	if err != nil {
		t.Fatal("expected no error, but an error returned")
	}
	if want, got := wantBody, string(res); got != want {
		t.Fatalf("unexpected body:\n- want: %v\n-  got: %v", want, got)
	}

	if want, got := wantHeader, req.Header.Get("Content-Type"); got != want {
		t.Fatalf("unexpected body:\n- want: %v\n-  got: %v", want, got)
	}

	req, err = c.NewJSONRequest(http.MethodPost, "/", nil, nil)
	if err == nil {
		t.Fatal("expected an error, but there was none")
	}
	if req != nil {
		t.Fatalf("expected a nil request, but got %v", req)
	}
}

func TestClientQueryParameters(t *testing.T) {
	c := &Client{
		u:      &url.URL{},
		client: &http.Client{},
	}

	const (
		wantFoo = "foo"
		wantBar = 1
	)

	req, err := c.NewRequest(http.MethodGet, "/", testValuer{
		Foo: wantFoo,
		Bar: wantBar,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	q := req.URL.Query()
	if want, got := 2, len(q); want != got {
		t.Fatalf("unexpected number of query parameters:\n- want: %v\n-  got: %v",
			want, got)
	}

	if want, got := wantFoo, q.Get("foo"); want != got {
		t.Fatalf("unexpected foo:\n- want: %v\n-  got: %v", want, got)
	}

	if want, got := strconv.Itoa(wantBar), q.Get("bar"); want != got {
		t.Fatalf("unexpected bar:\n- want: %v\n-  got: %v", want, got)
	}
}

func TestClientPrependBaseURLPath(t *testing.T) {
	u, err := url.Parse("http://example.com/netbox/")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	c := &Client{
		u:      u,
		client: &http.Client{},
	}

	req, err := c.NewRequest(http.MethodGet, "/api/ipam/vlans", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if want, got := "/netbox/api/ipam/vlans", req.URL.Path; want != got {
		t.Fatalf("unexpected URL path:\n- want: %q\n-  got: %q",
			want, got)
	}
}

func TestAuthenticatingClientAddsHeader(t *testing.T) {
	c, err := NewClient("localhost", "auth-token", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	req, err := c.NewRequest(http.MethodGet, "/api/ipam/vlans", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if want, got := "Token auth-token", req.Header.Get("Authorization"); want != got {
		t.Fatalf("unexpected Authorization header:\n- want: %q\n-  got: %q",
			want, got)
	}
}

type testValuer struct {
	Foo string
	Bar int
}

func (q testValuer) Values() (url.Values, error) {
	v := url.Values{}

	if q.Foo != "" {
		v.Set("foo", q.Foo)
	}

	if q.Bar != 0 {
		v.Set("bar", strconv.Itoa(q.Bar))
	}

	return v, nil
}

func testClient(t *testing.T, fn func(w http.ResponseWriter, r *http.Request)) (*Client, func()) {
	s := httptest.NewServer(http.HandlerFunc(fn))

	c, err := NewClient(s.URL, "", nil)
	if err != nil {
		t.Fatalf("error creating Client: %v", err)
	}

	return c, func() { s.Close() }
}

func testHandler(t *testing.T, method string, path string, v interface{}) http.HandlerFunc {
	return testStatusHandler(t, method, path, v, 0)
}

func testStatusHandler(t *testing.T, method string, path string, v interface{}, statusCode int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if want, got := method, r.Method; want != got {
			t.Fatalf("unexpected HTTP method:\n- want: %v\n-  got: %v", want, got)
		}

		if want, got := path, r.URL.Path; want != got {
			t.Fatalf("unexpected URL path:\n- want: %v\n-  got: %v", want, got)
		}

		if statusCode > 0 {
			w.WriteHeader(statusCode)
		}

		if err := json.NewEncoder(w).Encode(v); err != nil {
			t.Fatalf("error while encoding JSON: %v", err)
		}
	}
}

func testTenantGroupCreate(n int) *TenantGroup {
	return &TenantGroup{
		Name: fmt.Sprintf("Tenant Group %d", n),
		Slug: fmt.Sprintf("tenant-group-%d", n),
	}
}

func testTenantGroup(n int) *TenantGroup {
	return &TenantGroup{
		ID:   n,
		Name: fmt.Sprintf("Tenant Group %d", n),
		Slug: fmt.Sprintf("tenant-group-%d", n),
	}
}

func testTenant(n int) *Tenant {
	return testTenantWithGroup(n, testTenantGroup(n))
}

func testTenantCreate(n int) *Tenant {
	return testTenantWithGroupCreate(n, testTenantGroup(n))
}

func testTenantWithGroupCreate(n int, t *TenantGroup) *Tenant {
	return &Tenant{
		Name:        fmt.Sprintf("Tenant %d", n),
		Slug:        fmt.Sprintf("tenant-%d", n),
		Description: fmt.Sprintf("Tenant %d Description", n),
		Comments:    fmt.Sprintf("Tenant %d Comments", n),
		Group:       t,
	}
}

func testTenantWithGroup(n int, t *TenantGroup) *Tenant {
	return &Tenant{
		ID:          n,
		Name:        fmt.Sprintf("Tenant %d", n),
		Slug:        fmt.Sprintf("tenant-%d", n),
		Description: fmt.Sprintf("Tenant %d Description", n),
		Comments:    fmt.Sprintf("Tenant %d Comments", n),
		Group:       t,
	}
}

func testTenantIdentifier(n int) *Tenant {
	return &Tenant{
		ID:   n,
		Name: fmt.Sprintf("Tenant %d", n),
		Slug: fmt.Sprintf("tenant-%d", n),
	}
}

func testVRF(n int) *VRF {
	return &VRF{
		ID:            n,
		Name:          fmt.Sprintf("VRF %d", n),
		RD:            fmt.Sprintf("vrf-%d", n),
		EnforceUnique: true,
		Description:   fmt.Sprintf("VRF %d Description", n),
		Tenant:        testTenantWithGroup(n, nil),
	}
}

func testVRFCreate(n int) *VRF {
	return &VRF{
		Name:          fmt.Sprintf("VRF %d", n),
		RD:            fmt.Sprintf("vrf-%d", n),
		EnforceUnique: true,
		Description:   fmt.Sprintf("VRF %d Description", n),
		Tenant:        testTenantWithGroup(n, nil),
	}
}

func testVRFWithTenant(n int, tenant *Tenant) *VRF {
	return &VRF{
		ID:            n,
		Name:          fmt.Sprintf("VRF %d", n),
		RD:            fmt.Sprintf("vrf-%d", n),
		EnforceUnique: true,
		Description:   fmt.Sprintf("VRF %d Description", n),
		Tenant:        tenant,
	}
}
