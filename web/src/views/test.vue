<template>
  <div class="card">
    <p>当前主题：{{ theme }}</p>
    <select v-model="selected" @change="setTheme(selected)">
      <option value="light">浅色</option>
      <option value="dark">深色</option>
      <option value="solarized">Solarized</option>
      <option value="dracula">Dracula</option>
    </select>
  </div>
</template>

<script setup lang="ts">
import { inject, ref, watch } from 'vue'
import { ThemeSymbol, Theme } from '@/theme-context'

const themeContext = inject(ThemeSymbol)
if (!themeContext) throw new Error('Theme context not provided')

const { theme, setTheme } = themeContext
const selected = ref<Theme>(theme.value)

watch(theme, (val) => {
  selected.value = val
})
</script>
