package opentsdb

import (
	"github.com/maodanp/go-log"
	"testing"
)

func TestPut(t *testing.T) {
	c := NewClient("10.10.100.48", 4242, 3)
	metric := &UniMetric{
		MetricName: "test.opentsdb",
		TimeStamp:  1473302029,
		Value:      11111,
		Tags: map[string]interface{}{
			"host": "localhost",
			"port": "4240",
		},
	}
	rr, err := c.Put([]*UniMetric{metric})
	if err != nil {
		log.Logger.Debugf("errInfo: %+v", err)
	} else {
		log.Logger.Debug("put.resp.info", rr.RespInfo)
	}
	c.Close()
}

func TestQueryGet(t *testing.T) {
	c := NewClient("10.10.100.48", 4242, 3)
	query := &QueryRequestGet{
		Start:      "1473302020",
		Aggregator: "sum",
		MetricName: "test.opentsdb",
	}
	queryRsp, errRsp, err := c.QueryByGet(query)
	if err != nil {
		log.Logger.Warnf("err: %+v", err)
	} else if errRsp != nil {
		log.Logger.Warnf("err: %+v", *errRsp)
	} else {
		log.Logger.Debugf("query.get.info: %+v", *queryRsp)
	}
	c.Close()
}

func TestQueryPost(t *testing.T) {
	c := NewClient("10.10.100.48", 4242, 3)
	query := &QueryRequestPost{
		Start: "1473302020",
		Queries: []SubQueryRequest{
			SubQueryRequest{
				Aggregator: "sum",
				Metric:     "test.opentsdb",
			},
		},
	}
	queryRsp, errRsp, err := c.QueryByPost(query)
	if err != nil {
		log.Logger.Warnf("err: %+v", err)
	} else if errRsp != nil {
		log.Logger.Warnf("err: %+v", *errRsp)
	} else {
		log.Logger.Debugf("query.post.info: %+v", *queryRsp)
	}
	c.Close()
}
