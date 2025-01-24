package netbox

import (
	"os"
	"testing"
)

func HGetClient(t *testing.T) *APIClient {
	t.Helper()
	srv := os.Getenv("NETBOX_URL")
	token := os.Getenv("NETBOX_TOKEN")

	if srv == "" {
		t.Fatal("NETBOX_URL is not set")
	}
	if token == "" {
		t.Fatal("NETBOX_TOKEN is not set")
	}

	cfg := NewConfiguration()
	cfg.Servers[0].URL = srv
	cfg.AddDefaultHeader("Authorization", "Token "+token)

	client := NewAPIClient(cfg)
	return client

}
