interface CardProps {
  id: React.Key
  children: string
}

export const Card = ({id, children}: CardProps) => {
  return (
    <div key={id} className="block p-6 bg-white border border-gray-200 rounded-lg my-2">
      <p className="font-normal text-xl text-gray-700">
        {children}
      </p>
    </div>
  )
}