interface ButtonProps {
  children: string
  type?: "success" | "warning"
  onClick?: React.MouseEventHandler<HTMLButtonElement>
}

export const PrimaryButton = ({children, onClick}: ButtonProps) => {
  return (
    <button
      type="button"
      className="mt-4 text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300
        font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2"
      onClick={onClick}
    >
      {children}
    </button>
  )
}

export const SecondaryButton = ({children, onClick}: ButtonProps) => {
  return (
    <button
      type="button"
      className="mt-4 text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100
          focus:ring-4 focus:ring-gray-100 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2"
      onClick={onClick}
    >
      {children}
    </button>
  )
}

export const SmallButton = ({children, type, onClick}: ButtonProps) => {
  return (
    <button
      type="button"
      className={`px-3 py-2 text-xs font-medium text-center bg-blue-100 text-blue-800 rounded-lg
        hover:bg-blue-200 focus:ring-4 focus:outline-none focus:ring-blue-300
        ${type == "success" ? "bg-green-100 text-green-800 hover:bg-blue-200" :
        type == "warning" ? "bg-red-100 text-red-800 hover:bg-red-200" :
          "bg-gray-100 text-gray-800 hover:bg-gray-200"
      }`}
      onClick={onClick}
    >
      {children}
    </button>
  )
}
