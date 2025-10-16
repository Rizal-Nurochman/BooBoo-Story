import z from "zod";

export const bookSearchSchema = z.object({
  q: z
    .string()
    .optional()
    .transform((val) => (val && val.trim() !== "" ? val : undefined)),
  page: z
    .union([z.string(), z.number()])
    .optional()
    .transform((val) => (val ? Number(val) : 1))
    .default(1),
  limit: z
    .union([z.string(), z.number()])
    .optional()
    .transform((val) => (val ? Number(val) : 10))
    .default(10),
});
