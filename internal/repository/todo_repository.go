package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	firestoredb "github.com/billykore/todolist/internal/database/firestore"
	"github.com/billykore/todolist/internal/model"
)

type TodoRepository struct {
	firestore *firestore.Client
}

func NewTodoRepository(firestore *firestore.Client) *TodoRepository {
	return &TodoRepository{firestore: firestore}
}

func (r *TodoRepository) GetTodos(ctx context.Context, query *model.Query) ([]*model.Todo, error) {
	var iter *firestore.DocumentIterator
	if !query.IsEmpty() {
		iter = r.firestore.Collection("todos").Where(query.Key, "==", query.Value).Documents(ctx)
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

func (r *TodoRepository) GetTodoById(ctx context.Context, id string) (*model.Todo, error) {
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

func (r *TodoRepository) SaveTodo(ctx context.Context, todo *model.Todo) error {
	_, err := r.firestore.Collection(firestoredb.TodoCollectionPath).Doc(todo.Id).Set(ctx, todo)
	return err
}

func (r *TodoRepository) SetDoneTodo(ctx context.Context, id string) error {
	_, err := r.firestore.Collection(firestoredb.TodoCollectionPath).Doc(id).
		Update(ctx, []firestore.Update{
			{Path: "IsDone", Value: true},
		})
	return err
}

func (r *TodoRepository) DeleteTodo(ctx context.Context, id string) error {
	_, err := r.firestore.Collection(firestoredb.TodoCollectionPath).Doc(id).Delete(ctx)
	return err
}
