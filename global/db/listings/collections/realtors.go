package collections

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Realtor struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name" `
	Email       string             `bson:"email,omitempty" `
	PhoneNumber string             `bson:"phone_number,omitempty" `
	Company     string             `bson:"company,omitempty" `
}
