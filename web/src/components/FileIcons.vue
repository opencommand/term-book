<template>
  <!-- IPYNB 文件图标 -->
  <svg v-if="type === 'ipynb'" t="1749549621309" class="icon" viewBox="0 0 1024 1024" version="1.1"
    xmlns="http://www.w3.org/2000/svg" p-id="19113" width="16" height="16" fill="currentColor">
    <path
      d="M93.1072 422.2976h28.672v174.592H93.1072v-174.592z m116.224 148.3776v92.5696h-28.7744V422.2976h28.8256v26.4704c6.0416-10.3936 13.6192-18.0736 22.7328-23.04 9.2672-5.12 20.2752-7.68 33.0752-7.68 21.1968 0 38.4 8.448 51.5584 25.2928 13.312 16.8448 19.968 38.912 19.968 66.4064 0 27.392-6.656 49.5616-19.968 66.3552a62.464 62.464 0 0 1-51.5584 25.2928c-12.8 0-23.808-2.56-33.0752-7.5264a60.416 60.416 0 0 1-22.7328-23.1936z m97.6384-60.928c0-21.0944-4.4032-37.632-13.1072-49.5616a41.472 41.472 0 0 0-35.6864-18.1248 41.8816 41.8816 0 0 0-35.84 18.1248c-8.6528 11.9296-12.9536 28.4672-12.9536 49.5616 0 21.0944 4.3008 37.632 12.9536 49.664a41.984 41.984 0 0 0 35.84 17.92 41.6256 41.6256 0 0 0 35.6864-17.92c8.704-12.032 13.1072-28.5696 13.1072-49.664z m149.9136 103.3216c-8.0896 20.7872-16.0256 34.304-23.7056 40.6528-7.68 6.3488-17.92 9.5232-30.8224 9.5232h-22.9376v-24.0128h16.8448a27.392 27.392 0 0 0 18.3808-5.632c4.352-3.6864 9.216-12.544 14.4896-26.4704l5.12-13.1072-70.5536-171.7248h30.3616l54.5792 136.4992 54.528-136.4992h30.4128l-76.6976 190.7712z m261.376-121.6v105.3696h-28.672v-104.448c0-16.4864-3.2256-28.8256-9.6768-37.0688-6.4512-8.192-16.0768-12.288-28.9792-12.288-15.4624 0-27.648 4.9152-36.608 14.7968-8.96 9.8816-13.4144 23.296-13.4144 40.3456v98.6624h-28.8256v-174.592h28.8256v27.136c6.8608-10.4448 14.8992-18.3296 24.1664-23.552 9.3184-5.12 20.0704-7.7312 32.256-7.7312 20.0192 0 35.2256 6.1952 45.5168 18.688 10.24 12.3392 15.4112 30.6176 15.4112 54.6816z m182.784 18.2784c0-21.0944-4.3008-37.632-13.056-49.5616a41.472 41.472 0 0 0-35.6864-18.1248 41.8816 41.8816 0 0 0-35.84 18.1248c-8.6528 11.9296-12.9536 28.4672-12.9536 49.5616 0 21.0944 4.3008 37.632 12.9536 49.664a41.984 41.984 0 0 0 35.84 17.92 41.6256 41.6256 0 0 0 35.6864-17.92c8.704-12.032 13.1072-28.5696 13.1072-49.664z m-97.536-60.928c6.0416-10.4448 13.6192-18.1248 22.784-23.0912 9.216-5.12 20.224-7.68 33.024-7.68 21.1968 0 38.4 8.448 51.6096 25.2928 13.312 16.8448 19.968 38.912 19.968 66.4064 0 27.392-6.656 49.5616-19.968 66.3552a62.464 62.464 0 0 1-51.6096 25.2928 68.608 68.608 0 0 1-33.024-7.5264 60.416 60.416 0 0 1-22.784-23.1936v26.1632h-28.8256V354.3552h28.8256v94.4128z"
      p-id="19114"></path>
    <path
      d="M507.6736 82.2272a413.696 413.696 0 0 1 388.9152 272.9984c-75.776-88.8832-221.6448-148.8896-388.9152-148.8896-167.2192 0-312.832 60.2624-388.608 148.8896a413.0816 413.0816 0 0 1 388.608-272.9984z m-344.4736 7.424a40.0896 40.0896 0 1 1 0 80.1792 40.0896 40.0896 0 0 1 0-80.1792zM788.096 0a54.6304 54.6304 0 1 1 0 109.312 54.6304 54.6304 0 0 1 0-109.312zM228.0704 892.2112a65.8944 65.8944 0 1 1 0 131.7888 65.8944 65.8944 0 0 1 0-131.7888z m668.5184-256.6144a413.696 413.696 0 0 1-388.9152 272.9984 413.696 413.696 0 0 1-388.9152-272.9984c75.8784 88.8832 221.696 148.8896 388.9152 148.8896 167.2704 0 313.088-60.0064 388.9152-148.8896z"
      p-id="19115"></path>
  </svg>
  <!-- <svg v-if="type === 'ipynb'" class="file-icon" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
    <path
      d="M19 3H5C3.89543 3 3 3.89543 3 5V19C3 20.1046 3.89543 21 5 21H19C20.1046 21 21 20.1046 21 19V5C21 3.89543 20.1046 3 19 3Z"
      stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
    <path d="M8 12H16" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
    <path d="M12 8V16" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
  </svg> -->

  <!-- Python 文件图标 -->
  <svg v-else-if="type === 'py'" t="1749549874788" class="icon" viewBox="0 0 1024 1024" version="1.1"
    xmlns="http://www.w3.org/2000/svg" p-id="21303" width="16" height="16" fill="currentColor">
    <path
      d="M943.58718 401c-15.4-61.8-44.6-108.4-106.8-108.4h-80.2v94.8c0 73.6-62.4 135.6-133.6 135.6H409.38718c-58.4 0-106.8 50-106.8 108.6v203.6c0 58 50.4 92 106.8 108.6 67.6 19.8 132.6 23.4 213.6 0 53.8-15.6 106.8-47 106.8-108.6v-81.4H516.38718v-27.2h320.4c62.2 0 85.2-43.4 106.8-108.4 22.4-67 21.4-131.4 0-217.2zM636.38718 808c22.2 0 40.2 18.2 40.2 40.6 0 22.6-18 40.8-40.2 40.8-22 0-40.2-18.4-40.2-40.8 0.2-22.6 18.2-40.6 40.2-40.6zM399.58718 496.2h213.6c59.4 0 106.8-49 106.8-108.6V183.8c0-58-48.8-101.4-106.8-111.2-71.6-11.8-149.4-11.2-213.6 0.2-90.4 16-106.8 49.4-106.8 111.2v81.4h213.8v27.2h-294c-62.2 0-116.6 37.4-133.6 108.4-19.6 81.4-20.4 132.2 0 217.2 15.2 63.2 51.4 108.4 113.6 108.4H265.98718v-97.6c0-70.6 61-132.8 133.6-132.8z m-13.4-285.2c-22.2 0-40.2-18.2-40.2-40.6 0.2-22.6 18-40.8 40.2-40.8 22 0 40.2 18.4 40.2 40.8s-18 40.6-40.2 40.6z"
      p-id="21304"></path>
  </svg>

  <!-- <svg v-else-if="type === 'py'" class="file-icon" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
    <path d="M12 2C6.48 2 2 6.48 2 12C2 17.52 6.48 22 12 22C17.52 22 22 17.52 22 12C22 6.48 17.52 2 12 2Z"
      fill="currentColor" />
    <path d="M12 6V18C16.42 18 20 14.42 20 10C20 5.58 16.42 2 12 2Z" fill="currentColor" opacity="0.5" />
  </svg> -->

  <!-- JavaScript/TypeScript 文件图标 -->
  <svg v-else-if="type === 'js' || type === 'ts'" class="file-icon" viewBox="0 0 24 24" fill="none"
    xmlns="http://www.w3.org/2000/svg">
    <path d="M3 3H21V21H3V3Z" fill="currentColor" />
    <path d="M12 12V18C12 19.6569 13.3431 21 15 21H17" stroke="white" stroke-width="2" stroke-linecap="round" />
    <path d="M12 12V6C12 4.34315 10.6569 3 9 3H7" stroke="white" stroke-width="2" stroke-linecap="round" />
  </svg>

  <!-- JSON 文件图标 -->
  <svg v-else-if="type === 'json'" class="file-icon" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
    <path
      d="M21 15V19C21 20.1046 20.1046 21 19 21H5C3.89543 21 3 20.1046 3 19V5C3 3.89543 3.89543 3 5 3H19C20.1046 3 21 3.89543 21 5V9"
      stroke="currentColor" stroke-width="2" stroke-linecap="round" />
    <path d="M7 8L9 12L7 16" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
    <path d="M17 8L15 12L17 16" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
    <path d="M13 7L11 17" stroke="currentColor" stroke-width="2" stroke-linecap="round" />
  </svg>

  <!-- Markdown 文件图标 -->
  <svg v-else-if="type === 'md'" class="file-icon" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
    <path
      d="M21 15V19C21 20.1046 20.1046 21 19 21H5C3.89543 21 3 20.1046 3 19V5C3 3.89543 3.89543 3 5 3H19C20.1046 3 21 3.89543 21 5V9"
      stroke="currentColor" stroke-width="2" stroke-linecap="round" />
    <path d="M7 15V9L9 12L11 9V15" stroke="currentColor" stroke-width="2" stroke-linecap="round"
      stroke-linejoin="round" />
    <path d="M13 9H15C16.1046 9 17 9.89543 17 11V13C17 14.1046 16.1046 15 15 15H13V9Z" stroke="currentColor"
      stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
  </svg>

  <!-- 图片文件图标 -->
  <svg v-else-if="type === 'jpg' || type === 'png' || type === 'gif'" class="file-icon" viewBox="0 0 24 24" fill="none"
    xmlns="http://www.w3.org/2000/svg">
    <path d="M4 16L8.58579 11.4142C9.36683 10.6332 10.6332 10.6332 11.4142 11.4142L16 16" stroke="currentColor"
      stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
    <path d="M14 14L15.5858 12.4142C16.3668 11.6332 17.6332 11.6332 18.4142 12.4142L20 14" stroke="currentColor"
      stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
    <path
      d="M3 20H21C21.5523 20 22 19.5523 22 19V5C22 4.44772 21.5523 4 21 4H3C2.44772 4 2 4.44772 2 5V19C2 19.5523 2.44772 20 3 20Z"
      stroke="currentColor" stroke-width="2" />
    <path d="M9 9C9.55228 9 10 8.55228 10 8C10 7.44772 9.55228 7 9 7C8.44772 7 8 7.44772 8 8C8 8.55228 8.44772 9 9 9Z"
      fill="currentColor" />
  </svg>

  <!-- 默认文件图标 -->
  <svg v-else class="file-icon" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
    <path d="M14 2H6C4.89543 2 4 2.89543 4 4V20C4 21.1046 4.89543 22 6 22H18C19.1046 22 20 21.1046 20 20V8L14 2Z"
      stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
    <path d="M14 2V8H20" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
  </svg>
</template>

<script setup lang="ts">
defineProps({
  type: {
    type: String,
    required: true
  }
});
</script>

<style scoped>
.file-icon {
  width: 16px;
  height: 16px;
  margin-right: 4px;
  flex-shrink: 0;
}
</style>