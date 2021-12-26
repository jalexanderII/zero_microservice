package main

import (
	"fmt"
	"log"
	"net"

	"github.com/hashicorp/go-hclog"
	config "github.com/jalexanderII/zero_microservice"
	applicationDB "github.com/jalexanderII/zero_microservice/backend/services/application/database"
	"github.com/jalexanderII/zero_microservice/backend/services/application/server"
	applicationPB "github.com/jalexanderII/zero_microservice/gen/application"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	l := hclog.Default()
	l.Debug("Application Service")

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", config.SERVERPORT))
	if err != nil {
		l.Error("failed to listen", "error", err)
		panic(err)
	}

	db, err := applicationDB.ConnectToDB()
	appDB := applicationDB.NewApplicationDB(db)

	grpcServer := grpc.NewServer()
	applicationPB.RegisterApplicationServer(grpcServer, server.NewApplicationServer(appDB, l))

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(grpcServer)

	l.Info("Server started", "port", lis.Addr().String())

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("Serving gRPC: ", err)
	}

}
