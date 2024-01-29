package repository

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"github.com/billykore/todolist/libs/model"
	"github.com/billykore/todolist/libs/repository"
)

type userRepo struct {
	firestore *firestore.Client
}

func NewUserRepository(firestore *firestore.Client) repository.User {
	return &userRepo{firestore: firestore}
}

func (r *userRepo) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	iter := r.firestore.Collection("users").
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
