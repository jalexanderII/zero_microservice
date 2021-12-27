package main

import (
	"fmt"
	"log"
	"net"

	"github.com/hashicorp/go-hclog"
	config "github.com/jalexanderII/zero_microservice"
	fileServiceDB "github.com/jalexanderII/zero_microservice/backend/services/file_service/database"
	"github.com/jalexanderII/zero_microservice/backend/services/file_service/server"
	fileServicePB "github.com/jalexanderII/zero_microservice/gen/file_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	l := hclog.Default()
	l.Debug("FileService Service")

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", config.FIlESERVICESERVERPORT))
	if err != nil {
		l.Error("failed to listen", "error", err)
		panic(err)
	}

	db, err := fileServiceDB.InitiateMongoClient()
	if err != nil {
		l.Error("failed to connect to DB", "error", err)
		panic(err)
	}
	grpcServer := grpc.NewServer()
	fileServicePB.RegisterFileServiceServer(grpcServer, server.NewFileServiceServer(db, l))

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(grpcServer)

	l.Info("Server started", "port", lis.Addr().String())

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("Serving gRPC: ", err)
	}
}
