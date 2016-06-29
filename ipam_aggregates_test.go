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
	"bytes"
	"encoding/json"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestClientIPAMGetAggregate(t *testing.T) {
	wantAggregate := testAggregate(FamilyIPv4, 1)

	c, done := testClient(t, testHandler(t, http.MethodGet, "/api/ipam/aggregates/1", wantAggregate))
	defer done()

	gotAggregate, err := c.IPAM.GetAggregate(wantAggregate.ID)
	if err != nil {
		t.Fatalf("unexpected error from Client.IPAM.GetAggregate: %v", err)
	}

	if want, got := *wantAggregate, *gotAggregate; !aggregatesEqual(want, got) {
		t.Fatalf("unexpected Aggregate:\n- want: %v\n-  got: %v", want, got)
	}
}

func TestClientIPAMListAggregates(t *testing.T) {
	wantAggregates := []*Aggregate{
		testAggregate(FamilyIPv4, 1),
		testAggregate(FamilyIPv6, 2),
	}

	c, done := testClient(t, testHandler(t, http.MethodGet, "/api/ipam/aggregates/", wantAggregates))
	defer done()

	gotAggregates, err := c.IPAM.ListAggregates(nil)
	if err != nil {
		t.Fatalf("unexpected error from Client.IPAM.ListAggregates: %v", err)
	}

	want := derefAggregates(wantAggregates)
	got := derefAggregates(gotAggregates)
	if !aggregateSlicesEqual(want, got) {
		t.Fatalf("unexpected Aggregates:\n- want: %v\n-  got: %v", want, got)
	}
}

func TestListAggregatesOptionsValues(t *testing.T) {
	var tests = []struct {
		desc string
		o    *ListAggregatesOptions
		v    url.Values
	}{
		{
			desc: "empty options",
		},
		{
			desc: "family only",
			o: &ListAggregatesOptions{
				Family: FamilyIPv4,
			},
			v: url.Values{
				"family": []string{strconv.Itoa(int(FamilyIPv4))},
			},
		},
		{
			desc: "1 rir_id only",
			o: &ListAggregatesOptions{
				RIRID: []int{1},
			},
			v: url.Values{
				"rir_id": []string{"1"},
			},
		},
		{
			desc: "3 rir_ids only",
			o: &ListAggregatesOptions{
				RIRID: []int{1, 2, 3},
			},
			v: url.Values{
				"rir_id": []string{"1", "2", "3"},
			},
		},
		{
			desc: "1 rir only",
			o: &ListAggregatesOptions{
				RIR: []string{"rir"},
			},
			v: url.Values{
				"rir": []string{"rir"},
			},
		},
		{
			desc: "3 rirs only",
			o: &ListAggregatesOptions{
				RIR: []string{"rirfoo", "rirbar", "rirbaz"},
			},
			v: url.Values{
				"rir": []string{"rirfoo", "rirbar", "rirbaz"},
			},
		},
		{
			desc: "rir and rir_id, rir_id preferred",
			o: &ListAggregatesOptions{
				RIR:   []string{"rir"},
				RIRID: []int{1},
			},
			v: url.Values{
				"rir_id": []string{"1"},
			},
		},
		{
			desc: "date_added only",
			o: &ListAggregatesOptions{
				DateAdded: time.Date(2016, time.January, 22, 0, 0, 0, 0, time.UTC),
			},
			v: url.Values{
				"date_added": []string{"2016-01-22"},
			},
		},
		{
			desc: "all options",
			o: &ListAggregatesOptions{
				Family:    FamilyIPv4,
				RIRID:     []int{1},
				RIR:       []string{"rir"},
				DateAdded: time.Date(2016, time.January, 22, 0, 0, 0, 0, time.UTC),
			},
			v: url.Values{
				"family":     []string{strconv.Itoa(int(FamilyIPv4))},
				"rir_id":     []string{"1"},
				"date_added": []string{"2016-01-22"},
			},
		},
	}

	for i, tt := range tests {
		t.Logf("[%02d] test %q", i, tt.desc)

		v, err := tt.o.values()
		if err != nil {
			t.Fatalf("unexpected Values error: %v", err)
		}

		if want, got := tt.v, v; !reflect.DeepEqual(want, got) {
			t.Fatalf("unexpected url.Values map:\n- want: %v\n-  got: %v",
				want, got)
		}
	}
}

func TestAggregateMarshalJSON(t *testing.T) {
	var tests = []struct {
		desc string
		a    *Aggregate
		b    []byte
	}{
		{
			desc: "IPv4 aggregate",
			a:    testAggregate(FamilyIPv4, 1),
			b:    []byte(`{"id":1,"family":4,"prefix":"8.0.0.0/8","rir":{"id":1,"name":"RIRIdentifier 1","slug":"riridentifier1"},"date_added":"2016-01-01","description":"description 1"}`),
		},
		{
			desc: "IPv6 aggregate",
			a:    testAggregate(FamilyIPv6, 2),
			b:    []byte(`{"id":2,"family":6,"prefix":"2001::/16","rir":{"id":2,"name":"RIRIdentifier 2","slug":"riridentifier2"},"date_added":"2016-01-01","description":"description 2"}`),
		},
	}

	for i, tt := range tests {
		t.Logf("[%02d] test %q", i, tt.desc)

		b, err := json.Marshal(tt.a)
		if err != nil {
			t.Fatalf("unexpected JSON marshal error: %v", err)
		}

		if want, got := tt.b, b; !bytes.Equal(want, got) {
			t.Fatalf("unexpected JSON bytes:\n- want: %v\n-  got: %v",
				string(want), string(got))
		}
	}
}

func TestAggregateUnmarshalJSON(t *testing.T) {
	var tests = []struct {
		desc string
		b    []byte
		a    *Aggregate
		err  error
	}{
		{
			desc: "invalid aggregate due to prefix",
			b:    []byte(`{"prefix":"foo"}`),
			err: &net.ParseError{
				Type: "CIDR address",
				Text: "foo",
			},
		},
		{
			desc: "invalid aggregate due to date_added",
			b:    []byte(`{"prefix":"5.101.96.0/20","date_added":"foo"}`),
			err: &time.ParseError{
				Layout:     dateFormat,
				Value:      "foo",
				LayoutElem: "2006",
				ValueElem:  "foo",
			},
		},
		{
			desc: "IPv4 aggregate",
			b:    []byte(`{"id":1,"family":4,"prefix":"8.0.0.0/8","rir":{"id":1,"name":"RIRIdentifier 1","slug":"riridentifier1"},"date_added":"2016-01-01","description":"description 1"}`),
			a:    testAggregate(FamilyIPv4, 1),
		},
		{
			desc: "IPv6 aggregate",
			b:    []byte(`{"id":2,"family":6,"prefix":"2001::/16","rir":{"id":2,"name":"RIRIdentifier 2","slug":"riridentifier2"},"date_added":"2016-01-01","description":"description 2"}`),
			a:    testAggregate(FamilyIPv6, 2),
		},
	}

	for i, tt := range tests {
		t.Logf("[%02d] test %q", i, tt.desc)

		a := new(Aggregate)
		err := json.Unmarshal(tt.b, a)

		if want, got := tt.err, err; !reflect.DeepEqual(want, got) {
			t.Fatalf("unexpected error:\n- want: %v\n-  got: %v",
				want, got)
		}
		if err != nil {
			continue
		}

		if want, got := *tt.a, *a; !aggregatesEqual(want, got) {
			t.Fatalf("unexpected Aggregate:\n- want: %v\n-  got: %v",
				want, got)
		}
	}
}

func aggregateSlicesEqual(a, b []Aggregate) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !aggregatesEqual(a[i], b[i]) {
			return false
		}
	}

	return true
}

func aggregatesEqual(a, b Aggregate) bool {
	if a.ID != b.ID {
		return false
	}

	if a.Family != b.Family {
		return false
	}

	if a.Prefix.String() != b.Prefix.String() {
		return false
	}

	if !a.DateAdded.Equal(b.DateAdded) {
		return false
	}

	if a.Description != b.Description {
		return false
	}

	return true
}

// Used to print values of Aggregates in slice, instead of memory addresses.
func derefAggregates(aggregates []*Aggregate) []Aggregate {
	a := make([]Aggregate, len(aggregates))
	for i := range aggregates {
		a[i] = *aggregates[i]
	}

	return a
}
