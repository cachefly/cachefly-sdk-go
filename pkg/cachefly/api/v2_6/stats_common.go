package v2_6

import (
	"net/url"
	"strconv"
)

// StatsMeta contains metadata about a stats response.
type StatsMeta struct {
	Offset  int      `json:"offset,omitempty"`
	Limit   int      `json:"limit,omitempty"`
	Total   int      `json:"total,omitempty"`
	SortBy  []string `json:"sortBy,omitempty"`
	GroupBy []string `json:"groupBy,omitempty"`
}

// StatsDataPoint represents a single stats row. Fields differ by endpoint,
// so we keep it flexible as a generic map.
type StatsDataPoint map[string]interface{}

// StatsResponse is a generic response shape returned by stats endpoints.
type StatsResponse struct {
	Meta StatsMeta        `json:"meta"`
	Data []StatsDataPoint `json:"data"`
}

// StatsQueryOptions defines common query parameters for stats endpoints.
// Refer to API docs for the precise field availability per endpoint.
type StatsQueryOptions struct {
	Offset      int
	Limit       int
	From        string
	To          string
	GroupBy     []string
	SortBy      []string
	IncludeInfo bool
	Account     []string
	UID         []int
	POP         []string
	Country     []string
	Status      []int
	TldOnly     bool
}

func (o StatsQueryOptions) toURLValues() url.Values {
	v := url.Values{}

	if o.Offset >= 0 {
		v.Set("offset", strconv.Itoa(o.Offset))
	}
	if o.Limit > 0 {
		v.Set("limit", strconv.Itoa(o.Limit))
	}
	if o.From != "" {
		v.Set("from", o.From)
	}
	if o.To != "" {
		v.Set("to", o.To)
	}

	for _, g := range o.GroupBy {
		if g != "" {
			v.Add("groupBy", g)
		}
	}
	for _, s := range o.SortBy {
		if s != "" {
			v.Add("sortBy", s)
		}
	}

	if o.IncludeInfo {
		v.Set("includeInfo", "true")
	}

	for _, acc := range o.Account {
		v.Add("account", acc)
	}

	for _, id := range o.UID {
		v.Add("uid", strconv.Itoa(id))
	}

	for _, pop := range o.POP {
		v.Add("pop", pop)
	}

	for _, country := range o.Country {
		v.Add("country", country)
	}

	for _, st := range o.Status {
		v.Add("status", strconv.Itoa(st))
	}

	if o.TldOnly {
		v.Set("tldOnly", "true")
	}

	return v
}
