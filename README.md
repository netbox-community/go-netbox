go-netbox 
=========

[![GoDoc](http://godoc.org/github.com/netbox-community/go-netbox?status.svg)](http://godoc.org/github.com/netbox-community/go-netbox) [![Build Status](https://github.com/netbox-community/go-netbox/workflows/main/badge.svg?branch=master)](https://github.com/netbox-community/go-netbox/actions) [![Report Card](https://goreportcard.com/badge/github.com/netbox-community/go-netbox)](https://goreportcard.com/report/github.com/netbox-community/go-netbox)

Package `netbox` provides an API 2.0 client for [netbox-community's NetBox](https://github.com/netbox-community/netbox)
IPAM and DCIM service.

This package assumes you are using NetBox 2.0, as the NetBox 1.0 API no longer exists.

Using the client
================

The `github.com/go-netbox/netbox` package has some convenience functions for creating clients with the most common
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
cp new_swagger_file.json swagger.json
make generate
```
