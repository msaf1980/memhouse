package main

import (
	"flag"
	"log"

	memsrv "github.com/msaf1980/memhouse/server/grpc_v3"
)

var (
	addr         = flag.String("addr", ":8130", "TCP address to listen to")
	maxBatchSize = flag.Int("batch", 1024, "maximum metrics batch size")
)

func main() {
	flag.Parse()

	server, err := memsrv.NewMemhouseGRPCV3Server(*addr, *maxBatchSize)
	if err != nil {
		log.Fatal(err)
	}
	err = server.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
