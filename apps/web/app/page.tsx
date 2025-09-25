import FeaturesSection from '@/components/landingPage/FeaturesSection'
import HeroSection from '@/components/landingPage/HeroSection'
import React from 'react'

const page = () => {
  return (
    <div className='w-full pt-12 pb-8 px-6'>
      <HeroSection />
      <FeaturesSection />
    </div>
  )
}

export default page