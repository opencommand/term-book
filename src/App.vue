<template>
  <ThemeProvider>
    <!-- 显示路由内容 -->
    <router-view />
  </ThemeProvider>
</template>

<script setup lang="ts">
import ThemeProvider from './components/ThemeProvider.vue'
import Test from '../src/views/test.vue'
</script>
