package database

import (
	userPB "github.com/jalexanderII/zero_microservice/gen/users"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role int32

const (
	Undefined Role = iota
	Admin
	Renter
	Realtor
	Owner
)

func (r Role) String() string {
	switch r {
	case Admin:
		return "admin"
	case Renter:
		return "renter"
	case Realtor:
		return "realtor"
	case Owner:
		return "owner"
	}
	return "unknown"
}

// NilUser is the nil value of our user
var NilUser User

// User is the default user struct
type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	Role     Role               `bson:"role"`
	MetaData userPB.Metadata    `bson:"metadata"`
}
