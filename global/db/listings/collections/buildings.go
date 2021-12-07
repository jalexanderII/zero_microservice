package collections

import (
	listingDB "github.com/jalexanderII/zero_microservice/global/db/listings"
	"github.com/jalexanderII/zero_microservice/global/db/listings/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var BuildingCollection = *listingDB.DB.Collection("buildings")

type Building struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name" `
	Address     models.Place       `bson:"address,omitempty"`
	Description string             `bson:"description,omitempty"`
	Amenities   []string           `bson:"amenities,omitempty"`
	Uploads     []models.Content   `bson:"uploads,omitempty"`
	RealtorRef  primitive.ObjectID `bson:"realtor_id"`
	Realtor     Realtor            `bson:"realtor,omitempty"`
}
