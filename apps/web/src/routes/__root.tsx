import {
  HeadContent,
  Scripts,
  createRootRouteWithContext,
} from '@tanstack/react-router'
import { TanStackRouterDevtoolsPanel } from '@tanstack/react-router-devtools'
import { TanstackDevtools } from '@tanstack/react-devtools'

import Header from '../components/layout/Header'

import TanStackQueryDevtools from '../integrations/tanstack-query/devtools'

import appCss from '../styles.css?url'

import type { QueryClient } from '@tanstack/react-query'

import BackToTop from '@/components/layout/BackToTop'
import NotfoundPage from '@/components/layout/NotfoundPage'
import { Toaster } from '@/components/ui/sonner'
import { AuthService } from '@/services/auth.service'
import { createServerFn } from '@tanstack/react-start'
import { getCookie } from '@tanstack/react-start/server'

const fetchUserLogin=createServerFn({ method:'GET' }).handler(async()=>{
  const cookie = getCookie('access_token') 
  console.log('ck', cookie)
  const res=await AuthService.me(cookie)
  console.log(res)
  return{
    user:res
  }
})
interface MyRouterContext {
  queryClient: QueryClient
}



export const Route = createRootRouteWithContext<MyRouterContext>()({
  beforeLoad : async ()=>{
    const user=fetchUserLogin()
    return {user}
  },
  head: () => ({
    meta: [
      {
        charSet: 'utf-8',
      },
      {
        name: 'viewport',
        content: 'width=device-width, initial-scale=1',
      },
      {
        title: 'Bobo app',
      },
    ],
    links: [
      {
        rel: 'stylesheet',
        href: appCss,
      },
    ],
  }),

  shellComponent: RootDocument,
  notFoundComponent: () => <NotfoundPage />
})

function RootDocument({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <head>
        <HeadContent />
      </head>
      <body>
        <Header />
        <main>
          {children}
        </main>
        <TanstackDevtools
          config={{
            position: 'bottom-left',
          }}
          plugins={[
            {
              name: 'Tanstack Router',
              render: <TanStackRouterDevtoolsPanel />,
            },
            TanStackQueryDevtools,
          ]}
        />
        <BackToTop />
        <Scripts />
        <Toaster position='top-center'  />
      </body>
    </html>
  )
}
