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
	fileServicePB "github.com/jalexanderII/zero_microservice/gen/file_service"
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

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", config.FIlESERVICESERVERPORT), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	db, err := applicationDB.ConnectToDB()
	appDB := applicationDB.NewApplicationDB(db)
	fileServiceClient := fileServicePB.NewFileServiceClient(conn)

	grpcServer := grpc.NewServer()
	applicationPB.RegisterApplicationServer(grpcServer, server.NewApplicationServer(appDB, fileServiceClient, l))

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(grpcServer)

	l.Info("Server started", "port", lis.Addr().String())

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("Serving gRPC: ", err)
	}

}
