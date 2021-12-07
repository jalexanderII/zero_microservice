package collections

import (
	"github.com/jalexanderII/zero_microservice/global/db/listings/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Apartment struct {
	ID             primitive.ObjectID    `bson:"_id"`
	Name           string                `bson:"name" `
	Address        models.Place          `bson:"address,omitempty"`
	Rent           int32                 `bson:"rent,omitempty"`
	Sqft           int32                 `bson:"sqft,omitempty"`
	Beds           int32                 `bson:"beds,omitempty"`
	Baths          int32                 `bson:"baths,omitempty"`
	ListingMetrics models.ListingMetrics `bson:"listing_metrics,omitempty"`
	Description    string                `bson:"description,omitempty"`
	Amenities      []string              `bson:"amenities,omitempty"`
	Uploads        []models.Content      `bson:"uploads,omitempty"`
	IsArchived     bool
	BuildingRef    primitive.ObjectID `bson:"building_id"`
	Building       Building           `bson:"building,omitempty"`
	RealtorRef     primitive.ObjectID `bson:"realtor_id"`
	Realtor        Realtor            `bson:"realtor,omitempty"`
}
