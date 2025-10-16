import * as z from "zod"

//schema

export const loginSchema = z.object({
  email: z
    .string()
    .min(1, { message: "Email wajib diisi" })
    .email({ message: "Format email tidak valid" }),
  password: z
    .string()
    .min(6, { message: "Password minimal 6 karakter" })
    .max(100, { message: "Password terlalu panjang" }),
})

export const registerSchema = z
  .object({
    name: z
      .string()
      .min(3, { message: "Username minimal 3 karakter" })
      .max(20, { message: "Username maksimal 20 karakter" })
      .regex(/^[a-zA-Z0-9_]+$/, {
        message: "Username hanya boleh huruf, angka, dan underscore",
      }),
    email: z
      .string()
      .min(1, { message: "Email wajib diisi" })
      .email({ message: "Format email tidak valid" }),
    password: z
      .string()
      .min(6, { message: "Password minimal 6 karakter" })
      .max(100, { message: "Password terlalu panjang" }),
    confirmPassword: z.string().min(1, { message: "Konfirmasi password wajib diisi" }),
  })
  .refine((data) => data.password === data.confirmPassword, {
    message: "Password tidak cocok",
    path: ["confirmPassword"],
  })


// type

export type registerSchemaType=z.infer<typeof registerSchema>
export type LoginSchemaType = z.infer<typeof loginSchema>
