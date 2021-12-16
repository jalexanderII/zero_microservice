package server

import (
	"context"
	"database/sql"

	"github.com/jalexanderII/zero_microservice/backend/services/listings/database/genDB"
	"github.com/jalexanderII/zero_microservice/backend/services/listings/external_apis/geocensus"
	listingsPB "github.com/jalexanderII/zero_microservice/gen/listings"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s listingsServer) CreateApartment(ctx context.Context, in *listingsPB.CreateApartmentRequest) (*listingsPB.Apartment, error) {
	s.l.Debug("CreateApartment")
	var apartmentpb = in.Apartment

	coords, err := geocensus.GetGeoCodeZip(
		apartmentpb.Street, apartmentpb.City, apartmentpb.State, geocensus.FastStringConv(apartmentpb.ZipCode),
	)
	if err != nil {
		s.l.Error("Could not fetch coordinates", "error", err)
	}

	apartment, err := s.DB.CreateApartment(ctx, genDB.CreateApartmentParams{
		ApartmentID:  apartmentpb.Id,
		Name:         apartmentpb.Name,
		FullAddress:  apartmentpb.FullAddress,
		Street:       apartmentpb.Street,
		City:         apartmentpb.City,
		State:        apartmentpb.State,
		ZipCode:      apartmentpb.ZipCode,
		Neighborhood: apartmentpb.Neighborhood,
		Unit:         sql.NullString{String: apartmentpb.Unit, Valid: true},
		Lat:          int32(coords.X),
		Lng:          int32(coords.Y),
		Rent:         apartmentpb.Rent,
		Sqft:         sql.NullInt32{Int32: apartmentpb.Sqft, Valid: true},
		Beds:         apartmentpb.Beds,
		Baths:        apartmentpb.Baths,
		AvailableOn:  apartmentpb.AvailableOn.AsTime(),
		DaysOnMarket: sql.NullInt32{Int32: apartmentpb.DaysOnMarket, Valid: true},
		Description:  sql.NullString{String: apartmentpb.Description, Valid: true},
		Amenities:    apartmentpb.Amenities,
		UploadIds:    apartmentpb.UploadIds,
		IsArchived:   apartmentpb.IsArchived,
		BuildingID:   apartmentpb.BuildingRef,
		RealtorID:    apartmentpb.RealtorRef,
	})
	if err != nil {
		s.l.Error("[DB] Error creating apartment", "error", err)
		return nil, err
	}
	return ApartmentDBtoPB(apartment), nil
}

func (s listingsServer) GetApartment(ctx context.Context, in *listingsPB.GetApartmentRequest) (*listingsPB.Apartment, error) {
	s.l.Debug("GetApartment")
	apartment, err := s.DB.GetApartment(ctx, in.Id)
	if err != nil {
		s.l.Error("[DB] Error getting apartment by id", "error", err)
		return nil, err
	}
	return ApartmentDBtoPB(apartment), nil
}

func (s listingsServer) ListApartments(ctx context.Context, in *listingsPB.ListApartmentRequest) (*listingsPB.ListApartmentResponse, error) {
	s.l.Debug("ListApartments")
	apartments, err := s.DB.ListApartments(ctx)
	if err != nil {
		s.l.Error("[DB] Error getting all apartments", "error", err)
		return nil, err
	}
	res := make([]*listingsPB.Apartment, len(apartments))
	for idx, apartment := range apartments {
		res[idx] = ApartmentDBtoPB(apartment)
	}
	return &listingsPB.ListApartmentResponse{Apartments: res}, nil
}

func (s listingsServer) UpdateApartment(ctx context.Context, in *listingsPB.UpdateApartmentRequest) (*listingsPB.Apartment, error) {
	s.l.Debug("UpdateApartment")
	var apartmentpb = in.Apartment
	err := s.DB.UpdateApartment(ctx, genDB.UpdateApartmentParams{
		ApartmentID:  apartmentpb.Id,
		Name:         apartmentpb.Name,
		FullAddress:  apartmentpb.FullAddress,
		Street:       apartmentpb.Street,
		City:         apartmentpb.City,
		State:        apartmentpb.State,
		ZipCode:      apartmentpb.ZipCode,
		Neighborhood: apartmentpb.Neighborhood,
		Unit:         sql.NullString{String: apartmentpb.Unit, Valid: true},
		Lat:          int32(apartmentpb.Lat),
		Lng:          int32(apartmentpb.Lng),
		Rent:         apartmentpb.Rent,
		Sqft:         sql.NullInt32{Int32: apartmentpb.Sqft, Valid: true},
		Beds:         apartmentpb.Beds,
		Baths:        apartmentpb.Baths,
		AvailableOn:  apartmentpb.AvailableOn.AsTime(),
		DaysOnMarket: sql.NullInt32{Int32: apartmentpb.DaysOnMarket, Valid: true},
		Description:  sql.NullString{String: apartmentpb.Description, Valid: true},
		Amenities:    apartmentpb.Amenities,
		UploadIds:    apartmentpb.UploadIds,
		IsArchived:   apartmentpb.IsArchived,
		BuildingID:   apartmentpb.BuildingRef,
		RealtorID:    apartmentpb.RealtorRef,
	})
	if err != nil {
		s.l.Error("[DB] Error updating apartment", "error", err)
		return nil, err
	}
	return apartmentpb, nil
}

func (s listingsServer) DeleteApartment(ctx context.Context, in *listingsPB.DeleteApartmentRequest) (*listingsPB.DeleteApartmentResponse, error) {
	s.l.Debug("DeleteApartment")
	apartment, err := s.DB.GetApartment(ctx, in.Id)
	if err != nil {
		s.l.Error("[DB] Error fetching apartment to delete", "error", err)
		return nil, err
	}

	err = s.DB.DeleteApartment(ctx, in.Id)
	if err != nil {
		s.l.Error("[DB] Error deleting apartment by ID", "error", err)
		return nil, err
	}
	return &listingsPB.DeleteApartmentResponse{Status: listingsPB.STATUS_SUCCESS, Apartment: ApartmentDBtoPB(apartment)}, nil
}

func ApartmentDBtoPB(apartment genDB.Apartment) *listingsPB.Apartment {
	return &listingsPB.Apartment{
		Id:           apartment.ApartmentID,
		Name:         apartment.Name,
		FullAddress:  apartment.FullAddress,
		Street:       apartment.Street,
		City:         apartment.City,
		State:        apartment.State,
		ZipCode:      apartment.ZipCode,
		Neighborhood: apartment.Neighborhood,
		Unit:         apartment.Unit.String,
		Lat:          float64(apartment.Lat),
		Lng:          float64(apartment.Lng),
		Rent:         apartment.Rent,
		Sqft:         apartment.Sqft.Int32,
		Beds:         apartment.Beds,
		Baths:        apartment.Baths,
		AvailableOn:  timestamppb.New(apartment.AvailableOn),
		DaysOnMarket: apartment.DaysOnMarket.Int32,
		Description:  apartment.Description.String,
		Amenities:    apartment.Amenities,
		UploadIds:    apartment.UploadIds,
		IsArchived:   apartment.IsArchived,
		BuildingRef:  apartment.BuildingID,
		RealtorRef:   apartment.RealtorID,
	}
}
