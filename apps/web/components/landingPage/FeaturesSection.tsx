import { Dancing_Script } from 'next/font/google'
import React from 'react'


const fontThin=Dancing_Script({
    subsets:['latin'],
    weight:['400', '700',],
})

const informations = [
  {
    text: "#funny",
    bgColor: "bg-purple-200",
    textColor: "text-black"
  },
  {
    text: "#enjoy",
    bgColor: "bg-yellow-300",
    textColor: "text-black"
  },
  {
    text: "#happy",
    bgColor: "bg-purple-600",
    textColor: "text-white"
  }
]

const InformationView = () => {
  return (
    <div className="w-56 h-64 rounded-[9999px] border-[12px] border-[#C7B7A3]/70 -rotate-45">
        {informations.map((info, idx) => (
            <div
            key={idx}
            className={`${info.bgColor} ${info.textColor} px-8 font-semibold py-3 rounded-full mt-6 shadow w-fit ${(idx+1)%2==0 && 'ml-[80%]'} rotate-45`}
            >
                {info.text}
            </div>
        ))}
    </div>
  )
}



const features = [
  {
    title: "Kuis Seru",
    desc: "Uji pemahamanmu dengan kuis singkat yang menyenangkan!",
    bg: "bg-purple-200",
    img1: "/images/icons/quiz-icon.svg",
    img2: "/images/shapes/circle-rings.svg",
  },
  {
    title: "Aktivitas Kreatif",
    desc: "Temukan aktivitas menyenangkan seperti mewarnai, membuat kerajinan, dan eksperimen sains.",
    bg: "bg-purple-500",
    img1: "/images/icons/idea-icon.svg",
    img2: "/images/shapes/wave-shape.svg",
  },
  {
    title: "Belajar dengan Game",
    desc: "Pelajari sesuatu yang baru sambil anak-anak bermain game!",
    bg: "bg-yellow-300",
    img1: "/images/icons/game-icon.svg",
    img2: "/images/shapes/dots-pattern.svg",
  },
]



const FeaturesSection = () => {
  return (
    <div className='w-full max-w-7xl mx-auto space-y-6'>
        <div className="flex justify-between items-center">
            <div className={`space-y-4 ${fontThin.className} font-extrabold`}>
                <p className='text-black/80 italic text-6xl font-medium'>Berbagai Fitur</p>
                <p className='text-violet-600 italic text-6xl font-medium'>Menarik</p>
            </div>
            <InformationView />
        </div>
        <div>
           halo
        </div>
    </div>
  )
}

export default FeaturesSection;
