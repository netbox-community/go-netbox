// Copyright 2017 The go-netbox Authors.
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
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
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

	// Append trailing slash there is none. This is necessary
	// to be able to concat url parts in a correct manner.
	// See NewRequest
	if !strings.HasSuffix(addr, "/") {
		addr = addr + "/"
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
	return c.NewDataRequest(method, endpoint, options, nil)
}

// NewDataRequest creates a HTTP request using the input HTTP method, URL
// endpoint, a Valuer which creates URL parameters for the request, and
// a io.Reader as the body of the request.
//
// If a nil Valuer is specified, no query parameters will be sent with the
// request.
//
// If a nil io.Reader is specified, no body will be sent with the request.
func (c *Client) NewDataRequest(method string, endpoint string, options Valuer, body io.Reader) (*http.Request, error) {
	// Allow specifying a base path for API requests, so if a NetBox server
	// resides at a path like http://example.com/netbox/, API requests will
	// be sent to http://example.com/netbox/api/...
	//
	// Enables support of: https://github.com/digitalocean/netbox/issues/212.
	//
	// Remove leading slash if there is one. This is necessary to be able to
	// concat url parts in a correct manner. We can not use path.Join here,
	// because this always trims the trailing slash, which causes the
	// Do function to always run into 301 and then retry the correct
	// Location. With GET, it does work with one useless request, but it breaks
	// each other http method.
	// Doing this, because out-of-tree extensions are more robust. If someone
	// implements an own API-call, we do not override parts of c.u, even if
	// the caller uses "/api/...".
	rel, err := url.Parse(strings.TrimLeft(endpoint, "/"))
	if err != nil {
		return nil, err
	}

	u := c.u.ResolveReference(rel)

	// If no valuer specified, create a request with no query parameters
	if options == nil {
		return http.NewRequest(method, u.String(), body)
	}

	values, err := options.Values()
	if err != nil {
		return nil, err
	}
	u.RawQuery = values.Encode()

	return http.NewRequest(method, u.String(), body)
}

// NewJSONRequest creates a HTTP request using the input HTTP method, URL
// endpoint, a Valuer which creates URL parameters for the request, and
// an io.Reader as the body of the request.
//
// If a nil Valuer is specified, no query parameters will be sent with the
// request.
//
// The body parameter is marshaled to JSON and sent as a HTTP request body.
// Body must not be nil.
func (c *Client) NewJSONRequest(method string, endpoint string, options Valuer, body interface{}) (*http.Request, error) {
	if body == nil {
		return nil, errors.New("expected body to be not nil")
	}

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(body)
	if err != nil {
		return nil, err
	}

	req, err := c.NewDataRequest(
		method,
		endpoint,
		options,
		b,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	return req, nil
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
	// Test if the request was successful.
	// If not, do not try to decode the body. This would result in
	// misleading error messages
	err = httpStatusOK(res)
	if err != nil {
		return err
	}

	if v == nil {
		return nil
	}

	return json.NewDecoder(res.Body).Decode(v)
}

// httpStatusOK tests if the StatusCode of res is smaller than 300. Tries to
// Unmarshal the response into json, and returns only the detail
// if this exists. Otherwise returns the status code with the raw Body data.
func httpStatusOK(res *http.Response) error {
	if res.StatusCode >= http.StatusMultipleChoices {
		errDetail := struct {
			Detail string `json:"detail"`
		}{}
		bodyData, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("%d - %v", res.StatusCode, err)
		}
		err = json.Unmarshal(bodyData, &errDetail)
		if err == nil && errDetail.Detail != "" {
			return fmt.Errorf("%d - %s", res.StatusCode, errDetail.Detail)
		}

		return fmt.Errorf("%d - %s", res.StatusCode, bodyData)
	}
	return nil
}
