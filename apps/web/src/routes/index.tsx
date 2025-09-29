import { createFileRoute } from '@tanstack/react-router'
import FeatureSection from '@/components/home/FeatureSection'
import HeroSection from '@/components/home/HeroSection'
import HighlightSection from '@/components/home/HighlightSection'

export const Route = createFileRoute('/')({
  component: App,
})

function App() {
  return (
    <div className="w-full mx-auto ">
      <HeroSection />
      <FeatureSection />
      <HighlightSection />
    </div>
  )
}
