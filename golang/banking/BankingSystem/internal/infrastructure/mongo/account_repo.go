package mongo

import (
	"banking-hexagonal/internal/domain"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountRepo struct {
	collection *mongo.Collection
}

func NewAccountRepo(db *mongo.Database) domain.AccountRepository {
	return &AccountRepo{
		collection: db.Collection("account"),
	}
}

func (r *AccountRepo) Create(account *domain.Account) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := primitive.NewObjectID()
	account.ID = id.Hex()

	_, err := r.collection.InsertOne(ctx, bson.M{
		"_id":        id,
		"owner":      account.Owner,
		"balance":    account.Balance,
		"created_at": account.CreatedAt,
	})
	return err
}
