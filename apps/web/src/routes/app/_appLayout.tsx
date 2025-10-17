import HeaderApp from '@/components/app/HeaderApp'
import { createFileRoute, Outlet } from '@tanstack/react-router'
import { Client } from '@/components/ui/Client'

export const Route = createFileRoute('/app/_appLayout')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>
    <Client>
      <HeaderApp />
    </Client>
    <Outlet />
  </div>
}
