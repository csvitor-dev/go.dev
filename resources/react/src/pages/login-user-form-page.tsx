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
  loginUserSchema,
  type LoginUserSchema,
} from "@/schemas/login-user-schema";
import { Button } from "@/components/ui/button";
import { AlertCircle, SendHorizonal } from "lucide-react";
import { useState } from "react";
import { Alert, AlertDescription } from "@/components/ui/alert";
import Cookies from "js-cookie";

export default function LoginUserFormPage() {
  const {
    register,
    handleSubmit,
    formState: { errors },
    reset,
  } = useForm<LoginUserSchema>({
    resolver: zodResolver(loginUserSchema),
  });
  const [error, setError] = useState("");

  const submit: SubmitHandler<LoginUserSchema> = async ({
    email,
    password,
  }) => {
    setError("");
    reset();

    const response = await fetch("/auth/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ email, password }),
    });

    const json = await response.json();

    if (response.ok === false) {
      setError(json.error);
      return;
    }
    Cookies.set("auth_token", json.token, { path: "/" });
  };

  return (
    <AuthLayout>
      <Card className="max-w-md w-full mx-auto mt-10">
        <CardHeader>
          <CardTitle className="text-2xl font-bold">Login</CardTitle>
        </CardHeader>

        <CardContent className="space-y-4">
          {error && (
            <Alert className="text-red-600">
              <AlertCircle className="size-4" />
              <AlertDescription>{error}</AlertDescription>
            </Alert>
          )}

          <form onSubmit={handleSubmit(submit)} className="space-y-4">
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

            <Button type="submit" className="w-full mt-4">
              <SendHorizonal className="size-4 mr-2" />
              Submit
            </Button>
          </form>
        </CardContent>

        <CardFooter className="flex flex-col gap-3 mt-6">
          <p className="text-sm text-center text-muted-foreground">
            Still don't have an account?{" "}
            <a
              href="/auth/register"
              className="underline underline-offset-4 hover:text-primary"
            >
              Sign up
            </a>
          </p>
          <p className="text-sm text-center text-muted-foreground">
            Forgot your password?{" "}
            <a
              href="/auth/forgot-password"
              className="underline underline-offset-4 hover:text-primary"
            >
              Recover password
            </a>
          </p>
        </CardFooter>
      </Card>
    </AuthLayout>
  );
}
