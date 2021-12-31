package server

import (
	"context"
	"database/sql"

	"github.com/jalexanderII/zero_microservice/backend/services/listings/database/genDB"
	listingsPB "github.com/jalexanderII/zero_microservice/gen/listings"
)

func (s listingsServer) CreateOwner(ctx context.Context, in *listingsPB.CreateOwnerRequest) (*listingsPB.Owner, error) {
	s.l.Debug("CreateOwner")
	var ownerpb = in.Owner
	owner, err := s.DB.CreateOwner(ctx, genDB.CreateOwnerParams{
		Name:        ownerpb.Name,
		UserID:      ownerpb.UserRef,
		Email:       sql.NullString{String: ownerpb.Email, Valid: true},
		PhoneNumber: sql.NullString{String: ownerpb.PhoneNumber, Valid: true},
		Company:     sql.NullString{String: ownerpb.Company, Valid: true},
	})
	if err != nil {
		s.l.Error("[DB] Error creating owner", "error", err)
		return nil, err
	}
	return OwnerDBtoPB(owner), nil
}

func (s listingsServer) GetOwner(ctx context.Context, in *listingsPB.GetOwnerRequest) (*listingsPB.Owner, error) {
	s.l.Debug("GetOwner")
	owner, err := s.DB.GetOwner(ctx, in.Id)
	if err != nil {
		s.l.Error("[DB] Error getting owner by id", "error", err)
		return nil, err
	}
	return OwnerDBtoPB(owner), nil
}

func (s listingsServer) ListOwners(ctx context.Context, in *listingsPB.ListOwnerRequest) (*listingsPB.ListOwnerResponse, error) {
	s.l.Debug("ListOwners")
	owners, err := s.DB.ListOwners(ctx)
	if err != nil {
		s.l.Error("[DB] Error getting all owners", "error", err)
		return nil, err
	}
	res := make([]*listingsPB.Owner, len(owners))
	for idx, owner := range owners {
		res[idx] = OwnerDBtoPB(owner)
	}
	return &listingsPB.ListOwnerResponse{Owners: res}, nil
}

func (s listingsServer) UpdateOwner(ctx context.Context, in *listingsPB.UpdateOwnerRequest) (*listingsPB.Owner, error) {
	s.l.Debug("UpdateOwner")
	var ownerpb = in.Owner
	err := s.DB.UpdateOwner(ctx, genDB.UpdateOwnerParams{
		OwnerID:     ownerpb.Id,
		Name:        ownerpb.Name,
		Email:       sql.NullString{String: ownerpb.Email, Valid: true},
		PhoneNumber: sql.NullString{String: ownerpb.PhoneNumber, Valid: true},
		Company:     sql.NullString{String: ownerpb.Company, Valid: true},
	})
	if err != nil {
		s.l.Error("[DB] Error updating owner", "error", err)
		return nil, err
	}
	return ownerpb, nil
}

func (s listingsServer) DeleteOwner(ctx context.Context, in *listingsPB.DeleteOwnerRequest) (*listingsPB.DeleteOwnerResponse, error) {
	s.l.Debug("DeleteOwner")
	owner, err := s.DB.GetOwner(ctx, in.Id)
	if err != nil {
		s.l.Error("[DB] Error fetching owner to delete", "error", err)
		return nil, err
	}

	err = s.DB.DeleteOwner(ctx, owner.OwnerID)
	if err != nil {
		s.l.Error("[DB] Error deleting owner by ID", "error", err)
		return nil, err
	}
	return &listingsPB.DeleteOwnerResponse{Status: listingsPB.STATUS_SUCCESS, Owner: OwnerDBtoPB(owner)}, nil
}

func OwnerDBtoPB(owner genDB.Owner) *listingsPB.Owner {
	return &listingsPB.Owner{
		Id:          owner.OwnerID,
		UserRef:     owner.UserID,
		Name:        owner.Name,
		Email:       owner.Email.String,
		PhoneNumber: owner.PhoneNumber.String,
		Company:     owner.Company.String,
	}
}
