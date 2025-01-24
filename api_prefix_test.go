package netbox

import (
	"context"
	"testing"
)

func TestListPrefixes(t *testing.T) {
	client := HGetClient(t)

	_, r, err := client.IpamAPI.IpamPrefixesList(context.TODO()).Execute()
	if err != nil {
		if r != nil {
			t.Fatal(r)
		}
		t.Fatal(err)
	}
}
