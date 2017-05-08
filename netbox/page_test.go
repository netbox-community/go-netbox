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
	"errors"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"testing"
)

type testOptions struct{}

func (t *testOptions) Values() (url.Values, error) {
	return url.Values{
		"Hello": []string{"World"},
		"limit": []string{"30"},
	}, nil
}

func TestPageValues(t *testing.T) {

	var tests = []struct {
		desc string
		page *Page
		want url.Values
		err  error
	}{
		{
			desc: "nil page",
			page: nil,
			want: nil,
			err:  errors.New("page not defined"),
		},
		{
			desc: "options nil",
			page: NewPage(nil, "/", nil),
			want: url.Values{
				"limit":  []string{"50"},
				"offset": []string{"0"},
			},
			err: nil,
		},
		{
			desc: "options set",
			page: NewPage(nil, "/", &testOptions{}),
			want: url.Values{
				"Hello":  []string{"World"},
				"limit":  []string{"50"},
				"offset": []string{"0"},
			},
			err: nil,
		},
	}

	for i, tt := range tests {
		res, err := tt.page.Values()
		if want, got := tt.err, err; !reflect.DeepEqual(want, got) {
			t.Fatalf("[%d] %s - unecpected error:\n- want: %v\n-  got: %v", i, tt.desc, want, got)
		}
		if want, got := tt.want, res; !reflect.DeepEqual(want, got) {
			t.Fatalf("[%d] %s - unexpected values:\n- want: %v\n-  got: %v", i, tt.desc, want, got)
		}
	}
}

func TestNext(t *testing.T) {
	var pages = []pageData{
		{
			Count:       99,
			NextURL:     "http://example.com/?limit=50&offset=50",
			PreviousURL: "",
			Results:     []byte("{\"Hello\": \"World\"}"),
		},
		{
			Count:       99,
			NextURL:     "",
			PreviousURL: "http://example.com/?limit=50&offset=0",
			Results:     []byte("{\"Hello\": \"World\"}"),
		},
	}

	c, done := testClient(t, testHandler(t, http.MethodGet, "/", pages[0]))
	defer done()

	p := NewPage(c, "/", nil)
	if !p.Next() {
		t.Fatal("Expected another page, got none.")
	}
	p.c, done = testClient(t, testHandler(t, http.MethodGet, "/", pages[1]))
	defer done()
	if !p.Next() {
		t.Fatal("Expected another page, got none.")
	}
	if p.Next() {
		t.Fatal("Did not expect another page, but got one.")
	}
	if err := p.Err(); err != nil {
		t.Fatalf("Did not expect an error: %v", err)
	}
}

func TestPageWithError(t *testing.T) {
	c, done := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("foo"))
	})
	defer done()

	p := NewPage(c, "/", &testOptions{})
	if p.Next() {
		t.Fatal("Did not expect any page, but got one.")
	}
	if want, got := errors.New("403 - foo"), p.Err(); !reflect.DeepEqual(want, got) {
		t.Fatalf("unexpected error:\n- want: %v\n-  got: %v", want, got)
	}
}

func TestErr(t *testing.T) {
	var tests = []struct {
		desc string
		set  error
		want error
	}{
		{
			desc: "Err nil",
			set:  nil,
			want: nil,
		},
		{
			desc: "Err set",
			set:  errors.New("This is broken"),
			want: errors.New("This is broken"),
		},
	}
	for i, tt := range tests {
		p := NewPage(nil, "/", nil)
		p.setErr(tt.set)
		if want, got := tt.want, p.Err(); !reflect.DeepEqual(want, got) {
			t.Fatalf("[%d] %s - unexpected error:\n- want: %v\n-  got: %v", i, tt.desc, want, got)
		}
	}
}

func TestSetNext(t *testing.T) {
	p := NewPage(nil, "/", nil)
	p.setNext(99, 88)

	if want, got := 99, p.limit; want != got {
		t.Fatalf("unexpected limit:\n- want: %v\n-  got: %v", want, got)
	}
	if want, got := 88, p.offset; want != got {
		t.Fatalf("unexpected offset:\n- want: %v\n-  got: %v", want, got)
	}
}

func TestSetNextURL(t *testing.T) {
	var tests = []struct {
		desc       string
		url        string
		wantOffset int
		wantLimit  int
		err        error
	}{
		{
			desc:       "No Params returned",
			url:        "http://example.com",
			wantOffset: 0,
			wantLimit:  50,
			err:        errors.New("no such query parameter limit"),
		},
		{
			desc:       "No Offset returned",
			url:        "http://example.com?limit=20",
			wantOffset: 0,
			wantLimit:  50,
			err:        errors.New("no such query parameter offset"),
		},
		{
			desc:       "Limit not an int",
			url:        "http://example.com?limit=hello&offset=world",
			wantOffset: 0,
			wantLimit:  50,
			err: &strconv.NumError{
				Func: "Atoi",
				Num:  "hello",
				Err:  strconv.ErrSyntax,
			},
		},
		{
			desc:       "Offset not an int",
			url:        "http://example.com?limit=50&offset=world",
			wantOffset: 0,
			wantLimit:  50,
			err: &strconv.NumError{
				Func: "Atoi",
				Num:  "world",
				Err:  strconv.ErrSyntax,
			},
		},
		{
			desc:       "Correct limit and offset",
			url:        "http://example.com?limit=99&offset=88",
			wantOffset: 88,
			wantLimit:  99,
			err:        nil,
		},
	}

	for i, tt := range tests {
		p := NewPage(nil, "/", nil)
		p.setNextURL(tt.url)
		if want, got := tt.wantOffset, p.offset; want != got {
			t.Fatalf("[%d] %s - unexpected offset:\n- want: %v\n-  got: %v", i, tt.desc, want, got)
		}
		if want, got := tt.wantLimit, p.limit; want != got {
			t.Fatalf("[%d] %s - unexpected limit:\n- want: %v\n-  got: %v", i, tt.desc, want, got)
		}
		if want, got := tt.err, p.Err(); reflect.TypeOf(want) != reflect.TypeOf(got) {
			t.Fatalf("[%d] %s - unexpected error:\n- want: %v\n-  got: %v", i, tt.desc, want, got)
		}
	}
}
