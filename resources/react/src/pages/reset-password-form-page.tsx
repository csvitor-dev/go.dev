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
import { Button } from "@/components/ui/button";
import { AlertCircle, Lock } from "lucide-react";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { useEffect, useState } from "react";
import {
  resetPasswordSchema,
  type ResetPasswordSchema,
} from "@/schemas/reset-password-schema";

export default function ResetPasswordFormPage() {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<ResetPasswordSchema>({
    resolver: zodResolver(resetPasswordSchema),
  });

  const [error, setError] = useState("");
  const token = new URLSearchParams(document.location.search).get("token");

  useEffect(() => {
    if (token === null) {
      setError("Token has no provided!");
      return;
    }

    if (
      token !==
      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3NTY4NjAzNDAsImp0aSI6ImYxZWQ4YmE4LTZmY2UtNGFlOC04NzZjLTdmMTgwM2IxMDE5NSIsInVzZXJfaWQiOjR9.Q6FziiOIGty5YNIuM6TYNzO0BThUOM85ziwfhyM34a0"
    ) {
      setError("Token has expired or is invalid!");
      return;
    }
    // call action -> API -> validate token
  }, [token]);

  const submit: SubmitHandler<ResetPasswordSchema> = (data) => {
    console.log(data, token);
  };

  return (
    <AuthLayout>
      <Card className="max-w-md w-full">
        <CardHeader>
          <CardTitle className="text-2xl font-bold">Reset password</CardTitle>
        </CardHeader>

        <CardContent className="space-y-4">
          {error && (
            <Alert className="text-red-600">
              <AlertCircle className="h-4 w-4" />
              <AlertDescription>{error}</AlertDescription>
            </Alert>
          )}

          <form onSubmit={handleSubmit(submit)} className="space-y-4">
            <FormField
              type="text"
              target="password"
              register={register}
              error={errors.password}
              disabled={error.length > 0}
            >
              Password
            </FormField>

            <FormField
              type="text"
              target="confirmPassword"
              register={register}
              error={errors.confirmPassword}
              disabled={error.length > 0}
            >
              Confirm password
            </FormField>

            <Button
              type="submit"
              className="w-full mt-4"
              disabled={error.length > 0}
            >
              <Lock className="h-4 w-4 mr-2" />
              Reset
            </Button>
          </form>
        </CardContent>

        <CardFooter>
          <p className="text-xs text-muted-foreground text-center w-full">
            Make sure your new password is strong and unique.
          </p>
        </CardFooter>
      </Card>
    </AuthLayout>
  );
}
