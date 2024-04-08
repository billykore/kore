interface TypographyProps {
  children: React.ReactNode
}

export const Title = ({children}: TypographyProps) => {
  return (
    <h1 className="text-4xl mt-4 mb-12 font-bold">{children}</h1>
  )
}

export const BackgroundText = ({children}: TypographyProps) => {
  return (
    <h1 className="mt-12 text-2xl text-center text-gray-400">{children}</h1>
  )
}