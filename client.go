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
)

// A Client is a NetBox client.  It can be used to retrieve network and
// datacenter infrastructure information from a NetBox server.
type Client struct {
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

	c.IPAM = &IPAMService{c: c}

	return c, nil
}

// newRequest creates a HTTP request using the input HTTP method, URL
// endpoint, and a valuer which creates URL parameters for the request.
//
// If a nil valuer is specified, no query parameters will be sent with the
// request.
func (c *Client) newRequest(method string, endpoint string, options valuer) (*http.Request, error) {
	rel, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	u := c.u.ResolveReference(rel)

	// If no valuer specified, create a request with no query parameters
	if options == nil {
		return http.NewRequest(method, u.String(), nil)
	}

	values, err := options.values()
	if err != nil {
		return nil, err
	}
	u.RawQuery = values.Encode()

	return http.NewRequest(method, u.String(), nil)
}

// do executes an HTTP request and unmarshals result JSON onto v.
func (c *Client) do(req *http.Request, v interface{}) error {
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
