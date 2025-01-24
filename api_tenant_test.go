package netbox

import (
	"context"
	"testing"
)

func TestListTenants(t *testing.T) {
	client := HGetClient(t)

	_, r, err := client.TenancyAPI.TenancyTenantsList(context.TODO()).Execute()
	if err != nil {
		if r != nil {
			t.Fatal(r)
		}
		t.Fatal(err)
	}
}
