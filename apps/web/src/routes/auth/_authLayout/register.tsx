import RegisterForm from '@/components/auth/RegisterForm'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/auth/_authLayout/register')({
  component: RouteComponent,
})

function RouteComponent() {
  return (
    <div className="w-full max-w-6xl flex items-center min-h-screen justify-center mx-auto pt-8 relative z-10 gap-10 px-4">
      <div className="flex-1 space-y-3 text-center w-full max-w-lg mx-auto pb-6 h-full">
        <div className="flex justify-center items-center">
          <div className="relative inline-flex items-center">
            <h1 className="text-lg sm:text-xl md:text-2xl font-bold text-primary">
              Mulai Petualanganmu Bersama Bobo!
            </h1>
            <img
              src="/images/core/bobo.png"
              alt="bobo img"
              className="absolute right-[-3.5rem] md:right-[-4rem] top-1/2 -translate-y-1/2 w-14 md:w-20"
            />
          </div>
        </div>

        <p className="text-muted-foreground text-xs sm:text-sm text-center opacity-80 font-medium leading-relaxed max-w-[95%] mx-auto lg:mx-0">
          Yuk, jadi bagian dari dunia penuh cerita dan warna!  
          Dengan bergabung bersama Bobo, kamu bisa membaca kisah seru,  
          belajar hal-hal baru, dan menjelajahi dunia imajinasi yang tak terbatas!
        </p>

        <RegisterForm />
      </div>
    </div>
  )
}
