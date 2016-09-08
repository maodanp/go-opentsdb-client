package opentsdb

import (
	"bytes"
	"fmt"
	"math"
)

func PackQueryString(query *QueryRequestGet) (queryStr string) {
	var buffer bytes.Buffer
	first := true

	buffer.WriteString("&m=")
	// aggregator
	buffer.WriteString(query.Aggregator + ":")
	// rate
	if query.Rate {
		buffer.WriteString("rate")
		if query.Counter {
			buffer.WriteString("{counter,")
			if query.CounterMax != math.MaxInt64 {
				buffer.WriteString(fmt.Sprintf("%d", query.CounterMax))
			}
			buffer.WriteString(",")
			if query.ResetValue != 0 {
				buffer.WriteString(fmt.Sprintf("%d", query.ResetValue))
			}
			buffer.WriteString("}:")
		}
	}
	// down sampler
	if query.DownSampler != "" {
		buffer.WriteString(query.DownSampler + ":")
	}
	//explicit_tags
	if query.ExplicitTags {
		buffer.WriteString("explicit_tags")
	}
	//metric_name
	buffer.WriteString(query.MetricName)
	//grouping filetr
	buffer.WriteString("{")
	first = true
	for tagk, tagv := range query.GroupTagFilters {
		if !first {
			buffer.WriteString(",")
		}
		buffer.WriteString(tagk + "=" + tagv)
		first = false
	}
	buffer.WriteString("}")
	//no group filter
	if len(query.NonGroupTagFilters) > 0 {
		buffer.WriteString("{")
		first = true
		for tagk, tagv := range query.NonGroupTagFilters {
			if !first {
				buffer.WriteString(",")
			}
			buffer.WriteString(tagk + "=" + tagv)
			first = false
		}
		buffer.WriteString("}")
	}

	return buffer.String()
}
