package server

import (
	"testing"
	"time"

	"github.com/hashicorp/go-hclog"
	applicationDB "github.com/jalexanderII/zero_microservice/backend/services/application/database"
	applicationPB "github.com/jalexanderII/zero_microservice/gen/application"
)

var L = hclog.Default()

func Test_applicationServer_CreateApplication(t *testing.T) {
	ctx, cancel := applicationDB.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := applicationDB.ConnectToDB()
	appDB := applicationDB.NewApplicationDB(db)
	server := applicationServer{DB: appDB, l: L}

	in := &applicationPB.ApplicationRequest{
		Id:                     1,
		Name:                   "example",
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
	application, err := server.CreateApplication(ctx, &applicationPB.CreateApplicationRequest{Application: in})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if application.Id != in.Id {
		t.Errorf("2: Failed to create new application: %+v", application)
	}
}

func Test_applicationServer_GetApplication(t *testing.T) {
	ctx, cancel := applicationDB.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := applicationDB.ConnectToDB()
	appDB := applicationDB.NewApplicationDB(db)
	server := applicationServer{DB: appDB, l: L}

	application, err := server.GetApplication(ctx, &applicationPB.GetApplicationRequest{Id: 1})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if application.Id != 1 {
		t.Errorf("2: Failed to fetch correct application: %+v", application)
	}
}

func Test_applicationServer_ListApplications(t *testing.T) {
	ctx, cancel := applicationDB.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := applicationDB.ConnectToDB()
	appDB := applicationDB.NewApplicationDB(db)
	server := applicationServer{DB: appDB, l: L}

	applications, err := server.ListApplications(ctx, &applicationPB.ListApplicationRequest{})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if applications.Applications[0].Id != 1 {
		t.Errorf("2: Failed to fetch applications: %+v", applications.Applications[0])
	}
}

func Test_applicationServer_UpdateApplication(t *testing.T) {
	ctx, cancel := applicationDB.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := applicationDB.ConnectToDB()
	appDB := applicationDB.NewApplicationDB(db)
	server := applicationServer{DB: appDB, l: L}
	in := &applicationPB.ApplicationRequest{
		Name:                   "Updated",
		PreviousAddress:        "Updated",
		PreviousLandlord:       "Updated",
		PreviousLandlordNumber: "Updated",
		Employer:               "Updated",
		Salary:                 175000,
	}
	application, err := server.UpdateApplication(ctx, &applicationPB.UpdateApplicationRequest{Id: 1, Application: in})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if application.Name != "Updated" {
		t.Errorf("2: Failed to fetch correct application: %+v", application)
	}
}

func Test_applicationServer_DeleteApplication(t *testing.T) {
	ctx, cancel := applicationDB.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := applicationDB.ConnectToDB()
	appDB := applicationDB.NewApplicationDB(db)
	server := applicationServer{DB: appDB, l: L}

	in := &applicationPB.ApplicationRequest{
		Id:                     2,
		Name:                   "to_delete",
		SocialSecurity:         "to_delete",
		DateOfBirth:            "11/19/1990",
		DriversLicense:         "000666953",
		PreviousAddress:        "to_delete",
		PreviousLandlord:       "to_delete",
		PreviousLandlordNumber: "to_delete",
		Employer:               "to_delete",
		Salary:                 150000,
		UserRef:                1,
		ApartmentRef:           1,
	}
	application, err := server.CreateApplication(ctx, &applicationPB.CreateApplicationRequest{Application: in})
	if err != nil {
		t.Errorf("1: An error was returned creating a temp application: %v", err)
	}
	applications, err := server.ListApplications(ctx, &applicationPB.ListApplicationRequest{})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if len(applications.Applications) != 2 {
		t.Errorf("1: An error adding a temp application, number of applications in DB: %v", len(applications.Applications))
	}

	deleted, err := server.DeleteApplication(ctx, &applicationPB.DeleteApplicationRequest{Id: application.Id})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if deleted.Status != applicationPB.STATUS_SUCCESS {
		t.Errorf("2: Failed to delete application: %+v\n, %+v", deleted.Status, deleted.GetApplication())
	}
}
