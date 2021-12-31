package server

import (
	"testing"
	"time"

	listingsDB "github.com/jalexanderII/zero_microservice/backend/services/listings/database"
	"github.com/jalexanderII/zero_microservice/config"
	listingsPB "github.com/jalexanderII/zero_microservice/gen/listings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_listingsServer_CreateOwner(t *testing.T) {
	ctx, cancel := config.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, FileServiceClient: MockFileServiceClient(), l: L}

	in := &listingsPB.Owner{
		Name:        "Fitore Abazaga",
		Email:       "f.abazaga@platinum.com",
		PhoneNumber: "(646) 339-3247",
		Company:     "Platinum Properties",
		UserRef:     primitive.NewObjectID().Hex(),
	}

	owner, err := server.CreateOwner(ctx, &listingsPB.CreateOwnerRequest{Owner: in})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if owner.Name != in.Name {
		t.Errorf("1: An error creating a owner: %+v", owner)
	}
}

func Test_listingsServer_GetOwner(t *testing.T) {
	ctx, cancel := config.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, FileServiceClient: MockFileServiceClient(), l: L}

	owner, err := server.GetOwner(ctx, &listingsPB.GetOwnerRequest{Id: 1})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if owner.Name != "Fitore Abazaga" {
		t.Errorf("2: Failed to fetch correct owner: %+v", owner)
	}
}

func Test_listingsServer_ListOwners(t *testing.T) {
	ctx, cancel := config.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, FileServiceClient: MockFileServiceClient(), l: L}

	owners, err := server.ListOwners(ctx, &listingsPB.ListOwnerRequest{})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if len(owners.Owners) < 1 {
		t.Errorf("2: Failed to fetch owners: %+v", owners.Owners[0])
	}
}

func Test_listingsServer_UpdateOwner(t *testing.T) {
	ctx, cancel := config.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, FileServiceClient: MockFileServiceClient(), l: L}
	in := &listingsPB.Owner{
		Id:          1,
		Name:        "Fitore Abazaga",
		Email:       "f.abazaga@gmail.com",
		PhoneNumber: "(646) 339-3247",
		Company:     "Platinum Properties",
		UserRef:     primitive.NewObjectID().Hex(),
	}
	owner, err := server.UpdateOwner(ctx, &listingsPB.UpdateOwnerRequest{Id: 1, Owner: in})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if owner.Email != in.Email {
		t.Errorf("2: Failed to fetch correct owner: %+v", owner)
	}
}

func Test_listingsServer_DeleteOwner(t *testing.T) {
	ctx, cancel := config.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, FileServiceClient: MockFileServiceClient(), l: L}

	in := &listingsPB.Owner{
		Name:        "to_delete",
		Email:       "to_delete",
		PhoneNumber: "to_delete",
		Company:     "to_delete",
		UserRef:     primitive.NewObjectID().Hex(),
	}
	owner, err := server.CreateOwner(ctx, &listingsPB.CreateOwnerRequest{Owner: in})
	if err != nil {
		t.Errorf("1: An error was returned creating a temp owner: %v", err)
	}
	owners, err := server.ListOwners(ctx, &listingsPB.ListOwnerRequest{})
	if err != nil {
		t.Errorf("2: An error was returned: %v", err)
	}
	if len(owners.Owners) < 2 {
		t.Errorf("3: An error adding a temp owner, number of owners in DB: %v", len(owners.Owners))
	}

	deleted, err := server.DeleteOwner(ctx, &listingsPB.DeleteOwnerRequest{Id: owner.Id})
	if err != nil {
		t.Errorf("4: An error was returned: %v", err)
	}
	if deleted.Status != listingsPB.STATUS_SUCCESS {
		t.Errorf("5: Failed to delete owner: %+v\n, %+v", deleted.Status, deleted.GetOwner())
	}
}
