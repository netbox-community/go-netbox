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

// A DCIMService is used in a Client to access NetBox's DCIM API methods.
type DCIMService struct {
	c              *Client
	InventoryItems *InventoryItemsService
	Devices        *DevicesService
	Interfaces     *InterfacesService
}

// NewDCIMService returns a DCIMService initialized with all sub-services.
func NewDCIMService(client *Client) *DCIMService {
	return &DCIMService{
		c: client,
		InventoryItems: &InventoryItemsService{
			c: client,
		},
		Devices: &DevicesService{
			c: client,
		},
		Interfaces: &InterfacesService{
			c: client,
		},
	}
}

// SimpleIdentifier represents a simple object that consists of only an ID,
// name, and slug.
type SimpleIdentifier struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
