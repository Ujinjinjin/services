package main

import (
	"flag"
	"fmt"
	"github.com/ujinjinjin/user_service/factories"
	"github.com/ujinjinjin/user_service/repository"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/ujinjinjin/user_service/interface"
	"github.com/ujinjinjin/user_service/services"
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

