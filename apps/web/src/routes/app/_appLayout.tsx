import Header from '@/components/layout/Header'
import HeaderApp from '@/components/app/HeaderApp'
import { createFileRoute, Outlet } from '@tanstack/react-router'

export const Route = createFileRoute('/app/_appLayout')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>
    <HeaderApp />
    <Outlet />
  </div>
}
