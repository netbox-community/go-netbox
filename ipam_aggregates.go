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
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// An Aggregate is an IPv4 or IPv6 address aggregate.
type Aggregate struct {
	ID          int
	Family      Family
	Prefix      *net.IPNet
	RIR         *RIRIdentifier
	DateAdded   time.Time
	Description string
}

// GetAggregate retrieves an Aggregate object from NetBox by its ID.
func (s *IPAMService) GetAggregate(id int) (*Aggregate, error) {
	req, err := s.c.NewRequest(
		http.MethodGet,
		fmt.Sprintf("/api/ipam/aggregates/%d", id),
		nil,
	)
	if err != nil {
		return nil, err
	}

	a := new(Aggregate)
	err = s.c.Do(req, a)
	return a, err
}

// ListAggregates retrives a list of Aggregate objects from NetBox, filtered
// according to the parameters specified in options.
//
// If options is nil, all Aggregates will be retrieved.
func (s *IPAMService) ListAggregates(options *ListAggregatesOptions) ([]*Aggregate, error) {
	req, err := s.c.NewRequest(http.MethodGet, "/api/ipam/aggregates/", options)
	if err != nil {
		return nil, err
	}

	var as []*Aggregate
	err = s.c.Do(req, &as)
	return as, err
}

// ListAggregatesOptions is a used as an argument for Client.IPAM.ListAggregates.
// Integer fields with an *ID suffix are preferred over their string counterparts,
// and if both are set, only the *ID field will be used.
type ListAggregatesOptions struct {
	Family    Family
	RIRID     []int
	RIR       []string
	DateAdded time.Time
}

// Values generates a url.Values map from the data in ListAggregatesOptions.
func (o *ListAggregatesOptions) Values() (url.Values, error) {
	if o == nil {
		return nil, nil
	}

	v := url.Values{}

	if o.Family != 0 {
		v.Set("family", strconv.Itoa(int(o.Family)))
	}

	// IDs should always be preferred over string names

	switch {
	case len(o.RIRID) > 0:
		for _, r := range o.RIRID {
			v.Add("rir_id", strconv.Itoa(r))
		}
	case len(o.RIR) > 0:
		for _, r := range o.RIR {
			v.Add("rir", r)
		}
	}

	if !o.DateAdded.IsZero() {
		v.Set("date_added", o.DateAdded.Format(dateFormat))
	}

	return v, nil
}

// dateFormat is the package time equivalent of the date format used by NetBox.
const dateFormat = "2006-01-02"

// An aggregate is the raw JSON representation of an Aggregate.
type aggregate struct {
	ID          int            `json:"id"`
	Family      Family         `json:"family"`
	Prefix      string         `json:"prefix"`
	RIR         *RIRIdentifier `json:"rir"`
	DateAdded   string         `json:"date_added"`
	Description string         `json:"description"`
}

// MarshalJSON marshals an Aggregate into JSON bytes.
func (a *Aggregate) MarshalJSON() ([]byte, error) {
	var date string
	if !a.DateAdded.IsZero() {
		date = a.DateAdded.Format(dateFormat)
	}

	return json.Marshal(aggregate{
		ID:          a.ID,
		Family:      a.Family,
		Prefix:      a.Prefix.String(),
		RIR:         a.RIR,
		DateAdded:   date,
		Description: a.Description,
	})
}

// UnmarshalJSON unmarshals JSON bytes into an Aggregate, and verifies that
// the contained IP address and date are valid.
func (a *Aggregate) UnmarshalJSON(b []byte) error {
	var raw aggregate
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	_, prefix, err := net.ParseCIDR(raw.Prefix)
	if err != nil {
		return err
	}

	*a = Aggregate{
		ID:          raw.ID,
		Family:      raw.Family,
		Prefix:      prefix,
		RIR:         raw.RIR,
		Description: raw.Description,
	}

	if raw.DateAdded == "" {
		return nil
	}

	t, err := time.Parse(dateFormat, raw.DateAdded)
	if err != nil {
		return err
	}
	a.DateAdded = t

	return nil
}
