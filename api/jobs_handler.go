package api

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/thulani196/recruits-hub/database"
	"github.com/thulani196/recruits-hub/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collectionName = "jobs"
)

type JobRepository interface {
	CreateJob(job *types.Job) (*types.Job, error)
	GetJobByID(id string) (*types.Job, error)
	GetAll() ([]*types.Job, error)
	UpdateJob(id string, job *types.Job) error
}

type MongoJobRepository struct {
	collection *mongo.Collection
}

func NewMongoJobRepository() *MongoJobRepository {
	collection := database.DBInstance.DB.Collection(collectionName)
	return &MongoJobRepository{collection}
}

func (r *MongoJobRepository) CreateJob(job *types.Job) (*types.Job, error) {
	job.TotalApplicants = 0
	job.IsActive = true

	res, err := r.collection.InsertOne(context.Background(), job)
	if err != nil {
		fmt.Println("Error occured: ", err)
		return nil, err
	}

	// Check if the inserted ID is of type primitive.ObjectID
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		job.ID = oid.Hex()
	} else {
		fmt.Println("failed to convert InsertedID to ObjectID ")
		return nil, errors.New("failed to convert InsertedID to ObjectID")
	}

	return job, nil
}

func (r *MongoJobRepository) GetAll() ([]*types.Job, error) {
	var jobs []*types.Job
	cursor, err := r.collection.Find(context.Background(), bson.M{})

	if err != nil {
		fmt.Println("Error occured: ", err)
		return nil, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var job types.Job
		if err := cursor.Decode(&job); err != nil {
			fmt.Println("Error occured: ", err)
			return nil, err
		}

		jobs = append(jobs, &job)
	}

	return jobs, nil
}

func (r *MongoJobRepository) GetJobByID(id string) (*types.Job, error) {
	// Convert the ID string to a MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		fmt.Println("Error occured: ", err)
		return nil, err
	}

	var job types.Job
	err = r.collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&job)

	if err != nil {
		fmt.Println("Error occured: ", err)
		if err == mongo.ErrNoDocuments {
			return nil, errors.New(("no documents with this ID found"))
		}
		return nil, err
	}

	return &job, nil
}

func (r *MongoJobRepository) UpdateJob(id string, job *types.Job) error {
	// Convert the ID string to a MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		fmt.Println("Error occured: ", err)
		return err
	}

	// Define the filter to find the Todo by ID
	filter := bson.M{"_id": objectID}

	// Define the update to apply
	update := bson.M{
		"$set": bson.M{
			"title":       job.Title,
			"description": job.Description,
			"updatedAt":   time.Now(),
		},
	}

	// Perform the update operation
	_, err = r.collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		fmt.Println("Error occured: ", err)
		return err
	}

	return nil
}
