import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/creators/dashboard')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/creators/dashboard"!</div>
}
