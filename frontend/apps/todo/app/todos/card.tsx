import { SmallButton } from "@repo/ui/button"

interface CardProps {
  title: string
  body: string
  onDone: React.MouseEventHandler<HTMLButtonElement>
  onDelete: React.MouseEventHandler<HTMLButtonElement>
  isDone?: boolean
}

export const Card = ({title, body, isDone, onDone, onDelete}: CardProps) => {
  return (
    <div className="block p-6 bg-white border border-gray-200 rounded-lg my-2">
      <h5 className="font-bold tracking-tight text-gray-900">
        {title}
      </h5>
      {isDone && <span className="bg-green-100 text-green-800 text-xs font-medium me-2 px-2.5 py-0.5 rounded">
        Done
      </span>}
      <p className="font-normal mt-3 text-xl text-gray-700">
        {body}
      </p>
      <div className="flex mt-10">
        {!isDone && <div>
          <SmallButton type="success" onClick={onDone}>Done</SmallButton>
        </div>}
        <div className="ml-1">
          <SmallButton type="warning" onClick={onDelete}>Delete</SmallButton>
        </div>
      </div>
    </div>
  )
}
