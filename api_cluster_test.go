package netbox

import (
	"context"
	"fmt"
	"testing"
)

func TestVirtualizationClustersList(t *testing.T) {
	client := HGetClient(t)

	_, r, err := client.VirtualizationAPI.VirtualizationClustersList(context.TODO()).Execute()
	if err != nil {
		if r != nil {
			fmt.Println("body")
		}
		t.Fatal(err)
	}
}
