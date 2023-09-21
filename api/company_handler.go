package api

import (
	"context"
	"errors"
	"fmt"

	"github.com/thulani196/recruits-hub/database"
	"github.com/thulani196/recruits-hub/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	companyCollection = "companies"
)

type CompanyRepository interface {
	Create(company *types.Company) error
	GetAll() (*types.Company, error)
	GetById(id string) (*types.Company, error)
}

type MongoCompanyRepository struct {
	collection *mongo.Collection
}

func NewMongoCompanyRepository() *MongoCompanyRepository {
	collection := database.DBInstance.DB.Collection(companyCollection)
	return &MongoCompanyRepository{collection}
}

func (r *MongoCompanyRepository) Create(company *types.Company) (*types.Company, error) {
	res, err := r.collection.InsertOne(context.Background(), company)

	if err != nil {
		fmt.Println("Error occured: ", err)
		return nil, err
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		company.ID = oid.Hex()
	} else {
		fmt.Println("failed to convert InsertedID to ObjectID ")
		return nil, errors.New("failed to convert InsertedID to ObjectID")
	}

	return company, nil
}

func (r *MongoCompanyRepository) GetAll() ([]*types.Company, error) {
	var companies []*types.Company
	cursor, err := r.collection.Find(context.Background(), bson.M{})

	if err != nil {
		fmt.Println("Error occured: ", err)
		return nil, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var company types.Company
		if err := cursor.Decode(&company); err != nil {
			fmt.Println("Error occured: ", err)
			return nil, err
		}

		companies = append(companies, &company)
	}

	return companies, nil
}

func (r *MongoCompanyRepository) GetById(id string) (*types.Company, error) {
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		fmt.Println("Error occured: ", err)
		return nil, err
	}

	var company types.Company
	err = r.collection.FindOne(context.Background(), bson.M{"_id": oid}).Decode(&company)

	if err != nil {
		fmt.Println("Error occured: ", err)
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no documents with given id found")
		}
		return nil, err
	}

	return &company, nil
}
