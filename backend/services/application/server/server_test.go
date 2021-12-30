package server

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/go-hclog"
	config "github.com/jalexanderII/zero_microservice"
	applicationDB "github.com/jalexanderII/zero_microservice/backend/services/application/database"
	applicationPB "github.com/jalexanderII/zero_microservice/gen/application"
	fileServicePB "github.com/jalexanderII/zero_microservice/gen/file_service"
	"google.golang.org/grpc"
)

var L = hclog.Default()

func MockFileServiceClient() fileServicePB.FileServiceClient {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", config.FIlESERVICESERVERPORT), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return fileServicePB.NewFileServiceClient(conn)
}

func Test_applicationServer_Apply(t *testing.T) {
	ctx, cancel := config.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := applicationDB.ConnectToDB()
	appDB := applicationDB.NewApplicationDB(db)
	server := applicationServer{DB: appDB, FileServiceClient: MockFileServiceClient(), l: L}

	applications, err := server.ListApplications(ctx, &applicationPB.ListApplicationRequest{})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	originalCount := len(applications.Applications)

	in := &applicationPB.ApplicationRequest{
		Name:                   "from application",
		SocialSecurity:         "example",
		DateOfBirth:            "11/19/1990",
		DriversLicense:         "000666953",
		PreviousAddress:        "example",
		PreviousLandlord:       "example",
		PreviousLandlordNumber: "example",
		Employer:               "example",
		Salary:                 150000,
		UserRef:                1,
		ApartmentRef:           1,
	}
	response, err := server.Apply(ctx, in)
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}

	applications, err = server.ListApplications(ctx, &applicationPB.ListApplicationRequest{})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	newCount := len(applications.Applications)
	if newCount != originalCount+1 {
		t.Errorf("2: Failed to create new application: %+v", response)
	}
	if applications.Applications[newCount-1].Name != in.Name {
		t.Errorf("3: Wrong application created: %+v", applications.Applications[newCount-1])
	}
	if response.Status != "PENDING" {
		t.Errorf("4: Failed to submit new application: %+v", response)
	}
}

func Test_applicationServer_Upload(t *testing.T) {
	ctx, cancel := config.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := applicationDB.ConnectToDB()
	appDB := applicationDB.NewApplicationDB(db)
	server := applicationServer{DB: appDB, FileServiceClient: MockFileServiceClient(), l: L}

	md := &applicationPB.MetaData{
		Name:          "icon.png",
		SourceId:      5,
		ContentType:   applicationPB.ContentType_IMAGE,
		ContentSource: "APPLICATION",
	}

	in := &applicationPB.FileUploadRequest{Metadata: md, FilePath: "/Users/joel/Downloads/icon.png"}

	upload, err := server.Upload(ctx, in)
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if upload.Status != applicationPB.STATUS_SUCCESS {
		t.Errorf("2: An error uploading new content: %v", upload)
	}
	application, err := server.DB.GetApplicationRequest(ctx, md.SourceId)
	if err != nil {
		t.Errorf("3: An error fetching application: %v", err)
	}
	if application.ApplicationRequestID != md.SourceId {
		t.Errorf("4: fetched wrong application: %v", application)
	}
	if !contains(application.Attachments, md.Name) {
		t.Errorf("5: application was not updated with new upload: %v", application.Attachments)
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
