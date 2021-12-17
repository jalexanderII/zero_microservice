package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	listingsDB "github.com/jalexanderII/zero_microservice/backend/services/listings/database"
	contentStore "github.com/jalexanderII/zero_microservice/backend/services/listings/store"
	listingsPB "github.com/jalexanderII/zero_microservice/gen/listings"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type listingsServer struct {
	listingsPB.UnimplementedListingsServer
	DB           *listingsDB.ListingsDB
	ContentStore contentStore.ContentStore
	l            hclog.Logger
}

func NewListingsServer(db *listingsDB.ListingsDB, cs contentStore.ContentStore, l hclog.Logger) *listingsServer {
	return &listingsServer{DB: db, ContentStore: cs, l: l}
}

func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return status.Error(codes.Canceled, "request is canceled")
	case context.DeadlineExceeded:
		return status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	default:
		return nil
	}
}
