package server

import (
	"context"
	"testing"
	"time"

	"github.com/hashicorp/go-hclog"
	config "github.com/jalexanderII/zero_microservice"
	userDB "github.com/jalexanderII/zero_microservice/backend/services/users/database"
	"github.com/jalexanderII/zero_microservice/backend/services/users/middleware"
	userPB "github.com/jalexanderII/zero_microservice/gen/users"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var L = hclog.Default()

func Test_authServer_Login(t *testing.T) {
	ctx, cancel := config.NewDBContext(5 * time.Second)
	defer cancel()

	DB := userDB.InitiateMongoClient()
	userCollection := *DB.Collection(config.USERCOLLECTIONNAME)
	jwtManager := middleware.NewJWTManager(config.JWTSecret, config.TokenDuration)
	server := NewAuthServer(userCollection, jwtManager, L)

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
	server := NewAuthServer(userCollection, jwtManager, L)

	u := &userPB.SignupRequest{
		Username: "example2",
		Email:    "example2@gmail.com",
		Password: "password2",
		Role:     int32(userDB.Admin),
	}

	// _, err := server.SignUp(context.Background(), &userPB.SignupRequest{Username: "Carl", Email: "example@gmail.com", Password: "examplestring"})
	// if err.Error() != "Username already taken" {
	// 	t.Error("2: No or wrong error returned for username already taken")
	// }
	//
	// _, err := server.SignUp(context.Background(), &userPB.SignupRequest{Username: "Carlo", Email: "test@gmail.com", Password: "examplestring"})
	// if err.Error() != "Email already used" {
	// 	t.Error("3: No or wrong error returned for email already taken")
	// }

	_, err := server.SignUp(context.Background(), u)
	if err != nil {
		t.Errorf("4: Error creating new user: %v", err)
	}

	_, err = server.SignUp(context.Background(), &userPB.SignupRequest{Username: "Example", Email: "e", Password: "pass"})
	if err.Error() != "email validation failed" {
		t.Error("5: No or wrong error returned for email validation")
	}
}
