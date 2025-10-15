import { createFileRoute, Outlet } from '@tanstack/react-router'

export const Route = createFileRoute('/auth/_authLayout')({
  component: AuthLayout,
})


const borders=[
    {
        id:1,
        img:'/images/core/border.png',
        style:'w-[30%] md:w-[15%] absolute top-0 right-0 z-20'
    },
    {
        id:1,
        img:'/images/core/border.png',
        style:'w-[30%] md:w-[15%] rotate-[180deg] absolute bottom-0 left-0 z-20'
    },
]

function AuthLayout() {
    const leaves = Array.from({ length: 6 }).map((_) => ({
      src: '/images/core/leaf.png',
      left: `${Math.random() * 100}%`,
      size: `${Math.floor(Math.random() * 16) + 12}`,
      duration: `${Math.random() * 5 + 6}s`,
      delay: `${Math.random() * 5}s`,
      rotate: Math.random() > 0.5 ? '30deg' : '-25deg',
    }))

  return (
    <div className="relative min-h-screen overflow-hidden bg-[#fef7e6]">
      <Outlet />
      {leaves.map((leaf, i) => (
        <img
          key={i}
          src={leaf.src}
          alt={`leaf-${i}`}
          className="absolute animate-fall"
          style={{
            top: '-60px',
            left: leaf.left,
            width: `${leaf.size}px`,
            animationDuration: leaf.duration,
            animationDelay: leaf.delay,
            transform: `rotate(${leaf.rotate})`,
            opacity: 0.9,
          }}
        />
      ))}

      {borders.map((img)=>(
          <img src={img.img} className={img.style} key={img.id}  />
      ))}

      <div className="absolute inset-0 bg-gradient-to-b from-transparent via-[#5b7652]/10 to-[#3f5c43]/30 pointer-events-none blur-2xl" />
    </div>
  )
}
