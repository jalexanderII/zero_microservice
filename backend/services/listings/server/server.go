package server

import (
	"github.com/hashicorp/go-hclog"
	listingsDB "github.com/jalexanderII/zero_microservice/backend/services/listings/database"
	fileServicePB "github.com/jalexanderII/zero_microservice/gen/file_service"
	listingsPB "github.com/jalexanderII/zero_microservice/gen/listings"
)

type listingsServer struct {
	listingsPB.UnimplementedListingsServer
	DB                *listingsDB.ListingsDB
	FileServiceClient fileServicePB.FileServiceClient
	l                 hclog.Logger
}

func NewListingsServer(db *listingsDB.ListingsDB, f fileServicePB.FileServiceClient, l hclog.Logger) *listingsServer {
	return &listingsServer{DB: db, FileServiceClient: f, l: l}
}
