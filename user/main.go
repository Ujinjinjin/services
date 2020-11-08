package main

import (
	"flag"
	"fmt"
	"github.com/ujinjinjin/services/user/factories"
	"github.com/ujinjinjin/services/user/repository"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/ujinjinjin/services/user/interface"
	"github.com/ujinjinjin/services/user/services"
)

var (
	address = flag.String("address", "", "The server address")
	port = flag.Int("port", 10000, "The server port")
)

func main() {
	flag.Parse()

	// Init dependencies
	_dbContextFactory := factories.NewDbContextFactory()
	_repository := repository.NewUserRepository(_dbContextFactory)

	log.Printf("Starting server at %s:%d", *address, *port)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *address, *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, services.InitService(_repository))

	var serveResult = grpcServer.Serve(lis)
	if serveResult != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

