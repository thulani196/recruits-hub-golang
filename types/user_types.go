package types

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// User model
type User struct {
	Username         string    `bson:"username"`
	Password         string    `bson:"password"`
	IsActive         bool      `bson:"is_active"`
	Companies        []string  `bson:"companies,omitempty"`
	RegistrationDate time.Time `bson:"registration_date"`
	UserImage        string    `bson:"user_image,omitempty"`
}

type UserDetails struct {
	ID        string  `bson:"_id,omitempty"`
	UserID    string  `bson:"user_id,omitempty"`
	FirstName string  `bson:"first_name,omitempty"`
	LastName  string  `bson:"last_name,omitempty"`
	Hobbies   []Hobby `bson:"hobbies,omitempty"`
}

// Claims represents the JWT claims
type Claims struct {
	Username  string   `bson:"username"`
	Companies []string `bson:"companies"`
	jwt.StandardClaims
}

type Token struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}
