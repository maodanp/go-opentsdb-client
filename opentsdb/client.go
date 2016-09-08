package opentsdb

import (
	"crypto/tls"
	"fmt"
	_ "io"
	"net"
	"net/http"
	_ "net/url"
	_ "strconv"
	"strings"
	"time"
)

type UniMetric struct {
	MetricName string                 `json:"metric"`
	TimeStamp  int64                  `json:"timestamp"`
	Value      float64                `json:"value"`
	Tags       map[string]interface{} `json:"tags"`
}

type Client struct {
	metric UniMetric

	httpClient  *http.Client
	transport   *http.Transport
	dialTimeout time.Duration
	putUrl      string
	queryMethod string
	queryUrl    string

	batchPutLen int
	putChan     *UniMetric
}

func NewClient(host string, port int, dialTimeout int64) *Client {

	var putUrl, queryUrl string
	if strings.Contains(host, "://") {
		putUrl = fmt.Sprintf("%s:%d/api/put", host, port)
		queryUrl = fmt.Sprintf("%s:%d/api/query", host, port)
	} else {
		putUrl = fmt.Sprintf("http://%s:%d/api/put", host, port)
		queryUrl = fmt.Sprintf("http://%s:%d/api/query", host, port)
	}
	client := &Client{
		putUrl:   putUrl,
		queryUrl: queryUrl,
	}

	client.dialTimeout = time.Duration(dialTimeout)
	client.initHTTPClient()
	return client
}

// DefaultDial attempts to open a TCP connection to the provided address, explicitly
// enabling keep-alives with a one-second interval.
func (c *Client) DefaultDial(network, addr string) (net.Conn, error) {
	dialer := net.Dialer{
		Timeout:   c.dialTimeout * time.Second,
		KeepAlive: time.Second,
	}

	return dialer.Dial(network, addr)
}

func (c *Client) initHTTPClient() {
	c.transport = &http.Transport{
		Dial: c.DefaultDial,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	c.httpClient = &http.Client{Transport: c.transport}
}

// SetMaxBatchPutLen can set the length of buffer
// when buffer len reaches the max, it will put to opentsdb
func (c *Client) SetMaxBatchPutLen(len int) {
	c.batchPutLen = len
}

// SetQueryMethod set the query method.
// the method you can choose PUT or POST
func (c *Client) SetQueryMethod(Method string) {
	//c.batchPutLen = len
}

// Put put metrics which is given with []*UniMetric type
func (c *Client) Put(metrics []*UniMetric) (*PutResponse, error) {
	return c.putMetrics(metrics)
}

/*
// PutFromFile put metrics from a given path.
// The given file is excepted to use the JSON format
func (c *Client) PutFromFile(fpath string) (putResp *PutResponse, err error) {
	return
}

// PutFromReader put metrics from a given reader.
// Bytes from reader is excepted to use the JSON format
func (c *Client) PutFromReader(reader io.Reader) (putResp *PutResponse, err error) {
	return
}
*/

//Close close the client
func (c *Client) Close() {
	c.transport.DisableKeepAlives = true
	c.transport.CloseIdleConnections()
}

// QueryByGet query with GET method
// if err is not nil, it means there was error to call this function
// in this case queryResp and errResp are nil
// else if err is nil, you need judge whther errResp is nil
// if errResp is not nil, it means opentsdb query is failed
// if errResp is nil, we success to query from opentsdb
func (c *Client) QueryByGet(query *QueryRequestGet) (queryResp *QueryResponse, errResp *ErrorResponse, err error) {
	return c.queryGet(query)
}

// QueryByPost query with GET method
// see QueryByGet() API
func (c *Client) QueryByPost(query *QueryRequestPost) (queryResp *QueryResponse, errResp *ErrorResponse, err error) {
	return c.queryPost(query)
}

/*
// QueryFromFile query response info from a given path
func (c *Client) QueryFromFile(fpath string) (queryResp *QueryResponse, errResp *ErrorResponse, err error) {
	return
}

// QueryFromReader query response info from a given reader
func (c *Client) QueryFromReader(reader io.Reader) (queryResp *QueryResponse, errResp *ErrorResponse, err error) {
	return
}
*/
