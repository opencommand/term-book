<template>
  <div :class="theme">
    <slot />
  </div>
</template>
<script setup lang="ts">
import { ref, provide, watchEffect } from 'vue'
import { ThemeSymbol, Theme } from '../theme-context'

const theme = ref<Theme>('light')

// 主题切换函数
const toggleTheme = () => {
  const themes: Theme[] = ['light', 'dark', 'solarized', 'dracula']
  const index = themes.indexOf(theme.value)
  theme.value = themes[(index + 1) % themes.length]
}

const setTheme = (t: Theme) => {
  theme.value = t
}

// 初始化
const stored = localStorage.getItem('theme')
if (stored && ['light', 'dark', 'solarized', 'dracula'].includes(stored)) {
  theme.value = stored as Theme
}

// 更新 DOM class 和本地存储
watchEffect(() => {
  localStorage.setItem('theme', theme.value)
  document.documentElement.className = theme.value
})

provide(ThemeSymbol, {
  theme,
  toggleTheme,
  setTheme,
})
</script>
