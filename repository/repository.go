package repository

import (
	"context"
	"redis/model"
	"redis/mongodb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	GetHumans(ctx context.Context) (*model.Humans, error)
	GetHuman(ctx context.Context, id string) (*model.Human, error)
}

type repository struct {
	mongodb mongodb.MongoDB
}

func NewRepository(db mongodb.MongoDB) Repository {
	return &repository{mongodb: db}
}

func (r *repository) GetHumans(ctx context.Context) (*model.Humans, error) {
	var humans model.Humans
	cursor, err := r.mongodb.GetHumansCollection().Find(ctx, bson.D{})
	if err != nil {
		return &humans, err
	}
	err = cursor.All(ctx, &humans)
	if err != nil {
		return &humans, err
	}

	return &humans, err
}

func (r *repository) GetHuman(ctx context.Context, id string) (*model.Human, error) {
	var human model.Human
	var err error

	objectID, _ := primitive.ObjectIDFromHex(id)
	err = r.mongodb.GetHumansCollection().FindOne(
		ctx,
		bson.M{"_id": objectID},
	).Decode(&human)

	if err != nil {
		return &human, err
	}
	return &human, err
}
