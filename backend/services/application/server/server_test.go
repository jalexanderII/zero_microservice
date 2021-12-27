package server

import (
	"testing"
	"time"

	"github.com/hashicorp/go-hclog"
	applicationDB "github.com/jalexanderII/zero_microservice/backend/services/application/database"
	applicationPB "github.com/jalexanderII/zero_microservice/gen/application"
)

var L = hclog.Default()

func Test_applicationServer_Apply(t *testing.T) {
	ctx, cancel := applicationDB.NewDBContext(5 * time.Second)
	defer cancel()

	db, _ := applicationDB.ConnectToDB()
	appDB := applicationDB.NewApplicationDB(db)
	server := applicationServer{DB: appDB, l: L}

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
