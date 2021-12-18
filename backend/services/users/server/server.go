package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"regexp"
	"time"

	config "github.com/jalexanderII/zero_microservice"
	database2 "github.com/jalexanderII/zero_microservice/backend/services/users/database"
	userPB "github.com/jalexanderII/zero_microservice/gen/users"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type authServer struct {
	userPB.UnimplementedAuthServiceServer
	userCollection mongo.Collection
}

func NewServer(c mongo.Collection) *authServer {
	return &authServer{userCollection: c}
}

func (server authServer) Login(ctx context.Context, in *userPB.LoginRequest) (*userPB.AuthResponse, error) {
	login, password := in.GetLogin(), in.GetPassword()
	ctx, cancel := database2.NewDBContext(5 * time.Second)
	defer cancel()

	var user database2.User
	err := server.userCollection.FindOne(ctx, bson.M{"$or": []bson.M{{"username": login}, {"email": login}}}).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("something went wrong: %v", err)
	}
	if user == database2.NilUser || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, errors.New("wrong login credentials provided")
	}
	return &userPB.AuthResponse{Token: user.GetToken()}, nil
}

func (server authServer) SignUp(ctx context.Context, in *userPB.SignupRequest) (*userPB.AuthResponse, error) {
	username, email, password := in.GetUsername(), in.GetEmail(), in.GetPassword()
	match, _ := regexp.MatchString(config.EmailRegex, email)
	if len(username) < 4 || len(username) > 20 || len(email) < 7 || len(email) > 35 || len(password) < 8 || len(password) > 128 || !match {
		return nil, errors.New("validation failed")
	}
	res, err := server.UsernameUsed(context.Background(), &userPB.UsernameUsedRequest{Username: username})
	if err != nil {
		log.Printf("Error returned from UsernameUsed: %v\n", err)
		return nil, errors.New("something went wrong")
	}
	if res.GetUsed() {
		return nil, errors.New("username already taken")
	}

	res, err = server.EmailUsed(context.Background(), &userPB.EmailUsedRequest{Email: email})
	if err != nil {
		log.Printf("Error returned from EmailUsed: %v\n", err)
		return nil, errors.New("EmailUsed")
	}
	if res.GetUsed() {
		return nil, errors.New("email already used")
	}

	pw, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	newUser := database2.User{ID: primitive.NewObjectID(), Email: email, Username: username, Password: string(pw)}

	ctx, cancel := database2.NewDBContext(5 * time.Second)
	defer cancel()
	_, err = server.userCollection.InsertOne(ctx, newUser)
	if err != nil {
		log.Printf("Error inserting new user: %v\n", err)
		return nil, errors.New("something went wrong")
	}
	return &userPB.AuthResponse{Token: newUser.GetToken()}, nil
}

func (server authServer) EmailUsed(ctx context.Context, in *userPB.EmailUsedRequest) (*userPB.UsedResponse, error) {
	var email = in.GetEmail()
	ctx, cancel := database2.NewDBContext(5 * time.Second)
	defer cancel()
	var result database2.User
	err := server.userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("something went wrong: %v", err)
	}
	return &userPB.UsedResponse{Used: result != database2.NilUser}, nil
}

func (server authServer) UsernameUsed(ctx context.Context, in *userPB.UsernameUsedRequest) (*userPB.UsedResponse, error) {
	var username = in.GetUsername()
	ctx, cancel := database2.NewDBContext(5 * time.Second)
	defer cancel()
	var result database2.User
	err := server.userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("something went wrong: %v", err)
	}
	return &userPB.UsedResponse{Used: result != database2.NilUser}, nil
}

func (server authServer) AuthUser(_ context.Context, in *userPB.AuthUserRequest) (*userPB.AuthUserResponse, error) {
	var token = in.GetToken()
	user := database2.UserFromToken(token)
	return &userPB.AuthUserResponse{ID: user.ID.Hex(), Username: user.Username, Email: user.Email}, nil
}
