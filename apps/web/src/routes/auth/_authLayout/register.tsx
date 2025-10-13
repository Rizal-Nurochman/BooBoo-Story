import RegisterForm from '@/components/auth/RegisterForm'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/auth/_authLayout/register')({
  component: RouteComponent,
})

function RouteComponent() {
  return (
    <div className="w-full max-w-7xl flex items-center mx-auto pt-28 relative z-10 gap-10">
      <div className="flex-1 flex justify-center items-center">
        <img src="/images/auth/auth-child.png" alt="auth-child" />
      </div>
      <div className="flex-1 space-y-3 text-center">
        <div className="flex justify-center">
          <div className="relative inline-flex items-center">
            <h1 className="text-2xl md:text-3xl font-bold text-primary">
              Mulai Petualanganmu Bersama Bobo!
            </h1>
            <img
              src="/images/core/bobo.png"
              alt="bobo img"
              className="absolute right-[-3.5rem] md:right-[-4rem] top-1/2 -translate-y-1/2 w-14 md:w-20"
            />
          </div>
        </div>

        <p className="text-muted-foreground text-base font-medium">
          Yuk, jadi bagian dari dunia penuh cerita dan warna!  
          Dengan bergabung bersama Bobo, kamu bisa membaca kisah seru,  
          belajar hal-hal baru, dan menjelajahi dunia imajinasi yang tak terbatas!
        </p>

        <RegisterForm />
      </div>
    </div>
  )
}
