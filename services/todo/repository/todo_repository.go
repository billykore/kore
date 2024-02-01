package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	firestoredb "github.com/billykore/kore/libs/database/firestore"
	"github.com/billykore/kore/libs/model"
	"github.com/billykore/kore/libs/repository"
)

type todoRepo struct {
	firestore *firestore.Client
}

func NewTodoRepository(firestore *firestore.Client) repository.Todo {
	return &todoRepo{firestore: firestore}
}

func (r *todoRepo) GetTodos(ctx context.Context, isDone string) ([]*model.Todo, error) {
	var iter *firestore.DocumentIterator
	if isDone == "true" {
		iter = r.firestore.Collection("todos").Where("IsDone", "==", true).Documents(ctx)
	} else if isDone == "false" {
		iter = r.firestore.Collection("todos").Where("IsDone", "==", false).Documents(ctx)
	} else {
		iter = r.firestore.Collection("todos").Documents(ctx)
	}

	docs, err := iter.GetAll()
	if err != nil {
		return nil, err
	}

	var todos []*model.Todo
	for _, doc := range docs {
		todo := new(model.Todo)
		err = doc.DataTo(todo)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (r *todoRepo) GetTodoById(ctx context.Context, id string) (*model.Todo, error) {
	doc, err := r.firestore.Collection(firestoredb.TodoCollectionPath).Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}
	todo := new(model.Todo)
	err = doc.DataTo(todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *todoRepo) SaveTodo(ctx context.Context, todo *model.Todo) error {
	_, err := r.firestore.Collection(firestoredb.TodoCollectionPath).Doc(todo.Id).Set(ctx, todo)
	return err
}

func (r *todoRepo) UpdateTodo(ctx context.Context, id string) error {
	_, err := r.firestore.Collection(firestoredb.TodoCollectionPath).Doc(id).
		Update(ctx, []firestore.Update{
			{Path: "IsDone", Value: true},
		})
	return err
}

func (r *todoRepo) DeleteTodo(ctx context.Context, id string) error {
	_, err := r.firestore.Collection(firestoredb.TodoCollectionPath).Doc(id).Delete(ctx)
	return err
}
