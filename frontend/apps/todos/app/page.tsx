import RootLayout from "@/app/layout"
import { Title } from "@repo/ui/typography"

export default function Home() {
  return (
    <RootLayout>
      <div className="max-w-4xl py-6 mx-auto">
        <Title>Todos</Title>
      </div>
    </RootLayout>
  );
}
