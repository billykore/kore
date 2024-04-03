import Todo from "@repo/models/todo"

enum Status {
  OK = "OK"
}

export const getTodos = async (isDone?: boolean) => {
  const res = await fetch(`http://localhost:8000/todos?isDone=${isDone}`, {
    method: "GET",
    cache: "no-store",
  })
  const resJson = await res.json()
  if (resJson.status != Status.OK) {
    throw new Error(resJson.message)
  }
  return resJson.data
}

export const addTodo = async (todo: Todo) => {
  const res = await fetch("http://localhost:8000/todos", {
    method: "POST",
    mode: "cors",
    headers: {"Content-Type": "application/json"},
    body: JSON.stringify(todo)
  })
  const resJson = await res.json()
  if (resJson.status != Status.OK) {
    throw new Error(resJson.message)
  }
  return resJson.data
}

export const editTodo = async (id?: string) => {
  const res = await fetch(`http://localhost:8000/todos/${id}`, {
    method: "PUT",
    mode: "cors"
  })
  const resJson = await res.json()
  if (resJson.status != Status.OK) {
    throw new Error(resJson.message)
  }
  return resJson.data
}

export const deleteTodo = async (id?: string) => {
  const res = await fetch(`http://localhost:8000/todos/${id}`, {
    method: "DELETE",
    mode: "cors"
  })
  const resJson = await res.json()
  if (resJson.status != Status.OK) {
    throw new Error(resJson.message)
  }
  return resJson.data
}
