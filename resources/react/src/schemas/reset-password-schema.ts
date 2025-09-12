import z from "zod";

export const resetPasswordSchema = z
  .object({
    password: z.string().min(8).max(25),
    confirmPassword: z.string().min(8).max(25),
  })
  .superRefine(({ password, confirmPassword }, ctx) => {
    if (password !== confirmPassword) {
      ctx.addIssue({
        code: "custom",
        message: "Passwords do not match",
        path: ["confirmPassword"],
      });
    }
  });

export type ResetPasswordSchema = z.infer<typeof resetPasswordSchema>;
