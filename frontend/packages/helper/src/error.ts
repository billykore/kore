export const handleError = (err: any) => {
  if (err instanceof Error) {
    return err.message
  }
  return "Ops! Look's like something goes wrong."
}