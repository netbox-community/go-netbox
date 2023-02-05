go-netbox 
=========

[![GoDoc](http://godoc.org/github.com/perimeter-81/go-netbox?status.svg)](http://godoc.org/github.com/perimeter-81/go-netbox) [![Build Status](https://github.com/perimeter-81/go-netbox/workflows/main/badge.svg?branch=master)](https://github.com/perimeter-81/go-netbox/actions) [![Report Card](https://goreportcard.com/badge/github.com/perimeter-81/go-netbox)](https://goreportcard.com/report/github.com/perimeter-81/go-netbox)

Package `netbox` provides an API 3.2 client for [perimeter-81's Netbox](https://github.com/perimeter-81/netbox) IPAM and DCIM service.

This package assumes you are using Netbox 3.2, as the Netbox 1.0 API no longer exists. If you need support for previous Netbox versions, you can still import the corresponding release of the library. For example, run `go get github.com/perimeter-81/go-netbox@netbox_v2.11` if you need compatibility with Netbox 2.11.

Using the client
================

The `github.com/perimeter-81/go-netbox/netbox` package has some convenience functions for creating clients with the most common
configurations you are likely to need while connecting to NetBox. `NewNetboxAt` allows you to specify a hostname
(including port, if you need it), and `NewNetboxWithAPIKey` allows you to specify both a hostname:port and API token.
```golang
import (
    "github.com/perimeter-81/go-netbox/netbox"
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
	"github.com/perimeter-81/go-netbox/netbox/client"
	"github.com/perimeter-81/go-netbox/netbox/client/dcim"

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

Go Module support
================

Go 1.16+

`go get github.com/perimeter-81/go-netbox`


More complex client configuration
=================================

The client is generated using [go-swagger](https://github.com/go-swagger/go-swagger). This means the generated client
makes use of [github.com/go-openapi/runtime/client](https://godoc.org/github.com/go-openapi/runtime/client). If you need
a more complex configuration, it is probably possible with a combination of this generated client and the runtime
options.

The [godocs for the go-openapi/runtime/client module](https://godoc.org/github.com/go-openapi/runtime/client) explain
the client options in detail, including different authentication and debugging options. One thing I want to flag because
it is so useful: setting the `DEBUG` environment variable will dump all requests to standard out.

Regenerating the client
=======================

To regenerate the client with a new or different swagger schema, first clean the existing client, then replace
swagger.json and finally re-generate:
```
make clean
./scripts/swagger_modifier.py new_swagger_file.json
mv swagger_transformed.json swagger.json
make generate
```
