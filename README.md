# go-netbox

[![GoDoc](https://pkg.go.dev/badge/github.com/netbox-community/go-netbox/v4)](https://pkg.go.dev/github.com/netbox-community/go-netbox/v4) [![Build Status](https://github.com/netbox-community/go-netbox/workflows/main/badge.svg?branch=master)](https://github.com/netbox-community/go-netbox/actions) [![Report Card](https://goreportcard.com/badge/github.com/netbox-community/go-netbox)](https://goreportcard.com/report/github.com/netbox-community/go-netbox)

_go-netbox_ is —to nobody's surprise— the official [Go](https://go.dev) API client for the [Netbox](https://github.com/netbox-community/netbox) IPAM and DCIM service.

This project follows [Semantic Versioning](https://semver.org). The version of the library built for a Netbox version has the same tag, followed by a hyphen and the build number (an incremental integer), as several versions of the library may exist for the same version of Netbox.

## Installation

Use `go get` to add the library as a dependency to your project. Do not forget to run `go mod init` first if necessary.

```shell
go get github.com/netbox-community/go-netbox/v4

# Or install a specific version
go get github.com/netbox-community/go-netbox/v4@v4.2.2-0
```

**Note:** dependencies should be managed with [Go modules](https://go.dev/doc/modules/managing-dependencies).

## Usage

### Instantiate the client

The package has a constructor for creating a client by providing a URL and an authentication token.

```golang
package main

import (
	"context"
	"log"

	"github.com/netbox-community/go-netbox/v4"
)

func main() {
	ctx := context.Background()

	c := netbox.NewAPIClientFor("https://demo.netbox.dev", "v3ry$3cr3t")

	log.Printf("%+v", c)
}

```

### Use the client

With the client already instantiated, it is possible to consume any API feature.

For example, to list the first 10 active virtual machines:

```golang
package main

import (
	"context"
	"log"

	"github.com/netbox-community/go-netbox/v4"
)

func main() {
	ctx := context.Background()

	c := netbox.NewAPIClientFor("https://demo.netbox.dev", "v3ry$3cr3t")

	res, _, err := c.VirtualizationAPI.
		VirtualizationVirtualMachinesList(ctx).
		Status([]string{"active"}).
		Limit(10).
		Execute()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v", res.Results)
}
```

See [docs](docs) or [reference](https://pkg.go.dev/github.com/netbox-community/go-netbox) for more information on all possible usages.

## Development

The project comes with a containerized development environment that may be used on any platform. It is only required to have [Git](https://git-scm.com) and [Docker Desktop](https://www.docker.com/products/docker-desktop/) (or, separately, [Docker](https://docs.docker.com/engine/install) and [Docker Compose](https://docs.docker.com/compose/install/)) installed on the machine.

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

The library is entirely generated from the Netbox [OpenAPI](https://www.openapis.org/) specification using _[openapi-generator](https://github.com/OpenAPITools/openapi-generator)_. Therefore, files listed [here](.openapi-generator/files) should not be directly modified, as they will be overwritten in the next regeneration (see next section).

In order to fix a bug in the generated code, the corresponding error in the OpenAPI spec must be fixed. To do so, the following steps may be followed:

1. Optional. Patch the OpenAPI spec in this repo by editing [this script](scripts/fix-spec.py), so that a corrected version can be published as soon as possible.
2. Fix the OpenAPI spec in the [Netbox repository](https://github.com/netbox-community/netbox), either by reporting an issue or by creating a pull request.

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
