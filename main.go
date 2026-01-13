package netbox

import (
	"fmt"
	"strings"
)

const authHeaderName = "Authorization"
const languageHeaderName = "Accept-Language"
const languageHeaderValue = "en-US"
const v1HeaderFormat = "Token %s"
const v2HeaderFormat = "Bearer %s"
const v2Prefix = "nbt_"

func NewAPIClientFor(host string, token string) *APIClient {
	cfg := NewConfiguration()

	cfg.Servers[0].URL = host

	cfg.AddDefaultHeader(
		authHeaderName,
		GetFormattedAPIToken(token),
	)

	cfg.AddDefaultHeader(
		languageHeaderName,
		languageHeaderValue,
	)

	return NewAPIClient(cfg)
}

func GetFormattedAPIToken(token string) string {
	var authHeaderFormat string
	if strings.HasPrefix(token, v2Prefix) {
		authHeaderFormat = v2HeaderFormat
	} else {
		authHeaderFormat = v1HeaderFormat
	}
	return fmt.Sprintf(authHeaderFormat, token)
}
