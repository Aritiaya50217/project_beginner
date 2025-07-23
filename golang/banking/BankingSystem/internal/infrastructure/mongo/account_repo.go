package mongo

import (
	"banking-hexagonal/internal/domain"
	"context"
	"errors"
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

func (r *AccountRepo) FindByID(id string) (*domain.Account, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid id")
	}
	var doc struct {
		ID        primitive.ObjectID `bson:"_id"`
		Owner     string             `bson:"owner"`
		Balance   float64            `bson:"balance"`
		CreatedAt time.Time          `bson:"created_at"`
	}
	err = r.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&doc)
	if err != nil {
		return nil, err
	}
	return &domain.Account{
		ID:        doc.ID.Hex(),
		Owner:     doc.Owner,
		Balance:   doc.Balance,
		CreatedAt: doc.CreatedAt,
	}, nil

}

func (r *AccountRepo) Update(account *domain.Account) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	oid, err := primitive.ObjectIDFromHex(account.ID)
	if err != nil {
		return errors.New("invalid id")
	}
	update := bson.M{
		"$set": bson.M{
			"balance": account.Balance,
		},
	}
	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": oid}, update)
	return err
}
