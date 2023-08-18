go-netbox 
=========

[![GoDoc](http://godoc.org/github.com/fbreckle/go-netbox?status.svg)](http://godoc.org/github.com/fbreckle/go-netbox) [![Build Status](https://github.com/fbreckle/go-netbox/workflows/main/badge.svg?branch=master)](https://github.com/fbreckle/go-netbox/actions) [![Report Card](https://goreportcard.com/badge/github.com/fbreckle/go-netbox)](https://goreportcard.com/report/github.com/fbreckle/go-netbox)

Package `netbox` provides an API 2.0 client for [netbox-community's NetBox](https://github.com/netbox-community/netbox)
IPAM and DCIM service.

This package assumes you are using NetBox 2.0, as the NetBox 1.0 API no longer exists.


Why this fork exists
====================

This fork exists solely to support [e-breuninger/terraform-provider-netbox](https://github.com/e-breuninger/terraform-provider-netbox). As such, some changes in this fork do only make sense in that context.


About the state of this fork
============================

For a long time, the approach of taking the original netbox-provided swaggerfile, processing it and then generating a client from the processed swagger file was the only acceptable way for me to manage changes in this client.
The main advantages of this was that changes in the netbox api in newer versions could mostly be implemented fast, while the fixes that were necessary to make the terraform provider work were always present in the code.

With NetBox 3.5, the NetBox authors decided to change the api documentation from openapi2 to openapi3. This means that the code generator that is being used in this project can no longer be used.

Switching the code generator will also induce the need to change every single api call in the terraform provider. Furthermore, I am no expert whatsoever on openapi (or even Go), so evaluating the different available openapi generators and then choosing one to use is not an endeavour I can to pursue right now. This means we cannot re-generate the client with newer api specs right now or in the immediate future.

In conclusion, at the time of writing this, NetBox 3.5 is out quite some time already with NetBox 3.6 around the corner, so I decided that the show must go on and manual adjustments to this client are now the fastest way to allow the terraform provider to support newer NetBox versions.

This will be re-evaluated if the list of to-dos in the netbox provider gets smaller (e.g. all newish versions are supported, custom fields are solved globally).

Links:

[Original announcement of openapi3 change](https://github.com/netbox-community/netbox/discussions/11808)

[Issue](https://github.com/netbox-community/go-netbox/issues/155) and [discussion](https://github.com/netbox-community/go-netbox/discussions/156) in https://github.com/netbox-community/go-netbox


Versioning
==========

tbd. Meanwhile, look at branches and tags.


Changes in this fork
====================

NOTE: The list of changes is becoming too inconvenient to list here. See `preprocess.py` for all changes. The list below is incomplete at best or flat-out wrong!

Change `models.ip_address.AssignedObject` type to prevent json marshalling errors since [this change](https://github.com/netbox-community/netbox/pull/4781)

Add `x-omitempty: false` to some attributes, allowing them to be set to their empty value. [issue](https://github.com/netbox-community/go-netbox/issues/107)

Change ConfigContext type for VMs and Devices [#2](https://github.com/fbreckle/go-netbox/pull/2)

Fix LocalConfigContext to support arbitrary JSON object [#4](https://github.com/fbreckle/go-netbox/pull/4)

Changes to allow `available_ips` to be created. (2b418d0d6d13d1edd0320d9424096f8f7d2cbbec)
1. model returned by paths /*/available-ips/ from AvailableIP to IPAddress
2. model IPAddress's assigned_object property from string to object otherwise it fails to unmarshal


Using the client
================

The `github.com/fbreckle/go-netbox/netbox` package has some convenience functions for creating clients with the most common
configurations you are likely to need while connecting to NetBox. `NewNetboxAt` allows you to specify a hostname
(including port, if you need it), and `NewNetboxWithAPIKey` allows you to specify both a hostname:port and API token.
```golang
import (
    "github.com/fbreckle/go-netbox/netbox"
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
```
package main

import (
	"os"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"

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

Go 1.13+

`go get github.com/fbreckle/go-netbox`


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
swagger.json, run the json preprocessor (requires python3) and finally re-generate:
```
make clean
make preprocess
make generate
```
