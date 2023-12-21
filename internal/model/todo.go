package model

import "reflect"

type Todo struct {
	Id          string
	Title       string
	Description string
	IsDone      bool
}

type Query struct {
	Key   string
	Value any
}

func (q *Query) IsEmpty() bool {
	return reflect.ValueOf(q.Value).IsZero()
}
