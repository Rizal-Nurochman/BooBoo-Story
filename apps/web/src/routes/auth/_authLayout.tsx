import HeaderAuth from '@/components/layout/HeaderAuth'
import { createFileRoute, Outlet } from '@tanstack/react-router'

export const Route = createFileRoute('/auth/_authLayout')({
  component: AuthLayout,
})

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
      <HeaderAuth />
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
    </div>
  )
}
