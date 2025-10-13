import ForgotPasswordForm from '@/components/auth/ForgotPasswordForm'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/auth/_authLayout/forgot-password')({
  component: RouteComponent,
})

function RouteComponent() {
  return (
    <div className="w-full max-w-7xl mx-auto pt-28 relative z-10 flex items-center gap-10">
      <div className="flex-1 space-y-3 text-center">
        <div className="flex justify-center">
          <div className="relative inline-flex items-center">
            <h1 className="text-2xl md:text-3xl font-bold text-primary">
              Ups! Lupa Sandi? Bobo Datang Menolong!
            </h1>
            <img
              src="/images/core/bobo.png"
              alt="bobo img"
              className="absolute right-[-3.5rem] md:right-[-4rem] top-1/2 -translate-y-1/2 w-14 md:w-20"
            />
          </div>
        </div>

        <p className="text-muted-foreground text-base font-medium max-w-[95%] mx-auto">
          Tenang aja, petualanganmu belum berakhir!  
          Masukkan email kamu di bawah ini, dan Bobo akan membantu  
          mengirimkan tautan untuk mengatur ulang kata sandimu. 🌻
        </p>

        <ForgotPasswordForm />
      </div>

      <div className="flex-1 flex justify-center items-center">
        <img src="/images/auth/auth-child.png" alt="auth-child" />
      </div>
    </div>
  )
}
