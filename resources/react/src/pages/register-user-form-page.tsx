import { zodResolver } from "@hookform/resolvers/zod";
import { useForm, type SubmitHandler } from "react-hook-form";
import FormField from "@/components/form/form-field";
import {
  Card,
  CardHeader,
  CardTitle,
  CardContent,
  CardFooter,
} from "@/components/ui/card";
import AuthLayout from "@/components/layouts/auth-layout";
import {
  registerUserSchema,
  type RegisterUserSchema,
} from "@/schemas/register-user-schema";
import { Button } from "@/components/ui/button";

export default function RegisterUserFormPage() {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<RegisterUserSchema>({
    resolver: zodResolver(registerUserSchema),
  });

  const submit: SubmitHandler<RegisterUserSchema> = async (data) => {
    const { email, name, nickname, password } = data;

    const response = await fetch("/auth/register", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ email, name, nickname, password }),
    });
    const json = await response.json();

    console.log(json);
  };

  return (
    <AuthLayout>
      <Card className="max-w-md mx-auto mt-10">
        <CardHeader>
          <CardTitle className="text-2xl font-bold">
            Create an account
          </CardTitle>
        </CardHeader>

        <form onSubmit={handleSubmit(submit)}>
          <CardContent className="flex flex-col gap-4">
            <FormField
              type="text"
              target="name"
              register={register}
              error={errors.name}
            >
              Name
            </FormField>
            <FormField
              type="text"
              target="nickname"
              register={register}
              error={errors.nickname}
            >
              Nickname
            </FormField>
            <FormField
              type="email"
              target="email"
              register={register}
              error={errors.email}
            >
              Email
            </FormField>
            <FormField
              type="password"
              target="password"
              register={register}
              error={errors.password}
            >
              Password
            </FormField>
            <FormField
              type="password"
              target="confirmPassword"
              register={register}
              error={errors.confirmPassword}
            >
              Confirm password
            </FormField>
          </CardContent>

          <CardFooter className="flex flex-col gap-3 mt-6">
            <Button type="submit" className="w-full">
              Submit
            </Button>
            <p className="text-sm text-center text-muted-foreground">
              Already have an account?{" "}
              <a
                href="/auth/login"
                className="underline underline-offset-4 hover:text-primary"
              >
                Login
              </a>
            </p>
          </CardFooter>
        </form>
      </Card>
    </AuthLayout>
  );
}
