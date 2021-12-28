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
	server := listingsServer{DB: listingDB, FileServiceClient: MockFileServiceClient(), l: L}

	in := &listingsPB.Building{
		Name:         "135 William Street",
		FullAddress:  "135 william st, new york, ny, 10038",
		Street:       "135 William St",
		City:         "New York",
		State:        "NY",
		ZipCode:      10038,
		Neighborhood: "Fulton/Seaport",
		Description:  "Experience the best of New York’s past and present in this prewar converted building, located in the prime Financial District. 16 floors offer 30 sun-drenched loft-style apartments, each outfitted with marble baths, stainless steel appliances, high ceilings, and a washer/dryer. Ride the elevator down before heading mere blocks over to the iconic Pier 17 of South Street Seaport, FiDi’s premier waterfront destination for shopping and dining; Brookfield Place, boasting luxury retail such as Hermès, Paul Smith, and Salvatore Ferragamo and chef-driven restaurants like Le District, New York’s first and only large-format French marketplace; or Stone Street, everyone’s favorite cobblestoned strip for outdoor happy hours. With the 2,3 trains conveniently located on the corner at Fulton Street and the A,C and J,Z only two blocks away, gain unlimited access to all of New York.",
		Amenities:    []string{"Elevator"},
		UploadIds:    []string{},
		RealtorRef:   1,
	}

	building, err := server.CreateBuilding(ctx, &listingsPB.CreateBuildingRequest{Building: in})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if building.Name != in.Name {
		t.Errorf("2: An error creating a building: %+v", building)
	}
}

func Test_listingsServer_GetBuilding(t *testing.T) {
	ctx, cancel := listingsDB.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, FileServiceClient: MockFileServiceClient(), l: L}

	building, err := server.GetBuilding(ctx, &listingsPB.GetBuildingRequest{Id: 1})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if building.Name != "135 William Street" {
		t.Errorf("2: Failed to fetch correct building: %+v", building)
	}
}

func Test_listingsServer_ListBuildings(t *testing.T) {
	ctx, cancel := listingsDB.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, FileServiceClient: MockFileServiceClient(), l: L}

	buildings, err := server.ListBuildings(ctx, &listingsPB.ListBuildingRequest{})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if len(buildings.Buildings) < 1 {
		t.Errorf("2: Failed to fetch buildings: %+v", buildings.Buildings[0])
	}
}

func Test_listingsServer_UpdateBuilding(t *testing.T) {
	ctx, cancel := listingsDB.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, FileServiceClient: MockFileServiceClient(), l: L}
	in := &listingsPB.Building{
		Id:           1,
		Name:         "135 William Street",
		FullAddress:  "135 William Street, New York, NY, 10038",
		Street:       "135 William St",
		City:         "New York",
		State:        "NY",
		ZipCode:      10038,
		Neighborhood: "Fulton/Seaport",
		Description:  "Experience the best of New York’s past and present in this prewar converted building, located in the prime Financial District. 16 floors offer 30 sun-drenched loft-style apartments, each outfitted with marble baths, stainless steel appliances, high ceilings, and a washer/dryer. Ride the elevator down before heading mere blocks over to the iconic Pier 17 of South Street Seaport, FiDi’s premier waterfront destination for shopping and dining; Brookfield Place, boasting luxury retail such as Hermès, Paul Smith, and Salvatore Ferragamo and chef-driven restaurants like Le District, New York’s first and only large-format French marketplace; or Stone Street, everyone’s favorite cobblestoned strip for outdoor happy hours. With the 2,3 trains conveniently located on the corner at Fulton Street and the A,C and J,Z only two blocks away, gain unlimited access to all of New York.",
		Amenities:    []string{"Elevator"},
		UploadIds:    []string{},
		RealtorRef:   1,
	}
	building, err := server.UpdateBuilding(ctx, &listingsPB.UpdateBuildingRequest{Id: 1, Building: in})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if building.FullAddress != in.FullAddress {
		t.Errorf("2: Failed to fetch correct building: %+v", building)
	}
}

func Test_listingsServer_DeleteBuilding(t *testing.T) {
	ctx, cancel := listingsDB.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, FileServiceClient: MockFileServiceClient(), l: L}

	in := &listingsPB.Building{
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
		t.Errorf("2: An error was returned: %v", err)
	}
	if len(buildings.Buildings) < 2 {
		t.Errorf("3: An error adding a temp building, number of buildings in DB: %v", len(buildings.Buildings))
	}

	deleted, err := server.DeleteBuilding(ctx, &listingsPB.DeleteBuildingRequest{Id: building.Id})
	if err != nil {
		t.Errorf("4: An error was returned: %v", err)
	}
	if deleted.Status != listingsPB.STATUS_SUCCESS {
		t.Errorf("5: Failed to delete building: %+v\n, %+v", deleted.Status, deleted.GetBuilding())
	}
}
