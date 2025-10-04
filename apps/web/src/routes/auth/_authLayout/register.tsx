import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/auth/_authLayout/register')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/auth/register"!</div>
}
