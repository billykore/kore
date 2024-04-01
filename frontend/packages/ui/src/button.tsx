interface ButtonProps {
  children: string
  onClick?:  React.MouseEventHandler<HTMLButtonElement>
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
