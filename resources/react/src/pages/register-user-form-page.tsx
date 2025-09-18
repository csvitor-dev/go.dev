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
import AuthLayout from "@/layouts/auth-layout";
import {
  registerUserSchema,
  type RegisterUserSchema,
} from "@/schemas/register-user-schema";
import { Button } from "@/components/ui/button";
import { AlertCircle, CheckCircle2, SendHorizonal } from "lucide-react";
import { useState } from "react";
import { Alert, AlertDescription } from "@/components/ui/alert";

export default function RegisterUserFormPage() {
  const {
    register,
    handleSubmit,
    formState: { errors },
    reset,
  } = useForm<RegisterUserSchema>({
    resolver: zodResolver(registerUserSchema),
  });
  const [error, setError] = useState("");
  const [success, setSuccess] = useState(false);

  const submit: SubmitHandler<RegisterUserSchema> = async ({
    email,
    name,
    nickname,
    password,
  }) => {
    setError("");
    setSuccess(false);
    reset();

    const response = await fetch("/auth/register", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ email, name, nickname, password }),
    });

    if (response.ok) {
      setSuccess(true);
      return;
    }
    const { error } = (await response.json()) as { error: string };
    setError(error);
  };

  return (
    <AuthLayout>
      <Card className="max-w-md w-full mx-auto mt-10">
        <CardHeader>
          <CardTitle className="text-2xl font-bold">
            Create an account
          </CardTitle>
        </CardHeader>

        <CardContent className="space-y-4">
          {error && (
            <Alert className="text-red-600">
              <AlertCircle className="size-4" />
              <AlertDescription>{error}</AlertDescription>
            </Alert>
          )}

          {success && (
            <Alert className="text-emerald-600">
              <CheckCircle2 className="size-4" />
              <AlertDescription>
                <p>Registered successfully!</p>
                <a
                  href="/auth/login"
                  className="underline underline-offset-4 hover:text-primary"
                >
                  Go to Login!
                </a>
              </AlertDescription>
            </Alert>
          )}

          <form onSubmit={handleSubmit(submit)} className="space-y-4">
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

            <Button type="submit" className="w-full mt-4">
              <SendHorizonal className="size-4 mr-2" />
              Submit
            </Button>
          </form>
        </CardContent>

        <CardFooter className="flex flex-col gap-2">
          <p className="text-sm text-center text-muted-foreground">
            Already have an account?{" "}
            <a
              href="/auth/login"
              className="underline underline-offset-4 hover:text-primary"
            >
              Log in
            </a>
          </p>
        </CardFooter>
      </Card>
    </AuthLayout>
  );
}
