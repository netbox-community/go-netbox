# go-netbox

[![GoDoc](http://godoc.org/github.com/netbox-community/go-netbox?status.svg)](http://godoc.org/github.com/netbox-community/go-netbox) [![Build Status](https://github.com/netbox-community/go-netbox/workflows/main/badge.svg?branch=master)](https://github.com/netbox-community/go-netbox/actions) [![Report Card](https://goreportcard.com/badge/github.com/netbox-community/go-netbox)](https://goreportcard.com/report/github.com/netbox-community/go-netbox)

Package `netbox` provides an API 3.2 client for [netbox-community's Netbox](https://github.com/netbox-community/netbox) IPAM and DCIM service.

This package assumes you are using Netbox 3.2, as the Netbox 1.0 API no longer exists. If you need support for previous Netbox versions, you can still import the corresponding release of the library. For example, run `go get github.com/netbox-community/go-netbox@netbox_v2.11` if you need compatibility with Netbox 2.11.

## Using the client

The `github.com/netbox-community/go-netbox/netbox` package has some convenience functions for creating clients with the most common
configurations you are likely to need while connecting to NetBox. `NewNetboxAt` allows you to specify a hostname
(including port, if you need it), and `NewNetboxWithAPIKey` allows you to specify both a hostname:port and API token.
```golang
import (
    "github.com/netbox-community/go-netbox/netbox"
)
...
    c := netbox.NewNetboxAt("your.netbox.host:8000")
    // OR
    c := netbox.NewNetboxWithAPIKey("your.netbox.host:8000", "your_netbox_token")
```

If you specify the API key, you do not need to pass an additional `authInfo` to operations that need authentication, and
can pass `nil`:
```golang
    c.Dcim.DcimDeviceTypesCreate(createRequest, nil)
```

If you connect to netbox via HTTPS you have to create an HTTPS configured transport:
```golang
package main

import (
	"os"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/dcim"

	log "github.com/sirupsen/logrus"
)

func main() {
	token := os.Getenv("NETBOX_TOKEN")
	if token == "" {
		log.Fatalf("Please provide netbox API token via env var NETBOX_TOKEN")
	}

	netboxHost := os.Getenv("NETBOX_HOST")
	if netboxHost == "" {
		log.Fatalf("Please provide netbox host via env var NETBOX_HOST")
	}

	transport := httptransport.New(netboxHost, client.DefaultBasePath, []string{"https"})
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "Token "+token)

	c := client.New(transport, nil)

	req := dcim.NewDcimSitesListParams()
	res, err := c.Dcim.DcimSitesList(req, nil)
	if err != nil {
		log.Fatalf("Cannot get sites list: %v", err)
	}
	log.Infof("res: %v", res)
}
```

## Go Module support

Go 1.16+

`go get github.com/netbox-community/go-netbox`


## More complex client configuration

The client is generated using [go-swagger](https://github.com/go-swagger/go-swagger). This means the generated client
makes use of [github.com/go-openapi/runtime/client](https://godoc.org/github.com/go-openapi/runtime/client). If you need
a more complex configuration, it is probably possible with a combination of this generated client and the runtime
options.

The [godocs for the go-openapi/runtime/client module](https://godoc.org/github.com/go-openapi/runtime/client) explain
the client options in detail, including different authentication and debugging options. One thing I want to flag because
it is so useful: setting the `DEBUG` environment variable will dump all requests to standard out.

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

The library is almost entirely generated from the Netbox [OpenAPI](https://www.openapis.org/) specification. Therefore, files under directories `netbox/client` and `netbox/models` should not be directly modified, as they will be overwritten in the next regeneration (see next section).

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
