package server

import (
	"testing"
	"time"

	listingsDB "github.com/jalexanderII/zero_microservice/backend/services/listings/database"
	listingsPB "github.com/jalexanderII/zero_microservice/gen/listings"
)

func Test_listingsServer_CreateBuilding(t *testing.T) {
	ctx, cancel := listingsDB.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, ContentStore: Store, l: L}

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

func Test_listingsServer_GetBuilding(t *testing.T) {
	ctx, cancel := listingsDB.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, ContentStore: Store, l: L}

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
	server := listingsServer{DB: listingDB, ContentStore: Store, l: L}

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
	server := listingsServer{DB: listingDB, ContentStore: Store, l: L}
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
	server := listingsServer{DB: listingDB, ContentStore: Store, l: L}

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
