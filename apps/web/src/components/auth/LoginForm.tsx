import { useState } from "react"
import { useForm } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import { loginSchema, LoginSchemaType } from "@/schemas/auth.schema"
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from "../ui/form"
import { Input } from "../ui/input"
import { Button } from "../ui/button"
import { Eye, EyeOff } from "lucide-react"
import { Link } from "@tanstack/react-router"

const GoogleIcon = (props: React.SVGProps<SVGSVGElement>) => (
  <svg {...props} viewBox="0 0 24 24" fill="none">
    <path
      d="M22.48 12.27c0-.79-.07-1.54-.19-2.27H12v4.54h5.88a5 5 0 0 1-2.17 3.28v2.72h3.5c2.05-1.89 3.27-4.69 3.27-8.27Z"
      fill="#4285F4"
    />
    <path
      d="M12 23c2.97 0 5.46-.98 7.27-2.68l-3.5-2.72c-.98.65-2.23 1.04-3.77 1.04-2.9 0-5.36-1.96-6.24-4.6H2.13v2.82C3.93 20.53 7.64 23 12 23Z"
      fill="#34A853"
    />
    <path
      d="M5.76 14.04a7.05 7.05 0 0 1 0-4.08V7.14H2.13a10.96 10.96 0 0 0 0 9.72l3.63-2.82Z"
      fill="#FBBC05"
    />
    <path
      d="M12 4.75c1.62 0 3.07.56 4.22 1.67l3.14-3.14C17.46 1.26 14.97.25 12 .25 7.64.25 3.93 2.72 2.13 6.28l3.63 2.82C6.64 6.71 9.1 4.75 12 4.75Z"
      fill="#EA4335"
    />
  </svg>
)

const LoginForm = () => {
  const [showPassword, setShowPassword] = useState(false)

  const form = useForm<LoginSchemaType>({
    resolver: zodResolver(loginSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  })

  const fields = [
    {
      name: "email",
      label: "Alamat Email",
      placeholder: "Masukkan email ajaibmu",
      type: "email",
      showToggle: false,
    },
    {
      name: "password",
      label: "Kata Sandi Rahasia",
      placeholder: "Tulis password rahasiamu",
      type: showPassword ? "text" : "password",
      showToggle: true,
    },
  ] as const

  function onSubmit(data: LoginSchemaType) {
    console.log("Selamat datang, penjelajah dunia digital! 🌈", data)
  }

  return (
    <div>
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8 w-full">
          {fields.map((field) => (
            <FormField
              key={field.name}
              control={form.control}
              name={field.name}
              render={({ field: f }) => (
                <FormItem>
                  {field.type==='password'? (
                    <div className="flex justify-between items-center">
                        <FormLabel className="text-xl font-semibold">{field.label}</FormLabel>
                        <Link to="/auth/forgot-password" className="text-sm text-blue-600 hover:underline">
                            Lupa kata sandi?
                        </Link>
                    </div>
                  ):
                    <FormLabel className="text-xl font-semibold">{field.label}</FormLabel>
                  }
                  <FormControl>
                    <div className="relative">
                      <Input
                        type={field.type}
                        placeholder={field.placeholder}
                        className="py-2.5"
                        {...f}
                      />
                      {field?.showToggle && (
                        <button
                          type="button"
                          onClick={() => setShowPassword((prev) => !prev)}
                          className="absolute right-3 top-2.5 text-muted-foreground"
                        >
                          {showPassword ? <EyeOff size={18} /> : <Eye size={18} />}
                        </button>
                      )}
                    </div>
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          ))}

          <Button type="submit" className="w-full cursor-pointer">Masuk</Button>

          <div className="flex w-full items-center justify-center gap-2 text-sm text-muted-foreground">
            <div className="h-px bg-border flex-1 rounded-full" />
            <span>Atau</span>
            <div className="h-px bg-border flex-1 rounded-full" />
          </div>

          <Button
            type="button"
            variant="outline"
            className="w-full flex cursor-pointer items-center justify-center gap-2"
          >
            <GoogleIcon className="h-4 w-4" />
            Masuk dengan Google
          </Button>

          <div className="text-center text-sm text-muted-foreground">
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
