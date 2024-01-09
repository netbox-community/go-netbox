package netbox

import (
	"fmt"
)

const authHeaderName = "Authorization"
const authHeaderFormat = "Token %v"

func NewAPIClientFor(host string, token string) *APIClient {
	cfg := NewConfiguration()
	cfg.Servers[0].URL = host
	cfg.DefaultHeader[authHeaderName] = fmt.Sprintf(authHeaderFormat, token)

	return NewAPIClient(cfg)
}
