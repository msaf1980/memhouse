package database

import (
	"context"

	"github.com/msaf1980/memhouse/protocol/memhouse"
)

type DatabasePool interface {
	StoreMetrics(ctx context.Context, metrics []memhouse.Metric) error
	FetchMetrics(ctx context.Context, req *memhouse.FetchMetricsRequest) (*memhouse.FetchMetricResponce, error)
	Close() error

	StatMetricsStore() uint64     // called from single thread (stat collector)
	StatMetricsStoreErrs() uint64 // called from single thread (stat collector)
	StatMetricsFetch() uint64     // called from single thread (stat collector)
	StatMetricsFetchErrs() uint64 // called from single thread (stat collector)
}
