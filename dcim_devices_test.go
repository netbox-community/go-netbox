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

import "reflect"

func deviceEqual(a, b *Device) bool {
	var obsAB = []struct {
		a interface{}
		b interface{}
	}{
		{a: a.DeviceType, b: b.DeviceType},
		{a: a.DeviceRole, b: b.DeviceRole},
		{a: a.Platform, b: b.Platform},
		{a: a.Rack, b: b.Rack},
		{a: a.ParentDevice, b: b.ParentDevice},
		{a: a.ID, b: b.ID},
		{a: a.Name, b: b.Name},
		{a: a.DisplayName, b: b.DisplayName},
		{a: a.Serial, b: b.Serial},
		{a: a.Position, b: b.Position},
		{a: a.Face, b: b.Face},
		{a: a.Status, b: b.Status},
		{a: a.Comments, b: b.Comments},
		{a: a.PrimaryIP, b: b.PrimaryIP},
		{a: a.PrimaryIP4, b: b.PrimaryIP4},
		{a: a.PrimaryIP6, b: b.PrimaryIP6},
	}
	for _, o := range obsAB {

		switch o.a.(type) {
		case *IPAddressIdentifier:
			i, j := o.a.(*IPAddressIdentifier), o.b.(*IPAddressIdentifier)
			if ok := ipAddressIdentifiersEqual(*i, *j); !ok {
				return false
			}
		default:
			if !reflect.DeepEqual(o.a, o.b) {
				return false
			}
		}
	}

	return true
}
