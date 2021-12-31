package main

import (
	"fmt"
	"net"

	"github.com/hashicorp/go-hclog"
	userDB "github.com/jalexanderII/zero_microservice/backend/services/users/database"
	"github.com/jalexanderII/zero_microservice/backend/services/users/server"
	"github.com/jalexanderII/zero_microservice/config"
	"github.com/jalexanderII/zero_microservice/config/middleware"
	listingsPB "github.com/jalexanderII/zero_microservice/gen/listings"
	userPB "github.com/jalexanderII/zero_microservice/gen/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	l := hclog.Default()
	l.Debug("Users Service")

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", config.USERPORT))
	if err != nil {
		l.Error("failed to listen", "error", err)
		panic(err)
	}

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", config.LISTINGSERVICESERVERPORT), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	listingServiceClient := listingsPB.NewListingsClient(conn)

	DB := userDB.InitiateMongoClient()
	userCollection := *DB.Collection(config.USERCOLLECTIONNAME)
	jwtManager := middleware.NewJWTManager(config.JWTSecret, config.TokenDuration)

	grpcServer := grpc.NewServer()
	userPB.RegisterAuthServiceServer(grpcServer, server.NewAuthServer(userCollection, jwtManager, listingServiceClient, l))
	methods := config.ListGRPCResources(grpcServer)
	l.Info("Methods on this server", "methods", methods)

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(grpcServer)

	l.Info("Server started", "port", lis.Addr().String())
	if err := grpcServer.Serve(lis); err != nil {
		l.Error("Error starting server", "error", err)
	}
}
