package server_grps_v3

import (
	"context"
	"io"
	"log"
	"net"

	"github.com/msaf1980/memhouse/protocol/memhouse"
	grpc_v3 "github.com/msaf1980/memhouse/protocol/memhouse_grpc_v3"

	"google.golang.org/grpc"
)

type MemhouseGRPCV3Server struct {
	listener     net.Listener
	server       *grpc.Server
	maxBatchSize int
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
		//_ = metric
		log.Printf("%+v\n", metric)
	}
}
