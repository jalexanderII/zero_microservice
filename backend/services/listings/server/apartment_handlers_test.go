package server

import (
	"testing"
	"time"

	listingsDB "github.com/jalexanderII/zero_microservice/backend/services/listings/database"
	listingsPB "github.com/jalexanderII/zero_microservice/gen/listings"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_listingsServer_CreateApartment(t *testing.T) {
	ctx, cancel := listingsDB.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, ContentStore: Store, l: L}

	in := &listingsPB.Apartment{
		Id:           1,
		Name:         "example",
		FullAddress:  "example",
		Street:       "example",
		City:         "example",
		State:        "example",
		ZipCode:      10000,
		Neighborhood: "example",
		Unit:         "example",
		Lat:          21343,
		Lng:          32434,
		Rent:         1000,
		Sqft:         1000,
		Beds:         4,
		Baths:        1,
		AvailableOn:  timestamppb.Now(),
		DaysOnMarket: 1,
		Description:  "example",
		Amenities:    []string{"example"},
		UploadIds:    []string{"example"},
		IsArchived:   false,
		BuildingRef:  1,
		RealtorRef:   1,
	}
	apartment, err := server.CreateApartment(ctx, &listingsPB.CreateApartmentRequest{Apartment: in})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if apartment.Id != in.Id {
		t.Errorf("2: Failed to create new apartment: %+v", apartment)
	}
}

func Test_listingsServer_GetApartment(t *testing.T) {
	ctx, cancel := listingsDB.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, ContentStore: Store, l: L}

	apartment, err := server.GetApartment(ctx, &listingsPB.GetApartmentRequest{Id: 1})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if apartment.Id != 1 {
		t.Errorf("2: Failed to fetch correct apartment: %+v", apartment)
	}
}

func Test_listingsServer_ListApartments(t *testing.T) {
	ctx, cancel := listingsDB.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, ContentStore: Store, l: L}

	apartments, err := server.ListApartments(ctx, &listingsPB.ListApartmentRequest{})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if apartments.Apartments[0].Id != 1 {
		t.Errorf("2: Failed to fetch apartments: %+v", apartments.Apartments[0])
	}
}

func Test_listingsServer_UpdateApartment(t *testing.T) {
	ctx, cancel := listingsDB.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, ContentStore: Store, l: L}
	in := &listingsPB.Apartment{
		Id:           1,
		Name:         "Updated",
		FullAddress:  "Updated",
		Street:       "Updated",
		City:         "Updated",
		State:        "Updated",
		ZipCode:      10000,
		Neighborhood: "Updated",
		Unit:         "Updated",
		Lat:          21343,
		Lng:          32434,
		Rent:         1000,
		Sqft:         1000,
		Beds:         4,
		Baths:        1,
		AvailableOn:  timestamppb.Now(),
		DaysOnMarket: 1,
		Description:  "Updated",
		Amenities:    []string{"example", "Updated"},
		UploadIds:    []string{"example", "Updated"},
		IsArchived:   false,
		BuildingRef:  1,
		RealtorRef:   1,
	}
	apartment, err := server.UpdateApartment(ctx, &listingsPB.UpdateApartmentRequest{Id: 1, Apartment: in})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if apartment.Name != "Updated" {
		t.Errorf("2: Failed to fetch correct apartment: %+v", apartment)
	}
}

func Test_listingsServer_DeleteApartment(t *testing.T) {
	ctx, cancel := listingsDB.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, ContentStore: Store, l: L}

	in := &listingsPB.Apartment{
		Id:           2,
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
		t.Errorf("1: An error was returned: %v", err)
	}
	if len(apartments.Apartments) != 2 {
		t.Errorf("1: An error adding a temp apartment, number of apartments in DB: %v", len(apartments.Apartments))
	}

	deleted, err := server.DeleteApartment(ctx, &listingsPB.DeleteApartmentRequest{Id: apartment.Id})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if deleted.Status != listingsPB.STATUS_SUCCESS {
		t.Errorf("2: Failed to delete apartment: %+v\n, %+v", deleted.Status, deleted.GetApartment())
	}
}
