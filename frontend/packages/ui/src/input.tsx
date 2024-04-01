import { ChangeEventHandler } from "react"

interface InputProps {
  id: string
  label: string
  value?: string
  onChange?: ChangeEventHandler<HTMLInputElement>
}

export const Input = ({id, label, value, onChange}: InputProps) => {
  return (
    <div className="mt-4">
      <label
        htmlFor={id}
        className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
      >
        {label}
      </label>
      <input
        type="text"
        id={id}
        value={value}
        className="bg-gray-50 border border-gray-300 text-gray-900 text-sm 
          rounded-lg focus:ring-blue-500 focus:border-blue-500 block 
          w-full p-2.5 "
        placeholder="Jonny Sins"
        onChange={onChange}
      />
    </div>
  )
}