package main

import (
	"fmt"
	"net"

	"github.com/hashicorp/go-hclog"
	config "github.com/jalexanderII/zero_microservice"
	listingsDB "github.com/jalexanderII/zero_microservice/backend/services/listings/database"
	"github.com/jalexanderII/zero_microservice/backend/services/listings/server"
	contentStore "github.com/jalexanderII/zero_microservice/backend/services/listings/store"
	"github.com/jalexanderII/zero_microservice/gen/listings"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	l := hclog.Default()
	l.Debug("Listings Service")

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", config.SERVERPORT))
	if err != nil {
		l.Error("failed to listen", "error", err)
		panic(err)
	}

	store := contentStore.NewDiskImageStore("./store/tmp", l)
	db, err := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)

	grpcServer := grpc.NewServer()
	listings.RegisterListingsServer(grpcServer, server.NewListingsServer(listingDB, store, l))

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(grpcServer)

	l.Info("Server started", "port", lis.Addr().String())
	if err := grpcServer.Serve(lis); err != nil {
		l.Error("Error starting server", "error", err)
	}
}
