package server

import (
	"testing"
	"time"

	listingsDB "github.com/jalexanderII/zero_microservice/backend/services/listings/database"
	"github.com/jalexanderII/zero_microservice/config"
	listingsPB "github.com/jalexanderII/zero_microservice/gen/listings"
)

func Test_listingsServer_CreateRealtor(t *testing.T) {
	ctx, cancel := config.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, FileServiceClient: MockFileServiceClient(), l: L}

	in := &listingsPB.Realtor{
		Name:        "Fitore Abazaga",
		Email:       "f.abazaga@platinum.com",
		PhoneNumber: "(646) 339-3247",
		Company:     "Platinum Properties",
	}

	realtor, err := server.CreateRealtor(ctx, &listingsPB.CreateRealtorRequest{Realtor: in})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if realtor.Name != in.Name {
		t.Errorf("1: An error creating a realtor: %+v", realtor)
	}
}

func Test_listingsServer_GetRealtor(t *testing.T) {
	ctx, cancel := config.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, FileServiceClient: MockFileServiceClient(), l: L}

	realtor, err := server.GetRealtor(ctx, &listingsPB.GetRealtorRequest{Id: 1})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if realtor.Name != "Fitore Abazaga" {
		t.Errorf("2: Failed to fetch correct realtor: %+v", realtor)
	}
}

func Test_listingsServer_ListRealtors(t *testing.T) {
	ctx, cancel := config.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, FileServiceClient: MockFileServiceClient(), l: L}

	realtors, err := server.ListRealtors(ctx, &listingsPB.ListRealtorRequest{})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if len(realtors.Realtors) < 1 {
		t.Errorf("2: Failed to fetch realtors: %+v", realtors.Realtors[0])
	}
}

func Test_listingsServer_UpdateRealtor(t *testing.T) {
	ctx, cancel := config.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, FileServiceClient: MockFileServiceClient(), l: L}
	in := &listingsPB.Realtor{
		Id:          1,
		Name:        "Fitore Abazaga",
		Email:       "f.abazaga@gmail.com",
		PhoneNumber: "(646) 339-3247",
		Company:     "Platinum Properties",
	}
	realtor, err := server.UpdateRealtor(ctx, &listingsPB.UpdateRealtorRequest{Id: 1, Realtor: in})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if realtor.Email != in.Email {
		t.Errorf("2: Failed to fetch correct realtor: %+v", realtor)
	}
}

func Test_listingsServer_DeleteRealtor(t *testing.T) {
	ctx, cancel := config.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, FileServiceClient: MockFileServiceClient(), l: L}

	in := &listingsPB.Realtor{
		Name:        "to_delete",
		Email:       "to_delete",
		PhoneNumber: "to_delete",
		Company:     "to_delete",
	}
	realtor, err := server.CreateRealtor(ctx, &listingsPB.CreateRealtorRequest{Realtor: in})
	if err != nil {
		t.Errorf("1: An error was returned creating a temp realtor: %v", err)
	}
	realtors, err := server.ListRealtors(ctx, &listingsPB.ListRealtorRequest{})
	if err != nil {
		t.Errorf("2: An error was returned: %v", err)
	}
	if len(realtors.Realtors) < 2 {
		t.Errorf("3: An error adding a temp realtor, number of realtors in DB: %v", len(realtors.Realtors))
	}

	deleted, err := server.DeleteRealtor(ctx, &listingsPB.DeleteRealtorRequest{Id: realtor.Id})
	if err != nil {
		t.Errorf("4: An error was returned: %v", err)
	}
	if deleted.Status != listingsPB.STATUS_SUCCESS {
		t.Errorf("5: Failed to delete realtor: %+v\n, %+v", deleted.Status, deleted.GetRealtor())
	}
}
