package server

import (
	"context"
	"database/sql"
	"testing"
	"time"

	listingsDB "github.com/jalexanderII/zero_microservice/backend/services/listings/database"
	listingsPB "github.com/jalexanderII/zero_microservice/gen/listings"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_listingsServer_CreateApartment(t *testing.T) {
	ctx, cancel := listingsDB.NewDBContext(5 * time.Second)
	defer cancel()

	db, err := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)

	realtor, err := listingDB.CreateRealtor(ctx, listingsDB.CreateRealtorParams{
		RealtorID:   1,
		Name:        "example",
		Email:       sql.NullString{},
		PhoneNumber: sql.NullString{},
		Company:     sql.NullString{},
	})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}

	building, err := listingDB.CreateBuilding(ctx, listingsDB.CreateBuildingParams{
		BuildingID:   1,
		Name:         "example",
		FullAddress:  "example",
		Street:       "example",
		City:         "example",
		State:        "example",
		ZipCode:      10000,
		Neighborhood: "example",
		Lat:          sql.NullInt32{},
		Lng:          sql.NullInt32{},
		Description:  sql.NullString{},
		Amenities:    []string{"example"},
		UploadIds:    []string{"example"},
		RealtorID:    realtor.RealtorID,
	})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}

	server := listingsServer{listingDB}
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
		BuildingRef:  building.BuildingID,
		RealtorRef:   realtor.RealtorID,
	}
	apartment, err := server.CreateApartment(context.Background(), &listingsPB.CreateApartmentRequest{Apartment: in})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if apartment.Id != in.Id {
		t.Errorf("2: Failed to create new apartment: %+v", apartment)
	}
}
