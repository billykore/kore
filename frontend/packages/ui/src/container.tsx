interface ContainerProps {
  children: React.ReactNode
}

export const Container = ({children}: ContainerProps) => {
  return (
    <section className="max-w-2xl py-6 mx-auto">
      {children}
    </section>
  )
}