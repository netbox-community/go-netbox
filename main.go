package netbox

//go:generate oapi-codegen -package netbox -generate types -o types.go api/openapi.yaml
//go:generate oapi-codegen -package netbox -generate client -o client.go api/openapi.yaml
