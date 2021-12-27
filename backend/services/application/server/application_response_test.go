package server

import (
	"testing"
	"time"

	"github.com/google/uuid"
	applicationDB "github.com/jalexanderII/zero_microservice/backend/services/application/database"
	"github.com/jalexanderII/zero_microservice/backend/services/application/database/genDB"
	applicationPB "github.com/jalexanderII/zero_microservice/gen/application"
)

func Test_applicationServer_CreateResponse(t *testing.T) {
	ctx, cancel := applicationDB.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := applicationDB.ConnectToDB()
	appDB := applicationDB.NewApplicationDB(db)
	server := applicationServer{DB: appDB, FileServiceClient: MockFileServiceClient(), l: L}

	in := &applicationPB.ApplicationResponse{
		Id:             1,
		ReferenceId:    &applicationPB.UUID{Value: uuid.NewString()},
		Status:         string(genDB.ApplicationStatusAPPROVED),
		ApplicationRef: 1,
	}
	application, err := server.CreateResponse(ctx, &applicationPB.CreateApplicationResponse{Application: in})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if application.Id != in.Id {
		t.Errorf("2: Failed to create new application: %+v", application)
	}
}

func Test_applicationServer_GetResponse(t *testing.T) {
	ctx, cancel := applicationDB.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := applicationDB.ConnectToDB()
	appDB := applicationDB.NewApplicationDB(db)
	server := applicationServer{DB: appDB, FileServiceClient: MockFileServiceClient(), l: L}

	application, err := server.GetResponse(ctx, &applicationPB.GetApplicationResponse{Id: 1})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if application.Id != 1 {
		t.Errorf("2: Failed to fetch correct application: %+v", application)
	}
}

func Test_applicationServer_ListResponses(t *testing.T) {
	ctx, cancel := applicationDB.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := applicationDB.ConnectToDB()
	appDB := applicationDB.NewApplicationDB(db)
	server := applicationServer{DB: appDB, FileServiceClient: MockFileServiceClient(), l: L}

	applications, err := server.ListResponses(ctx, &applicationPB.ListApplicationResponse{})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if applications.Applications[0].Id != 1 {
		t.Errorf("2: Failed to fetch applications: %+v", applications.Applications[0])
	}
}

func Test_applicationServer_UpdateResponse(t *testing.T) {
	ctx, cancel := applicationDB.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := applicationDB.ConnectToDB()
	appDB := applicationDB.NewApplicationDB(db)
	server := applicationServer{DB: appDB, FileServiceClient: MockFileServiceClient(), l: L}
	in := &applicationPB.ApplicationResponse{
		Id:     1,
		Status: string(genDB.ApplicationStatusPENDING),
	}
	application, err := server.UpdateResponse(ctx, &applicationPB.UpdateApplicationResponse{Id: 1, Application: in})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if application.Status != "PENDING" {
		t.Errorf("2: Failed to fetch correct application: %+v", application)
	}
}

func Test_applicationServer_DeleteResponse(t *testing.T) {
	ctx, cancel := applicationDB.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := applicationDB.ConnectToDB()
	appDB := applicationDB.NewApplicationDB(db)
	server := applicationServer{DB: appDB, FileServiceClient: MockFileServiceClient(), l: L}

	in := &applicationPB.ApplicationResponse{
		Id:             2,
		ReferenceId:    &applicationPB.UUID{Value: uuid.NewString()},
		Status:         string(genDB.ApplicationStatusAPPROVED),
		ApplicationRef: 1,
	}
	application, err := server.CreateResponse(ctx, &applicationPB.CreateApplicationResponse{Application: in})
	if err != nil {
		t.Errorf("1: An error was returned creating a temp application: %v", err)
	}
	applications, err := server.ListResponses(ctx, &applicationPB.ListApplicationResponse{})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if len(applications.Applications) != 2 {
		t.Errorf("1: An error adding a temp application, number of applications in DB: %v", len(applications.Applications))
	}

	deleted, err := server.DeleteResponse(ctx, &applicationPB.DeleteApplicationResponse{Id: application.Id})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if deleted.Status != applicationPB.STATUS_SUCCESS {
		t.Errorf("2: Failed to delete application: %+v\n, %+v", deleted.Status, deleted.GetApplication())
	}
}
