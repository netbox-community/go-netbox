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

// Package netbox provides an API client for DigitalOcean's NetBox IPAM and
// DCIM service.
package netbox

import "net/url"

// A Valuer is an object which can generate a url.Values map from itself.
// Valuer implementations are used to generate request URL parameters
// that NetBox can use to filter data.
type Valuer interface {
	values() (url.Values, error)
}

// A Family is an IP address family, used by NetBox to filter either IPv4
// or IPv6 addresses.  Use its Valid method or compare against the predefined
// Family constants to determine if the contained value is a valid Family.
type Family int

// Family constants which can be used with NetBox.
const (
	FamilyIPv4 Family = 4
	FamilyIPv6 Family = 6
)

// String returns the string representation of a Family.
func (f Family) String() string {
	switch f {
	case FamilyIPv4:
		return "IPv4"
	case FamilyIPv6:
		return "IPv6"
	default:
		return "Unknown"
	}
}

// Valid determines if a Family is valid.
func (f Family) Valid() bool {
	return f == FamilyIPv4 || f == FamilyIPv6
}

// A Status is an operational status of an object.
type Status int

const (
	// StatusContainer indicates an object is a summary of child prefixes.
	StatusContainer Status = 0

	// StatusActive indicates an object is active and in use.
	StatusActive Status = 1

	// StatusReserved indicates an object is reserved for future use.
	StatusReserved Status = 2

	// StatusDeprecated indicates an object is no longer in use.
	StatusDeprecated Status = 3
)

// String returns the string representation of a Status.
func (s Status) String() string {
	switch s {
	case StatusContainer:
		return "Container"
	case StatusActive:
		return "Active"
	case StatusReserved:
		return "Reserved"
	case StatusDeprecated:
		return "Deprecated"
	default:
		return "Unknown"
	}
}
