# go-opentsdb-client

## What is opentsdb?
OpenTSDB is a distributed, scalable Time Series Database (TSDB). OpenTSDB allows you to collect thousands of metrics from tens of thousands of hosts and applications, at a high rate (every few seconds). OpenTSDB will never delete or downsample data and can easily store hundreds of billions of data points.

I develop the `go-opentsdb-client` to simplify the process in your golang project. Before you use these apis, You'd better read [OpenTSDB Rest API Doc](http://opentsdb.net/docs/build/html/index.html).

## Example Usage
~~~go
func TestPut(t *testing.T) {
        // create a new client.
        // host: 127.0.0.1
        // port: 4242
        // connectionTimeOut: 3s
        c := NewClient("127.0.0.1", 4242, 3)

        // construct UniMetric data
        metric := &UniMetric{
                MetricName: "test.opentsdb",
                TimeStamp:  1473302029,
                Value:      11111,
                Tags: map[string]interface{}{
                        "host": "localhost",
                        "port": "4240",
                },
        }

        // put metrics to opentsdb
        rr, err := c.Put([]*UniMetric{metric})
        if err != nil {
                log.Logger.Debugf("errInfo: %+v", err)
        }else{
                log.Logger.Debug("put.resp.info", rr.RespInfo)
        }
}
~~~

## Usually used APIs

~~~go
// NewClient create a Client obj with host、port、connTimeout
NewClient(host string, port int, dialTimeout int64) *Client

// Put put metrics to oentsdb
func (c *Client) Put(metrics []*UniMetric) (*PutResponse, error)

// QueryByGet query with GET method
func (c *Client) QueryByGet(query *QueryRequestGet) (queryResp *QueryResponse, errResp *ErrorResponse, err error)

// QueryByPost query with GET method
func (c *Client) QueryByPost(query *QueryRequestPost) (queryResp *QueryResponse, errResp *ErrorResponse, err error)
~~~

## Development Plan
* Add connection pool
* Add RESTful API
* Add connTimeOut and rwTimeOut
* Read metrics by different way (file, io etc.)
