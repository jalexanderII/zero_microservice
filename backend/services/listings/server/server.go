package server

import (
	"context"
	"database/sql"

	listingsDB "github.com/jalexanderII/zero_microservice/backend/services/listings/database"
	listingsPB "github.com/jalexanderII/zero_microservice/gen/listings"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type listingsServer struct {
	DB *listingsDB.ListingsDB
}

func NewListingsServer(db *listingsDB.ListingsDB) *listingsServer {
	return &listingsServer{db}
}

func (s listingsServer) CreateApartment(ctx context.Context, in *listingsPB.CreateApartmentRequest) (*listingsPB.Apartment, error) {
	var apartmentpb = in.Apartment
	apartment, err := s.DB.CreateApartment(ctx, listingsDB.CreateApartmentParams{
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
		return nil, err
	}
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
		Lng:          float64(int32(apartment.Lng)),
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
	}, nil
}

func (listingsServer) GetApartment(context.Context, *listingsPB.GetApartmentRequest) (*listingsPB.Apartment, error) {
	return nil, nil
}

func (listingsServer) ListApartments(context.Context, *listingsPB.ListApartmentRequest) (*listingsPB.ListApartmentResponse, error) {
	return nil, nil
}

func (listingsServer) UpdateApartment(context.Context, *listingsPB.UpdateApartmentRequest) (*listingsPB.Apartment, error) {
	return nil, nil
}

func (listingsServer) DeleteApartment(context.Context, *listingsPB.DeleteApartmentRequest) (*listingsPB.DeleteApartmentResponse, error) {
	return nil, nil
}

func (listingsServer) CreateBuilding(context.Context, *listingsPB.CreateBuildingRequest) (*listingsPB.Building, error) {
	return nil, nil
}

func (listingsServer) GetBuilding(context.Context, *listingsPB.GetBuildingRequest) (*listingsPB.Building, error) {
	return nil, nil
}

func (listingsServer) ListBuildings(context.Context, *listingsPB.ListBuildingRequest) (*listingsPB.ListBuildingResponse, error) {
	return nil, nil
}

func (listingsServer) UpdateBuilding(context.Context, *listingsPB.UpdateBuildingRequest) (*listingsPB.Building, error) {
	return nil, nil
}

func (listingsServer) DeleteBuilding(context.Context, *listingsPB.DeleteBuildingRequest) (*listingsPB.DeleteBuildingResponse, error) {
	return nil, nil
}

func (listingsServer) CreateRealtor(context.Context, *listingsPB.CreateRealtorRequest) (*listingsPB.Realtor, error) {
	return nil, nil
}

func (listingsServer) GetRealtor(context.Context, *listingsPB.GetRealtorRequest) (*listingsPB.Realtor, error) {
	return nil, nil
}

func (listingsServer) ListRealtors(context.Context, *listingsPB.ListRealtorRequest) (*listingsPB.ListRealtorResponse, error) {
	return nil, nil
}

func (listingsServer) UpdateRealtor(context.Context, *listingsPB.UpdateRealtorRequest) (*listingsPB.Realtor, error) {
	return nil, nil
}

func (listingsServer) DeleteRealtor(context.Context, *listingsPB.DeleteRealtorRequest) (*listingsPB.DeleteRealtorResponse, error) {
	return nil, nil
}

func (listingsServer) UploadPhoto(listingsPB.Listings_UploadPhotoServer) error {
	return nil
}

func (listingsServer) StreamPhotos(listingsPB.Listings_StreamPhotosServer) error {
	return nil
}

func (listingsServer) DeletePhoto(context.Context, *listingsPB.DeletePhotoRequest) (*listingsPB.DeletePhotoResponse, error) {
	return nil, nil
}
