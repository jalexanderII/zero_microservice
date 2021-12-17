package database

import (
	"encoding/json"

	"github.com/golang-jwt/jwt"
	config "github.com/jalexanderII/zero_microservice"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// NilUser is the nil value of our user
var NilUser User

// User is the default user struct
type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}

// GetToken returns the Users JWT token
func (u User) GetToken() string {
	byteSlc, _ := json.Marshal(u)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data": string(byteSlc),
	})
	tokenString, _ := token.SignedString(config.JWTSecret)
	return tokenString
}

// UserFromToken returns a user which is authenticated with this token
func UserFromToken(token string) User {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return config.JWTSecret, nil
	})
	if err != nil {
		return User{}
	}
	var result User
	err = json.Unmarshal([]byte(claims["data"].(string)), &result)
	if err != nil {
		return User{}
	}
	return result
}
