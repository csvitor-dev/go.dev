import z from "zod";

export const forgotPasswordSchema = z.object({
  email: z.email().min(12).max(50),
});

export type ForgotPasswordSchema = z.infer<typeof forgotPasswordSchema>;
