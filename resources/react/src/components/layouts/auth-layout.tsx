import ThemeToggleButton from "../theme-toggle-button";

interface Props {
  children: React.ReactNode;
}

export default function AuthLayout({ children }: Props) {
  return (
    <>
      <header className="flex justify-end px-2 py-1">
        <ThemeToggleButton />
      </header>
      <main>{children}</main>
    </>
  );
}
