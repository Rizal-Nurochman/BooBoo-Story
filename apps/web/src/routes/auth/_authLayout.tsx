import HeaderAuth from '@/components/layout/HeaderAuth'
import { createFileRoute, Outlet } from '@tanstack/react-router'

export const Route = createFileRoute('/auth/_authLayout')({
  component: AuthLayout,
})

function AuthLayout() {
  return (
    <div>
      <HeaderAuth />
      <div className='w-full max-w-7xl mx-auto pt-28'>
         <Outlet /> 
      </div>
    </div>
  )
}
