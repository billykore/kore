package repo

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"github.com/billykore/kore/libs/db"
	"github.com/billykore/kore/libs/model"
	"github.com/billykore/kore/libs/repo"
)

type userRepo struct {
	firestore *firestore.Client
}

func NewUserRepository(firestore *firestore.Client) repo.UserRepository {
	return &userRepo{firestore: firestore}
}

func (r *userRepo) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	iter := r.firestore.Collection(db.UsersCollectionPath).
		Where("username", "==", username).
		Limit(1).
		Documents(ctx)

	docs, err := iter.GetAll()
	if err != nil {
		return nil, err
	}
	if len(docs) < 1 {
		return nil, errors.New("user not found")
	}

	user := new(model.User)
	err = docs[0].DataTo(user)
	return user, err
}
