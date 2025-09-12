import z from "zod";

export const loginUserSchema = z.object({
  email: z.email().min(12).max(50),
  password: z.string().min(8).max(25),
});

export type LoginUserSchema = z.infer<typeof loginUserSchema>;
