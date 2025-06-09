<template>
  <div :class="['vscode-container', theme]" v-if="!loading">
    <!-- 活动栏 -->
    <div class="activity-bar">
      <div class="activity-bar-items">
        <button class="activity-item active" title="文件资源管理器">
          <span class="icon">📁</span>
        </button>
        <button class="activity-item" title="搜索">
          <span class="icon">🔍</span>
        </button>
        <button class="activity-item" title="源代码管理">
          <span class="icon">🔃</span>
        </button>
        <button class="activity-item" title="运行和调试">
          <span class="icon">▶</span>
        </button>
        <button class="activity-item" title="扩展">
          <span class="icon">🧩</span>
        </button>
      </div>
      <div class="theme-selector">
        <select v-model="selected" @change="setTheme(selected)">
          <option value="light">浅色</option>
          <option value="dark">深色</option>
          <option value="solarized">Solarized</option>
          <option value="dracula">Dracula</option>
        </select>
      </div>
    </div>

    <!-- 侧边栏 -->
    <div class="sidebar">
      <div class="sidebar-header">
        <h3>NOTEBOOK.ipynb</h3>
      </div>
      <div class="sidebar-content">
        <div class="outline-item" v-for="(cell, index) in cells" :key="cell.id" @click="scrollToCell(index)">
          <span class="outline-icon">[ ]</span>
          <span class="outline-text">单元格 {{ index + 1 }}</span>
        </div>
      </div>
    </div>

    <!-- 编辑器区域 -->
    <div class="editor-container">
      <!-- 状态栏 -->
      <div class="status-bar">
        <div class="status-bar-left">
          <span class="status-item">main</span>
          <span class="status-item">UTF-8</span>
          <span class="status-item">Python 3.9.7</span>
        </div>
        <div class="status-bar-right">
          <span class="status-item">Ln 1, Col 1</span>
          <span class="status-item">Spaces: 4</span>
        </div>
      </div>

      <!-- 编辑器内容 -->
      <!-- 编辑器内容 -->
      <div class="editor-content">
        <div class="cell code-cell" v-for="(cell, index) in cells" :key="cell.id" ref="cellElements">
          <div class="cell-toolbar">
            <span class="cell-index">In [{{ cell.executionCount || ' ' }}]:</span>
            <div class="cell-actions">
              <button @click="executeCell(index)" class="run-button" title="运行单元格" :disabled="cell.isRunning">
                <span class="icon">{{ cell.isRunning ? '⏳' : '▶' }}</span>
              </button>
              <button @click="addCell(index)" class="icon-button" title="添加单元格" :disabled="cell.isRunning">+</button>
              <button @click="removeCell(index)" class="icon-button" title="删除单元格" :disabled="cell.isRunning">×</button>
            </div>
          </div>
          <div class="code-editor">
            <textarea v-model="cell.content" class="code-input"
              :placeholder="index === 0 ? '输入代码并按Shift+Enter运行' : '输入代码...'" :disabled="cell.isRunning"></textarea>
          </div>
          <div class="cell-output" v-if="cell.output || cell.isRunning">
            <div class="output-content">
              <pre>{{ cell.output }}</pre>
              <div class="execution-info">
                <div v-if="cell.isRunning" class="execution-progress">
                  <span class="progress-icon">⏳</span>
                  <span>正在运行... {{ cell.currentExecutionTime }}ms</span>
                  <progress :value="cell.progress" max="100"></progress>
                </div>
                <div v-if="!cell.isRunning && cell.executionTime" class="execution-time">
                  <span class="time-icon">⏱️</span>
                  <span>执行时间: {{ cell.executionTime }}ms</span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- 添加新单元格按钮 -->
        <div class="add-cell-container">
          <button @click="addCell(cells.length)" class="add-cell-button">
            + 添加代码单元格
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { inject, ref, watch, nextTick, onMounted } from 'vue'
import { ThemeSymbol, Theme } from '../../theme-context'

import { getFileListApi } from '../../api/Document.ts'

const themeContext = inject(ThemeSymbol)
if (!themeContext) throw new Error('Theme context not provided')

const { theme, setTheme } = themeContext
const selected = ref<Theme>(theme.value)

watch(theme, (val) => {
  selected.value = val
})

const loading = ref(false)
// 单元格数据
interface Cell {
  id: string;
  content: string;
  output?: string;
  executionCount?: number;
  executionTime?: number;
  isRunning?: boolean;
  currentExecutionTime?: number;
  progress?: number;
  timer?: number;
}

const cells = ref<Cell[]>([
  { id: generateId(), content: "print('Hello, World!')", output: "Hello, World!", executionCount: 1 },
  { id: generateId(), content: "import numpy as np\n\nx = np.arange(10)\nx", executionCount: 2 },
  { id: generateId(), content: "# 在这里输入你的代码", output: "" }
])

const cellElements = ref<HTMLElement[]>([])

function generateId(): string {
  return Math.random().toString(36).substring(2, 9)
}

function addCell(index: number) {
  cells.value.splice(index + 1, 0, {
    id: generateId(),
    content: "",
    output: "",
    isRunning: false
  });

  nextTick(() => {
    if (cellElements.value[index + 1]) {
      cellElements.value[index + 1].scrollIntoView({ behavior: 'smooth' });
    }
  });
}

function removeCell(index: number) {
  // 如果单元格正在运行，清除定时器
  if (cells.value[index].timer) {
    clearInterval(cells.value[index].timer);
  }
  if (cells.value.length > 1) {
    cells.value.splice(index, 1);
  }
}

function executeCell(index: number) {
  const cell = cells.value[index];

  // 如果已经在运行，则不做任何操作
  if (cell.isRunning) return;

  // 初始化运行状态
  cell.isRunning = true;
  cell.currentExecutionTime = 0;
  cell.progress = 0;
  cell.output = "正在执行...";

  // 模拟10秒执行过程
  const startTime = Date.now();
  const totalDuration = 10000; // 10秒

  // 更新计时器和进度
  cell.timer = window.setInterval(() => {
    const elapsed = Date.now() - startTime;
    cell.currentExecutionTime = elapsed;
    cell.progress = Math.min(100, (elapsed / totalDuration) * 100);

    // 模拟执行过程中的输出变化
    if (elapsed > 3000 && !cell.output?.includes("加载数据...")) {
      cell.output = "正在执行...\n加载数据...";
    }
    if (elapsed > 6000 && !cell.output?.includes("处理中")) {
      cell.output = "正在执行...\n加载数据...\n处理中...";
    }

    // 执行完成
    if (elapsed >= totalDuration) {
      clearInterval(cell.timer);
      cell.isRunning = false;
      cell.executionTime = elapsed;
      cell.executionCount = (cell.executionCount || 0) + 1;

      // 模拟执行结果
      const code = cell.content;
      if (code.includes("print(")) {
        cell.output = code.match(/print\((.*)\)/)?.[1] || "执行完成";
      } else if (code.includes("np.arange")) {
        cell.output = "array([0, 1, 2, 3, 4, 5, 6, 7, 8, 9])";
      } else {
        cell.output = "执行完成 (模拟输出)";
      }
    }
  }, 100);
}

function scrollToCell(index: number) {
  if (cellElements.value[index]) {
    cellElements.value[index].scrollIntoView({ behavior: 'smooth' })
  }
}

const loadPageData = async () => {
  loading.value = true;
  try {
    const [fileRes] = await Promise.all([
      getFileListApi(),
      // getUserInfoApi(),
      // getPermissionListApi()
    ]);

    // 处理文件列表
    console.log(fileRes, 'sdadas');

    // if (fileRes.success) {
    //   // fileList.value = fileRes.data || [];
    //   console.log(fileRes);

    // } else {
    //   console.warn('获取文件列表失败:', fileRes.message);
    // }

    // // 示例：处理用户信息
    // if (userRes.success) {
    //   userInfo.value = userRes.data;
    // } else {
    //   console.warn('获取用户信息失败:', userRes.message);
    // }

    // // 示例：处理权限列表
    // if (permRes.success) {
    //   permissions.value = permRes.data || [];
    // } else {
    //   console.warn('获取权限失败:', permRes.message);
    // }

  } catch (err) {
    console.error('页面加载失败:', err);
    // errorMessage.value = '加载失败，请稍后再试';
  } finally {
    loading.value = false;
  }
};

onMounted(() => {
  loadPageData();
})
</script>

<style scoped>
.vscode-container {
  display: flex;
  height: 100vh;
  width: 100vw;
  overflow: hidden;
}

/* 活动栏样式 */
.activity-bar {
  width: 48px;
  background-color: var(--card-bg);
  border-right: 1px solid var(--border-color);
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
  padding: 10px 0;
}

.activity-bar-items {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 12px;
}



.activity-item {
  width: 48px;
  height: 48px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: none;
  border: none;
  cursor: pointer;
  color: var(--text-color);
  opacity: 0.6;
  transition: all 0.2s;
  border-radius: 4px;
  box-shadow: none;
}


.activity-item:hover {
  opacity: 0.8;
  background-color: var(--hover-bg);
}

.activity-item.active {
  opacity: 1;
  background-color: var(--highlight-color);
  border-left: 5px solid var(--border-color);
}

.theme-selector {
  /* font-size: 10px; */
  /* padding: 10px; */
  /* padding-left: 0; */
}

.theme-selector select {
  width: 100%;
  background-color: var(--input-bg);
  color: var(--text-color);
  border: 1px solid var(--border-color);
  border-radius: 4px;
  padding: 4px;
  font-size: 12px;
}

/* 侧边栏样式 */
.sidebar {
  width: 250px;
  background-color: var(--card-bg);
  border-right: 1px solid var(--border-color);
  display: flex;
  flex-direction: column;
}

.sidebar-header {
  padding: 12px;
  border-bottom: 1px solid var(--border-color);
  font-weight: bold;
}

.sidebar-content {
  flex: 1;
  overflow-y: auto;
  padding: 8px 0;
}

.outline-item {
  padding: 6px 12px;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
}

.outline-item:hover {
  background-color: var(--hover-bg);
}

.outline-icon {
  color: var(--accent-color);
  font-family: monospace;
}

/* 编辑器区域样式 */
.editor-container {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.editor-content {
  flex: 1;
  overflow-y: auto;
  padding: 20px;
}

/* 状态栏样式 */
.status-bar {
  height: 22px;
  background-color: var(--card-bg);
  border-top: 1px solid var(--border-color);
  display: flex;
  justify-content: space-between;
  padding: 0 10px;
  font-size: 12px;
  color: var(--text-color);
}

.status-item {
  margin-left: 8px;
  opacity: 0.8;
}

/* 单元格样式 */
.cell {
  border-radius: 4px;
  overflow: hidden;
  margin-bottom: 20px;
  background-color: var(--card-bg);
  border: 1px solid var(--border-color);
}

.code-cell {
  position: relative;
}

.cell-toolbar {
  display: flex;
  align-items: center;
  padding: 4px 8px;
  background-color: var(--hover-bg);
  border-bottom: 1px solid var(--border-color);
}

.cell-index {
  font-family: monospace;
  font-size: 13px;
  color: var(--text-color);
  opacity: 0.7;
  margin-right: auto;
}

.cell-actions {
  display: flex;
  gap: 4px;
}

.run-button {
  padding: 2px 8px;
  font-size: 12px;
  min-width: 0;
}

.icon-button {
  background: none;
  border: none;
  color: var(--text-color);
  cursor: pointer;
  padding: 2px 6px;
  font-size: 14px;
  border-radius: 4px;
  min-width: 0;
  width: 24px;
  height: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.icon-button:hover {
  background-color: var(--hover-bg);
}

.code-editor {
  padding: 8px;
}

.code-input {
  width: 100%;
  min-height: 80px;
  font-family: 'Consolas', 'Courier New', monospace;
  font-size: 13px;
  line-height: 1.5;
  resize: none;
  border: 1px solid var(--border-color);
  border-radius: 4px;
  background-color: var(--input-bg);
  color: var(--text-color);
  tab-size: 4;
}

.cell-output {
  padding: 8px;
  background-color: var(--bg-color);
  border-top: 1px solid var(--border-color);
}

.output-content {
  font-family: monospace;
  white-space: pre-wrap;
  background-color: var(--highlight-color);
  padding: 8px;
  border-radius: 4px;
  overflow-x: auto;
  font-size: 13px;
}

.add-cell-container {
  display: flex;
  justify-content: center;
  margin: 20px 0;
}

.add-cell-button {
  background-color: transparent;
  color: var(--text-color);
  border: 1px dashed var(--border-color);
  padding: 8px 16px;
  width: 100%;
  max-width: 300px;
  font-size: 13px;
}

.add-cell-button:hover {
  background-color: var(--hover-bg);
  border-style: solid;
}

/* 主题特定的调整 */
.dark .code-input,
.dark .output-content {
  background-color: var(--input-bg);
}

.solarized .code-input {
  color: #586e75;
}

.dracula .cell-toolbar {
  background-color: var(--card-bg);
}

.dracula .activity-item.active {
  background-color: var(--accent-color);
  color: white;
}

.execution-info {
  margin-top: 8px;
  font-size: 12px;
}

.execution-progress {
  display: flex;
  flex-direction: column;
  gap: 4px;
  color: var(--text-color);
}

.execution-progress progress {
  width: 100%;
  height: 4px;
  border-radius: 2px;
  overflow: hidden;
}

.execution-progress progress::-webkit-progress-bar {
  background-color: var(--border-color);
}

.execution-progress progress::-webkit-progress-value {
  background-color: var(--accent-color);
  transition: width 0.3s ease;
}

.progress-icon {
  margin-right: 4px;
}

.execution-time {
  display: flex;
  align-items: center;
  gap: 4px;
  color: var(--text-color);
  opacity: 0.7;
}

.time-icon {
  font-size: 14px;
}

/* 禁用状态样式 */
button:disabled,
textarea:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

/* 主题特定的样式调整 */
.dark .execution-progress {
  color: #e0e0e0;
}

.solarized .execution-progress {
  color: #586e75;
}

.dracula .execution-progress {
  color: #f8f8f2;
}
</style>