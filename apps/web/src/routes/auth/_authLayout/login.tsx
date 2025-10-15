import LoginForm from '@/components/auth/LoginForm'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/auth/_authLayout/login')({
  component: RouteComponent,
})

function RouteComponent() {
  return (
    <div className="w-full max-w-6xl mx-auto px-4 p-6 sm:px-6 pt-20 sm:pt-24 relative min-h-screen z-10 flex justify-center items-center gap-8 lg:gap-10">
      <div className="flex-1 space-y-2 text-center lg:text-left w-full max-w-lg mx-auto">
        <div className="flex flex-col justify-center items-center lg:justify-center">
          <div className="relative inline-flex items-center mx-auto">
            <h1 className="text-lg sm:text-xl md:text-2xl font-bold text-primary leading-tight">
              Yuk, Masuk ke Dunia Cerita Bobo!
            </h1>
            <img
              src="/images/core/bobo.png"
              alt="bobo img"
              className="absolute right-[-2.5rem] sm:right-[-2rem] md:right-[-4rem] top-1/2 -translate-y-1/2 w-10 sm:w-14 md:w-20"
            />
          </div>

          <p className="text-muted-foreground text-xs sm:text-sm text-center opacity-80 font-medium leading-relaxed max-w-[95%] mx-auto lg:mx-0">
            Ayo bertualang bersama Bobo dan teman-teman di negeri dongeng penuh warna!  
            Di sini kamu bisa membaca kisah seru, belajar hal-hal baru, dan menemukan  
            keajaiban di setiap halaman cerita.
          </p>
        </div>
        <div className="mt-3 sm:mt-4">
          <LoginForm />
        </div>
      </div>
    </div>
  )
}
