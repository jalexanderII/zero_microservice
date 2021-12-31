package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"regexp"

	"github.com/hashicorp/go-hclog"
	userDB "github.com/jalexanderII/zero_microservice/backend/services/users/database"
	"github.com/jalexanderII/zero_microservice/config"
	"github.com/jalexanderII/zero_microservice/config/middleware"
	listingsPB "github.com/jalexanderII/zero_microservice/gen/listings"
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
	DB             mongo.Collection
	jwtm           *middleware.JWTManager
	ListingsClient listingsPB.ListingsClient
	l              hclog.Logger
}

func NewAuthServer(DB mongo.Collection, jwtm *middleware.JWTManager, lc listingsPB.ListingsClient, l hclog.Logger) *AuthServer {
	return &AuthServer{DB: DB, jwtm: jwtm, ListingsClient: lc, l: l}
}

func (server AuthServer) Login(ctx context.Context, in *userPB.LoginRequest) (*userPB.AuthResponse, error) {
	username, password := in.GetUsername(), in.GetPassword()
	var user userDB.User
	err := server.DB.FindOne(ctx, bson.M{"$or": []bson.M{{"username": username}, {"email": username}}}).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("cannot find user: %v", err)
	}
	if user.ID.IsZero() || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
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

	newId, err := server.createUserFromMetadata(ctx, in.GetMetadata(), newUser.ID.Hex())
	if err != nil {
		return nil, err
	}
	if newId == 0 {
		return nil, fmt.Errorf("new user not created, id is < 1 %v", newId)
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

func (server AuthServer) GetUser(ctx context.Context, in *userPB.GetUserRequest) (*userPB.User, error) {
	var user userDB.User
	filter := bson.D{{"_id", in.GetId()}}
	err := server.DB.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return userDBToPB(&user), nil
}

func (server AuthServer) ListUsers(ctx context.Context, in *userPB.ListUserRequest) (*userPB.ListUserResponse, error) {
	var results []userDB.User
	cursor, err := server.DB.Find(ctx, bson.D{})
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	res := make([]*userPB.User, len(results))
	for idx, user := range results {
		res[idx] = userDBToPB(&user)
	}
	return &userPB.ListUserResponse{Users: res}, nil
}

func (server AuthServer) UpdateUser(ctx context.Context, in *userPB.UpdateUserRequest) (*userPB.User, error) {
	username, email := in.GetUser().GetUsername(), in.GetUser().GetEmail()
	filter := bson.D{{"_id", in.GetId()}}
	update := bson.D{{"$set", bson.D{{"username", username}, {"email", email}}}}
	_, err := server.DB.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	var user userDB.User
	err = server.DB.FindOne(ctx, filter).Decode(&user)
	return userDBToPB(&user), nil
}

func (server AuthServer) DeleteUser(ctx context.Context, in *userPB.DeleteUserRequest) (*userPB.DeleteUserResponse, error) {
	filter := bson.D{{"_id", in.GetId()}}
	_, err := server.DB.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}
	var user userDB.User
	err = server.DB.FindOne(ctx, filter).Decode(&user)
	return &userPB.DeleteUserResponse{Status: userPB.STATUS_SUCCESS, User: userDBToPB(&user)}, nil
}

func (server AuthServer) createUserFromMetadata(ctx context.Context, md *userPB.Metadata, userID string) (int32, error) {
	switch userType := md.UserType.(type) {
	case *userPB.Metadata_AdminMetadata:
		return 0, nil
	case *userPB.Metadata_RenterMetadata:
		return 0, nil
	case *userPB.Metadata_OwnerMetadata:
		o := &listingsPB.Owner{
			Name:        md.GetOwnerMetadata().GetName(),
			Email:       md.GetOwnerMetadata().GetEmail(),
			PhoneNumber: md.GetOwnerMetadata().GetPhoneNumber(),
			Company:     md.GetOwnerMetadata().GetCompany(),
			UserRef:     userID,
		}
		owner, err := server.ListingsClient.CreateOwner(ctx, &listingsPB.CreateOwnerRequest{Owner: o})
		if err != nil {
			return 0, fmt.Errorf("[Error] creating owner : %v", err)
		}
		return owner.GetId(), nil
	case *userPB.Metadata_RealtorMetadata:
		r := &listingsPB.Realtor{
			Name:        md.GetOwnerMetadata().GetName(),
			Email:       md.GetOwnerMetadata().GetEmail(),
			PhoneNumber: md.GetOwnerMetadata().GetPhoneNumber(),
			Company:     md.GetOwnerMetadata().GetCompany(),
			UserRef:     userID,
		}
		realtor, err := server.ListingsClient.CreateRealtor(ctx, &listingsPB.CreateRealtorRequest{Realtor: r})
		if err != nil {
			return 0, fmt.Errorf("[Error] creating owner : %v", err)
		}
		return realtor.GetId(), nil
	default:
		return 0, fmt.Errorf("[Error] incorrect user type: %v", userType)
	}
}

func userDBToPB(user *userDB.User) *userPB.User {
	return &userPB.User{
		Id:       user.ID.Hex(),
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
}
