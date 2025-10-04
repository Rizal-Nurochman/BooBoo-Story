import LoginForm from '@/components/auth/LoginForm'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/auth/_authLayout/login')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div className='w-full flex items-center gap-10'>
    <div className="flex-1 space-y-3 text-center">
      <h1 className="text-3xl md:text-4xl font-bold text-primary">
        Yuk, Masuk ke Dunia Cerita Bobo!
      </h1>
      <p className="text-muted-foreground text-base font-medium">
        Ayo bertualang bersama Bobo dan teman-teman di negeri dongeng penuh warna!  
        Di sini kamu bisa membaca kisah seru, belajar hal-hal baru, dan menemukan  
        keajaiban di setiap halaman cerita.
      </p>
      <LoginForm />
    </div>

    <div className="flex-1 flex justify-center items-center">
      <img src='/images/auth/auth-child.png'alt='auth-child' />
    </div>
  </div>
}