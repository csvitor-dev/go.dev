import ThemeToggleButton from "@/components/theme-toggle-button";

interface Props {
  children: React.ReactNode;
}

export default function AuthLayout({ children }: Props) {
  return (
    <div className="flex flex-col min-h-screen">
      <header className="flex justify-end px-2 py-1">
        <ThemeToggleButton />
      </header>
      <main className="flex flex-1 items-center justify-center bg-background px-4">
        {children}
      </main>
    </div>
  );
}
