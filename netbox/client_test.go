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
			want:       errors.New("some error occurred"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
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

func TestClientPing(t *testing.T) {
	wantEndpoints := map[string]string{
		"dcim":     "https://netbox.example.com/api/dcim/",
		"circuits": "https://netbox.example.com/api/circuits/",
		"ipam":     "https://netbox.example.com/api/ipam/",
		"secrets":  "https://netbox.example.com/api/secrets/",
		"tenancy":  "https://netbox.example.com/api/tenancy/",
		"extras":   "https://netbox.example.com/api/extras/",
	}
	c, done := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		method := http.MethodGet
		path := "/api/"
		if want, got := method, r.Method; want != got {
			t.Fatalf("unexpected HTTP method:\n- want: %v\n-  got: %v", want, got)
		}
		if want, got := path, r.URL.Path; want != got {
			t.Fatalf("unexpected URL path:\n- want: %v\n-  got: %v", want, got)
		}
		w.Header().Add("API-Version", "2.0")

		err := json.NewEncoder(w).Encode(struct {
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
		})
		if err != nil {
			t.Fatalf("error while encoding JSON: %v", err)
		}
	})
	defer done()

	api, err := c.Ping()
	if err != nil {
		t.Fatalf("expected non error:\n-  got: %v", err)
	}
	if want, got := "2.0", api.Version; want != got {
		t.Fatalf("unexpected api.Version:\n- want: %v\n-  got: %v", want, got)
	}
	if want, got := wantEndpoints, api.Endpoints; !reflect.DeepEqual(want, got) {
		t.Fatalf("unexpected api.Endpoints\n- want: %v\n-  got: %v", want, got)
	}
}

func TestPingUnauthorized(t *testing.T) {
	c, done := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("{\"detail\":\"Invalid token\"}"))
	})
	defer done()

	res, err := c.Ping()
	if want, got := errors.New("Invalid token"), err; !reflect.DeepEqual(want, got) {
		t.Fatalf("expected error:\n- want: %v\n-  got: %v", want, got)
	}
	if res != nil {
		t.Fatalf("unexpected result:\n- want: %v\n-  got: %v", nil, res)
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
		t.Fatal("expected an error, but no error returned")
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
		t.Fatal("expected an error, but no error returned")
	}

	if want, got := "/netbox/api/ipam/vlans", req.URL.Path; want != got {
		t.Fatalf("unexpected URL path:\n- want: %q\n-  got: %q",
			want, got)
	}
}

type testValuer struct {
	Foo string
	Bar int
}

func (q testValuer) values() (url.Values, error) {
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

	c, err := NewClient(s.URL, nil)
	if err != nil {
		t.Fatalf("error creating Client: %v", err)
	}

	return c, func() { s.Close() }
}

func testHandler(t *testing.T, method string, path string, v interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if want, got := method, r.Method; want != got {
			t.Fatalf("unexpected HTTP method:\n- want: %v\n-  got: %v", want, got)
		}

		if want, got := path, r.URL.Path; want != got {
			t.Fatalf("unexpected URL path:\n- want: %v\n-  got: %v", want, got)
		}

		if err := json.NewEncoder(w).Encode(v); err != nil {
			t.Fatalf("error while encoding JSON: %v", err)
		}
	}
}
