import { useState } from "react"
import { useForm } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import { registerSchema, registerSchemaType } from "@/schemas/auth.schema"
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from "../ui/form"
import { Input } from "../ui/input"
import { Button } from "../ui/button"
import { Eye, EyeOff } from "lucide-react"
import { Link } from "@tanstack/react-router"
import GoogleIcon from "../ui/GoogleIcon"

const RegisterForm = () => {
  const [showPassword, setShowPassword] = useState(false)
  const [showConfirmPassword, setShowConfirmPassword] = useState(false)

  const form = useForm<registerSchemaType>({
    resolver: zodResolver(registerSchema),
    defaultValues: {
      username: "",
      email: "",
      password: "",
      confirmPassword: "",
    },
  })

  const fields = [
    {
      name: "username",
      label: "Username",
      placeholder: "Masukkan username kamu",
      type: "text",
      showToggle: false,
    },
    {
      name: "email",
      label: "Email",
      placeholder: "Masukkan email ajaibmu",
      type: "email",
      showToggle: false,
    },
    {
      name: "password",
      label: "Kata Sandi",
      placeholder: "Tulis password rahasiamu",
      type: showPassword ? "text" : "password",
      showToggle: true,
      toggleFn: () => setShowPassword((prev) => !prev),
      isShown: showPassword,
    },
    {
      name: "confirmPassword",
      label: "Konfirmasi Sandi",
      placeholder: "Ulangi password rahasiamu",
      type: showConfirmPassword ? "text" : "password",
      showToggle: true,
      toggleFn: () => setShowConfirmPassword((prev) => !prev),
      isShown: showConfirmPassword,
    },
  ] as const

  function onSubmit(data: registerSchemaType) {
    console.log("Selamat datang, penjelajah dunia digital! 🌈", data)
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
                  <FormLabel className="text-sm sm:text-base font-semibold">{field.label}</FormLabel>
                  <FormControl>
                    <div className="relative">
                      <Input
                        type={field.type}
                        placeholder={field.placeholder}
                        className="py-2 text-sm"
                        {...f}
                      />
                      {field.showToggle && (
                        <button
                          type="button"
                          onClick={field.toggleFn}
                          className="absolute right-3 top-2.5 text-muted-foreground"
                        >
                          {field.isShown ? <EyeOff size={16} /> : <Eye size={16} />}
                        </button>
                      )}
                    </div>
                  </FormControl>
                  <FormMessage className="text-xs" />
                </FormItem>
              )}
            />
          ))}

          <Button type="submit" className="w-full text-sm py-2 cursor-pointer">
            Daftar
          </Button>

          <div className="flex w-full items-center justify-center gap-2 text-xs text-muted-foreground">
            <div className="h-px bg-border flex-1 rounded-full" />
            <span>Atau</span>
            <div className="h-px bg-border flex-1 rounded-full" />
          </div>

          <Button
            type="button"
            variant="outline"
            className="w-full flex items-center justify-center gap-2 text-sm py-2 cursor-pointer"
          >
            <GoogleIcon className="h-4 w-4" />
            Bergabung dengan Google
          </Button>

          <div className="text-center text-xs sm:text-sm text-muted-foreground">
            Sudah punya akun?{" "}
            <Link to="/auth/login" className="text-blue-600 font-medium hover:underline">
              Masuk Sekarang
            </Link>
          </div>
        </form>
      </Form>
    </div>
  )
}

export default RegisterForm
