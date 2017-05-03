// Copyright 2016 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package netbox

import (
	"encoding/json"
	"net/http"
	"net/url"
	"path"
)

// A Client is a NetBox client.  It can be used to retrieve network and
// datacenter infrastructure information from a NetBox server.
type Client struct {
	// DCIM provides access to methods in NetBox's DCIM API.
	DCIM *DCIMService

	// IPAM provides access to methods in NetBox's IPAM API.
	IPAM *IPAMService

	u      *url.URL
	client *http.Client
}

// NewClient returns a new instance of a NetBox client.  addr specifies the address
// of the NetBox server, and client specifies an optional HTTP client to use
// for requests.
//
// If client is nil, a default HTTP client will be used.
func NewClient(addr string, client *http.Client) (*Client, error) {
	if client == nil {
		client = &http.Client{}
	}

	u, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}

	c := &Client{
		u:      u,
		client: client,
	}

	c.DCIM = &DCIMService{c: c}
	c.IPAM = &IPAMService{c: c}

	return c, nil
}

// NewRequest creates a HTTP request using the input HTTP method, URL
// endpoint, and a Valuer which creates URL parameters for the request.
//
// If a nil Valuer is specified, no query parameters will be sent with the
// request.
func (c *Client) NewRequest(method string, endpoint string, options Valuer) (*http.Request, error) {
	rel, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	// Allow specifying a base path for API requests, so if a NetBox server
	// resides at a path like http://example.com/netbox/, API requests will
	// be sent to http://example.com/netbox/api/...
	//
	// Enables support of: https://github.com/digitalocean/netbox/issues/212.
	if c.u.Path != "" {
		rel.Path = path.Join(c.u.Path, rel.Path)
	}

	u := c.u.ResolveReference(rel)

	// If no valuer specified, create a request with no query parameters
	if options == nil {
		return http.NewRequest(method, u.String(), nil)
	}

	values, err := options.Values()
	if err != nil {
		return nil, err
	}
	u.RawQuery = values.Encode()

	return http.NewRequest(method, u.String(), nil)
}

// Do executes an HTTP request and if v is not nil, Do unmarshals result
// JSON onto v.
func (c *Client) Do(req *http.Request, v interface{}) error {
	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	if v == nil {
		return nil
	}

	return json.NewDecoder(res.Body).Decode(v)
}
