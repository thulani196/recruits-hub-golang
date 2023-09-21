package validation

import (
	"context"
	"log"

	"github.com/thulani196/recruits-hub/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	userCollection = "users"
)

// User model
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// IsUsernameTaken checks if a username is already taken in the database
func IsUsernameTaken(username string) bool {
	// Get a handle to the "users" collection in MongoDB
	collection := database.DBInstance.DB.Collection(userCollection)

	// Define a filter to find a user with the given username
	filter := bson.M{"username": username}

	// Try to find a document matching the filter
	var existingUser User
	err := collection.FindOne(context.Background(), filter).Decode(&existingUser)
	if err == mongo.ErrNoDocuments {
		// No matching user found, username is not taken
		return false
	} else if err != nil {
		log.Printf("Error checking username: %v", err)
		return true // An error occurred, assume the username is taken to be safe
	}

	// A user with the same username already exists
	return true
}
