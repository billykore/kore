import { SmallButton } from "@repo/ui/button"
import { useRouter } from "next/navigation"

interface CardProps {
  title: string
  body: string
  onDone: () => void
  onDelete: () => void
  isDone?: boolean
}

export const Card = ({title, body, isDone, onDone, onDelete}: CardProps) => {
  const router = useRouter()

  const doneFn = () => {
    onDone()
    router.refresh()
  }

  const deleteFn = () => {
    onDelete()
    router.refresh()
  }

  return (
    <div className="block p-6 bg-white border border-gray-200 rounded-lg my-2">
      <h5 className="font-bold tracking-tight text-gray-900">
        {title}
      </h5>
      {isDone && <span
        className="bg-green-100 text-green-800 text-xs font-medium me-2 px-2.5 py-0.5 rounded">
        Done
      </span>}
      <p className="font-normal mt-3 text-xl text-gray-700">
        {body}
      </p>
      <div className="flex mt-10">
        {!isDone && <div>
          <SmallButton type="success" onClick={() => doneFn()}>Done</SmallButton>
        </div>}
        <div className="ml-1">
          <SmallButton type="warning" onClick={() => deleteFn()}>Delete</SmallButton>
        </div>
      </div>
    </div>
  )
}
