package server

import (
	"context"

	"github.com/jalexanderII/zero_microservice/backend/services/application/database/genDB"
	applicationPB "github.com/jalexanderII/zero_microservice/gen/application"
)

func (s applicationServer) CreateResponse(ctx context.Context, in *applicationPB.CreateApplicationResponse) (*applicationPB.ApplicationResponse, error) {
	s.l.Debug("CreateResponse")
	var applicationResponsepb = in.Application

	application, err := s.DB.CreateApplicationResponse(ctx, genDB.CreateApplicationResponseParams{
		Status:        genDB.ApplicationStatus(applicationResponsepb.Status),
		Attachments:   applicationResponsepb.Attachments,
		ApplicationID: applicationResponsepb.ApplicationRef,
	})
	if err != nil {
		s.l.Error("[DB] Error creating application", "error", err)
		return nil, err
	}
	return ResponseDBtoPB(application), nil
}

func (s applicationServer) GetResponse(ctx context.Context, in *applicationPB.GetApplicationResponse) (*applicationPB.ApplicationResponse, error) {
	s.l.Debug("GetApplicationResponse")
	applicationResponse, err := s.DB.GetApplicationResponse(ctx, in.Id)
	if err != nil {
		s.l.Error("[DB] Error getting application by id", "error", err)
		return nil, err
	}
	return ResponseDBtoPB(applicationResponse), nil
}

func (s applicationServer) ListResponses(ctx context.Context, in *applicationPB.ListApplicationResponse) (*applicationPB.ListApplicationResponseResponse, error) {
	s.l.Debug("ListApplicationResponses")
	applications, err := s.DB.ListApplicationResponse(ctx)
	if err != nil {
		s.l.Error("[DB] Error getting all applications", "error", err)
		return nil, err
	}
	res := make([]*applicationPB.ApplicationResponse, len(applications))
	for idx, applicationResponse := range applications {
		res[idx] = ResponseDBtoPB(applicationResponse)
	}
	return &applicationPB.ListApplicationResponseResponse{Applications: res}, nil
}

func (s applicationServer) UpdateResponse(ctx context.Context, in *applicationPB.UpdateApplicationResponse) (*applicationPB.ApplicationResponse, error) {
	s.l.Debug("UpdateApplicationResponse")
	var applicationResponsepb = in.Application
	err := s.DB.UpdateApplicationResponse(ctx, genDB.UpdateApplicationResponseParams{
		ApplicationResponseID: in.Id,
		Status:                genDB.ApplicationStatus(applicationResponsepb.Status),
		Attachments:           applicationResponsepb.Attachments,
	})
	if err != nil {
		s.l.Error("[DB] Error updating application", "error", err)
		return nil, err
	}
	return applicationResponsepb, nil
}

func (s applicationServer) DeleteResponse(ctx context.Context, in *applicationPB.DeleteApplicationResponse) (*applicationPB.DeleteApplicationResponseResponse, error) {
	s.l.Debug("DeleteApplicationResponse")
	applicationResponse, err := s.DB.GetApplicationResponse(ctx, in.Id)
	if err != nil {
		s.l.Error("[DB] Error fetching application to delete", "error", err)
		return nil, err
	}

	err = s.DB.DeleteApplicationResponse(ctx, in.Id)
	if err != nil {
		s.l.Error("[DB] Error deleting application by ID", "error", err)
		return nil, err
	}
	return &applicationPB.DeleteApplicationResponseResponse{Status: applicationPB.STATUS_SUCCESS, Application: ResponseDBtoPB(applicationResponse)}, nil
}

func ResponseDBtoPB(applicationResponse genDB.ApplicationResponse) *applicationPB.ApplicationResponse {
	return &applicationPB.ApplicationResponse{
		Id:             applicationResponse.ApplicationResponseID,
		ReferenceId:    &applicationPB.UUID{Value: applicationResponse.ReferenceID.UUID.String()},
		Status:         string(applicationResponse.Status),
		Attachments:    applicationResponse.Attachments,
		ApplicationRef: applicationResponse.ApplicationID,
	}
}
