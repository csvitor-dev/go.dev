import type {
  FieldError,
  FieldValues,
  Path,
  UseFormRegister,
} from "react-hook-form";
import { Input } from "@/components/ui/input";

interface Props<T extends FieldValues> {
  type: string;
  target: Path<T>;
  register: UseFormRegister<T>;
  children: string;
  error?: FieldError | undefined;
  disabled?: boolean | undefined;
}

export default function FormField<T extends FieldValues>({
  type,
  target,
  register,
  children,
  error,
  disabled = false,
}: Props<T>) {
  return (
    <>
      <label className="block text-sm font-medium mb-1">{children}</label>
      <Input
        id={target}
        type={type}
        {...register(target)}
        disabled={disabled}
      />
      {error && (
        <span className="text-sm text-red-500 mt-1">{error.message}</span>
      )}
    </>
  );
}
