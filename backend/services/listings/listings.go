package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	listingsDB "github.com/jalexanderII/zero_microservice/backend/services/listings/database"
	"github.com/jalexanderII/zero_microservice/backend/services/listings/server"
	contentStore "github.com/jalexanderII/zero_microservice/backend/services/listings/store"
	"github.com/jalexanderII/zero_microservice/gen/listings"
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
	store := contentStore.NewDiskImageStore("./store/tmp")
	db, err := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	listings.RegisterListingsServer(grpcServer, server.NewListingsServer(listingDB, store))
	log.Printf("Server started at %v", lis.Addr().String())

	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}
