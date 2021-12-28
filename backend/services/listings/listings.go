package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/hashicorp/go-hclog"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	config "github.com/jalexanderII/zero_microservice"
	listingsDB "github.com/jalexanderII/zero_microservice/backend/services/listings/database"
	"github.com/jalexanderII/zero_microservice/backend/services/listings/server"
	fileServicePB "github.com/jalexanderII/zero_microservice/gen/file_service"
	"github.com/jalexanderII/zero_microservice/gen/listings"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
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

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", config.FIlESERVICESERVERPORT), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fileServiceClient := fileServicePB.NewFileServiceClient(conn)
	db, err := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)

	grpcServer := grpc.NewServer()
	listings.RegisterListingsServer(grpcServer, server.NewListingsServer(listingDB, fileServiceClient, l))

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(grpcServer)

	l.Info("Server started", "port", lis.Addr().String())
	go func() {
		log.Fatal("Serving gRPC: ", grpcServer.Serve(lis).Error())
	}()

	// From https://rogchap.com/2019/07/26/in-process-grpc-web-proxy/
	grpcWebServer := grpcweb.WrapServer(grpcServer)
	httpServer := &http.Server{
		Addr: config.WEBPROXYPORT,
		Handler: h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 {
				grpcWebServer.ServeHTTP(w, r)
			} else {
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
				w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-User-Agent, X-Grpc-Web")
				w.Header().Set("grpc-status", "")
				w.Header().Set("grpc-message", "")
				if grpcWebServer.IsGrpcWebRequest(r) {
					grpcWebServer.ServeHTTP(w, r)
				}
			}
		}), &http2.Server{}),
	}
	log.Fatal("Serving Proxy: ", httpServer.ListenAndServe().Error())
}
