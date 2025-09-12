import FormField from "@/components/form/form-field";
import {
  registerUserSchema,
  type RegisterUserSchema,
} from "@/schemas/register-user-schema";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm, type SubmitHandler } from "react-hook-form";

export default function RegisterUserForm() {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<RegisterUserSchema>({
    resolver: zodResolver(registerUserSchema),
  });

  const submit: SubmitHandler<RegisterUserSchema> = (data) => {
    console.log(data);
  };

  return (
    <main>
      <div className="max-w-md mx-auto mt-10 p-6 bg-white rounded-2xl shadow-lg">
        <h2 className="text-2xl font-bold mb-6 text-gray-800">
          Create an account
        </h2>

        <form onSubmit={handleSubmit(submit)} className="flex flex-col gap-4">
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

          <button
            type="submit"
            className="mt-4 w-full bg-indigo-600 text-white font-semibold py-2 px-4 rounded-lg transition hover:bg-indigo-700 hover:cursor-pointer"
          >
            Submit
          </button>
        </form>

        <div className="mt-4 text-center">
          <a
            href="/auth/login"
            className="text-sm text-indigo-500 hover:underline"
          >
            Already have an account?
          </a>
        </div>
      </div>
    </main>
  );
}
