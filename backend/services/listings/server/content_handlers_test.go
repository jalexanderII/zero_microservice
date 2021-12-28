package server

import (
	"testing"
	"time"

	listingsDB "github.com/jalexanderII/zero_microservice/backend/services/listings/database"
	listingsPB "github.com/jalexanderII/zero_microservice/gen/listings"
)

func Test_listingsServer_Upload(t *testing.T) {
	ctx, cancel := listingsDB.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := listingsDB.ConnectToDB()
	listingDB := listingsDB.NewListingsDB(db)
	server := listingsServer{DB: listingDB, FileServiceClient: MockFileServiceClient(), l: L}

	md := &listingsPB.MetaData{
		Name:          "icon.png",
		SourceId:      1,
		ContentType:   listingsPB.ContentType_IMAGE,
		ContentSource: listingsPB.ContentSource_APARTMENT,
	}

	in := &listingsPB.FileUploadRequest{Metadata: md, FilePath: "/Users/joel/Downloads/icon.png"}

	upload, err := server.Upload(ctx, in)
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if upload.Status != listingsPB.STATUS_SUCCESS {
		t.Errorf("2: An error uploading new content: %v", upload)
	}
	apartment, err := server.DB.GetApartment(ctx, md.SourceId)
	if err != nil {
		t.Errorf("3: An error fetching apartment: %v", err)
	}
	if apartment.ApartmentID != md.SourceId {
		t.Errorf("4: fetched wrong apartment: %v", apartment)
	}
	if !contains(apartment.UploadIds, md.Name) {
		t.Errorf("5: apartment was not updated with new upload: %v", apartment.UploadIds)
	}
}

// https://play.golang.org/p/Qg_uv_inCek
// contains checks if a string is present in a slice
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
