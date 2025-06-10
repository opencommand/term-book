<template>
  <div class="code-editor-container" ref="editorContainer"></div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch, onBeforeUnmount, nextTick } from 'vue'
import * as monaco from 'monaco-editor'
import { registerCustomThemes } from '../monaco-themes'

const props = defineProps({
  modelValue: {
    type: String,
    default: ''
  },
  language: {
    type: String,
    default: 'javascript'
  },
  readOnly: {
    type: Boolean,
    default: false
  },
  theme: {
    type: String,
    default: 'light'
  },
  options: {
    type: Object,
    default: () => ({})
  }
})

const emit = defineEmits(['update:modelValue', 'save', 'run'])

const editorContainer = ref<HTMLElement | null>(null)
let editor: monaco.editor.IStandaloneCodeEditor | null = null
let isEditorUpdating = false
let themesRegistered = false // 添加标志位记录主题是否已注册

const updateHeight = throttle(() => {
  if (!editor || !editorContainer.value) return;
  const contentHeight = editor.getContentHeight();
  editorContainer.value.style.height = `${contentHeight}px`;
  editor.layout();
}, 100);

// 节流实现
function throttle<T extends (...args: any[]) => void>(fn: T, delay: number): T {
  let lastCall = 0;
  return ((...args: any[]) => {
    const now = Date.now();
    if (now - lastCall >= delay) {
      fn(...args);
      lastCall = now;
    }
  }) as T;
}

const resizeObserver = new ResizeObserver(() => {
  editor?.layout()
})

onMounted(async () => {
  if (!editorContainer.value) return

  // 修改主题注册逻辑
  if (!themesRegistered) {
    registerCustomThemes()
    themesRegistered = true
  }

  try {
    editor = monaco.editor.create(editorContainer.value, {
      value: props.modelValue || '',
      language: props.language,
      theme: props.theme,
      readOnly: props.readOnly,
      automaticLayout: false,
      minimap: { enabled: false },
      fontSize: 14,
      scrollBeyondLastLine: false,
      ...props.options
    })

    resizeObserver.observe(editorContainer.value)

    editor.onDidChangeModelContent(() => {
      if (isEditorUpdating) return

      const value = editor?.getValue() || ''
      emit('update:modelValue', value)
      updateHeight()
    })

    editor.onDidContentSizeChange(updateHeight)

    editor.addCommand(monaco.KeyMod.CtrlCmd | monaco.KeyCode.KeyS, () => {
      emit('save')
    })

    editor.addCommand(monaco.KeyMod.Shift | monaco.KeyCode.Enter, () => {
      emit('run')
    })

    await nextTick()
    updateHeight()
  } catch (error) {
    console.error('Failed to initialize Monaco Editor:', error)
  }
})

onBeforeUnmount(() => {
  resizeObserver.disconnect()
  if (editor) {
    editor.dispose()
    editor = null
  }
})

watch(() => props.modelValue, (newVal) => {
  if (editor && editor.getValue() !== newVal) {
    isEditorUpdating = true
    editor.setValue(newVal || '')
    isEditorUpdating = false
    updateHeight()
  }
})

watch(() => props.theme, (newTheme) => {
  if (editor) {
    monaco.editor.setTheme(newTheme)
  }
})

watch(() => props.language, (newLang) => {
  const model = editor?.getModel()
  if (model) {
    monaco.editor.setModelLanguage(model, newLang)
  }
})

function debounce<T extends (...args: any[]) => void>(fn: T, delay: number): T {
  let timeoutId: number | null = null
  return ((...args: any[]) => {
    if (timeoutId) {
      clearTimeout(timeoutId)
    }
    timeoutId = setTimeout(() => {
      fn(...args)
      timeoutId = null
    }, delay)
  }) as T
}

defineExpose({
  getEditor: () => editor,
  getValue: () => editor?.getValue() || '',
  setValue: (val: string) => {
    if (editor) {
      isEditorUpdating = true
      editor.setValue(val || '')
      isEditorUpdating = false
      updateHeight()
    }
  },
  focus: () => editor?.focus(),
  layout: () => editor?.layout()
})
</script>

<style scoped>
.code-editor-container {
  width: 100%;
  min-height: 50px;
  max-height: 500px;
  border: 1px solid var(--border-color);
  border-radius: 4px;
  overflow: hidden;
  transition: height 0.2s ease;
}
</style>