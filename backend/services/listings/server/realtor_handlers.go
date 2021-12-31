package server

import (
	"context"
	"database/sql"

	"github.com/jalexanderII/zero_microservice/backend/services/listings/database/genDB"
	listingsPB "github.com/jalexanderII/zero_microservice/gen/listings"
)

func (s listingsServer) CreateRealtor(ctx context.Context, in *listingsPB.CreateRealtorRequest) (*listingsPB.Realtor, error) {
	s.l.Debug("CreateRealtor")
	var realtorpb = in.Realtor
	realtor, err := s.DB.CreateRealtor(ctx, genDB.CreateRealtorParams{
		Name:        realtorpb.Name,
		UserID:      realtorpb.UserRef,
		Email:       sql.NullString{String: realtorpb.Email, Valid: true},
		PhoneNumber: sql.NullString{String: realtorpb.PhoneNumber, Valid: true},
		Company:     sql.NullString{String: realtorpb.Company, Valid: true},
	})
	if err != nil {
		s.l.Error("[DB] Error creating realtor", "error", err)
		return nil, err
	}
	return RealtorDBtoPB(realtor), nil
}

func (s listingsServer) GetRealtor(ctx context.Context, in *listingsPB.GetRealtorRequest) (*listingsPB.Realtor, error) {
	s.l.Debug("GetRealtor")
	realtor, err := s.DB.GetRealtor(ctx, in.Id)
	if err != nil {
		s.l.Error("[DB] Error getting realtor by id", "error", err)
		return nil, err
	}
	return RealtorDBtoPB(realtor), nil
}

func (s listingsServer) ListRealtors(ctx context.Context, in *listingsPB.ListRealtorRequest) (*listingsPB.ListRealtorResponse, error) {
	s.l.Debug("ListRealtors")
	realtors, err := s.DB.ListRealtors(ctx)
	if err != nil {
		s.l.Error("[DB] Error getting all realtors", "error", err)
		return nil, err
	}
	res := make([]*listingsPB.Realtor, len(realtors))
	for idx, realtor := range realtors {
		res[idx] = RealtorDBtoPB(realtor)
	}
	return &listingsPB.ListRealtorResponse{Realtors: res}, nil
}

func (s listingsServer) UpdateRealtor(ctx context.Context, in *listingsPB.UpdateRealtorRequest) (*listingsPB.Realtor, error) {
	s.l.Debug("UpdateRealtor")
	var realtorpb = in.Realtor
	err := s.DB.UpdateRealtor(ctx, genDB.UpdateRealtorParams{
		RealtorID:   realtorpb.Id,
		Name:        realtorpb.Name,
		Email:       sql.NullString{String: realtorpb.Email, Valid: true},
		PhoneNumber: sql.NullString{String: realtorpb.PhoneNumber, Valid: true},
		Company:     sql.NullString{String: realtorpb.Company, Valid: true},
	})
	if err != nil {
		s.l.Error("[DB] Error updating realtor", "error", err)
		return nil, err
	}
	return realtorpb, nil
}

func (s listingsServer) DeleteRealtor(ctx context.Context, in *listingsPB.DeleteRealtorRequest) (*listingsPB.DeleteRealtorResponse, error) {
	s.l.Debug("DeleteRealtor")
	realtor, err := s.DB.GetRealtor(ctx, in.Id)
	if err != nil {
		s.l.Error("[DB] Error fetching realtor to delete", "error", err)
		return nil, err
	}

	err = s.DB.DeleteRealtor(ctx, realtor.RealtorID)
	if err != nil {
		s.l.Error("[DB] Error deleting realtor by ID", "error", err)
		return nil, err
	}
	return &listingsPB.DeleteRealtorResponse{Status: listingsPB.STATUS_SUCCESS, Realtor: RealtorDBtoPB(realtor)}, nil
}

func RealtorDBtoPB(realtor genDB.Realtor) *listingsPB.Realtor {
	return &listingsPB.Realtor{
		Id:          realtor.RealtorID,
		UserRef:     realtor.UserID,
		Name:        realtor.Name,
		Email:       realtor.Email.String,
		PhoneNumber: realtor.PhoneNumber.String,
		Company:     realtor.Company.String,
	}
}
