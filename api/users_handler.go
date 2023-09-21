package api

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/thulani196/recruits-hub/database"
	"github.com/thulani196/recruits-hub/types"
	"github.com/thulani196/recruits-hub/utils"
	"github.com/thulani196/recruits-hub/validation"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	// Define a secret key for JWT
	secretKey      = []byte(os.Getenv("JWT_SECRET_KEY"))
	userCollection = "users"
)

type UserRepository interface {
	RegisterHandler(user *types.User) error
	LoginHandler(user *types.User) (*types.Token, error)
	ProtectedHandler(c *fiber.Ctx) error
}

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepository() *MongoUserRepository {
	collection := database.DBInstance.DB.Collection(collectionName)
	return &MongoUserRepository{collection}
}

// RegisterHandler handles user registration
func (r *MongoUserRepository) RegisterHandler(user *types.User) error {
	// Hash and salt the user's password (use a password hashing library)
	hashedPassword, err := utils.HashAndSaltPassword(user.Password)
	if err != nil {
		fmt.Println("Error occured: ", err)
		return errors.New("password could not be hashed")
	}

	// Get a handle to the "users" collection in MongoDB
	collection := database.DBInstance.DB.Collection(userCollection)

	// Create a new user document
	userDocument := bson.M{
		"username":          user.Username,
		"password":          hashedPassword, // Store the hashed password
		"is_active":         true,
		"companies":         make([]string, 0),
		"registration_date": time.Now(),
		"user_image":        "",
	}

	// Insert the user document into the collection
	_, err = collection.InsertOne(context.Background(), userDocument)
	if err != nil {
		return errors.New("failed to register user")
	}

	return nil
}

// LoginHandler handles user login and issues a JWT token
func (r *MongoUserRepository) LoginHandler(user *types.User) (*types.Token, error) {

	// Authenticate the user (check the password)
	if !authenticateUser(user.Username, user.Password) {
		fmt.Println("invalid username or password")
		return nil, errors.New("invalid username or password")
	}

	// Get a handle to the "users" collection in MongoDB
	collection := database.DBInstance.DB.Collection(userCollection)

	// Find the user document by username
	var foundUser types.User
	filter := bson.M{"username": user.Username}
	err := collection.FindOne(context.Background(), filter).Decode(&foundUser)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, errors.New("invalid username or password")
	}

	// Create a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, types.Claims{
		Username:  foundUser.Username,
		Companies: foundUser.Companies,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}

	var loginResponse = &types.Token{
		Success: true,
		Token:   tokenString,
	}

	return loginResponse, nil
}

// AuthenticateUser checks the user's credentials and returns true if they are valid
func authenticateUser(username, password string) bool {
	// Get a handle to the "users" collection in MongoDB
	collection := database.DBInstance.DB.Collection(userCollection)

	// Define a filter to find the user by username
	filter := bson.M{"username": username}

	// Find the user document by username
	var foundUser types.User
	err := collection.FindOne(context.Background(), filter).Decode(&foundUser)
	if err != nil {
		// User not found or other error occurred
		fmt.Println("Error occured: ", err)
		return false
	}

	// Verify the provided password against the hashed and salted password in the database
	return utils.CheckPassword(password, foundUser.Password) == nil
}

// ProtectedHandler handles a protected endpoint
func ProtectedHandler(c *fiber.Ctx) error {

	unathorized := c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"success": false,
		"message": "Unauthorized",
	})

	// Extract the JWT token from the Authorization header
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return unathorized
	}

	// Parse and validate the token
	token, err := jwt.ParseWithClaims(tokenString, &types.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		fmt.Println("Error occured: ", err)
		return unathorized
	}

	// Check if the token is valid
	if !token.Valid {
		return unathorized
	}

	_, ok := token.Claims.(*types.Claims)
	if !ok {
		return unathorized
	}

	return c.Next()
}

func ValidateUserRegistration(c *fiber.Ctx, user *types.User) error {
	// Check if the username is already taken (you should implement this logic)
	if validation.IsUsernameTaken(user.Username) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Username is already in use.",
		})
	}

	return c.Next()
}
