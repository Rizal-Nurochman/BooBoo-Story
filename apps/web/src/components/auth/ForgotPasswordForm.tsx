import { useState } from "react"
import { useForm } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import { z } from "zod"
import { motion } from "motion/react"
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from "../ui/form"
import { Input } from "../ui/input"
import { Button } from "../ui/button"
import { Eye, EyeOff, Loader2 } from "lucide-react"
import { Link } from "@tanstack/react-router"

const emailSchema = z.object({
  email: z.string().email({ message: "Email tidak valid" }),
})

const otpSchema = z.object({
  otp: z.string().min(6, "Kode OTP harus 6 digit").max(6, 'Kode OTP harusnya hanya 6 digit'),
})

const resetPasswordSchema = z
  .object({
    password: z.string().min(6, "Password minimal 6 karakter"),
    confirmPassword: z.string().min(6, "Konfirmasi password minimal 6 karakter"),
  })
  .refine((data) => data.password === data.confirmPassword, {
    message: "Password dan konfirmasi tidak cocok",
    path: ["confirmPassword"],
  })

const ForgotPasswordForm = () => {
  const [step, setStep] = useState(1)
  const [showPassword, setShowPassword] = useState(false)
  const [showConfirmPassword, setShowConfirmPassword] = useState(false)

  const emailForm = useForm({
    resolver: zodResolver(emailSchema),
    defaultValues: { email: "" },
  })

  const otpForm = useForm({
    resolver: zodResolver(otpSchema),
    defaultValues: { otp: "" },
  })

  const resetForm = useForm({
    resolver: zodResolver(resetPasswordSchema),
    defaultValues: { password: "", confirmPassword: "" },
  })

  const handleEmailSubmit = async (data: z.infer<typeof emailSchema>) => {
    await new Promise((r) => setTimeout(r, 1200))
    console.log("📩 Email dikirim:", data)
    setStep(2)
  }

  const handleOtpSubmit = async (data: z.infer<typeof otpSchema>) => {
    await new Promise((r) => setTimeout(r, 1200))
    console.log("✅ OTP benar:", data)
    setStep(3)
  }

  const handleResetSubmit = async (data: z.infer<typeof resetPasswordSchema>) => {
    await new Promise((r) => setTimeout(r, 1200))
    console.log("🔑 Password berhasil diubah:", data)
    alert("Password kamu telah diperbarui!")
  }

  // 🔥 Helper motion button
  const MotionButton = motion(Button)

  return (
    <div className="w-full cursor-pointer space-y-6">
      <h2 className="text-2xl font-semibold text-center mb-4">Lupa Password</h2>

      {step === 1 && (
        <Form {...emailForm}>
          <form onSubmit={emailForm.handleSubmit(handleEmailSubmit)} className="space-y-6">
            <FormField
              control={emailForm.control}
              name="email"
              render={({ field }) => (
                <FormItem>
                  <FormLabel className="text-lg font-medium">Alamat Email</FormLabel>
                  <FormControl>
                    <Input type="email" placeholder="Masukkan email terdaftar kamu" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <MotionButton
              whileTap={{ scale: 0.92 }}
              type="submit"
              className="w-full cursor-pointer"
              disabled={emailForm.formState.isSubmitting}
            >
              {emailForm.formState.isSubmitting ? (
                <>
                  <Loader2 className="animate-spin mr-2 h-4 w-4" />
                  Mengirim...
                </>
              ) : (
                "Kirim OTP"
              )}
            </MotionButton>
          </form>
        </Form>
      )}

      {step === 2 && (
        <Form {...otpForm}>
          <form onSubmit={otpForm.handleSubmit(handleOtpSubmit)} className="space-y-6">
            <FormField
              control={otpForm.control}
              name="otp"
              render={({ field }) => (
                <FormItem>
                  <FormLabel className="text-lg font-medium">Kode OTP</FormLabel>
                  <FormControl>
                    <Input type="text" placeholder="Masukkan kode OTP" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <MotionButton
              whileTap={{ scale: 0.92 }}
              type="submit"
              className="w-full cursor-pointer"
              disabled={otpForm.formState.isSubmitting}
            >
              {otpForm.formState.isSubmitting ? (
                <>
                  <Loader2 className="animate-spin mr-2 h-4 w-4" />
                  Memverifikasi...
                </>
              ) : (
                "Verifikasi OTP"
              )}
            </MotionButton>

            <MotionButton
              whileTap={{ scale: 0.92 }}
              type="button"
              variant="ghost"
              onClick={() => setStep(1)}
              className="w-full cursor-pointer text-muted-foreground"
              disabled={otpForm.formState.isSubmitting}
            >
              Kembali ke Email
            </MotionButton>
          </form>
        </Form>
      )}

      {step === 3 && (
        <Form {...resetForm}>
          <form onSubmit={resetForm.handleSubmit(handleResetSubmit)} className="space-y-6">
            <FormField
              control={resetForm.control}
              name="password"
              render={({ field }) => (
                <FormItem>
                  <FormLabel className="text-lg font-medium">Password Baru</FormLabel>
                  <FormControl>
                    <div className="relative">
                      <Input
                        type={showPassword ? "text" : "password"}
                        placeholder="Masukkan password baru"
                        {...field}
                      />
                      <button
                        type="button"
                        onClick={() => setShowPassword((p) => !p)}
                        className="absolute right-3 top-2.5 text-muted-foreground"
                      >
                        {showPassword ? <EyeOff size={18} /> : <Eye size={18} />}
                      </button>
                    </div>
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={resetForm.control}
              name="confirmPassword"
              render={({ field }) => (
                <FormItem>
                  <FormLabel className="text-lg font-medium">Konfirmasi Password Baru</FormLabel>
                  <FormControl>
                    <div className="relative">
                      <Input
                        type={showConfirmPassword ? "text" : "password"}
                        placeholder="Ulangi password baru"
                        {...field}
                      />
                      <button
                        type="button"
                        onClick={() => setShowConfirmPassword((p) => !p)}
                        className="absolute right-3 top-2.5 text-muted-foreground"
                      >
                        {showConfirmPassword ? <EyeOff size={18} /> : <Eye size={18} />}
                      </button>
                    </div>
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <MotionButton
              whileTap={{ scale: 0.92 }}
              type="submit"
              className="w-full cursor-pointer"
              disabled={resetForm.formState.isSubmitting}
            >
              {resetForm.formState.isSubmitting ? (
                <>
                  <Loader2 className="animate-spin mr-2 h-4 w-4" />
                  Menyimpan...
                </>
              ) : (
                "Simpan Password Baru"
              )}
            </MotionButton>
          </form>
        </Form>
      )}

      <div className="text-center text-sm text-muted-foreground">
        Oh, ternyata sudah ingat password kamu?{" "}
        <Link to="/auth/login" className="text-blue-600 font-medium hover:underline">
          Coba Masuk Lagi
        </Link>
      </div>
    </div>
  )
}

export default ForgotPasswordForm
