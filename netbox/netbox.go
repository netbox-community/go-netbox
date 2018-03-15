package netbox

import (
	"fmt"

	"github.com/go-openapi/strfmt"
	runtimeclient "github.com/go-openapi/runtime/client"

	"github.com/digitalocean/go-netbox/netbox/client"
)

func NewNetboxAt(host string) *client.NetBox {
	t := client.DefaultTransportConfig().WithHost(host)
	return client.NewHTTPClientWithConfig(strfmt.Default, t)
}

const authHeaderName = "Authorization"
const authHeaderFormat = "Token %v"
func NewNetboxWithAPIKey(host string, apiToken string) *client.NetBox {
	t := runtimeclient.New(host, client.DefaultBasePath, client.DefaultSchemes)
	t.DefaultAuthentication = runtimeclient.APIKeyAuth(authHeaderName, "header", fmt.Sprintf(authHeaderFormat, apiToken))
	return client.New(t, strfmt.Default)
}
