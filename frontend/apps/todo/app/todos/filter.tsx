export default function Filter() {
  return (
    <div className="mb-2 inline-flex rounded-md shadow-sm" role="group">
      <button
        type="button"
        className="px-4 py-2 text-sm font-medium text-gray-900 bg-white border border-gray-200 rounded-h-lg
          hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-2 focus:ring-blue-700
          focus:text-blue-700"
      >
        All
      </button>
      <button
        type="button"
        className="px-4 py-2 text-sm font-medium text-gray-900 bg-white border-t border-b border-gray-200
          hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-2 focus:ring-blue-700
          focus:text-blue-700"
      >
        Done
      </button>
      <button
        type="button"
        className="px-4 py-2 text-sm font-medium text-gray-900 bg-white border border-gray-200
          rounded-e-lg hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-2
          focus:ring-blue-700 focus:text-blue-700"
      >
        Not Done
      </button>
    </div>
  )
}