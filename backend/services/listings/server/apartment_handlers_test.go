package server

import (
	"testing"
	"time"

	listingsDB "github.com/jalexanderII/zero_microservice/backend/services/listings/database"
	"github.com/jalexanderII/zero_microservice/config"
	listingsPB "github.com/jalexanderII/zero_microservice/gen/listings"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_listingsServer_CreateApartment(t *testing.T) {
	ctx, cancel := config.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, FileServiceClient: MockFileServiceClient(), l: L}

	in := &listingsPB.Apartment{
		Name:         "22 Eldridge St",
		FullAddress:  "22 Eldridge St, New York, NY 10002",
		Street:       "22 Eldridge St",
		City:         "New York",
		State:        "NY",
		ZipCode:      10002,
		Neighborhood: "ChinaTown",
		Unit:         "3",
		Rent:         5000,
		Sqft:         1000,
		Beds:         2,
		Baths:        1,
		AvailableOn:  timestamppb.Now(),
		DaysOnMarket: 1,
		Description:  "example",
		Amenities:    []string{"cats allowed"},
		UploadIds:    []string{},
		IsArchived:   true,
		BuildingRef:  1,
		RealtorRef:   1,
	}
	apartment, err := server.CreateApartment(ctx, &listingsPB.CreateApartmentRequest{Apartment: in})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if apartment.Name != in.Name {
		t.Errorf("2: Failed to create new apartment: %+v", apartment)
	}
	if apartment.Lat != -73.993546 && apartment.Lng != 40.71506 {
		t.Errorf("3: Failed to get coordinates using googleplaces api: %+v", apartment)
	}
}

func Test_listingsServer_GetApartment(t *testing.T) {
	ctx, cancel := config.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, FileServiceClient: MockFileServiceClient(), l: L}

	apartment, err := server.GetApartment(ctx, &listingsPB.GetApartmentRequest{Id: 1})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if apartment.Name != "22 Eldridge St" {
		t.Errorf("2: Failed to fetch correct apartment: %+v", apartment)
	}
}

func Test_listingsServer_ListApartments(t *testing.T) {
	ctx, cancel := config.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, FileServiceClient: MockFileServiceClient(), l: L}

	apartments, err := server.ListApartments(ctx, &listingsPB.ListApartmentRequest{})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if len(apartments.Apartments) < 1 {
		t.Errorf("2: Failed to fetch apartments: %+v", apartments.Apartments[0])
	}
}

func Test_listingsServer_UpdateApartment(t *testing.T) {
	ctx, cancel := config.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, FileServiceClient: MockFileServiceClient(), l: L}
	in := &listingsPB.Apartment{
		Id:           1,
		Name:         "22 Eldridge St",
		FullAddress:  "22 Eldridge St, New York, NY 10002",
		Street:       "22 Eldridge St",
		City:         "New York",
		State:        "NY",
		ZipCode:      10002,
		Neighborhood: "ChinaTown",
		Unit:         "3",
		Rent:         5000,
		Sqft:         1000,
		Beds:         2,
		Baths:        1,
		AvailableOn:  timestamppb.Now(),
		DaysOnMarket: 1,
		Description:  "- Unique Commercial Loft\n- 3rd Floor Walk-up\n- 1,000 sqft\n- 400 sqft outdoor (Roof deck has been redone will provide pict)\n- 2 Rooms set up street view\n- Full Kitchen\n- Nice Bathroom\n- Lot of Storage\n- Crossing Canal Street",
		Amenities:    []string{"cats allowed", "Private Outdoor Space"},
		UploadIds:    []string{},
		IsArchived:   true,
		BuildingRef:  1,
		RealtorRef:   1,
	}
	apartment, err := server.UpdateApartment(ctx, &listingsPB.UpdateApartmentRequest{Id: 1, Apartment: in})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if apartment.Description != in.Description {
		t.Errorf("2: Failed to fetch correct apartment: %+v", apartment)
	}
}

func Test_listingsServer_DeleteApartment(t *testing.T) {
	ctx, cancel := config.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, FileServiceClient: MockFileServiceClient(), l: L}

	in := &listingsPB.Apartment{
		Name:         "to_delete",
		FullAddress:  "to_delete",
		Street:       "to_delete",
		City:         "to_delete",
		State:        "to_delete",
		ZipCode:      10000,
		Neighborhood: "to_delete",
		Unit:         "to_delete",
		Lat:          21343,
		Lng:          32434,
		Rent:         1000,
		Sqft:         1000,
		Beds:         4,
		Baths:        1,
		AvailableOn:  timestamppb.Now(),
		DaysOnMarket: 1,
		Description:  "to_delete",
		Amenities:    []string{"to_delete"},
		UploadIds:    []string{"to_delete"},
		IsArchived:   false,
		BuildingRef:  1,
		RealtorRef:   1,
	}
	apartment, err := server.CreateApartment(ctx, &listingsPB.CreateApartmentRequest{Apartment: in})
	if err != nil {
		t.Errorf("1: An error was returned creating a temp apartment: %v", err)
	}
	apartments, err := server.ListApartments(ctx, &listingsPB.ListApartmentRequest{})
	if err != nil {
		t.Errorf("2: An error was returned: %v", err)
	}
	if len(apartments.Apartments) < 2 {
		t.Errorf("3: An error adding a temp apartment, number of apartments in DB: %v", len(apartments.Apartments))
	}

	deleted, err := server.DeleteApartment(ctx, &listingsPB.DeleteApartmentRequest{Id: apartment.Id})
	if err != nil {
		t.Errorf("4: An error was returned: %v", err)
	}
	if deleted.Status != listingsPB.STATUS_SUCCESS {
		t.Errorf("5: Failed to delete apartment: %+v\n, %+v", deleted.Status, deleted.GetApartment())
	}
}
