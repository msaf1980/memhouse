package client_grps_v3

import (
	"context"
	"fmt"
	"io"
	"net"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	"github.com/msaf1980/memhouse/protocol/memhouse"

	// only for fake memhouse GRPC v3 server
	avl "github.com/emirpasic/gods/trees/avltree"
	grpc_v3 "github.com/msaf1980/memhouse/protocol/memhouse_grpc_v3"
)

func TestMemhouseGRPCV3Client(t *testing.T) {
	server, err := NewMemhouseGRPCV3Server(":0", 1024) // create fake memhouse GRPC v3 server
	require.NoError(t, err)

	go server.Serve()

	time.Sleep(10 * time.Millisecond) // wait for start server in background

	address := server.Address().String()
	ctx := context.Background()
	client, close, err := NewMemhouseGRPCV3Client(ctx, address, 100*time.Millisecond)
	require.NoError(t, err)
	if close != nil {
		defer close()
	}

	// Test for store metrics
	stream, err := client.StoreMetrics(context.Background())
	if err != nil {
		panic(err)
	}

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

	for i := range metrics {
		err = stream.Send(&metrics[i])
		assert.NoError(t, err)
	}

	reply, err := stream.CloseAndRecv()
	assert.NoError(t, err)
	assert.Equal(t, grpc_v3.StoreMetricsResponce{Code: 200, Status: "OK"}, *reply)

	// Verify store and test fetch

	// Fetch All
	request := memhouse.FetchMetricsRequest{
		From:  0,
		Until: 30,
		Path: []string{
			"test.metric.1",
			"test.metric.20",
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

	resp, err := client.FetchMetrics(context.Background(), &request)
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

	// Fetch Range
	request = memhouse.FetchMetricsRequest{
		From:  11,
		Until: 30,
		Path: []string{
			"test.metric.1",
			"test.metric.20",
			"test.metric",
		},
	}
	respExpect = memhouse.FetchMetricResponce{
		Metrics: []*memhouse.MetricValues{
			{
				Path: "test.metric.1",
				Values: []*memhouse.MetricValue{
					{Timestamp: 30, Value: 22.0},
				},
			},
			{
				Path: "test.metric.20", // no metrics for from/until timestamp range, but metric exist
			},
		},
	}

	resp, err = client.FetchMetrics(context.Background(), &request)
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

// Fake memhouse GRPC v3 server, only for simplify client test

//type values map[uint64]float64 // key is timestamp

type MemhouseGRPCV3Server struct {
	listener     net.Listener
	server       *grpc.Server
	maxBatchSize int
	lock         sync.RWMutex
	data         *avl.Tree // key is a metric name, store values map
}

func NewMemhouseGRPCV3Server(address string, maxBatchSize int) (*MemhouseGRPCV3Server, error) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	srv := &MemhouseGRPCV3Server{
		listener:     listener,
		server:       grpc.NewServer(),
		maxBatchSize: maxBatchSize,
		data:         avl.NewWithStringComparator(),
	}

	grpc_v3.RegisterMemhouseServer(srv.server, srv)

	return srv, nil
}

func (srv *MemhouseGRPCV3Server) Address() net.Addr {
	return srv.listener.Addr()
}

func (srv *MemhouseGRPCV3Server) Serve() error {
	return srv.server.Serve(srv.listener)
}

func (s *MemhouseGRPCV3Server) FetchMetrics(ctx context.Context, req *memhouse.FetchMetricsRequest) (*memhouse.FetchMetricResponce, error) {
	resp := &memhouse.FetchMetricResponce{}

	s.lock.Lock()

	for i := range req.Path {
		if valNode, ok := s.data.Get(req.Path[i]); ok {
			valTree := valNode.(*avl.Tree)
			metrics := &memhouse.MetricValues{Path: req.Path[i]}

			if node, ok := valTree.Ceiling(int(req.From)); ok {
				metrics.Values = append(metrics.Values, &memhouse.MetricValue{
					Timestamp: int64(node.Key.(int)), Value: node.Value.(float64),
				})
				for {
					node = node.Next()
					if node == nil {
						break
					}
					ts := node.Key.(int)
					if ts > int(req.Until) {
						break
					}
					metrics.Values = append(metrics.Values, &memhouse.MetricValue{
						Timestamp: int64(ts), Value: node.Value.(float64),
					})
				}
			}

			// 	Values: []*pb_v3.MetricValue{
			// 		{Timestamp: 0, Value: 10.0}, {Timestamp: 10, Value: 21.0}, {Timestamp: 30, Value: 22.0},
			// 	},
			// },

			resp.Metrics = append(resp.Metrics, metrics)
		}
	}

	s.lock.Unlock()

	return resp, nil
}

func (s *MemhouseGRPCV3Server) StoreMetrics(stream grpc_v3.Memhouse_StoreMetricsServer) error {
	for {
		metric, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return stream.SendAndClose(&grpc_v3.StoreMetricsResponce{
					Code:   200,
					Status: "OK",
				})
			}
			return err
		}

		s.lock.Lock()

		if node, ok := s.data.Get(metric.Path); ok {
			metrics := node.(*avl.Tree)
			metrics.Put(int(metric.Timestamp), metric.Value) // last win, no duplication check
		} else {
			metrics := avl.NewWithIntComparator()
			metrics.Put(int(metric.Timestamp), metric.Value)
			s.data.Put(metric.Path, metrics)
		}

		s.lock.Unlock()
	}
}
