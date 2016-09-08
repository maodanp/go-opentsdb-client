package opentsdb

//put response
type PutResponse struct {
	StatusCode int
	RespInfo   string
}

type QueryFilterRequest struct {
	Type   string `json:"type"`
	Tagk   string `json:"tagk"`
	Filter string `json:"filter"`
}

type SubQueryRequest struct {
	Aggregator string               `json:"aggregator"`
	Metric     string               `json:"metric"`
	Rate       string               `json:"rate,omitempty"`
	Counter    bool                 `json:"counter,omitempty"`
	Downsample string               `json:"downsample,omitempty"`
	Tags       map[string]string    `json:"tags,omitempty"`
	Filters    []QueryFilterRequest `json:"filters,omitempty"`
}

// QueryRequestPost query request by using post method
type QueryRequestPost struct {
	Start             string            `json:"start"`
	End               string            `json:"end,omitempty"`
	Queries           []SubQueryRequest `json:"queries"`
	NoAnnotations     string            `json:"noAnnotations,omitempty"`
	GlobalAnnotations string            `json:"globalAnnotations,omitempty"`
	MsResolution      string            `json:"msResolution,omitempty"`
	ShowTSUIDs        string            `json:"showTSUIDs,omitempty"`
}

// QueryRequestGet query request by using get method
type QueryRequestGet struct {
	// Required
	Start      string
	Aggregator string
	MetricName string
	// Optional
	Rate               bool
	Counter            bool
	CounterMax         int64
	ResetValue         int64
	DownSampler        string
	ExplicitTags       bool
	GroupTagFilters    map[string]string
	NonGroupTagFilters map[string]string
}

//query response
type MetricResponse struct {
	Metric         string                 `json:"metric"`
	Tags           map[string]interface{} `json:"tags"`
	AggregatedTags []string               `json:"aggregateTags"`
	Dps            map[string]float64     `json:"dps"`
	Tsuids         []string               `json:"tsuids,omitempty"`
}

type QueryResponse []MetricResponse

type ErrorResponse struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Details string `json:"details"`
		Trace   string `json:"trace"`
	} `json:"error"`
}
