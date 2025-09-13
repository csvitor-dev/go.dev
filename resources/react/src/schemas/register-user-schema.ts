import z from "zod";

export const registerUserSchema = z
  .object({
    name: z.string().min(3).max(50),
    nickname: z.string().min(3).max(50),
    email: z.email().min(12).max(50),
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

export type RegisterUserSchema = z.infer<typeof registerUserSchema>;
