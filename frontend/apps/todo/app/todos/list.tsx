"use client"

import Todo from "@repo/models/todo"
import { Card } from "@/app/todos/card"
import { getTodos, deleteTodo, editTodo } from "@/app/todos/api"
import { handleError } from "@repo/helper/error"
import Filter from "@/app/todos/filter"

export default async function TodoList() {
  const todos: Todo[] = await getTodos()

  const onDone = async (id?: string) => {
    try {
      const res = await editTodo(id)
      alert(res.message)
    } catch (err) {
      alert(handleError(err))
    }
  }

  const onDelete = async (id?: string) => {
    try {
      const res = await deleteTodo(id)
      alert(res.message)
    } catch (err) {
      alert(handleError(err))
    }
  }

  return (
    <>
      <Filter/>

      {todos.map(todo => (
        <Card
          key={todo.id}
          title={todo.title}
          body={todo.description}
          isDone={todo.isDone}
          onDone={() => onDone(todo.id)}
          onDelete={() => onDelete(todo.id)}
        />
      ))}
    </>
  )
}
