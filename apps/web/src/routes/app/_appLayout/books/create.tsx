import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/app/_appLayout/books/create')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/app/_appLayout/books/create"!</div>
}
