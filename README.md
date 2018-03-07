netbox [![GoDoc](http://godoc.org/github.com/digitalocean/go-netbox?status.svg)](http://godoc.org/github.com/digitalocean/go-netbox) [![Build Status](https://travis-ci.org/digitalocean/go-netbox.svg?branch=master)](https://travis-ci.org/digitalocean/go-netbox) [![Report Card](https://goreportcard.com/badge/github.com/digitalocean/go-netbox)](https://goreportcard.com/report/github.com/digitalocean/go-netbox)
======

Package `netbox` provides an API 2.0 client for [DigitalOcean's NetBox](https://github.com/digitalocean/netbox)
IPAM and DCIM service.

This package assumes you are using NetBox 2.0, as the NetBox 1.0 API no longer exists.

Using the client
================

The `github.com/go-netbox/netbox/client` package is the entry point for using the client. By default, the client will
connect to an api at `http://localhost:8000/api`. The simplest possible usage looks like:
```golang
    // Passing nil to this method results in all default values
	c := client.NewHTTPClient(nil)

	... work with the returned client ...
```

A more likely scenario is to connect to a remote netbox:
```golang
	t := client.DefaultTransportConfig().WithHost("your.netbox.host")
	c := client.NewHTTPClientWithConfig(nil, t)
```

The client is generated using [go-swagger](https://github.com/go-swagger/go-swagger). This means the generated client
makes use of [github.com/go-openapi/runtime/client](https://godoc.org/github.com/go-openapi/runtime/client). The [godocs
for that module](https://godoc.org/github.com/go-openapi/runtime/client) explain the client options in detail, including
different authentication and debugging options.

Setting the debug flag will print all requests and responses on standard out, which is great for debugging unexpected
results. It does require creating the client in the lower level `go-openapi` libraries:
```golang
import (
	"github.com/digitalocean/go-netbox/netbox/client"
	"github.com/go-openapi/strfmt"
	runtimeclient "github.com/go-openapi/runtime/client"
)

func main() {
	t := runtimeclient.New(client.DefaultHost, client.DefaultBasePath, client.DefaultSchemes)
	t.SetDebug(true)
	c := client.New(t, strfmt.Default)

	... work with c ...
)
```

Regenerating the client
=======================

To regenerate the client with a new or different swagger schema, first clean the existing client, then replace
swagger.json and finally re-generate:
```
make clean
cp new_swagger_file.json swagger.json
make generate
```
