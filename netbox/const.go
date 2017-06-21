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
