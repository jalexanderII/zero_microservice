package server

import (
	"context"
	"testing"
	"time"

	listingsPB "github.com/jalexanderII/zero_microservice/gen/listings"
	listingsDB "github.com/jalexanderII/zero_microservice/global/db/listings"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_listingsServer_CreateApartment(t *testing.T) {
	listingsDB.ConnectToTestDB("zero_listings")

	var unit = "example"
	exampleApartment := listingsPB.Apartment{
		Id:   primitive.NewObjectID().Hex(),
		Name: "V",
		Address: &listingsPB.Place{
			FullAddress:  "example",
			Street:       "example",
			City:         "example",
			State:        "example",
			ZipCode:      300083,
			Neighborhood: "example",
			Unit:         &unit,
			Coordinates:  &listingsPB.Coordinates{Lat: 32, Lng: 45},
		},
		Rent:  5000,
		Sqft:  5000,
		Beds:  4,
		Baths: 2,
		ListingMetrics: &listingsPB.ListingMetrics{
			AvailableOn:  &timestamppb.Timestamp{Seconds: time.Now().Unix()},
			DaysOnMarket: 1,
		},
		Description: "example",
		Amenities:   []string{"example"},
		Uploads:     []*listingsPB.Content{},
		IsArchived:  false,
		BuildingRef: primitive.NewObjectID().Hex(),
		RealtorRef:  primitive.NewObjectID().Hex(),
	}

	server := NewListingsServer()

	res, err := server.CreateApartment(context.Background(), &listingsPB.CreateApartmentRequest{Apartment: &exampleApartment})
	if err != nil {
		t.Errorf("1: Error calling service: %v", err)
	}
	if res.Id != exampleApartment.Id {
		t.Errorf("2: Wrong Apartment returned: %v", res)
	}
}
