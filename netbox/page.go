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
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
)

// A Page contains all necessary information about a the current position
// in a paged result. It is used to walk over all pages from List calls.
type Page struct {
	limit    int
	offset   int
	done     bool
	c        *Client
	data     pageData
	endpoint string
	options  Valuer
	err      error
}

// NewPage Returns a new Page to walk over List calls.
func NewPage(client *Client, endpoint string, options Valuer) *Page {
	return &Page{
		limit:    50, // netbox server default
		c:        client,
		endpoint: endpoint,
		options:  options,
	}
}

// pageData is the internal representation of a page.
type pageData struct {
	Count       int             `json:"count"`
	NextURL     string          `json:"next"`
	PreviousURL string          `json:"previous"`
	Results     json.RawMessage `json:"results"`
}

// Values implements the Valuer interface. One could pass options
// to a List call. These must be extended by limit and offset to
// get the appropriate next page.
func (p *Page) Values() (url.Values, error) {
	if p == nil {
		return nil, errors.New("page not defined")
	}
	v := url.Values{}
	if p.options != nil {
		opts, err := p.options.Values()
		if err != nil {
			return nil, err
		}
		if opts != nil {
			v = opts
		}
	}
	v.Set("limit", strconv.Itoa(p.limit))
	v.Set("offset", strconv.Itoa(p.offset))

	return v, nil
}

// Next advances to the next page of API results.  When Next returns false,
// no more results are available.
func (p *Page) Next() bool {
	if p == nil {
		return false
	}
	if p.done {
		return false
	}

	req, err := p.c.NewRequest(http.MethodGet, p.endpoint, p)
	if err != nil {
		p.err = err
		return false
	}

	// Clear NextURL, because it will not be overriden if it is empty in the
	// resultset. I(cglaubitz) do not want to allocate a new PageData on each call
	// to Next() to reduce the need for garbage collection. No idea if this still
	// works when I have to "empty" NextURL here.
	p.data.NextURL = ""
	err = p.c.Do(req, &p.data)
	if err != nil {
		p.err = err
		return false
	}

	if p.data.NextURL == "" {
		// We are on the last page, so we still need to return true to
		// indicate that there is still data to process. But we set
		// p.done to true, so the next "Next()" returns false.
		p.done = true
	} else {
		p.setNextURL(p.data.NextURL)
	}
	return true

}

// setNext sets limit and offset parameter for the next page.
func (p *Page) setNext(limit int, offset int) {
	p.limit = limit
	p.offset = offset
}

// setNextURL extracts limit and offset from the nextURL, obtained from the result.
// Under the hood, it uses setNext to finally set those parameters.
func (p *Page) setNextURL(urlStr string) {
	nextURL, err := url.Parse(urlStr)
	if err != nil {
		// We dont want to cancel this run, since there is data,
		// but do not want to run into Next
		p.setErr(err)
		return
	}

	query, err := url.ParseQuery(nextURL.RawQuery)
	if err != nil {
		// Same like above
		p.setErr(err)
		return
	}

	limits := query["limit"]
	offsets := query["offset"]
	if len(limits) == 0 {
		p.setErr(errors.New("no such query parameter limit"))
		return
	}
	if len(offsets) == 0 {
		p.setErr(errors.New("no such query parameter offset"))
		return
	}

	limit, err := strconv.Atoi(limits[0])
	if err != nil {
		p.setErr(err)
		return
	}

	offset, err := strconv.Atoi(offsets[0])
	if err != nil {
		p.setErr(err)
		return
	}
	p.setNext(limit, offset)
}

// Err returns the internal err field. This should be called right after
// the for to get any errors occured during Next()
func (p *Page) Err() error {
	return p.err
}

// setErr sets an internal err field.
func (p *Page) setErr(err error) {
	if p.err == nil {
		p.err = err
	}
}
