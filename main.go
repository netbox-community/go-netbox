package netbox

const authHeaderName = "Authorization"
const languageHeaderName = "Accept-Language"
const languageHeaderValue = "en-US"

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
