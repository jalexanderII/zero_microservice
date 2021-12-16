package server

import (
	"testing"
	"time"

	"github.com/hashicorp/go-hclog"
	listingsDB "github.com/jalexanderII/zero_microservice/backend/services/listings/database"
	contentStore "github.com/jalexanderII/zero_microservice/backend/services/listings/store"
	listingsPB "github.com/jalexanderII/zero_microservice/gen/listings"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var l = hclog.Default()
var store = contentStore.NewDiskImageStore("./store/tmp", l)

func Test_listingsServer_CreateRealtor(t *testing.T) {
	ctx, cancel := listingsDB.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, ContentStore: store, l: l}

	in := &listingsPB.Realtor{
		Id:          2,
		Name:        "example2",
		Email:       "example2",
		PhoneNumber: "example2",
		Company:     "example2",
	}

	realtor, err := server.CreateRealtor(ctx, &listingsPB.CreateRealtorRequest{Realtor: in})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if realtor.Id != 2 && realtor.Name != "example2" {
		t.Errorf("1: An error creating a realtor: %+v", realtor)
	}
}

func Test_listingsServer_CreateBuilding(t *testing.T) {
	ctx, cancel := listingsDB.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, ContentStore: store, l: l}

	in := &listingsPB.Building{
		Id:           2,
		Name:         "example2",
		FullAddress:  "example2",
		Street:       "example2",
		City:         "example2",
		State:        "example2",
		ZipCode:      10000,
		Neighborhood: "example2",
		Lat:          2143,
		Lng:          4345,
		Description:  "example2",
		Amenities:    []string{"example2"},
		UploadIds:    []string{"example2"},
		RealtorRef:   2,
	}

	building, err := server.CreateBuilding(ctx, &listingsPB.CreateBuildingRequest{Building: in})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if building.Id != 2 && building.Name != "example2" {
		t.Errorf("1: An error creating a building: %+v", building)
	}
}

func Test_listingsServer_CreateApartment(t *testing.T) {
	ctx, cancel := listingsDB.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, ContentStore: store, l: l}

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
	server := listingsServer{DB: listingDB, ContentStore: store, l: l}

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
	server := listingsServer{DB: listingDB, ContentStore: store, l: l}

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
	server := listingsServer{DB: listingDB, ContentStore: store, l: l}
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
	server := listingsServer{DB: listingDB, ContentStore: store, l: l}

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

func Test_listingsServer_GetBuilding(t *testing.T) {
	ctx, cancel := listingsDB.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, ContentStore: store, l: l}

	building, err := server.GetBuilding(ctx, &listingsPB.GetBuildingRequest{Id: 2})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if building.Name != "example2" {
		t.Errorf("2: Failed to fetch correct building: %+v", building)
	}
}

func Test_listingsServer_ListBuildings(t *testing.T) {
	ctx, cancel := listingsDB.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, ContentStore: store, l: l}

	buildings, err := server.ListBuildings(ctx, &listingsPB.ListBuildingRequest{})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if buildings.Buildings[0].Name != "Updated" {
		t.Errorf("2: Failed to fetch buildings: %+v", buildings.Buildings[0])
	}
}

func Test_listingsServer_UpdateBuilding(t *testing.T) {
	ctx, cancel := listingsDB.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, ContentStore: store, l: l}
	in := &listingsPB.Building{
		Id:           1,
		Name:         "Updated",
		FullAddress:  "Updated",
		Street:       "Updated",
		City:         "Updated",
		State:        "Updated",
		ZipCode:      10000,
		Neighborhood: "Updated",
		Lat:          2143,
		Lng:          4345,
		Description:  "Updated",
		Amenities:    []string{"example", "Updated"},
		UploadIds:    []string{"example", "Updated"},
		RealtorRef:   1,
	}
	building, err := server.UpdateBuilding(ctx, &listingsPB.UpdateBuildingRequest{Id: 1, Building: in})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if building.Name != "Updated" {
		t.Errorf("2: Failed to fetch correct building: %+v", building)
	}
}

func Test_listingsServer_DeleteBuilding(t *testing.T) {
	ctx, cancel := listingsDB.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, ContentStore: store, l: l}

	in := &listingsPB.Building{
		Id:           3,
		Name:         "to_delete",
		FullAddress:  "to_delete",
		Street:       "to_delete",
		City:         "to_delete",
		State:        "to_delete",
		ZipCode:      10000,
		Neighborhood: "to_delete",
		Lat:          2143,
		Lng:          4345,
		Description:  "to_delete",
		Amenities:    []string{"to_delete"},
		UploadIds:    []string{"to_delete"},
		RealtorRef:   1,
	}
	building, err := server.CreateBuilding(ctx, &listingsPB.CreateBuildingRequest{Building: in})
	if err != nil {
		t.Errorf("1: An error was returned creating a temp building: %v", err)
	}
	buildings, err := server.ListBuildings(ctx, &listingsPB.ListBuildingRequest{})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if len(buildings.Buildings) != 3 {
		t.Errorf("1: An error adding a temp building, number of buildings in DB: %v", len(buildings.Buildings))
	}

	deleted, err := server.DeleteBuilding(ctx, &listingsPB.DeleteBuildingRequest{Id: building.Id})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if deleted.Status != listingsPB.STATUS_SUCCESS {
		t.Errorf("2: Failed to delete building: %+v\n, %+v", deleted.Status, deleted.GetBuilding())
	}
}

func Test_listingsServer_GetRealtor(t *testing.T) {
	ctx, cancel := listingsDB.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, ContentStore: store, l: l}

	realtor, err := server.GetRealtor(ctx, &listingsPB.GetRealtorRequest{Id: 2})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if realtor.Name != "example2" {
		t.Errorf("2: Failed to fetch correct realtor: %+v", realtor)
	}
}

func Test_listingsServer_ListRealtors(t *testing.T) {
	ctx, cancel := listingsDB.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, ContentStore: store, l: l}

	realtors, err := server.ListRealtors(ctx, &listingsPB.ListRealtorRequest{})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if realtors.Realtors[0].Name != "example" {
		t.Errorf("2: Failed to fetch realtors: %+v", realtors.Realtors[0])
	}
}

func Test_listingsServer_UpdateRealtor(t *testing.T) {
	ctx, cancel := listingsDB.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, ContentStore: store, l: l}
	in := &listingsPB.Realtor{
		Id:          1,
		Name:        "Updated",
		Email:       "Updated",
		PhoneNumber: "Updated",
		Company:     "Updated",
	}
	realtor, err := server.UpdateRealtor(ctx, &listingsPB.UpdateRealtorRequest{Id: 1, Realtor: in})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if realtor.Name != "Updated" {
		t.Errorf("2: Failed to fetch correct realtor: %+v", realtor)
	}
}

func Test_listingsServer_DeleteRealtor(t *testing.T) {
	ctx, cancel := listingsDB.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, ContentStore: store, l: l}

	in := &listingsPB.Realtor{
		Id:          3,
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
		t.Errorf("1: An error was returned: %v", err)
	}
	if len(realtors.Realtors) != 3 {
		t.Errorf("1: An error adding a temp realtor, number of realtors in DB: %v", len(realtors.Realtors))
	}

	deleted, err := server.DeleteRealtor(ctx, &listingsPB.DeleteRealtorRequest{Id: realtor.Id})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if deleted.Status != listingsPB.STATUS_SUCCESS {
		t.Errorf("2: Failed to delete realtor: %+v\n, %+v", deleted.Status, deleted.GetRealtor())
	}
}
