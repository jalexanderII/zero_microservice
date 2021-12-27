package server

import (
	"context"

	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	applicationDB "github.com/jalexanderII/zero_microservice/backend/services/application/database"
	"github.com/jalexanderII/zero_microservice/backend/services/application/database/genDB"
	applicationPB "github.com/jalexanderII/zero_microservice/gen/application"
)

type applicationServer struct {
	applicationPB.UnimplementedApplicationServer
	DB *applicationDB.ApplicationDB
	l  hclog.Logger
}

func NewApplicationServer(db *applicationDB.ApplicationDB, l hclog.Logger) *applicationServer {
	return &applicationServer{DB: db, l: l}
}

func (s applicationServer) Apply(ctx context.Context, in *applicationPB.ApplicationRequest) (*applicationPB.ApplicationResponse, error) {
	s.l.Debug("Applying")

	application, err := s.CreateApplication(ctx, &applicationPB.CreateApplicationRequest{Application: in})
	if err != nil {
		s.l.Error("Error applying", "error", err)
		return nil, err
	}
	ar := &applicationPB.ApplicationResponse{
		ReferenceId:    &applicationPB.UUID{Value: uuid.NewString()},
		Status:         string(genDB.ApplicationStatusPENDING),
		Attachments:    []string{},
		ApplicationRef: application.Id,
	}

	response, err := s.CreateResponse(ctx, &applicationPB.CreateApplicationResponse{Application: ar})
	if err != nil {
		s.l.Error("Error submitting application", "error", err)
		return nil, err
	}

	return response, nil
}
