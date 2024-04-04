interface CardProps {
  title: string
  body: string
}

export default function Card({title, body}: CardProps) {
  return (
    <div className="block p-6 bg-white border border-gray-200 rounded-lg my-2">
      <h5 className="font-bold tracking-tight text-gray-900">
        {title}
      </h5>
      <p className="font-normal mt-3 text-xl text-gray-700">
        {body}
      </p>
    </div>
  )
}