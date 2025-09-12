import type { ThemeProviderState } from "@/types/theme";
import { createContext } from "react";

const initialState: ThemeProviderState = {
  theme: "system",
  setTheme: () => {},
};

export const ThemeContext = createContext<ThemeProviderState>(initialState);
