export type Theme = "light" | "dark" | "system";

export type ThemeProviderState = {
  theme: Theme;
  setTheme: (theme: Theme) => void;
};
