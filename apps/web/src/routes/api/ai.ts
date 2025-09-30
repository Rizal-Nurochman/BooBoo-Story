import { createFileRoute } from '@tanstack/react-router'


export const Route = createFileRoute('/api/ai')({
  server: {
    handlers: {
      GET: () => {
        return new Response(JSON.stringify(['Alice', 'Bob', 'Charlie']), {
          headers: {
            'Content-Type': 'application/json',
          },
        })
      },
    },
  },
})