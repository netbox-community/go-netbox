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

func deviceEqual(i, j *Device) bool {
	if i.ID != j.ID {
		return false
	}
	if i.Name != j.Name {
		return false
	}
	if i.DisplayName != j.DisplayName {
		return false
	}
	if i.DeviceType.ID != j.DeviceType.ID {
		return false
	}
	if i.DeviceType.Model != j.DeviceType.Model {
		return false
	}
	if i.DeviceType.Slug != j.DeviceType.Slug {
		return false
	}
	if i.DeviceType.Manufacturer.ID != j.DeviceType.Manufacturer.ID {
		return false
	}
	if i.DeviceType.Manufacturer.Name != j.DeviceType.Manufacturer.Name {
		return false
	}
	if i.DeviceType.Manufacturer.Slug != j.DeviceType.Manufacturer.Slug {
		return false
	}
	if i.DeviceRole.ID != j.DeviceRole.ID {
		return false
	}
	if i.DeviceRole.Name != j.DeviceRole.Name {
		return false
	}
	if i.DeviceRole.Slug != j.DeviceRole.Slug {
		return false
	}
	if i.Platform.ID != j.Platform.ID {
		return false
	}
	if i.Platform.Name != j.Platform.Name {
		return false
	}
	if i.Platform.Slug != j.Platform.Slug {
		return false
	}
	if i.Serial != j.Serial {
		return false
	}
	if i.Rack.ID != j.Rack.ID {
		return false
	}
	if i.Rack.Name != j.Rack.Name {
		return false
	}
	if i.Rack.DisplayName != j.Rack.DisplayName {
		return false
	}
	if i.Rack.FacilityID != j.Rack.FacilityID {
		return false
	}
	if i.Position != j.Position {
		return false
	}
	if i.Face != j.Face {
		return false
	}
	if i.ParentDevice.ID != j.ParentDevice.ID {
		return false
	}
	if i.ParentDevice.Name != j.ParentDevice.Name {
		return false
	}
	if i.Status != j.Status {
		return false
	}
	if i.PrimaryIP.ID != j.PrimaryIP.ID {
		return false
	}
	if i.PrimaryIP.Family != j.PrimaryIP.Family {
		return false
	}
	if i.PrimaryIP.Address.String() != j.PrimaryIP.Address.String() {
		return false
	}
	if i.PrimaryIP4.ID != j.PrimaryIP4.ID {
		return false
	}
	if i.PrimaryIP4.Family != j.PrimaryIP4.Family {
		return false
	}
	if i.PrimaryIP4.Address.String() != j.PrimaryIP4.Address.String() {
		return false
	}
	if i.PrimaryIP6.ID != j.PrimaryIP6.ID {
		return false
	}
	if i.PrimaryIP6.Family != j.PrimaryIP6.Family {
		return false
	}
	if i.PrimaryIP6.Address.String() != j.PrimaryIP6.Address.String() {
		return false
	}
	if i.Comments != j.Comments {
		return false
	}
	return true
}
