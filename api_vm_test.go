package netbox

import (
	"context"
	"testing"
)

func TestVirtualMachineList(t *testing.T) {
	client := HGetClient(t)

	_, r, err := client.VirtualizationAPI.VirtualizationVirtualMachinesList(context.TODO()).Execute()
	if err != nil {
		if r != nil || r.StatusCode != 200 {
			t.Fatal(r)
		}
		t.Fatal(err)
	}
}
