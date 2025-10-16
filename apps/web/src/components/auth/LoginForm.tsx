import { useState } from "react"
import { useForm } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import { loginSchema, LoginSchemaType } from "@/schemas/auth.schema"
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from "../ui/form"
import { Input } from "../ui/input"
import { Button } from "../ui/button"
import { Eye, EyeOff } from "lucide-react"
import { Link } from "@tanstack/react-router"
import GoogleIcon from "../ui/GoogleIcon"
import { AuthService } from "@/services/auth.service"
import { toast } from "sonner"
import { sleep } from "@/lib/api"

const LoginForm = () => {
  const [showPassword, setShowPassword] = useState(false)

  const form = useForm<LoginSchemaType>({
    resolver: zodResolver(loginSchema),
    defaultValues: { email: "", password: "" },
  })

  const fields = [
    {
      name: "email",
      label: "Alamat Email",
      placeholder: "Masukkan email kamu",
      type: "email",
      showToggle: false,
    },
    {
      name: "password",
      label: "Kata Sandi",
      placeholder: "Masukkan password kamu",
      type: showPassword ? "text" : "password",
      showToggle: true,
    },
  ] as const

    async function onSubmit (data: LoginSchemaType) {
      await sleep();
     const res= await AuthService.login(data)
     if (res.status === 'error') {
      toast.error(res.message || 'Gagal masuk. Silakan coba lagi.')
     }else{
      toast.success('Berhasil masuk. Selamat datang kembali!')
     }
  }

  return (
    <div className="w-full">
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-5 w-full">
          {fields.map((field) => (
            <FormField
              key={field.name}
              control={form.control}
              name={field.name}
              render={({ field: f }) => (
                <FormItem className="text-start">
                  {field.type === "password" ? (
                    <div className="flex justify-between items-center">
                      <FormLabel className="text-sm font-semibold">{field.label}</FormLabel>
                      <Link
                        to="/auth/forgot-password"
                        className="text-xs text-blue-600 hover:underline"
                      >
                        Lupa sandi?
                      </Link>
                    </div>
                  ) : (
                    <FormLabel className="text-sm font-semibold">{field.label}</FormLabel>
                  )}
                  <FormControl>
                    <div className="relative">
                      <Input
                        type={field.type}
                        placeholder={field.placeholder}
                        className="py-2 text-sm"
                        {...f}
                      />
                      {field?.showToggle && (
                        <button
                          type="button"
                          onClick={() => setShowPassword((prev) => !prev)}
                          className="absolute right-3 top-2 text-muted-foreground"
                        >
                          {showPassword ? <EyeOff size={16} /> : <Eye size={16} />}
                        </button>
                      )}
                    </div>
                  </FormControl>
                  <FormMessage className="text-xs" />
                </FormItem>
              )}
            />
          ))}

          <Button disabled={form.formState.isSubmitting} type="submit" className="w-full py-2 text-sm font-medium cursor-pointer disabled:cursor-not-allowed">
            {form.formState.isSubmitting ? "Memproses..." : "Masuk"}
          </Button>

          <div className="flex w-full items-center justify-center gap-2 text-xs text-muted-foreground">
            <div className="h-px bg-border flex-1 rounded-full" />
            <span>Atau</span>
            <div className="h-px bg-border flex-1 rounded-full" />
          </div>

          <Button
            type="button"
            variant="outline"
            className="w-full flex items-center justify-center gap-2 py-2 text-sm cursor-pointer"
          >
            <GoogleIcon className="h-3.5 w-3.5" />
            Masuk dengan Google
          </Button>

          <div className="text-center text-xs text-muted-foreground">
            Belum punya akun?{" "}
            <Link to="/auth/register" className="text-blue-600 font-medium hover:underline">
              Daftar Sekarang
            </Link>
          </div>
        </form>
      </Form>
    </div>
  )
}

export default LoginForm
