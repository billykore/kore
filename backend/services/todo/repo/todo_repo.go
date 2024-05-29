package repo

import (
	"context"

	"github.com/billykore/kore/backend/pkg/model"
	"gorm.io/gorm"
)

type TodoRepo struct {
	db *gorm.DB
}

func NewTodoRepository(postgres *gorm.DB) *TodoRepo {
	return &TodoRepo{db: postgres}
}

func (r *TodoRepo) List(ctx context.Context, isDone string) ([]*model.Todo, error) {
	var todos []*model.Todo
	res := r.db.WithContext(ctx).Find(&todos)
	if err := res.Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *TodoRepo) GetById(ctx context.Context, id int64) (*model.Todo, error) {
	todo := new(model.Todo)
	res := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(todo)
	if err := res.Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *TodoRepo) Save(ctx context.Context, todo *model.Todo) error {
	res := r.db.WithContext(ctx).Save(todo)
	err := res.Error
	return err
}

func (r *TodoRepo) Update(ctx context.Context, id int64) error {
	todo := new(model.Todo)
	res := r.db.WithContext(ctx).
		Model(todo).
		Where("id = ?", id).
		UpdateColumn("is_done", true)
	err := res.Error
	return err
}

func (r *TodoRepo) Delete(ctx context.Context, id int64) error {
	todo := new(model.Todo)
	res := r.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(todo)
	err := res.Error
	return err
}
