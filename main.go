package netbox

//go:generate oapi-codegen -package netbox -generate types -o types.go api/openapi.yaml
//go:generate oapi-codegen -package netbox -generate client -o client.go api/openapi.yaml

import (
	"fmt"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
)

const authHeaderName = "Authorization"
const authHeaderFormat = "Token %v"

func NewClientWithToken(host string, token string) (*ClientWithResponses, error) {
	provider, err := securityprovider.NewSecurityProviderApiKey(
		"header",
		authHeaderName,
		fmt.Sprintf(authHeaderFormat, token),
	)

	if err != nil {
		return nil, err
	}

	return NewClientWithResponses(
		host,
		WithRequestEditorFn(provider.Intercept),
	)
}
