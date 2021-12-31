package main

import (
	"fmt"
	"log"
	"net"

	"github.com/hashicorp/go-hclog"
	applicationDB "github.com/jalexanderII/zero_microservice/backend/services/application/database"
	"github.com/jalexanderII/zero_microservice/backend/services/application/server"
	userDB "github.com/jalexanderII/zero_microservice/backend/services/users/database"
	"github.com/jalexanderII/zero_microservice/backend/services/users/middleware"
	authServer "github.com/jalexanderII/zero_microservice/backend/services/users/server"
	"github.com/jalexanderII/zero_microservice/config"
	applicationPB "github.com/jalexanderII/zero_microservice/gen/application"
	fileServicePB "github.com/jalexanderII/zero_microservice/gen/file_service"
	userPB "github.com/jalexanderII/zero_microservice/gen/users"
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

	jwtManager := middleware.NewJWTManager(config.JWTSecret, config.TokenDuration)
	interceptor := middleware.NewAuthInterceptor(jwtManager, config.AccessibleRoles(), l)

	db, err := applicationDB.ConnectToDB()
	appDB := applicationDB.NewApplicationDB(db)
	fileServiceClient := fileServicePB.NewFileServiceClient(conn)

	userdb := userDB.InitiateMongoClient()
	userCollection := *userdb.Collection(config.USERCOLLECTIONNAME)

	serverOptions := []grpc.ServerOption{grpc.UnaryInterceptor(interceptor.Unary())}
	grpcServer := grpc.NewServer(serverOptions...)

	userPB.RegisterAuthServiceServer(grpcServer, authServer.NewAuthServer(userCollection, jwtManager, l))
	applicationPB.RegisterApplicationServer(grpcServer, server.NewApplicationServer(appDB, fileServiceClient, l))
	methods := config.ListGRPCResources(grpcServer)
	l.Info("Methods on this server", "methods", methods)

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(grpcServer)

	l.Info("Server started", "port", lis.Addr().String())

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("Serving gRPC: ", err)
	}

}
