package repo

import (
	"context"

	"github.com/billykore/kore/backend/pkg/model"
	"github.com/billykore/kore/backend/pkg/repo"
	"gorm.io/gorm"
)

type todoRepo struct {
	db *gorm.DB
}

func NewTodoRepository(postgres *gorm.DB) repo.TodoRepository {
	return &todoRepo{db: postgres}
}

func (r *todoRepo) List(ctx context.Context, isDone string) ([]*model.Todo, error) {
	var todos []*model.Todo
	res := r.db.WithContext(ctx).Find(&todos)
	if err := res.Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *todoRepo) GetById(ctx context.Context, id int64) (*model.Todo, error) {
	todo := new(model.Todo)
	res := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(todo)
	if err := res.Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *todoRepo) Save(ctx context.Context, todo *model.Todo) error {
	res := r.db.WithContext(ctx).Save(todo)
	err := res.Error
	return err
}

func (r *todoRepo) Update(ctx context.Context, id int64) error {
	todo := new(model.Todo)
	res := r.db.WithContext(ctx).
		Model(todo).
		Where("id = ?", id).
		UpdateColumn("is_done", true)
	err := res.Error
	return err
}

func (r *todoRepo) Delete(ctx context.Context, id int64) error {
	todo := new(model.Todo)
	res := r.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(todo)
	err := res.Error
	return err
}
