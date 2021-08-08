package client_grps_v3

import (
	"context"
	"time"

	grpc_v3 "github.com/msaf1980/memhouse/protocol/memhouse_grpc_v3"

	"google.golang.org/grpc"
)

func NewMemhouseGRPCV3Client(ctx context.Context, addr string, conTimeout time.Duration) (grpc_v3.MemhouseClient, func(), error) {
	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	return grpc_v3.NewMemhouseClient(conn), func() { conn.Close() }, nil
}
