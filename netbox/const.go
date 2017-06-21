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

// Describe Interface Status
const (
	InterfaceStatusOffline   = 0
	InterfaceStatusActive    = 1
	InterfaceStatusPlanned   = 2
	InterfaceStatusStaged    = 3
	InterfaceStatusFailed    = 4
	InterfaceStatusInventory = 5
)

// Describe the face of a device in the rack
const (
	RackFaceFront = 0
	RackFaceRear  = 1
)

// Describe Interface's Form Factor
const (
	InterfaceFormFactorVirtual       = 0
	InterfaceFormFactorLAG           = 200
	InterfaceFormFactor100MEFixed    = 800
	InterfaceFormFactor1GEFixed      = 1000
	InterfaceFormFactor1GEGBIC       = 1050
	InterfaceFormFactor1GESFP        = 1100
	InterfaceFormFactor10GEFixed     = 1150
	InterfaceFormFactor10GESFPPlus   = 1200
	InterfaceFormFactor10GEXFP       = 1300
	InterfaceFormFactor10GEXENPAK    = 1310
	InterfaceFormFactor10GEX2        = 1320
	InterfaceFormFactor25GESFP28     = 1350
	InterfaceFormFactor40GEQSFPPlus  = 1400
	InterfaceFormFactor100GECFP      = 1500
	InterfaceFormFactor100GEQSFP28   = 1600
	InterfaceFormFactor1GFCSFP       = 3010
	InterfaceFormFactor2GFCSFP       = 3020
	InterfaceFormFactor4GFCSFP       = 3040
	InterfaceFormFactor8GFCSFPPlus   = 3080
	InterfaceFormFactor16GFCSFPPlus  = 3160
	InterfaceFormFactorT1            = 4000
	InterfaceFormFactorE1            = 4010
	InterfaceFormFactorT3            = 4040
	InterfaceFormFactorE3            = 4050
	InterfaceFormFactorStackWise     = 5000
	InterfaceFormFactorStackWisePlus = 5050
	InterfaceFormFactorFlexStack     = 5100
	InterfaceFormFactorFlexStackPlus = 5150
	InterfaceFormFactorJuniperVCP    = 5200
	InterfaceFormFactorOther         = 32767
)
