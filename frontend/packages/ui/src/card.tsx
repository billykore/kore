interface CardProps {
  title: string
  body: string
}

export const Card = ({title, body}: CardProps) => {
  return (
    <div key={title} className="block p-6 bg-white border border-gray-200 rounded-lg my-2">
      <h5 className="mb-2 font-bold tracking-tight text-gray-900">
        {title}
      </h5>
      <p className="font-normal text-xl text-gray-700">
        {body}
      </p>
    </div>
  )
}