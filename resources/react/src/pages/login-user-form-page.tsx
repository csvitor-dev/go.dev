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
import { SendHorizonal } from "lucide-react";

export default function LoginUserFormPage() {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginUserSchema>({
    resolver: zodResolver(loginUserSchema),
  });

  const submit: SubmitHandler<LoginUserSchema> = (data) => {
    console.log(data);
  };

  return (
    <AuthLayout>
      <Card className="max-w-md w-full mx-auto mt-10">
        <CardHeader>
          <CardTitle className="text-2xl font-bold">Login</CardTitle>
        </CardHeader>

        <form onSubmit={handleSubmit(submit)}>
          <CardContent className="flex flex-col gap-4">
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
          </CardContent>

          <CardFooter className="flex flex-col gap-3 mt-6">
            <Button type="submit" className="w-full">
              <SendHorizonal className="size-4 mr-2" />
              Submit
            </Button>
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
        </form>
      </Card>
    </AuthLayout>
  );
}
