package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"regexp"

	"github.com/hashicorp/go-hclog"
	userDB "github.com/jalexanderII/zero_microservice/backend/services/users/database"
	"github.com/jalexanderII/zero_microservice/backend/services/users/middleware"
	"github.com/jalexanderII/zero_microservice/config"
	userPB "github.com/jalexanderII/zero_microservice/gen/users"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	userPB.UnimplementedAuthServiceServer
	DB   mongo.Collection
	jwtm *middleware.JWTManager
	l    hclog.Logger
}

func NewAuthServer(DB mongo.Collection, jwtm *middleware.JWTManager, l hclog.Logger) *AuthServer {
	return &AuthServer{DB: DB, jwtm: jwtm, l: l}
}

func (server AuthServer) Login(ctx context.Context, in *userPB.LoginRequest) (*userPB.AuthResponse, error) {
	username, password := in.GetUsername(), in.GetPassword()
	var user userDB.User
	err := server.DB.FindOne(ctx, bson.M{"$or": []bson.M{{"username": username}, {"email": username}}}).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("cannot find user: %v", err)
	}
	if user == userDB.NilUser || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, errors.New("wrong login credentials provided")
	}

	token, err := server.jwtm.Generate(&user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	return &userPB.AuthResponse{Token: token}, nil
}

func (server AuthServer) SignUp(ctx context.Context, in *userPB.SignupRequest) (*userPB.AuthResponse, error) {
	username, email, password, role := in.GetUsername(), in.GetEmail(), in.GetPassword(), in.GetRole()
	match, _ := regexp.MatchString(config.EmailRegex, email)
	if !match {
		return nil, errors.New("email validation failed")
	}

	// res, err := server.UsernameUsed(ctx, username)
	// if err != nil {
	// 	log.Printf("Error returned from UsernameUsed: %v\n", err)
	// 	return nil, err
	// }
	// if res {
	// 	return nil, errors.New("username already taken")
	// }
	// res, err = server.EmailUsed(ctx, email)
	// if err != nil {
	// 	log.Printf("Error returned from EmailUsed: %v\n", err)
	// 	return nil, err
	// }
	// if res {
	// 	return nil, errors.New("email already used")
	// }

	pw, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	newUser := userDB.User{ID: primitive.NewObjectID(), Email: email, Username: username, Password: string(pw), Role: userDB.Role(role)}

	_, err := server.DB.InsertOne(ctx, newUser)
	if err != nil {
		log.Printf("Error inserting new user: %v\n", err)
		return nil, err
	}

	token, err := server.jwtm.Generate(&newUser)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	return &userPB.AuthResponse{Token: token}, nil
}

// func (server AuthServer) EmailUsed(ctx context.Context, email string) (bool, error) {
// 	var user userDB.User
// 	err := server.DB.FindOne(ctx, bson.M{"$or": []bson.M{{"username": email}, {"email": email}}}).Decode(&user)
// 	if err != nil {
// 		// ErrNoDocuments means that the filter did not match any documents in the collection.
// 		if err == mongo.ErrNoDocuments {
// 			return true, fmt.Errorf("not found: %v", err)
// 		}
// 		return true, fmt.Errorf("error fetching email: %v", err)
// 	}
// 	return user != userDB.NilUser, nil
// }

//
// func (server AuthServer) UsernameUsed(ctx context.Context, username string) (bool, error) {
// 	var user userDB.User
// 	err := server.DB.FindOne(ctx, bson.M{"$or": []bson.M{{"username": username}, {"email": username}}}).Decode(&user)
// 	if err != nil {
// 		// ErrNoDocuments means that the filter did not match any documents in the collection.
// 		if err == mongo.ErrNoDocuments {
// 			return true, fmt.Errorf("not found: %v", err)
// 		}
// 		return true, fmt.Errorf("error fetching username: %v", err)
// 	}
// 	return user != userDB.NilUser, nil
// }
