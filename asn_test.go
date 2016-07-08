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

func TestASNPrivate(t *testing.T) {
	var tests = []struct {
		a  ASN
		ok bool
	}{
		{
			a: 0,
		},
		{
			a: math.MinInt64,
		},
		{
			a: math.MaxInt64,
		},
		{
			a:  privateMin16Bit,
			ok: true,
		},
		{
			a:  privateMin16Bit + 1,
			ok: true,
		},
		{
			a:  privateMax16Bit - 1,
			ok: true,
		},
		{
			a:  privateMax16Bit,
			ok: true,
		},
		{
			a:  privateMin32Bit,
			ok: true,
		},
		{
			a:  privateMax32Bit,
			ok: true,
		},
	}

	for _, tt := range tests {
		if want, got := tt.ok, tt.a.Private(); want != got {
			t.Fatalf("unexpected ASN(%d).Private():\n- want: %v\n-  got: %v",
				tt.a, want, got)
		}
	}
}

func TestASNPublic(t *testing.T) {
	var tests = []struct {
		a  ASN
		ok bool
	}{
		{
			a: math.MinInt64,
		},
		{
			a: math.MaxInt64,
		},
		{
			a: privateMin16Bit,
		},
		{
			a: privateMin16Bit + 1,
		},
		{
			a: privateMax16Bit - 1,
		},
		{
			a: privateMax16Bit,
		},
		{
			a: privateMin32Bit,
		},
		{
			a: privateMax32Bit,
		},
		{
			a: reservedStart16Bit,
		},
		{
			a: reservedMin16Bit,
		},
		{
			a: reservedMin16Bit + 1,
		},
		{
			a: reservedMax16Bit - 1,
		},
		{
			a: reservedMax16Bit,
		},
		{
			a: reservedDocMin16Bit,
		},
		{
			a: reservedDocMin16Bit + 1,
		},
		{
			a: reservedDocMax16Bit - 1,
		},
		{
			a: reservedDocMax16Bit,
		},
		{
			a: reservedEnd16Bit,
		},
		{
			a: reservedDocMin32Bit,
		},
		{
			a: reservedDocMin32Bit + 1,
		},
		{
			a: reservedDocMax32Bit - 1,
		},
		{
			a: reservedDocMax32Bit,
		},
		{
			a: reservedMin32Bit,
		},
		{
			a: reservedMin32Bit + 1,
		},
		{
			a: reservedMax32Bit - 1,
		},
		{
			a: reservedMax32Bit,
		},
		{
			a: reservedEnd32Bit,
		},
		{
			a: asTrans,
		},
		{
			a:  reservedStart16Bit + 1,
			ok: true,
		},
		{
			a:  reservedMin16Bit - 1,
			ok: true,
		},
		{
			a:  reservedMax32Bit + 1,
			ok: true,
		},
		{
			a:  privateMin32Bit - 1,
			ok: true,
		},
	}

	for _, tt := range tests {
		if want, got := tt.ok, tt.a.Public(); want != got {
			t.Fatalf("unexpected ASN(%d).Public():\n- want: %v\n-  got: %v",
				tt.a, want, got)
		}
	}
}

func TestASNReserved(t *testing.T) {
	var tests = []struct {
		a  ASN
		ok bool
	}{
		{
			a: math.MinInt64,
		},
		{
			a: math.MaxInt64,
		},
		{
			a:  reservedStart16Bit,
			ok: true,
		},
		{
			a:  reservedMin16Bit,
			ok: true,
		},
		{
			a:  reservedMin16Bit + 1,
			ok: true,
		},
		{
			a:  reservedMax16Bit - 1,
			ok: true,
		},
		{
			a:  reservedMax16Bit,
			ok: true,
		},
		{
			a:  reservedDocMin16Bit,
			ok: true,
		},
		{
			a:  reservedDocMin16Bit + 1,
			ok: true,
		},
		{
			a:  reservedDocMax16Bit - 1,
			ok: true,
		},
		{
			a:  reservedDocMax16Bit,
			ok: true,
		},
		{
			a:  reservedEnd16Bit,
			ok: true,
		},
		{
			a:  reservedDocMin32Bit,
			ok: true,
		},
		{
			a:  reservedDocMin32Bit + 1,
			ok: true,
		},
		{
			a:  reservedDocMax32Bit - 1,
			ok: true,
		},
		{
			a:  reservedDocMax32Bit,
			ok: true,
		},
		{
			a:  reservedMin32Bit,
			ok: true,
		},
		{
			a:  reservedMin32Bit + 1,
			ok: true,
		},
		{
			a:  reservedMax32Bit - 1,
			ok: true,
		},
		{
			a:  reservedMax32Bit,
			ok: true,
		},
		{
			a:  reservedEnd32Bit,
			ok: true,
		},
		{
			a:  asTrans,
			ok: true,
		},
	}

	for _, tt := range tests {
		if want, got := tt.ok, tt.a.Reserved(); want != got {
			t.Fatalf("unexpected ASN(%d).Reserved():\n- want: %v\n-  got: %v",
				tt.a, want, got)
		}
	}
}
