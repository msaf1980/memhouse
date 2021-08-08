package redis

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/msaf1980/memhouse/protocol/memhouse"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRedisConnector(t *testing.T) {
	ctx := context.Background()

	s, err := miniredis.Run()
	require.NoError(t, err)
	defer s.Close()

	var options redis.Options
	connector := NewRedisConnector(s.Addr(), "", 0, options)

	metrics := []memhouse.Metric{
		{
			Path:      "test.metric.1",
			Timestamp: 0,
			Value:     10.0,
		},
		{
			Path:      "test.metric.20",
			Timestamp: 0,
			Value:     25.1,
		},
		{
			Path:      "test.metric.1",
			Timestamp: 10,
			Value:     21.0,
		},
		{
			Path:      "test.metric.20",
			Timestamp: 10,
			Value:     45.1,
		},
		{
			Path:      "test.metric.1",
			Timestamp: 30,
			Value:     22.0,
		},
	}

	err = connector.StoreMetrics(ctx, metrics)
	assert.NoError(t, err)

	assert.Equal(t, len(metrics), int(connector.StatMetricsStore()))
	assert.Equal(t, 0, int(connector.StatMetricsStore())) // stat for store reset in prev call
	assert.Equal(t, 0, int(connector.StatMetricsStoreErrs()))

	// fetch all
	req := memhouse.FetchMetricsRequest{
		From:  0,
		Until: 30,
		Path: []string{
			"test.metric.1",
			"test.metric.20",
			"test.metric.3",
		},
	}
	respExpect := memhouse.FetchMetricResponce{
		Metrics: []*memhouse.MetricValues{
			{
				Path: "test.metric.1",
				Values: []*memhouse.MetricValue{
					{Timestamp: 0, Value: 10.0}, {Timestamp: 10, Value: 21.0}, {Timestamp: 30, Value: 22.0},
				},
			},
			{
				Path: "test.metric.20",
				Values: []*memhouse.MetricValue{
					{Timestamp: 0, Value: 25.1}, {Timestamp: 10, Value: 45.1},
				},
			},
		},
	}

	resp, err := connector.FetchMetrics(ctx, &req)
	assert.NoError(t, err)
	for i := 0; i < max(len(resp.Metrics), len(respExpect.Metrics)); i++ {
		if i >= len(resp.Metrics) {
			t.Errorf("- %v\n", dumpMetricValues(respExpect.Metrics[i]))
		} else if i >= len(respExpect.Metrics) {
			t.Errorf("+ %s\n", dumpMetricValues(resp.Metrics[i]))
		} else if !reflect.DeepEqual(respExpect.Metrics[i], resp.Metrics[i]) {
			t.Errorf("- %v\n", dumpMetricValues(respExpect.Metrics[i]))
			t.Errorf("+ %s\n", dumpMetricValues(resp.Metrics[i]))
		}
	}

	// fetch range
	req = memhouse.FetchMetricsRequest{
		From:  10,
		Until: 29,
		Path: []string{
			"test.metric.1",
			"test.metric.20",
			"test.metric.3",
		},
	}
	respExpect = memhouse.FetchMetricResponce{
		Metrics: []*memhouse.MetricValues{
			{
				Path: "test.metric.1",
				Values: []*memhouse.MetricValue{
					{Timestamp: 10, Value: 21.0},
				},
			},
			{
				Path: "test.metric.20",
				Values: []*memhouse.MetricValue{
					{Timestamp: 10, Value: 45.1},
				},
			},
		},
	}

	resp, err = connector.FetchMetrics(ctx, &req)
	assert.NoError(t, err)
	for i := 0; i < max(len(resp.Metrics), len(respExpect.Metrics)); i++ {
		if i >= len(resp.Metrics) {
			t.Errorf("- %v\n", dumpMetricValues(respExpect.Metrics[i]))
		} else if i >= len(respExpect.Metrics) {
			t.Errorf("+ %s\n", dumpMetricValues(resp.Metrics[i]))
		} else if !reflect.DeepEqual(respExpect.Metrics[i], resp.Metrics[i]) {
			t.Errorf("- %v\n", dumpMetricValues(respExpect.Metrics[i]))
			t.Errorf("+ %s\n", dumpMetricValues(resp.Metrics[i]))
		}
	}
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func dumpMetricValues(values *memhouse.MetricValues) string {
	var sb strings.Builder
	sb.WriteString("Path: \"")
	sb.WriteString(values.Path)
	sb.WriteString("\", Values: [ ")
	for i := range values.Values {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("{ %d, %f }", values.Values[i].Timestamp, values.Values[i].Value))
	}
	sb.WriteString(" ]")

	return sb.String()
}

// Mock for RedisConnector

// type redisMockConnector struct {
// 	connector *RedisConnector
// 	mock      redismock.ClientMock
// }

// func newRedisMockConnector() *redisMockConnector {
// 	var client *redis.Client
// 	c := &redisMockConnector{}

// 	client, c.mock = redismock.NewClientMock()

// 	c.mock.ExpectZAdd("key", "field").SetVal("test value")

// 	c.connector = NewCustomRedisConnector(client)

// 	return c
// }

// func (c *redisMockConnector) StoreMetrics(ctx context.Context, metrics []memhouse.Metric) error {
// 	return c.connector.StoreMetrics(ctx, metrics)
// }

// func (c *redisMockConnector) FetchMetrics(ctx context.Context, req *memhouse.FetchMetricsRequest) (*memhouse.FetchMetricResponce, error) {
// 	return c.connector.FetchMetrics(ctx, req)
// }

// func (c *redisMockConnector) Close() error {
// 	return c.connector.Close()
// }
