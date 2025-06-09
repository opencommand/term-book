<template>
  <div :class="['vscode-container', theme]" v-if="!loading">
    <!-- 活动栏 -->
    <div class="activity-bar">
      <div class="activity-bar-items">
        <button class="activity-item active" title="文件资源管理器" @click="toggleFileExplorer">
          <span class="icon">🗂</span> <!-- 黑白文件夹图标 -->
        </button>
        <button class="activity-item" title="搜索">
          <span class="icon">🔍</span> <!-- 黑白放大镜 -->
        </button>
        <button class="activity-item" title="源代码管理">
          <span class="icon">↻</span> <!-- 更简洁的同步/刷新 -->
        </button>
        <button class="activity-item" title="运行和调试">
          <span class="icon">▶</span> <!-- 黑色播放符号 -->
        </button>
        <button class="activity-item" title="扩展">
          <span class="icon">＋</span> <!-- 黑白加号 -->
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

    <!-- 文件资源管理器 -->
    <!-- 文件资源管理器 -->
    <div class="file-explorer" v-if="showFileExplorer" :style="{ width: explorerWidth + 'px' }">
      <div class="explorer-header">
        <h3>文件资源管理器</h3>
        <button class="icon-button" @click="refreshFileList" title="刷新">
          <span class="icon">↻</span>
        </button>
      </div>
      <div class="explorer-content">
        <div class="file-tree">
          <div class="folder" v-for="folder in fileList" :key="folder.name">
            <div class="folder-header" @click="toggleFolder(folder)">
              <span class="icon">{{ folder.expanded ? '🗀' : '🗁' }}</span>
              <span class="folder-name">{{ folder.name }}</span>
            </div>
            <div class="folder-content" v-if="folder.expanded">
              <div class="file-item" v-for="file in folder.files" :key="file.name"
                :class="{ active: activeFile === file.name }" @click="openFile(file)">
                <span class="icon">{{ getFileIcon(file.name) }}</span>
                <span class="file-name">{{ file.name }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- 拖动句柄 -->
      <div class="resizer" @mousedown="startResizing"></div>
    </div>


    <!-- 侧边栏 -->
    <div v-if="isExplorerVisible" class="sidebar" :style="{ width: sidebarWidth + 'px' }">
      <div class="sidebar-header">
        <h3>{{ activeFile || '未选择文件' }}</h3>
      </div>
      <div class="sidebar-content">
        <div class="outline-item" v-for="(cell, index) in cells" :key="cell.id" @click="scrollToCell(index)">
          <span class="outline-icon">[ ]</span>
          <span class="outline-text">单元格 {{ index + 1 }}</span>
        </div>
      </div>
      <div class="resizer" @mousedown.prevent="beginResize"></div>
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
      <div class="editor-content">
        <div v-if="!activeFile" class="empty-editor">
          <div class="empty-message">
            <h3>欢迎使用 Notebook 编辑器</h3>
            <p>请从文件资源管理器中选择一个文件开始编辑</p>
            <button @click="createNewNotebook" class="new-notebook-button">
              + 创建新 Notebook
            </button>
          </div>
        </div>

        <template v-else>
          <div class="cell code-cell" v-for="(cell, index) in cells" :key="cell.id" ref="cellElements">
            <div class="cell-toolbar">
              <span class="cell-index">In [{{ cell.executionCount || ' ' }}]:</span>
              <div class="cell-actions">
                <button @click="executeCell(index)" class="run-button" title="运行单元格" :disabled="cell.isRunning">
                  <span class="icon">{{ cell.isRunning ? '⏳' : '▶' }}</span>
                </button>
                <button @click="addCell(index)" class="icon-button" title="添加单元格" :disabled="cell.isRunning">+</button>
                <button @click="removeCell(index)" class="icon-button" title="删除单元格"
                  :disabled="cell.isRunning">×</button>
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
                    <span class="progress-icon">⌛</span> <!-- 黑白沙漏 -->
                    <span>正在运行... {{ cell.currentExecutionTime }}ms</span>
                    <progress :value="cell.progress" max="100"></progress>
                  </div>
                  <div v-if="!cell.isRunning && cell.executionTime" class="execution-time">
                    <span class="time-icon">⏲</span> <!-- 黑白定时器 -->
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
        </template>
      </div>
    </div>
  </div>
  <div v-else>loading </div>
</template>

<script setup lang="ts">
import { inject, ref, watch, nextTick, onMounted, onBeforeUnmount } from 'vue'
import { ThemeSymbol, Theme } from '../../theme-context'
import { getFileListApi, openFileApi, runCellApi } from '../../api/Document.ts'
const themeContext = inject(ThemeSymbol)
if (!themeContext) throw new Error('Theme context not provided')

const { theme, setTheme } = themeContext
const selected = ref<Theme>(theme.value)

watch(theme, (val) => {
  selected.value = val
})

const loading = ref(false)

const activeFile = ref<string | null>(null)
const showFileExplorer = ref(true)





const isExplorerVisible = ref(true)
const sidebarWidth = ref(250)
const resizingActive = ref(false)
let initialMouseX = 0
let initialSidebarWidth = 250


function onResizeMove(event) {
  if (!resizingActive.value) return
  const movementX = event.clientX - initialMouseX
  sidebarWidth.value = Math.min(Math.max(initialSidebarWidth + movementX, 5), 800)
}

function onResizeEnd() {
  resizingActive.value = false
  document.removeEventListener('mousemove', onResizeMove)
  document.removeEventListener('mouseup', onResizeEnd)
}

function beginResize(event) {
  resizingActive.value = true
  initialMouseX = event.clientX
  initialSidebarWidth = sidebarWidth.value
  document.addEventListener('mousemove', onResizeMove)
  document.addEventListener('mouseup', onResizeEnd)
}



const currentFile = ref(null)



const explorerWidth = ref(250)
const isResizing = ref(false)
let startX = 0
let startWidth = 250





// 拖动逻辑
const startResizing = (e) => {
  isResizing.value = true
  startX = e.clientX
  startWidth = explorerWidth.value
  document.addEventListener('mousemove', resize)
  document.addEventListener('mouseup', stopResizing)
}

const resize = (e) => {
  if (!isResizing.value) return
  const delta = e.clientX - startX
  explorerWidth.value = Math.min(Math.max(startWidth + delta, 5), 1000)
}

const stopResizing = () => {
  isResizing.value = false
  document.removeEventListener('mousemove', resize)
  document.removeEventListener('mouseup', stopResizing)
}
// 文件列表数据
interface FileItem {
  name: string
  type: string
  path: string
  content?: string
}

interface FolderItem {
  name: string
  expanded: boolean
  files: FileItem[]
}

const fileList = ref<FolderItem[]>([
  {
    name: 'Notebooks',
    expanded: true,
    files: [
      { name: '数据分析111.ipynb', type: 'notebook', path: '/Notebooks/数据分析11111.ipynb' },
      { name: '机器学习.ipynb', type: 'notebook', path: '/Notebooks/机器学习.ipynb' },
      { name: '可视化.ipynb', type: 'notebook', path: '/Notebooks/可视化.ipynb' }
    ]
  },
  {
    name: '数据集',
    expanded: false,
    files: [
      { name: 'sales_data.csv', type: 'csv', path: '/数据集/sales_data.csv' },
      { name: 'customer_data.json', type: 'json', path: '/数据集/customer_data.json' }
    ]
  },
  {
    name: '脚本',
    expanded: false,
    files: [
      { name: 'utils.py', type: 'python', path: '/脚本/utils.py' },
      { name: 'config.py', type: 'python', path: '/脚本/config.py' }
    ]
  }
])

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

function toggleFileExplorer() {
  showFileExplorer.value = !showFileExplorer.value
}

function refreshFileList() {
  loading.value = true
  // 模拟API调用延迟
  setTimeout(() => {
    loading.value = false
  }, 800)
}

function toggleFolder(folder: FolderItem) {
  folder.expanded = !folder.expanded
}

function getFileIcon(filename: string): string {
  const ext = filename.split('.').pop()?.toLowerCase()
  switch (ext) {
    case 'ipynb': return '📓'
    case 'py': return '🐍'
    case 'csv': return '📊'
    case 'json': return '🔣'
    default: return '📄'
  }
}

const loadFile = async (filename) => {
  try {
    const res = await openFileApi(filename)
    return res.data
  } catch (err) {
    console.error('打开文件失败', err)
  }
}

const openFile = async (file: FileItem) => {
  activeFile.value = file.name

  try {
    const fileData = await loadFile(file.name)

    // 数据为空或格式不正确
    if (!fileData || !Array.isArray(fileData.cells)) {
      console.warn('文件内容无效或不包含 cells 字段', fileData)
      cells.value = []
      return
    }

    // 正常解析 cells
    cells.value = fileData.cells.map((c, index) => ({
      id: generateId(),
      content: c?.input ?? `# 第 ${index + 1} 单元格无内容`,
      output: c?.output ?? ''
    }))
  } catch (error) {
    console.error('读取文件内容失败', error)
    cells.value = [] // 确保界面不会残留旧内容
  }
}



function createNewNotebook() {
  const newName = `未命名-${new Date().getTime()}.ipynb`
  fileList.value[0].files.unshift({
    name: newName,
    type: 'notebook',
    path: `/Notebooks/${newName}`
  })
  fileList.value[0].expanded = true
  openFile(fileList.value[0].files[0])
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
  if (cells.value[index].timer) {
    clearInterval(cells.value[index].timer);
  }
  if (cells.value.length > 1) {
    cells.value.splice(index, 1);
  }
}

// function executeCell(index: number) {
//   const cell = cells.value[index];

//   if (cell.isRunning) return;

//   cell.isRunning = true;
//   cell.currentExecutionTime = 0;
//   cell.progress = 0;
//   cell.output = "正在执行...";

//   const startTime = Date.now();
//   const totalDuration = 10000;

//   cell.timer = window.setInterval(() => {
//     const elapsed = Date.now() - startTime;
//     cell.currentExecutionTime = elapsed;
//     cell.progress = Math.min(100, (elapsed / totalDuration) * 100);

//     if (elapsed > 3000 && !cell.output?.includes("加载数据...")) {
//       cell.output = "正在执行...\n加载数据...";
//     }
//     if (elapsed > 6000 && !cell.output?.includes("处理中")) {
//       cell.output = "正在执行...\n加载数据...\n处理中...";
//     }

//     if (elapsed >= totalDuration) {
//       clearInterval(cell.timer);
//       cell.isRunning = false;
//       cell.executionTime = elapsed;
//       cell.executionCount = (cell.executionCount || 0) + 1;

//       const code = cell.content;
//       if (code.includes("print(")) {
//         cell.output = code.match(/print\((.*)\)/)?.[1] || "执行完成";
//       } else if (code.includes("np.arange")) {
//         cell.output = "array([0, 1, 2, 3, 4, 5, 6, 7, 8, 9])";
//       } else {
//         cell.output = "执行完成 (模拟输出)";
//       }
//     }
//   }, 100);
// }


async function executeCell(index: number) {
  // 获取指定的 cell
  const cell = cells.value[index];

  // 如果正在执行，避免重复提交
  if (cell.isRunning) return;

  // 设置执行状态
  cell.isRunning = true;
  cell.output = "正在执行...";
  cell.progress = 0;
  cell.currentExecutionTime = 0;

  // 记录开始时间
  const startTime = Date.now();

  try {
    const startTime = Date.now();

    // 假设 cell.content 是纯字符串命令
    console.log(cell.content);

    const response = await runCellApi(cell.content);

    // response.data 是后端返回的data字段
    const result = response.data;

    cell.output = result.output || "执行完成";
    cell.executionTime = result.exec_time || (Date.now() - startTime);
    cell.executionCount = (cell.executionCount || 0) + 1;
  } catch (error) {
    cell.output = `执行出错：${(error as Error).message}`;
  } finally {
    cell.isRunning = false;
    cell.progress = 100;
  }

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
    // console.log(fileRes, 'sdadas');

    if (fileRes.success) {
      try {
        const rawFiles = fileRes.data
        console.log(rawFiles);

        if (Array.isArray(rawFiles) && fileList.value.length > 0) {
          fileList.value[0].files = rawFiles.map((item: any) => ({
            name: String(item),
            type: 'notebook',
            path: `/some/path/${item}`
          }))
        }
      } catch (err) {
        console.error('文件数据解析失败:', err)
      }
    } else {
      console.warn('获取文件列表失败:', fileRes.message)
    }


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
onBeforeUnmount(() => {
  stopResizing()
  document.removeEventListener('mousemove', onResizeMove)
  document.removeEventListener('mouseup', onResizeEnd)
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
  padding: 4px 0;
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

/* 文件资源管理器样式 */
.file-explorer {
  width: 250px;
  background-color: var(--card-bg);
  border-right: 1px solid var(--border-color);
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.explorer-header {
  padding: 12px;
  border-bottom: 1px solid var(--border-color);
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.explorer-header h3 {
  margin: 0;
  font-size: 14px;
  font-weight: bold;
}

.explorer-content {
  flex: 1;
  overflow-y: auto;
  padding: 8px 0;
}

.file-tree {
  padding: 0 8px;
}

.folder {
  margin-bottom: 4px;
}

.folder-header {
  padding: 6px 8px;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  border-radius: 4px;
}

.folder-header>.icon {
  font-size: 12px;
}

.folder-header:hover {
  background-color: var(--hover-bg);
}

.folder-name {
  flex: 1;
}

.folder-content {
  padding-left: 24px;
}

.file-item {
  padding: 6px 8px;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  border-radius: 4px;
}

.file-item>.icon {
  font-size: 12px;
}

.file-item:hover {
  background-color: var(--hover-bg);
}

.file-item.active {
  background-color: var(--highlight-color);
  font-weight: bold;
}

.file-name {
  flex: 1;
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
  font-size: 14px;
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

.empty-editor {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
  color: var(--text-color);
  opacity: 0.7;
}

.empty-message {
  text-align: center;
  max-width: 400px;
}

.empty-message h3 {
  margin-bottom: 16px;
}

.new-notebook-button {
  margin-top: 20px;
  padding: 10px 20px;
  background-color: var(--button-bg);
  color: var(--button-text-color);
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
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

.run-button>span {
  font-size: 10px;
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

/* 执行信息样式 */
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

.dark .execution-progress {
  color: #e0e0e0;
}

.solarized .execution-progress {
  color: #586e75;
}

.dracula .execution-progress {
  color: #f8f8f2;
}

/* 图标样式 */
.icon {
  font-size: 18px;
}

.resizer {
  width: 5px;
  cursor: ew-resize;
  position: absolute;
  right: 0;
  top: 0;
  bottom: 0;
  background: transparent;
}

.resizer:hover {
  background: rgba(0, 0, 0, 0.1);
}

.file-explorer {
  display: flex;
  flex-direction: column;
  height: 100%;
  min-width: 5px;
  max-width: 600px;

  position: relative;
  user-select: none;

}


.sidebar {
  height: 100vh;

  overflow: auto;
  position: relative;
  display: flex;
  flex-direction: column;
  min-width: 5px;
  max-width: 600px;

}


.resizer {
  position: absolute;
  top: 0;
  right: 0;
  width: 6px;
  height: 100%;
  cursor: ew-resize;
  user-select: none;
  background-color: transparent;
}
</style>