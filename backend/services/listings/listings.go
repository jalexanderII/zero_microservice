package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/jalexanderII/zero_microservice/backend/services/listings/server"
	"github.com/jalexanderII/zero_microservice/gen/listings"
	"github.com/jalexanderII/zero_microservice/global"
	listingDB "github.com/jalexanderII/zero_microservice/global/db/listings"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 9090, "The server port")
)

func main() {
	log.Println("Listings Service")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server.ApartmentCollection = *listingDB.DB.Collection(global.ApartmentCollection)
	server.BuildingCollection = *listingDB.DB.Collection(global.BuildingCollection)
	server.RealtorCollection = *listingDB.DB.Collection(global.RealtorCollection)

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	listings.RegisterListingsServer(grpcServer, server.NewListingsServer())
	log.Printf("Server started at %v", lis.Addr().String())

	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}
