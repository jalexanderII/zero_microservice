package server

import (
	"bytes"
	"context"
	"database/sql"
	"io"
	"log"

	listingsDB "github.com/jalexanderII/zero_microservice/backend/services/listings/database"
	"github.com/jalexanderII/zero_microservice/backend/services/listings/database/genDB"
	contentStore "github.com/jalexanderII/zero_microservice/backend/services/listings/store"
	listingsPB "github.com/jalexanderII/zero_microservice/gen/listings"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const maxImageSize = 1 << 20

type listingsServer struct {
	listingsPB.UnimplementedListingsServer
	DB           *listingsDB.ListingsDB
	ContentStore contentStore.ContentStore
}

func NewListingsServer(db *listingsDB.ListingsDB, cs contentStore.ContentStore) *listingsServer {
	return &listingsServer{DB: db, ContentStore: cs}
}

func (s listingsServer) CreateApartment(ctx context.Context, in *listingsPB.CreateApartmentRequest) (*listingsPB.Apartment, error) {
	var apartmentpb = in.Apartment
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
	return ApartmentDBtoPB(apartment), nil
}

func (s listingsServer) GetApartment(ctx context.Context, in *listingsPB.GetApartmentRequest) (*listingsPB.Apartment, error) {
	apartment, err := s.DB.GetApartment(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return ApartmentDBtoPB(apartment), nil
}

func (s listingsServer) ListApartments(ctx context.Context, in *listingsPB.ListApartmentRequest) (*listingsPB.ListApartmentResponse, error) {
	apartments, err := s.DB.ListApartments(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]*listingsPB.Apartment, len(apartments))
	for idx, apartment := range apartments {
		res[idx] = ApartmentDBtoPB(apartment)
	}
	return &listingsPB.ListApartmentResponse{Apartments: res}, nil
}

func (s listingsServer) UpdateApartment(ctx context.Context, in *listingsPB.UpdateApartmentRequest) (*listingsPB.Apartment, error) {
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
		return nil, err
	}
	return apartmentpb, nil
}

func (s listingsServer) DeleteApartment(ctx context.Context, in *listingsPB.DeleteApartmentRequest) (*listingsPB.DeleteApartmentResponse, error) {
	apartment, err := s.DB.GetApartment(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	err = s.DB.DeleteApartment(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &listingsPB.DeleteApartmentResponse{Status: listingsPB.STATUS_SUCCESS, Apartment: ApartmentDBtoPB(apartment)}, nil
}

func (s listingsServer) CreateBuilding(ctx context.Context, in *listingsPB.CreateBuildingRequest) (*listingsPB.Building, error) {
	var buildingpb = in.Building
	building, err := s.DB.CreateBuilding(ctx, genDB.CreateBuildingParams{
		BuildingID:   buildingpb.Id,
		Name:         buildingpb.Name,
		FullAddress:  buildingpb.FullAddress,
		Street:       buildingpb.Street,
		City:         buildingpb.City,
		State:        buildingpb.State,
		ZipCode:      buildingpb.ZipCode,
		Neighborhood: buildingpb.Neighborhood,
		Lat:          int32(buildingpb.Lat),
		Lng:          int32(buildingpb.Lng),
		Description:  sql.NullString{String: buildingpb.Description, Valid: true},
		Amenities:    buildingpb.Amenities,
		UploadIds:    buildingpb.UploadIds,
		RealtorID:    buildingpb.RealtorRef,
	})
	if err != nil {
		return nil, err
	}
	return BuildingDBtoPB(building), nil
}

func (s listingsServer) GetBuilding(ctx context.Context, in *listingsPB.GetBuildingRequest) (*listingsPB.Building, error) {
	building, err := s.DB.GetBuilding(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return BuildingDBtoPB(building), nil
}

func (s listingsServer) ListBuildings(ctx context.Context, in *listingsPB.ListBuildingRequest) (*listingsPB.ListBuildingResponse, error) {
	buildings, err := s.DB.ListBuildings(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]*listingsPB.Building, len(buildings))
	for idx, building := range buildings {
		res[idx] = BuildingDBtoPB(building)
	}
	return &listingsPB.ListBuildingResponse{Buildings: res}, nil
}

func (s listingsServer) UpdateBuilding(ctx context.Context, in *listingsPB.UpdateBuildingRequest) (*listingsPB.Building, error) {
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
		Lat:          int32(buildingpb.Lat),
		Lng:          int32(buildingpb.Lng),
		Description:  sql.NullString{String: buildingpb.Description, Valid: true},
		Amenities:    buildingpb.Amenities,
		UploadIds:    buildingpb.UploadIds,
		RealtorID:    buildingpb.RealtorRef,
	})
	if err != nil {
		return nil, err
	}
	return buildingpb, nil
}

func (s listingsServer) DeleteBuilding(ctx context.Context, in *listingsPB.DeleteBuildingRequest) (*listingsPB.DeleteBuildingResponse, error) {
	building, err := s.DB.GetBuilding(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	err = s.DB.DeleteBuilding(ctx, building.BuildingID)
	if err != nil {
		return nil, err
	}
	return &listingsPB.DeleteBuildingResponse{Status: listingsPB.STATUS_SUCCESS, Building: BuildingDBtoPB(building)}, nil
}

func (s listingsServer) CreateRealtor(ctx context.Context, in *listingsPB.CreateRealtorRequest) (*listingsPB.Realtor, error) {
	var realtorpb = in.Realtor
	realtor, err := s.DB.CreateRealtor(ctx, genDB.CreateRealtorParams{
		RealtorID:   realtorpb.Id,
		Name:        realtorpb.Name,
		Email:       sql.NullString{String: realtorpb.Email, Valid: true},
		PhoneNumber: sql.NullString{String: realtorpb.PhoneNumber, Valid: true},
		Company:     sql.NullString{String: realtorpb.Company, Valid: true},
	})
	if err != nil {
		return nil, err
	}
	return RealtorDBtoPB(realtor), nil
}

func (s listingsServer) GetRealtor(ctx context.Context, in *listingsPB.GetRealtorRequest) (*listingsPB.Realtor, error) {
	realtor, err := s.DB.GetRealtor(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return RealtorDBtoPB(realtor), nil
}

func (s listingsServer) ListRealtors(ctx context.Context, in *listingsPB.ListRealtorRequest) (*listingsPB.ListRealtorResponse, error) {
	realtors, err := s.DB.ListRealtors(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]*listingsPB.Realtor, len(realtors))
	for idx, realtor := range realtors {
		res[idx] = RealtorDBtoPB(realtor)
	}
	return &listingsPB.ListRealtorResponse{Realtors: res}, nil
}

func (s listingsServer) UpdateRealtor(ctx context.Context, in *listingsPB.UpdateRealtorRequest) (*listingsPB.Realtor, error) {
	var realtorpb = in.Realtor
	err := s.DB.UpdateRealtor(ctx, genDB.UpdateRealtorParams{
		RealtorID:   realtorpb.Id,
		Name:        realtorpb.Name,
		Email:       sql.NullString{String: realtorpb.Email, Valid: true},
		PhoneNumber: sql.NullString{String: realtorpb.PhoneNumber, Valid: true},
		Company:     sql.NullString{String: realtorpb.Company, Valid: true},
	})
	if err != nil {
		return nil, err
	}
	return realtorpb, nil
}

func (s listingsServer) DeleteRealtor(ctx context.Context, in *listingsPB.DeleteRealtorRequest) (*listingsPB.DeleteRealtorResponse, error) {
	realtor, err := s.DB.GetRealtor(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	err = s.DB.DeleteRealtor(ctx, realtor.RealtorID)
	if err != nil {
		return nil, err
	}
	return &listingsPB.DeleteRealtorResponse{Status: listingsPB.STATUS_SUCCESS, Realtor: RealtorDBtoPB(realtor)}, nil
}

// UploadContent is a streaming RPC to upload content
func (s listingsServer) UploadContent(stream listingsPB.Listings_UploadContentServer) error {
	req, err := stream.Recv()
	if err != nil {
		return status.Errorf(codes.Unknown, "cannot receive image info")
	}

	contentType := req.GetInfo().GetContentType()
	sourceID := req.GetInfo().GetSourceId()
	sourceType := req.GetInfo().GetContentSource()
	log.Printf("receive an upload request for content %v with type %v", sourceID, contentType)

	switch sourceType {
	case listingsPB.ContentInfo_APARTMENT:
		apartment, err := s.DB.GetApartment(stream.Context(), sourceID)
		if err != nil {
			return status.Errorf(codes.Internal, "cannot find apartment: %v", err)
		}
		if apartment.Name == "" {
			return status.Errorf(codes.InvalidArgument, "apartment id %v doesn't exist", sourceID)
		}
	case listingsPB.ContentInfo_BUILDING:
		building, err := s.DB.GetBuilding(stream.Context(), sourceID)
		if err != nil {
			return status.Errorf(codes.Internal, "cannot find building: %v", err)
		}
		if building.Name == "" {
			return status.Errorf(codes.InvalidArgument, "building id %v doesn't exist", sourceID)
		}
	default:
		return status.Errorf(codes.Internal, "incorrect source type or other error: %v", sourceType)
	}

	contentData := bytes.Buffer{}
	contentSize := 0

	for {
		err := contextError(stream.Context())
		if err != nil {
			return err
		}

		log.Print("waiting to receive more data")

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			return status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err)
		}

		chunk := req.GetChunkData()
		size := len(chunk)

		log.Printf("received a chunk with size: %d", size)

		contentSize += size
		if contentSize > maxImageSize {
			return status.Errorf(codes.InvalidArgument, "image is too large: %d > %d", contentSize, maxImageSize)
		}

		_, err = contentData.Write(chunk)
		if err != nil {
			return status.Errorf(codes.Internal, "cannot write chunk data: %v", err)
		}
	}

	contentID, contentInfo, err := s.ContentStore.Save(sourceID, contentType, contentData)
	if err != nil {
		return status.Errorf(codes.Internal, "cannot save content to the store: %v", err)
	}

	_, err = s.DB.UploadContent(stream.Context(), genDB.UploadContentParams{
		ContentID:     contentID,
		Filename:      sql.NullString{String: contentInfo.Path, Valid: true},
		ContentType:   genDB.ContentType(contentInfo.ContentType),
		ContentSource: genDB.ContentSource(sourceType),
		SourceID:      sourceID,
	})
	if err != nil {
		return err
	}

	switch sourceType {
	case listingsPB.ContentInfo_APARTMENT:
		err = s.DB.AppendContentApartment(stream.Context(), genDB.AppendContentApartmentParams{
			ApartmentID: sourceID,
			ArrayAppend: contentInfo.Path,
		})
		if err != nil {
			return status.Errorf(codes.Internal, "error appending content to apartment: %v", err)
		}
	case listingsPB.ContentInfo_BUILDING:
		err = s.DB.AppendContentBuilding(stream.Context(), genDB.AppendContentBuildingParams{
			BuildingID:  sourceID,
			ArrayAppend: contentInfo.Path,
		})
		if err != nil {
			return status.Errorf(codes.Internal, "error appending content to building: %v", err)
		}
	}

	res := &listingsPB.UploadContentResponse{
		ContentId: contentID,
		Size:      uint32(contentSize),
	}

	err = stream.SendAndClose(res)
	if err != nil {
		return status.Errorf(codes.Unknown, "cannot send response: %v", err)
	}

	log.Printf("saved content with id: %v, size: %d", contentID, contentSize)
	return nil
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
		Lat:          float64(building.Lat),
		Lng:          float64(building.Lng),
		Description:  building.Description.String,
		Amenities:    building.Amenities,
		UploadIds:    building.UploadIds,
		RealtorRef:   building.RealtorID,
	}
}

func RealtorDBtoPB(realtor genDB.Realtor) *listingsPB.Realtor {
	return &listingsPB.Realtor{
		Id:          realtor.RealtorID,
		Name:        realtor.Name,
		Email:       realtor.Email.String,
		PhoneNumber: realtor.PhoneNumber.String,
		Company:     realtor.Company.String,
	}
}

func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return status.Error(codes.Canceled, "request is canceled")
	case context.DeadlineExceeded:
		return status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	default:
		return nil
	}
}
