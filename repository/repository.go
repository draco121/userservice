package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/draco121/common/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserRepository interface {
	InsertOne(ctx context.Context, user *models.User) (string, error)
	UpdateOne(ctx context.Context, user *models.User) (*models.User, error)
	FindOneById(ctx context.Context, id string) (*models.User, error)
	FindOneByEmail(ctx context.Context, email string) (*models.User, error)
	DeleteOneById(ctx context.Context, id string) (*models.User, error)
}

type userRepository struct {
	IUserRepository
	db *mongo.Database
}

func NewUserRepository(database *mongo.Database) IUserRepository {
	repo := userRepository{db: database}
	return &repo
}

func (ur userRepository) InsertOne(ctx context.Context, user *models.User) (string, error) {
	result, _ := ur.FindOneByEmail(ctx, user.Email)
	if result != nil {
		return "", fmt.Errorf("record exists")
	} else {
		result, err := ur.db.Collection("users").InsertOne(ctx, user)
		if err != nil {
			return "", err
		} else {
			id := result.InsertedID.(primitive.ObjectID)
			return id.Hex(), nil
		}
	}
}

func (ur userRepository) UpdateOne(ctx context.Context, user *models.User) (*models.User, error) {
	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": bson.M{
		"password": user.Password,
	}}
	result := models.User{}
	err := ur.db.Collection("users").FindOneAndUpdate(ctx, filter, update).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else {
		return &result, nil
	}
}

func (ur userRepository) FindOneById(ctx context.Context, id string) (*models.User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	} else {
		filter := bson.D{{Key: "_id", Value: objectId}}
		result := models.User{}
		err := ur.db.Collection("users").FindOne(ctx, filter).Decode(&result)
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		} else {
			return &result, nil
		}
	}
}

func (ur userRepository) FindOneByEmail(ctx context.Context, email string) (*models.User, error) {
	filter := bson.D{{Key: "email", Value: email}}
	result := models.User{}
	err := ur.db.Collection("users").FindOne(ctx, filter).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else {
		return &result, nil
	}
}

func (ur userRepository) DeleteOneById(ctx context.Context, id string) (*models.User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	} else {
		filter := bson.D{{Key: "_id", Value: objectId}}
		result := models.User{}
		err := ur.db.Collection("users").FindOneAndDelete(ctx, filter).Decode(&result)
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		} else {
			return &result, nil
		}
	}
}
