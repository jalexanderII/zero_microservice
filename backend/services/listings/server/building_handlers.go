package server

import (
	"context"
	"database/sql"

	"github.com/jalexanderII/zero_microservice/backend/services/listings/database/genDB"
	"github.com/jalexanderII/zero_microservice/backend/services/listings/external_apis/geocensus"
	"github.com/jalexanderII/zero_microservice/backend/services/listings/utils"
	listingsPB "github.com/jalexanderII/zero_microservice/gen/listings"
)

func (s listingsServer) CreateBuilding(ctx context.Context, in *listingsPB.CreateBuildingRequest) (*listingsPB.Building, error) {
	s.l.Debug("CreateBuilding")
	var buildingpb = in.Building

	coords, err := geocensus.GetGeoCodeZip(
		buildingpb.Street, buildingpb.City, buildingpb.State, utils.FastStringConv(buildingpb.ZipCode), false,
	)
	if err != nil {
		s.l.Error("Could not fetch coordinates", "error", err)
	}

	building, err := s.DB.CreateBuilding(ctx, genDB.CreateBuildingParams{
		Name:         buildingpb.Name,
		FullAddress:  buildingpb.FullAddress,
		Street:       buildingpb.Street,
		City:         buildingpb.City,
		State:        buildingpb.State,
		ZipCode:      buildingpb.ZipCode,
		Neighborhood: buildingpb.Neighborhood,
		Lat:          coords.X,
		Lng:          coords.Y,
		Description:  sql.NullString{String: buildingpb.Description, Valid: true},
		Amenities:    buildingpb.Amenities,
		UploadIds:    buildingpb.UploadIds,
		OwnerID:      buildingpb.OwnerRef,
	})
	if err != nil {
		s.l.Error("[DB] Error creating building", "error", err)
		return nil, err
	}
	return BuildingDBtoPB(building), nil
}

func (s listingsServer) GetBuilding(ctx context.Context, in *listingsPB.GetBuildingRequest) (*listingsPB.Building, error) {
	s.l.Debug("GetBuilding")
	building, err := s.DB.GetBuilding(ctx, in.Id)
	if err != nil {
		s.l.Error("[DB] Error getting building by id", "error", err)
		return nil, err
	}
	return BuildingDBtoPB(building), nil
}

func (s listingsServer) ListBuildings(ctx context.Context, in *listingsPB.ListBuildingRequest) (*listingsPB.ListBuildingResponse, error) {
	s.l.Debug("ListBuildings")
	buildings, err := s.DB.ListBuildings(ctx)
	if err != nil {
		s.l.Error("[DB] Error getting all buildings", "error", err)
		return nil, err
	}
	res := make([]*listingsPB.Building, len(buildings))
	for idx, building := range buildings {
		res[idx] = BuildingDBtoPB(building)
	}
	return &listingsPB.ListBuildingResponse{Buildings: res}, nil
}

func (s listingsServer) UpdateBuilding(ctx context.Context, in *listingsPB.UpdateBuildingRequest) (*listingsPB.Building, error) {
	s.l.Debug("UpdateBuilding")
	var buildingpb = in.Building
	err := s.DB.UpdateBuilding(ctx, genDB.UpdateBuildingParams{
		BuildingID:   buildingpb.Id,
		Name:         buildingpb.Name,
		FullAddress:  buildingpb.FullAddress,
		Street:       buildingpb.Street,
		City:         buildingpb.City,
		State:        buildingpb.State,
		ZipCode:      buildingpb.ZipCode,
		Neighborhood: buildingpb.Neighborhood,
		Lat:          buildingpb.Lat,
		Lng:          buildingpb.Lng,
		Description:  sql.NullString{String: buildingpb.Description, Valid: true},
		Amenities:    buildingpb.Amenities,
		UploadIds:    buildingpb.UploadIds,
		OwnerID:      buildingpb.OwnerRef,
	})
	if err != nil {
		s.l.Error("[DB] Error updating building", "error", err)
		return nil, err
	}
	return buildingpb, nil
}

func (s listingsServer) DeleteBuilding(ctx context.Context, in *listingsPB.DeleteBuildingRequest) (*listingsPB.DeleteBuildingResponse, error) {
	s.l.Debug("DeleteBuilding")
	building, err := s.DB.GetBuilding(ctx, in.Id)
	if err != nil {
		s.l.Error("[DB] Error fetching building to delete", "error", err)
		return nil, err
	}

	err = s.DB.DeleteBuilding(ctx, building.BuildingID)
	if err != nil {
		s.l.Error("[DB] Error deleting building by ID", "error", err)
		return nil, err
	}
	return &listingsPB.DeleteBuildingResponse{Status: listingsPB.STATUS_SUCCESS, Building: BuildingDBtoPB(building)}, nil
}

func BuildingDBtoPB(building genDB.Building) *listingsPB.Building {
	return &listingsPB.Building{
		Id:           building.BuildingID,
		Name:         building.Name,
		FullAddress:  building.FullAddress,
		Street:       building.Street,
		City:         building.City,
		State:        building.State,
		ZipCode:      building.ZipCode,
		Neighborhood: building.Neighborhood,
		Lat:          building.Lat,
		Lng:          building.Lng,
		Description:  building.Description.String,
		Amenities:    building.Amenities,
		UploadIds:    building.UploadIds,
		OwnerRef:     building.OwnerID,
	}
}
