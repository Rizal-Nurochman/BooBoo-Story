import ForgotPasswordForm from '@/components/auth/ForgotPasswordForm'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/auth/_authLayout/forgot-password')({
  component: RouteComponent,
})

function RouteComponent() {
  return (
    <div className="w-full max-w-6xl px-8 mx-auto pt-28 relative z-10 min-h-screen flex items-center">
      <div className="flex-1 space-y-3 text-center w-full max-w-lg mx-auto">
        <div className="flex justify-center items-center">
          <div className="relative inline-flex items-center">
            <h1 className="text-md md:text-xl font-bold text-primary">
              Ups! Lupa Sandi? Bobo Datang Menolong!
            </h1>
            <img
              src="/images/core/bobo.png"
              alt="bobo img"
              className="absolute right-[-2.5rem] md:right-[-4rem] top-1/2 -translate-y-1/2 w-14 md:w-20"
            />
          </div>
        </div>

        <p className="text-muted-foreground text-xs opacity-80 md:Text-sm font-medium max-w-[95%] mx-auto">
          Tenang aja, petualanganmu belum berakhir!  
          Masukkan email kamu di bawah ini, dan Bobo akan membantu  
          mengirimkan tautan untuk mengatur ulang kata sandimu. 🌻
        </p>

        <ForgotPasswordForm />
      </div>
    </div>
  )
}
