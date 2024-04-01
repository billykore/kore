import { ChangeEventHandler } from "react"

interface TextAreaProps {
  id: string
  label: string
  value?: string
  onChange?: ChangeEventHandler<HTMLTextAreaElement>
}

export const TextArea = ({id, label, value, onChange}: TextAreaProps) => {
  return (
    <div className="mt-4">
      <label
        htmlFor={id}
        className="block mb-2 text-sm font-medium text-gray-900"
      >
        {label}
      </label>
      <textarea
        id={id}
        rows={4}
        className="block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border
            border-gray-300 focus:ring-blue-500 focus:border-blue-500"
        placeholder="Write your thoughts here..."
        value={value}
        onChange={onChange}></textarea>
    </div>
  )
}
