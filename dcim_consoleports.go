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

// ConsolePort represents a console port object.
type ConsolePort struct {
	ID               int                `json:"id"`
	Device           *DeviceIdentifier  `json:"device"`
	Name             string             `json:"name"`
	CSPort           *ConsoleServerPort `json:"cs_port"`
	ConnectionStatus bool               `json:"connection_status"`
}

// ConsoleServerPort represents a console server port object.
type ConsoleServerPort struct {
	ID     int               `json:"id"`
	Device *DeviceIdentifier `json:"device"`
	Name   string            `json:"name"`
}
