package server

import (
	"context"
	"database/sql"

	"github.com/jalexanderII/zero_microservice/backend/services/application/database/genDB"
	applicationPB "github.com/jalexanderII/zero_microservice/gen/application"
)

func (s applicationServer) CreateApplication(ctx context.Context, in *applicationPB.CreateApplicationRequest) (*applicationPB.ApplicationRequest, error) {
	s.l.Debug("CreateApplication")
	var applicationpb = in.Application

	application, err := s.DB.CreateApplicationRequest(ctx, genDB.CreateApplicationRequestParams{
		Name:                   applicationpb.Name,
		SocialSecurity:         applicationpb.SocialSecurity,
		DateOfBirth:            applicationpb.DateOfBirth,
		DriversLicense:         applicationpb.DriversLicense,
		PreviousAddress:        sql.NullString{String: applicationpb.PreviousAddress, Valid: true},
		PreviousLandlord:       sql.NullString{String: applicationpb.PreviousLandlord, Valid: true},
		PreviousLandlordNumber: sql.NullString{String: applicationpb.PreviousLandlordNumber, Valid: true},
		Employer:               sql.NullString{String: applicationpb.Employer, Valid: true},
		Salary:                 applicationpb.Salary,
		UserID:                 applicationpb.UserRef,
		ApartmentID:            applicationpb.ApartmentRef,
	})
	if err != nil {
		s.l.Error("[DB] Error creating application", "error", err)
		return nil, err
	}
	return ApplicationDBtoPB(application), nil
}

func (s applicationServer) GetApplication(ctx context.Context, in *applicationPB.GetApplicationRequest) (*applicationPB.ApplicationRequest, error) {
	s.l.Debug("GetApplication")
	application, err := s.DB.GetApplicationRequest(ctx, in.Id)
	if err != nil {
		s.l.Error("[DB] Error getting application by id", "error", err)
		return nil, err
	}
	return ApplicationDBtoPB(application), nil
}

func (s applicationServer) ListApplications(ctx context.Context, in *applicationPB.ListApplicationRequest) (*applicationPB.ListApplicationRequestResponse, error) {
	s.l.Debug("ListApplications")
	applications, err := s.DB.ListApplicationRequest(ctx)
	if err != nil {
		s.l.Error("[DB] Error getting all applications", "error", err)
		return nil, err
	}
	res := make([]*applicationPB.ApplicationRequest, len(applications))
	for idx, application := range applications {
		res[idx] = ApplicationDBtoPB(application)
	}
	return &applicationPB.ListApplicationRequestResponse{Applications: res}, nil
}

func (s applicationServer) UpdateApplication(ctx context.Context, in *applicationPB.UpdateApplicationRequest) (*applicationPB.ApplicationRequest, error) {
	s.l.Debug("UpdateApplication")
	var applicationpb = in.Application
	err := s.DB.UpdateApplicationRequest(ctx, genDB.UpdateApplicationRequestParams{
		ApplicationRequestID:   in.Id,
		Name:                   applicationpb.Name,
		PreviousAddress:        sql.NullString{String: applicationpb.PreviousAddress, Valid: true},
		PreviousLandlord:       sql.NullString{String: applicationpb.PreviousLandlord, Valid: true},
		PreviousLandlordNumber: sql.NullString{String: applicationpb.PreviousLandlordNumber, Valid: true},
		Employer:               sql.NullString{String: applicationpb.Employer, Valid: true},
		Salary:                 applicationpb.Salary,
	})
	if err != nil {
		s.l.Error("[DB] Error updating application", "error", err)
		return nil, err
	}
	return applicationpb, nil
}

func (s applicationServer) DeleteApplication(ctx context.Context, in *applicationPB.DeleteApplicationRequest) (*applicationPB.DeleteApplicationRequestResponse, error) {
	s.l.Debug("DeleteApplication")
	application, err := s.DB.GetApplicationRequest(ctx, in.Id)
	if err != nil {
		s.l.Error("[DB] Error fetching application to delete", "error", err)
		return nil, err
	}

	err = s.DB.DeleteApplicationRequest(ctx, in.Id)
	if err != nil {
		s.l.Error("[DB] Error deleting application by ID", "error", err)
		return nil, err
	}
	return &applicationPB.DeleteApplicationRequestResponse{Status: applicationPB.STATUS_SUCCESS, Application: ApplicationDBtoPB(application)}, nil
}

func ApplicationDBtoPB(application genDB.Application) *applicationPB.ApplicationRequest {
	return &applicationPB.ApplicationRequest{
		Id:                     application.ApplicationRequestID,
		Name:                   application.Name,
		UserRef:                application.UserID,
		SocialSecurity:         application.SocialSecurity,
		DateOfBirth:            application.DateOfBirth,
		DriversLicense:         application.DriversLicense,
		PreviousAddress:        application.PreviousAddress.String,
		PreviousLandlord:       application.PreviousLandlord.String,
		PreviousLandlordNumber: application.PreviousLandlordNumber.String,
		Employer:               application.Employer.String,
		Salary:                 application.Salary,
		ApartmentRef:           application.ApartmentID,
	}
}
