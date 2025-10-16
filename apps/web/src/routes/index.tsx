import { createFileRoute } from '@tanstack/react-router'
import FeatureSection from '@/components/home/FeatureSection'
import HeroSection from '@/components/home/HeroSection'
import HighlightSection from '@/components/home/HighlightSection'
import ReadSection from '@/components/home/ReadSection'
import FaqSection from '@/components/home/FaqSection'
import AuthorSection from '@/components/home/AuthorSection'

export const Route = createFileRoute('/')({
  component: App,
  ssr: false,
})

function App() {
  return (
    <div className="w-full mx-auto overflow-x-hidden">
      <HeroSection />
      <FeatureSection />
      <HighlightSection />
      <AuthorSection />
      <ReadSection />
      <FaqSection />
    </div>
  )
}
