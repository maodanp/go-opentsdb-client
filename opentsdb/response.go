package opentsdb

import (
	"encoding/json"
	"net/http"
)

var (
	validHttpStatusCode = map[int]bool{
		http.StatusCreated:            true,
		http.StatusOK:                 true,
		http.StatusNoContent:          true,
		http.StatusBadRequest:         true,
		http.StatusNotFound:           true,
		http.StatusPreconditionFailed: true,
		http.StatusForbidden:          true,
		http.StatusUnauthorized:       true,
	}
)

type RawResponse struct {
	StatusCode int
	Body       []byte
	Header     http.Header
}

// DecodeQueryResp decode query response
// if StatusCode is equal StatusOK, we can get error details from queryResp
// else we can get error details from errResp
func (rr *RawResponse) DecodeQueryResp() (queryResp *QueryResponse, errResp *ErrorResponse, err error) {
	if rr.StatusCode == http.StatusOK {
		queryResp = &QueryResponse{}
		err = json.Unmarshal(rr.Body, &queryResp)
		if err != nil {
			return queryResp, errResp, err
		}
	} else {
		errResp = &ErrorResponse{}
		err = json.Unmarshal(rr.Body, &errResp)
		if err != nil {
			return queryResp, errResp, err
		}
	}
	return
}
