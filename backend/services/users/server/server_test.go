package server

import (
	"context"
	"testing"

	config "github.com/jalexanderII/zero_microservice"
	database2 "github.com/jalexanderII/zero_microservice/backend/services/users/database"
	userPB "github.com/jalexanderII/zero_microservice/gen/users"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func Test_authServer_Login(t *testing.T) {
	db := database2.ConnectToTestDB()
	userCollection := *db.Collection(config.USERCOLLECTIONNAME)
	server := authServer{userCollection: userCollection}

	pw, _ := bcrypt.GenerateFromPassword([]byte("example"), bcrypt.DefaultCost)
	_, err := db.Collection("user").InsertOne(context.Background(), database2.User{ID: primitive.NewObjectID(), Email: "test@gmail.com", Username: "Carl", Password: string(pw)})
	if err != nil {
		t.Errorf("Something went wrong: %v", err)
	}

	_, err = server.Login(context.Background(), &userPB.LoginRequest{Login: "test@gmail.com", Password: "example"})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}

	_, err = server.Login(context.Background(), &userPB.LoginRequest{Login: "wrong", Password: "wrong"})
	if err == nil {
		t.Error("2: Error was nil")
	}

	_, err = server.Login(context.Background(), &userPB.LoginRequest{Login: "Carl", Password: "example"})
	if err != nil {
		t.Errorf("3: An error was returned: %v", err)
	}
}

func Test_authServer_UsernameUsed(t *testing.T) {
	db := database2.ConnectToTestDB()
	userCollection := *db.Collection(config.USERCOLLECTIONNAME)
	server := authServer{userCollection: userCollection}

	_, err := db.Collection("user").InsertOne(context.Background(), database2.User{Username: "Carl"})
	if err != nil {
		t.Errorf("Something went wrong: %v", err)
	}

	res, err := server.UsernameUsed(context.Background(), &userPB.UsernameUsedRequest{Username: "Carlo"})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if res.GetUsed() {
		t.Errorf("1: Wrong result")
	}
	res, err = server.UsernameUsed(context.Background(), &userPB.UsernameUsedRequest{Username: "Carl"})
	if err != nil {
		t.Errorf("2: An error was returned: %v", err)
	}
	if !res.GetUsed() {
		t.Errorf("2: Wrong result")
	}
}

func Test_authServer_EmailUsed(t *testing.T) {
	db := database2.ConnectToTestDB()
	userCollection := *db.Collection(config.USERCOLLECTIONNAME)
	server := authServer{userCollection: userCollection}

	_, err := db.Collection("user").InsertOne(context.Background(), database2.User{Email: "test@gmail.com"})
	if err != nil {
		t.Errorf("Something went wrong: %v", err)
	}

	res, err := server.EmailUsed(context.Background(), &userPB.EmailUsedRequest{Email: "wrong"})
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if res.GetUsed() {
		t.Errorf("1: Wrong result")
	}
	res, err = server.EmailUsed(context.Background(), &userPB.EmailUsedRequest{Email: "test@gmail.com"})
	if err != nil {
		t.Errorf("2: An error was returned: %v", err)
	}
	if !res.GetUsed() {
		t.Errorf("2: Wrong result")
	}
}

func Test_authServer_SignUp(t *testing.T) {
	db := database2.ConnectToTestDB()
	userCollection := *db.Collection(config.USERCOLLECTIONNAME)
	server := authServer{userCollection: userCollection}

	_, err := db.Collection("user").InsertOne(context.Background(), database2.User{Username: "Carl", Email: "test@gmail.com"})
	if err != nil {
		t.Errorf("Something went wrong: %v", err)
	}

	_, err = server.SignUp(context.Background(), &userPB.SignupRequest{Username: "Carl", Email: "example@gmail.com", Password: "examplestring"})
	if err.Error() != "Username already taken" {
		t.Error("1: No or wrong error returned")
	}
	_, err = server.SignUp(context.Background(), &userPB.SignupRequest{Username: "Carlo", Email: "test@gmail.com", Password: "examplestring"})
	if err.Error() != "Email already used" {
		t.Error("2: No or wrong error returned")
	}

	_, err = server.SignUp(context.Background(), &userPB.SignupRequest{Username: "Example", Email: "example@gmail.com", Password: "examplestring"})
	if err != nil {
		t.Errorf("3: Error creating new user: %v", err)
	}

	_, err = server.SignUp(context.Background(), &userPB.SignupRequest{Username: "Example", Email: "example@gmail.com", Password: "pass"})
	if err.Error() != "Validation failed" {
		t.Error("4: No or wrong error returned for validation")
	}
}

func Test_authServer_AuthUser(t *testing.T) {
	db := database2.ConnectToTestDB()
	userCollection := *db.Collection(config.USERCOLLECTIONNAME)
	server := authServer{userCollection: userCollection}

	res, err := server.AuthUser(context.Background(), &userPB.AuthUserRequest{Token: config.ExampleToken})
	if err != nil {
		t.Errorf("An error was returned: %v", err)
	}
	if res.GetID() != "61ad37c77b5cdb57f14c36e6" || res.GetUsername() != "Carl" || res.GetEmail() != "test@gmail.com" {
		t.Errorf("Wrong result returned: %v", res)
	}
}
