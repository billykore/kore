package repo

import (
	"context"

	"cloud.google.com/go/firestore"
	firestoredb "github.com/billykore/kore/backend/pkg/db"
	"github.com/billykore/kore/backend/pkg/model"
	"github.com/billykore/kore/backend/pkg/repo"
)

type todoRepo struct {
	firestore *firestore.Client
}

func NewTodoRepository(firestore *firestore.Client) repo.TodoRepository {
	return &todoRepo{firestore: firestore}
}

func (r *todoRepo) Get(ctx context.Context, isDone string) ([]*model.Todo, error) {
	var iter *firestore.DocumentIterator
	if isDone == "true" {
		iter = r.firestore.Collection(firestoredb.TodosCollectionPath).Where("IsDone", "==", true).Documents(ctx)
	} else if isDone == "false" {
		iter = r.firestore.Collection(firestoredb.TodosCollectionPath).Where("IsDone", "==", false).Documents(ctx)
	} else {
		iter = r.firestore.Collection(firestoredb.TodosCollectionPath).Documents(ctx)
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

func (r *todoRepo) GetById(ctx context.Context, id string) (*model.Todo, error) {
	doc, err := r.firestore.Collection(firestoredb.TodosCollectionPath).Doc(id).Get(ctx)
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

func (r *todoRepo) Save(ctx context.Context, todo *model.Todo) error {
	_, err := r.firestore.Collection(firestoredb.TodosCollectionPath).Doc(todo.Id).Set(ctx, todo)
	return err
}

func (r *todoRepo) Update(ctx context.Context, id string) error {
	_, err := r.firestore.Collection(firestoredb.TodosCollectionPath).Doc(id).
		Update(ctx, []firestore.Update{
			{Path: "IsDone", Value: true},
		})
	return err
}

func (r *todoRepo) Delete(ctx context.Context, id string) error {
	_, err := r.firestore.Collection(firestoredb.TodosCollectionPath).Doc(id).Delete(ctx)
	return err
}
