// theme-context.ts
import { InjectionKey, Ref } from 'vue'

export type Theme = 'light' | 'dark' | 'solarized' | 'dracula'

export interface ThemeContext {
  theme: Ref<Theme>
  toggleTheme: () => void
  setTheme: (t: Theme) => void
}

export const ThemeSymbol: InjectionKey<ThemeContext> = Symbol('ThemeContext')
