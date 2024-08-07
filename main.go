package netbox

import (
	"fmt"
)

const authHeaderName = "Authorization"
const authHeaderFormat = "Token %v"
const languageHeaderName = "Accept-Language"
const languageHeaderValue = "en-US"

func NewAPIClientFor(host string, token string) *APIClient {
	cfg := NewConfiguration()

	cfg.Servers[0].URL = host

	cfg.AddDefaultHeader(
		authHeaderName,
		fmt.Sprintf(authHeaderFormat, token),
	)

	cfg.AddDefaultHeader(
		languageHeaderName,
		languageHeaderValue,
	)

	return NewAPIClient(cfg)
}
