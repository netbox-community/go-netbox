package netbox

import (
	"fmt"
)

const (
	authHeaderName   = "Authorization"
	authHeaderFormat = "Token %v"
)

func NewAPIClientFor(host string, token string) *APIClient {
	cfg := NewConfiguration()

	cfg.Servers[0].URL = host

	cfg.AddDefaultHeader(
		authHeaderName,
		fmt.Sprintf(authHeaderFormat, token),
	)

	return NewAPIClient(cfg)
}
