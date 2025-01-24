package netbox

import (
	"context"
	"io"
	"testing"
)

func TestListNetboxSites(t *testing.T) {
	client := HGetClient(t)

	_, r, err := client.DcimAPI.DcimSitesList(context.TODO()).Execute()
	if err != nil {
		if r != nil {
			bodyBytes, _ := io.ReadAll(r.Body)
			t.Error(string(bodyBytes))
			t.Fatal(err)
		}
	}
}
