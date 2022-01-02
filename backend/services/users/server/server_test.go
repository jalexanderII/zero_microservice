package server

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/go-hclog"
	userDB "github.com/jalexanderII/zero_microservice/backend/services/users/database"
	"github.com/jalexanderII/zero_microservice/config"
	"github.com/jalexanderII/zero_microservice/config/middleware"
	listingsPB "github.com/jalexanderII/zero_microservice/gen/listings"
	userPB "github.com/jalexanderII/zero_microservice/gen/users"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

var L = hclog.Default()

func MockListingsClient() listingsPB.ListingsClient {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", config.LISTINGSERVICESERVERPORT), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return listingsPB.NewListingsClient(conn)
}

func TestAuthServer_GetUser(t *testing.T) {
	ctx, cancel := config.NewDBContext(5 * time.Second)
	defer cancel()

	DB := userDB.InitiateMongoClient()
	userCollection := *DB.Collection(config.USERCOLLECTIONNAME)
	jwtManager := middleware.NewJWTManager(config.JWTSecret, config.TokenDuration)
	server := NewAuthServer(userCollection, jwtManager, MockListingsClient(), L)

	user, err := server.GetUser(ctx, &userPB.GetUserRequest{Id: "61ccf8eaec2ee3f28ad48932"})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if user.Username != "guest" {
		t.Errorf("2: Failed to fetch correct user: %+v", user)
	}
}

func TestAuthServer_ListUsers(t *testing.T) {
	ctx, cancel := config.NewDBContext(5 * time.Second)
	defer cancel()

	DB := userDB.InitiateMongoClient()
	userCollection := *DB.Collection(config.USERCOLLECTIONNAME)
	jwtManager := middleware.NewJWTManager(config.JWTSecret, config.TokenDuration)
	server := NewAuthServer(userCollection, jwtManager, MockListingsClient(), L)

	users, err := server.ListUsers(ctx, &userPB.ListUserRequest{})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if len(users.Users) < 1 {
		t.Errorf("2: Failed to fetch realtors: %+v", users.Users[0])
	}
}

func TestAuthServer_UpdateUser(t *testing.T) {
	ctx, cancel := config.NewDBContext(5 * time.Second)
	defer cancel()

	DB := userDB.InitiateMongoClient()
	userCollection := *DB.Collection(config.USERCOLLECTIONNAME)
	jwtManager := middleware.NewJWTManager(config.JWTSecret, config.TokenDuration)
	server := NewAuthServer(userCollection, jwtManager, MockListingsClient(), L)

	u := &userPB.User{
		Username: "Carl",
		Email:    "updated@gmail.com",
		Password: "carlupdated",
	}

	user, err := server.UpdateUser(ctx, &userPB.UpdateUserRequest{Id: "000000000000000000000000", User: u})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if user.Email != u.Email {
		t.Errorf("2: Failed to fetch correct user: %+v", user)
	}
}

func TestAuthServer_DeleteUser(t *testing.T) {
	ctx, cancel := config.NewDBContext(5 * time.Second)
	defer cancel()

	DB := userDB.InitiateMongoClient()
	userCollection := *DB.Collection(config.USERCOLLECTIONNAME)
	jwtManager := middleware.NewJWTManager(config.JWTSecret, config.TokenDuration)
	server := NewAuthServer(userCollection, jwtManager, MockListingsClient(), L)

	u, err := server.ListUsers(ctx, &userPB.ListUserRequest{})
	originalLen := len(u.GetUsers())

	pw, _ := bcrypt.GenerateFromPassword([]byte("to_delete"), bcrypt.DefaultCost)
	newUser := userDB.User{ID: primitive.NewObjectID(), Email: "to_delete", Username: "to_delete", Password: string(pw), Role: userDB.Renter}
	_, err = server.DB.InsertOne(ctx, &newUser)
	if err != nil {
		t.Errorf("1: Error creating new user:: %v", err)
	}

	users, err := server.ListUsers(ctx, &userPB.ListUserRequest{})
	if err != nil {
		t.Errorf("2: An error was returned: %v", err)
	}
	newLen := len(users.GetUsers())
	if newLen != originalLen+1 {
		t.Errorf("3: An error adding a temp user, number of users in DB: %v", newLen)
	}

	deleted, err := server.DeleteUser(ctx, &userPB.DeleteUserRequest{Id: newUser.ID.Hex()})
	if err != nil {
		t.Errorf("4: An error was returned: %v", err)
	}
	if deleted.Status != userPB.STATUS_SUCCESS {
		t.Errorf("5: Failed to delete user: %+v\n, %+v", deleted.Status, deleted.GetUser())
	}
}

func Test_authServer_Login(t *testing.T) {
	ctx, cancel := config.NewDBContext(5 * time.Second)
	defer cancel()

	DB := userDB.InitiateMongoClient()
	userCollection := *DB.Collection(config.USERCOLLECTIONNAME)
	jwtManager := middleware.NewJWTManager(config.JWTSecret, config.TokenDuration)
	server := NewAuthServer(userCollection, jwtManager, MockListingsClient(), L)

	pw, _ := bcrypt.GenerateFromPassword([]byte("guest2"), bcrypt.DefaultCost)
	_, err := server.DB.InsertOne(context.Background(), userDB.User{ID: primitive.NewObjectID(), Username: "guest2", Email: "guest2@gmail.com", Password: string(pw), Role: 3})
	if err != nil {
		t.Errorf("1: Error inserting new user into db: %v", err)
	}

	_, err = server.Login(ctx, &userPB.LoginRequest{Username: "guest2", Password: "guest2"})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}

	_, err = server.Login(ctx, &userPB.LoginRequest{Username: "wrong", Password: "wrong"})
	if err == nil {
		t.Error("2: Error was nil")
	}

	_, err = server.Login(ctx, &userPB.LoginRequest{Username: "guest2@gmail.com", Password: "guest2"})
	if err != nil {
		t.Errorf("3: An error was returned: %v", err)
	}
}

func Test_authServer_SignUp(t *testing.T) {
	DB := userDB.InitiateMongoClient()
	userCollection := *DB.Collection(config.USERCOLLECTIONNAME)
	jwtManager := middleware.NewJWTManager(config.JWTSecret, config.TokenDuration)
	server := NewAuthServer(userCollection, jwtManager, MockListingsClient(), L)

	u := &userPB.SignupRequest{
		Username: "example2",
		Email:    "example2@gmail.com",
		Password: "password2",
		Role:     int32(userDB.Admin),
		Metadata: &userPB.Metadata{UserType: &userPB.Metadata_AdminMetadata{}},
	}

	_, err := server.SignUp(context.Background(), &userPB.SignupRequest{Username: "Carl", Email: "example@gmail.com", Password: "examplestring"})
	if err.Error() != "username already taken" {
		t.Error("1: No or wrong error returned for username already taken")
	}

	_, err = server.SignUp(context.Background(), &userPB.SignupRequest{Username: "Carlo", Email: "test@gmail.com", Password: "examplestring"})
	if err.Error() != "email already used" {
		t.Error("2: No or wrong error returned for email already taken")
	}

	_, err = server.SignUp(context.Background(), u)
	if err != nil {
		t.Errorf("3: Error creating new user: %v", err)
	}

	_, err = server.SignUp(context.Background(), &userPB.SignupRequest{Username: "Example", Email: "e", Password: "pass"})
	if err.Error() != "email validation failed" {
		t.Error("4: No or wrong error returned for email validation")
	}
}

func TestAuthServer_createUserFromMetadata(t *testing.T) {
	DB := userDB.InitiateMongoClient()
	userCollection := *DB.Collection(config.USERCOLLECTIONNAME)
	jwtManager := middleware.NewJWTManager(config.JWTSecret, config.TokenDuration)
	server := NewAuthServer(userCollection, jwtManager, MockListingsClient(), L)

	u := &userPB.SignupRequest{
		Username: "example3",
		Email:    "example3@gmail.com",
		Password: "password3",
		Role:     int32(userDB.Owner),
		Metadata: &userPB.Metadata{
			UserType: &userPB.Metadata_OwnerMetadata{
				OwnerMetadata: &userPB.OwnerMetadata{
					Name:        "from_signup3",
					Email:       "from_signup3@gmail.com",
					PhoneNumber: "325699844",
					Company:     "from_signup3",
				},
			},
		},
	}

	newId, err := server.createUserFromMetadata(context.Background(), u.Metadata, primitive.NewObjectID().Hex())
	if err != nil {
		t.Errorf("1: Error creating new user: %v", err)
	}
	owner, err := server.ListingsClient.GetOwner(context.Background(), &listingsPB.GetOwnerRequest{Id: newId})
	if err != nil {
		t.Errorf("2: Error creating new user: %v", err)
	}
	if owner.GetName() != u.GetMetadata().GetOwnerMetadata().GetName() {
		t.Errorf("3: Error owner has wrong metadata: %v", owner)
	}
}
