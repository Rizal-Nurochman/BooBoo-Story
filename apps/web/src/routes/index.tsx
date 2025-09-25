import { createFileRoute } from '@tanstack/react-router'
import HeroSection from '~/components/home/HeroSection'

export const Route = createFileRoute('/')({
  component: App,
})

function App() {
  return (
    <div className="w-full max-w-6xl mx-auto ">
      <HeroSection />
    </div>
  )
}
