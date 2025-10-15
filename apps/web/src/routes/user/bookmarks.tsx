import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/user/bookmarks')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/user/bookmarks"!</div>
}
