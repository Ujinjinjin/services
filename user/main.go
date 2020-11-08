package main

import (
	"flag"
	"fmt"
	"github.com/ujinjinjin/services/user/factories"
	pb "github.com/ujinjinjin/services/user/interface"
	"github.com/ujinjinjin/services/user/repository"
	"github.com/ujinjinjin/services/user/services"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strconv"
)

var (
	address *string
	port *int
	dbHost *string
	dbUser *string
	dbPassword *string
	dbName *string
)

func main() {
	// Init dependencies
	_dbContextFactory := factories.NewDbContextFactory(dbHost, dbUser, dbPassword, dbName)
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

func init() {
	if len(os.Args) > 1 {
		log.Print("Using command line argument")
		address = flag.String("address", "", "The server address")
		port = flag.Int("port", 10000, "The server port")
		dbHost = flag.String("db-host", "", "Database host")
		dbUser = flag.String("db-user", "", "Database user")
		dbPassword = flag.String("db-password", "", "Database password")
		dbName = flag.String("db-name", "", "Database name")

		flag.Parse()
	} else {
		log.Print("Using environment variables")
		_address := os.Getenv("SERVICE_ADDRESS")
		_port, _ := strconv.Atoi(os.Getenv("SERVICE_PORT"))
		_dbHost := os.Getenv("DB_HOST")
		_dbUser := os.Getenv("DB_USER")
		_dbPassword := os.Getenv("DB_PASSWORD")
		_dbName := os.Getenv("DB_NAME")

		address = &_address
		port = &_port
		dbHost = &_dbHost
		dbUser = &_dbUser
		dbPassword = &_dbPassword
		dbName = &_dbName
	}
}

