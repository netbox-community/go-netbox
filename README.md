# go-netbox

[![GoDoc](https://pkg.go.dev/badge/github.com/netbox-community/go-netbox/v3)](https://pkg.go.dev/github.com/netbox-community/go-netbox/v3) [![Build Status](https://github.com/netbox-community/go-netbox/workflows/main/badge.svg?branch=master)](https://github.com/netbox-community/go-netbox/actions) [![Report Card](https://goreportcard.com/badge/github.com/netbox-community/go-netbox)](https://goreportcard.com/report/github.com/netbox-community/go-netbox)

_go-netbox_ is —to nobody's surprise— the official [Go](https://go.dev) API client for [Netbox](https://github.com/netbox-community/netbox) IPAM and DCIM service.

This project follows [Semantic Versioning](https://semver.org). The version of the library built for a Netbox version has the same tag, followed by a hyphen and the build number (an incremental integer), as several versions of the library may exist for the same version of Netbox.

## Installation

Use `go get` to add the library as a project's dependency. Do not forget to run `go mod init` first if necessary.

```shell
go get github.com/netbox-community/go-netbox/v3

# Or install a specific version
go get github.com/netbox-community/go-netbox/v3@v3.4.5-1
```

**Note:** dependencies should be managed with [Go modules](https://go.dev/doc/modules/managing-dependencies).

## Usage

### Instantiate the client

The package has some convenience functions for creating clients with the most common configurations.

```golang
package main

import (
	"log"

	"github.com/netbox-community/go-netbox/v3/netbox"
)

func main() {
	c := netbox.NewNetboxAt("netbox.example.org:8000")
    
	// or:
	// c := netbox.NewNetboxWithAPIKey("netbox.example.org:8000", "<api-token>")

	log.Printf("%+v", c)
}
```

In order to consume the Netbox API with HTTP over TLS, a transport must be created.

```golang
package main

import (
	"fmt"
	"log"

	transport "github.com/go-openapi/runtime/client"
	"github.com/netbox-community/go-netbox/v3/netbox/client"
)

func main() {
	t := transport.New("netbox.example.org", client.DefaultBasePath, []string{"https"})

	t.DefaultAuthentication = transport.APIKeyAuth(
		"Authorization",
		"header",
		fmt.Sprintf("Token %v", "<api-token>"),
	)

	c := client.New(t, nil)

	log.Printf("%+v", c)
}
```

For more complex client configurations, see the documentation of _[github.com/go-openapi/runtime/client](https://pkg.go.dev/github.com/go-openapi/runtime/client)_, the library which in turn is used by the client.

**Note:** setting the `DEBUG` environment variable will dump all requests to standard error output.

### Use the client

With the client already instantiated, it is possible to consume any API feature.

For example, to list the first 100 active virtual machines:

```golang
package main

import (
	"log"

	"github.com/netbox-community/go-netbox/v3/netbox"
	"github.com/netbox-community/go-netbox/v3/netbox/client/virtualization"
)

var status = "active"
var pageLimit = int64(100)

func main() {
	c := netbox.NewNetboxWithAPIKey("netbox.example.org", "<api-token>")

	req := virtualization.
		NewVirtualizationVirtualMachinesListParams().
		WithStatus(&status).
		WithLimit(&pageLimit)

	// additional `authInfo` is `nil` because the API token has already been specified in the client 
	res, err := w.netbox.Virtualization.VirtualizationVirtualMachinesList(req, nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", res.Payload.Results)
}
```

See [reference](https://pkg.go.dev/github.com/netbox-community/go-netbox) for more information on all possible usages.

## Development

The project comes with a containerized development environment that can be used from any platform. It is only required to have [Git](https://git-scm.com) and [Docker Desktop](https://www.docker.com/products/docker-desktop/) (or, separately, [Docker](https://docs.docker.com/engine/install) and [Docker Compose](https://docs.docker.com/compose/install/)) installed on the machine.

To start the development environment, run the following command.

```bash
make
```

Then, to attach a shell in the container, run the command below.

```bash
make shell
```

Finally, to stop the development environment, run the following command.

```bash
make down
```

### Considerations

The library is almost entirely generated from the Netbox [OpenAPI](https://www.openapis.org/) specification using _[go-swagger](https://github.com/go-swagger/go-swagger)_. Therefore, files under directories `netbox/client` and `netbox/models` should not be directly modified, as they will be overwritten in the next regeneration (see next section).

To fix issues in generated code, there are two options:

- Change the OpenAPI spec with pre-generation hooks (see [`scripts/pre-generation`](scripts/pre-generation)).
- If the above is not possible, change the generated code with post-generation hooks (see [`scripts/post-generation`](scripts/post-generation)).

### Regenerate the library

To update the OpenAPI specification to the latest Netbox version and regenerate the library, run the following command.

```bash
make build
```

If regeneration of the library is needed for a specific Netbox version other than the latest one, pass the corresponding argument.

```bash
make build NETBOX_VERSION=3.0.0
```

In order to obtain the OpenAPI specification, the version of _[netbox-docker](https://github.com/netbox-community/netbox-docker)_ corresponding to the given Netbox version is used. However, it is also possible to provide a specific version of _netbox-docker_.

```bash
make build NETBOX_VERSION=3.0.0 NETBOX_DOCKER_VERSION=1.3.1
```
