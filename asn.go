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

// An ASN is an Autonomous System Number, used to designate that a collection of
// Internet Protocol routing prefixes are under the control of one or more
// network operators.
type ASN int

// Reserved and private ASN ranges, used in filters below.
// Reference: http://www.iana.org/assignments/as-numbers/as-numbers.xhtml.
const (
	reservedStart16Bit  = 0
	reservedMin16Bit    = 64297
	reservedMax16Bit    = 64495
	reservedDocMin16Bit = 64496
	reservedDocMax16Bit = 64511
	reservedEnd16Bit    = 65535

	reservedDocMin32Bit = 65536
	reservedDocMax32Bit = 65551
	reservedMin32Bit    = 65552
	reservedMax32Bit    = 131071
	reservedEnd32Bit    = 4294967295

	privateMin16Bit = 64512
	privateMax16Bit = 65534

	privateMin32Bit = 4200000000
	privateMax32Bit = 4294967294

	// RFC 6793, Section 9, AS_TRANS, reserved
	asTrans = 23456
)

// Private determines if an ASN resides within a private range.
func (a ASN) Private() bool {
	switch {
	case a >= privateMin16Bit && a <= privateMax16Bit:
		return true
	case a >= privateMin32Bit && a <= privateMax32Bit:
		return true
	default:
		return false
	}
}

// Public determines if an ASN does not reside within a private or reserved
// range.
func (a ASN) Public() bool {
	if !a.Valid() {
		return false
	}

	return !a.Private() && !a.Reserved()
}

// Reserved determines if an ASN resides within a reserved range.
func (a ASN) Reserved() bool {
	switch {
	case a == reservedStart16Bit:
		return true
	case a == asTrans:
		return true
	case a >= reservedMin16Bit && a <= reservedMax16Bit:
		return true
	case a >= reservedDocMin16Bit && a <= reservedDocMax16Bit:
		return true
	case a == reservedEnd16Bit:
		return true
	case a >= reservedDocMin32Bit && a <= reservedDocMax32Bit:
		return true
	case a >= reservedMin32Bit && a <= reservedMax32Bit:
		return true
	case a == reservedEnd32Bit:
		return true
	default:
		return false
	}
}

// Valid determines if an ASN is valid.
func (a ASN) Valid() bool {
	return a >= reservedStart16Bit && a <= reservedEnd32Bit
}
