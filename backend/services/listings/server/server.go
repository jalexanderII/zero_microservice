package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jalexanderII/zero_microservice/gen/listings"
	listingDB "github.com/jalexanderII/zero_microservice/global/db/listings"
	"github.com/jalexanderII/zero_microservice/global/db/listings/collections"
	"github.com/jalexanderII/zero_microservice/global/db/listings/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ApartmentCollection mongo.Collection
	BuildingCollection  mongo.Collection
	RealtorCollection   mongo.Collection
)

type listingsServer struct{}

func NewListingsServer() *listingsServer {
	return &listingsServer{}
}

func (listingsServer) CreateApartment(ctx context.Context, in *listings.CreateApartmentRequest) (*listings.Apartment, error) {
	apartment := in.Apartment
	address := apartment.GetAddress()
	uploads, err := getContentFromPb(apartment.GetUploads())
	if err != nil {
		log.Printf("Error getting uploads source: %v\n", err)
		return nil, errors.New("something went wrong")
	}

	ctx, cancel := listingDB.NewDBContext(5 * time.Second)
	defer cancel()

	buildingRef, err := primitive.ObjectIDFromHex(apartment.GetBuildingRef())
	if err != nil {
		log.Printf("Error getting building ref: %v\n", err)
		return nil, errors.New("something went wrong")
	}
	realtorRef, err := primitive.ObjectIDFromHex(apartment.GetRealtorRef())
	if err != nil {
		log.Printf("Error getting realtor ref: %v\n", err)
		return nil, errors.New("something went wrong")
	}

	newApartment := collections.Apartment{
		ID:   primitive.NewObjectID(),
		Name: apartment.GetName(),
		Address: models.Place{
			FullAddress:  address.FullAddress,
			Street:       address.Street,
			City:         address.City,
			State:        address.State,
			ZipCode:      address.ZipCode,
			Neighborhood: address.Neighborhood,
			Unit:         *address.Unit,
			Coordinates:  models.Coordinates{Lat: address.Coordinates.Lat, Lng: address.Coordinates.Lng},
		},
		Rent:  apartment.GetRent(),
		Sqft:  apartment.GetSqft(),
		Beds:  apartment.GetBeds(),
		Baths: apartment.GetBaths(),
		ListingMetrics: models.ListingMetrics{
			AvailableOn:  apartment.GetListingMetrics().AvailableOn.AsTime(),
			DaysOnMarket: apartment.GetListingMetrics().DaysOnMarket,
		},
		Description: apartment.GetDescription(),
		Amenities:   apartment.GetAmenities(),
		Uploads:     uploads,
		IsArchived:  apartment.GetIsArchived(),
		BuildingRef: buildingRef,
		RealtorRef:  realtorRef,
	}
	res, err := ApartmentCollection.InsertOne(ctx, newApartment)
	if err != nil {
		log.Printf("Error inserting new user: %v\n", err)
		return nil, errors.New("Something went wrong")
	}
	fmt.Printf("inserted document with ID %v\n", res.InsertedID)

	var result listings.Apartment
	err = ApartmentCollection.FindOne(ctx, bson.M{"ID": res.InsertedID}).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("something went wrong: %v", err)
	}
	return &result, nil
}

func (listingsServer) GetApartment(context.Context, *listings.GetApartmentRequest) (*listings.Apartment, error) {
	return nil, nil
}

func (listingsServer) ListApartments(context.Context, *listings.ListApartmentRequest) (*listings.ListApartmentResponse, error) {
	return nil, nil
}

func (listingsServer) UpdateApartment(context.Context, *listings.UpdateApartmentRequest) (*listings.Apartment, error) {
	return nil, nil
}

func (listingsServer) DeleteApartment(context.Context, *listings.DeleteApartmentRequest) (*listings.DeleteApartmentResponse, error) {
	return nil, nil
}

func (listingsServer) CreateBuilding(context.Context, *listings.CreateBuildingRequest) (*listings.Building, error) {
	return nil, nil
}

func (listingsServer) GetBuilding(context.Context, *listings.GetBuildingRequest) (*listings.Building, error) {
	return nil, nil
}

func (listingsServer) ListBuildings(context.Context, *listings.ListBuildingRequest) (*listings.ListBuildingResponse, error) {
	return nil, nil
}

func (listingsServer) UpdateBuilding(context.Context, *listings.UpdateBuildingRequest) (*listings.Building, error) {
	return nil, nil
}

func (listingsServer) DeleteBuilding(context.Context, *listings.DeleteBuildingRequest) (*listings.DeleteBuildingResponse, error) {
	return nil, nil
}

func (listingsServer) CreateRealtor(context.Context, *listings.CreateRealtorRequest) (*listings.Realtor, error) {
	return nil, nil
}

func (listingsServer) GetRealtor(context.Context, *listings.GetRealtorRequest) (*listings.Realtor, error) {
	return nil, nil
}

func (listingsServer) ListRealtors(context.Context, *listings.ListRealtorRequest) (*listings.ListRealtorResponse, error) {
	return nil, nil
}

func (listingsServer) UpdateRealtor(context.Context, *listings.UpdateRealtorRequest) (*listings.Realtor, error) {
	return nil, nil
}

func (listingsServer) DeleteRealtor(context.Context, *listings.DeleteRealtorRequest) (*listings.DeleteRealtorResponse, error) {
	return nil, nil
}

func (listingsServer) UploadPhoto(listings.Listings_UploadPhotoServer) error {
	return nil
}

func (listingsServer) StreamPhotos(listings.Listings_StreamPhotosServer) error {
	return nil
}

func (listingsServer) DeletePhoto(context.Context, *listings.DeletePhotoRequest) (*listings.DeletePhotoResponse, error) {
	return nil, nil
}

func getContentFromPb(contentPb []*listings.Content) ([]models.Content, error) {
	contents := make([]models.Content, len(contentPb))
	for idx, res := range contentPb {
		s, err := getContentSource(res)
		if err != nil {
			return nil, err
		}
		contents[idx] = models.Content{
			Id:       res.GetId(),
			Filename: res.GetFilename(),
			FileId:   res.GetFileId(),
			Source:   s,
			Type:     res.GetType(),
		}
	}
	return contents, nil
}

func getContentSource(res *listings.Content) (models.IsContentSource, error) {
	switch s := res.GetSource().(type) {
	case *listings.Content_ApartmentRef:
		return &models.ContentApartmentRef{ApartmentRef: s.ApartmentRef}, nil
	case *listings.Content_BuildingRef:
		return &models.ContentBuildingRef{BuildingRef: s.BuildingRef}, nil
	default:
		return nil, fmt.Errorf("incorrect content source")
	}
}
