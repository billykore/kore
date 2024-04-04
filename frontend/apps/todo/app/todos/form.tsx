"use client"

import { useState } from "react"
import { Input } from "@repo/ui/input"
import { TextArea } from "@repo/ui/textarea"
import { PrimaryButton } from "@repo/ui/button"
import { addTodo } from "@/app/todos/api"
import { handleError } from "@repo/helper/error"

export default function Form() {
  const [title, setTitle] = useState<string>("")
  const [description, setDescription] = useState<string>("")

  const saveTodo = async () => {
    try {
      const res = await addTodo({title: title, description: description})
      alert(res.message)
      setTitle("")
      setDescription("")
    } catch (err) {
      alert(handleError(err))
    }
  }

  return (
    <>
      <Input
        id="name"
        label="Title"
        placeholder="Todo name"
        value={title}
        onChange={(e) => setTitle(e.target.value)}
      />
      <TextArea
        id="desc"
        label="Description"
        placeholder="Todo description..."
        value={description}
        onChange={(e) => setDescription(e.target.value)}
      />
      <PrimaryButton onClick={() => saveTodo()}>
        Add
      </PrimaryButton>
    </>
  )
}