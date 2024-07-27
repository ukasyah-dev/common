package testkit

import (
	"net"

	"github.com/emitra-labs/common/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClientConn(server *grpc.Server) (*grpc.ClientConn, func()) {
	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		log.Errorf("Failed to listen localhost:0: %s", err)
	}

	go func() {
		if err := server.Serve(lis); err != nil {
			log.Errorf("Error serving test server: %s", err)
		}
	}()

	conn, err := grpc.NewClient(
		lis.Addr().String(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Errorf("Failed to connect to server: %s", err)
	}

	closer := func() {
		err := lis.Close()
		if err != nil {
			log.Errorf("Failed to close listener: %s", err)
		}
		server.Stop()
	}

	return conn, closer
}
