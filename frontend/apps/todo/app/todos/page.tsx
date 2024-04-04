import { Title } from "@repo/ui/typography"
import { Divider } from "@repo/ui/divider"
import { Container } from "@repo/ui/container"
import TodoList from "@/app/todos/list"
import TodoForm from "@/app/todos/form"

export default async function TodosPage() {
  return (
    <Container>
      <Title>Todos</Title>
      <TodoForm/>
      <Divider/>
      <TodoList/>
    </Container>
  )
}