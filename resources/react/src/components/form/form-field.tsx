import type {
  FieldError,
  FieldValues,
  Path,
  UseFormRegister,
} from "react-hook-form";

interface Props<T extends FieldValues> {
  type: string;
  target: Path<T>;
  register: UseFormRegister<T>;
  children: string;
  error?: FieldError | undefined;
}

export default function FormField<T extends FieldValues>({
  type,
  target,
  register,
  children,
  error,
}: Props<T>) {
  return (
    <>
      <label htmlFor={target} className="text-sm font-medium text-gray-700">
        {children}
      </label>
      <input
        id={target}
        type={type}
        className="w-full rounded-lg border border-gray-300 p-2 focus:ring-2 focus:ring-blue-500 outline-none"
        {...register(target)}
      />
      {error && (
        <span className="text-red-600 text-sm bg-red-100 p-2 rounded-lg">
          {error.message}
        </span>
      )}
    </>
  );
}
