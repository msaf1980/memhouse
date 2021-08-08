package redis

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/msaf1980/memhouse/protocol/memhouse"

	"github.com/go-redis/redis/v8"
)

type RedisConnector struct {
	client *redis.Client

	metrics_store      int64 // Counter for stored metrics
	metrics_store_prev int64

	metrics_store_errs      int64 // Errors for stored metrics
	metrics_store_errs_prev int64

	metrics_fetch      int64 // Counter for fetched metrics
	metrics_fetch_prev int64

	metrics_fetch_errs      int64 // Errors for fetched metrics
	metrics_fetch_errs_prev int64
}

func NewRedisConnector(addr, password string, db int, options redis.Options) *RedisConnector {
	connector := &RedisConnector{}

	options.Addr = addr
	options.Password = password
	options.DB = db

	connector.client = redis.NewClient(&options)

	return connector
}

func NewCustomRedisConnector(client *redis.Client) *RedisConnector {
	connector := &RedisConnector{}
	connector.client = client

	return connector
}

func (c *RedisConnector) StoreMetrics(ctx context.Context, metrics []memhouse.Metric) error {
	var err error

	var adds int64

	for i := range metrics {
		var add int64
		v := fmt.Sprintf("%f %d", metrics[i].Value, metrics[i].Timestamp)
		value := redis.Z{
			Score:  float64(metrics[i].Timestamp),
			Member: v,
		}
		add, err = c.client.ZAdd(ctx, metrics[i].Path, &value).Result()
		if err != nil {
			break
		}
		adds += int64(add)
	}

	atomic.AddInt64(&c.metrics_store, adds)
	if err != nil {
		atomic.AddInt64(&c.metrics_store_errs, 1)
	}
	return err
}

func (c *RedisConnector) FetchMetrics(ctx context.Context, req *memhouse.FetchMetricsRequest) (*memhouse.FetchMetricResponce, error) {
	var fetched int64
	var lastErr error
	var errs int64

	zRange := redis.ZRangeBy{
		Min: strconv.FormatInt(req.From, 10),
		Max: strconv.FormatInt(req.Until, 10),
	}
	resp := &memhouse.FetchMetricResponce{}
	for i := range req.Path {
		vals, err := c.client.ZRangeByScoreWithScores(ctx, req.Path[i], &zRange).Result()
		if err != nil {
			lastErr = err
			break
		}

		metric := &memhouse.MetricValues{Path: req.Path[i]}
		metric.Values = make([]*memhouse.MetricValue, len(vals))
		for j := range vals {
			var v float64
			s := vals[j].Member.(string)
			if indx := strings.IndexByte(s, ' '); indx > 0 {
				v, err = strconv.ParseFloat(s[:indx], 64)
			} else {
				err = fmt.Errorf("parse member value error")
			}

			if err != nil {
				lastErr = err
				errs++
				v = math.NaN()
			}
			metric.Values[j] = &memhouse.MetricValue{
				Timestamp: int64(vals[j].Score),
				Value:     v,
			}
		}

		fetched += int64(len(metric.Values))

		if len(vals) > 0 {
			resp.Metrics = append(resp.Metrics, metric)
		}
	}

	atomic.AddInt64(&c.metrics_fetch, fetched)
	if errs > 0 {
		atomic.AddInt64(&c.metrics_fetch_errs, errs)
	}

	return resp, lastErr
}

func (c *RedisConnector) Close() error {
	return c.client.Close()
}

func (c *RedisConnector) StatMetricsStore() int64 {
	n := atomic.SwapInt64(&c.metrics_store, 0) - c.metrics_store_prev
	c.metrics_store_prev = 0
	return n
}

func (c *RedisConnector) StatMetricsStoreErrs() int64 {
	n := atomic.SwapInt64(&c.metrics_store_errs, 0) - c.metrics_store_errs_prev
	c.metrics_store_errs_prev = 0
	return n
}

func (c *RedisConnector) StatMetricsFetched() int64 {
	n := atomic.SwapInt64(&c.metrics_fetch, 0) - c.metrics_fetch_prev
	c.metrics_fetch_prev = 0
	return n
}

func (c *RedisConnector) StatMetricsFetchedErrs() int64 {
	n := atomic.SwapInt64(&c.metrics_fetch_errs, 0) - c.metrics_fetch_errs_prev
	c.metrics_fetch_errs_prev = 0
	return n
}
