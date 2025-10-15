import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/user/my-stories')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/user/my-stories"!</div>
}
