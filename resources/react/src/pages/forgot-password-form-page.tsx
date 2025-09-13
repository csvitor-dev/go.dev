import { zodResolver } from "@hookform/resolvers/zod";
import { useForm, type SubmitHandler } from "react-hook-form";
import FormField from "@/components/form/form-field";
import {
  Card,
  CardHeader,
  CardTitle,
  CardContent,
  CardFooter,
  CardDescription,
} from "@/components/ui/card";
import AuthLayout from "@/layouts/auth-layout";
import { Button } from "@/components/ui/button";
import { AlertCircle, CheckCircle2, Mail } from "lucide-react";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { useState } from "react";
import {
  forgotPasswordSchema,
  type ForgotPasswordSchema,
} from "@/schemas/forgot-password-schema";

export default function ForgotPasswordFormPage() {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<ForgotPasswordSchema>({
    resolver: zodResolver(forgotPasswordSchema),
  });
  const [error, setError] = useState("");
  const [success, setSuccess] = useState(false);

  const submit: SubmitHandler<ForgotPasswordSchema> = ({ email }) => {
    setError("");
    setSuccess(false);

    if (email === "example@gmail.com") {
      setError("Fail: this email is invalid!");
      return;
    }
    setSuccess(true);
  };

  return (
    <AuthLayout>
      <Card className="max-w-md w-full">
        <CardHeader>
          <CardTitle className="text-2xl font-bold">Recover Password</CardTitle>
          <CardDescription>
            Enter your email. If we find an account, we will send you
            instructions.
          </CardDescription>
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
                If the email exists, you will receive the instructions shortly.
              </AlertDescription>
            </Alert>
          )}

          <form className="space-y-4" onSubmit={handleSubmit(submit)}>
            <FormField
              type="email"
              target="email"
              register={register}
              error={errors.email}
            >
              Email
            </FormField>

            <Button type="submit" className="w-full mt-4">
              <Mail className="size-4 mr-2" />
              Send instructions
            </Button>
          </form>
        </CardContent>

        <CardFooter>
          <p className="text-xs text-muted-foreground text-center w-full">
            Make sure to check your spam folder as well.
          </p>
        </CardFooter>
      </Card>
    </AuthLayout>
  );
}
