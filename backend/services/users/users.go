package main

import (
	"fmt"
	"net"

	"github.com/hashicorp/go-hclog"
	config "github.com/jalexanderII/zero_microservice"
	userDB "github.com/jalexanderII/zero_microservice/backend/services/users/database"
	"github.com/jalexanderII/zero_microservice/backend/services/users/server"
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

	db := userDB.ConnectToDB()
	userCollection := *db.Collection(config.USERCOLLECTIONNAME)

	grpcServer := grpc.NewServer()
	userPB.RegisterAuthServiceServer(grpcServer, server.NewServer(userCollection))

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(grpcServer)

	l.Info("Server started", "port", lis.Addr().String())
	if err := grpcServer.Serve(lis); err != nil {
		l.Error("Error starting server", "error", err)
	}
}
