<template>
  <div class="explorer-content">
    <div class="file-tree">
      <div class="folder" v-for="folder in fileList" :key="folder.name">
        <div class="folder-header" @click="toggleFolder(folder)">
          <span class="icon" style="display: flex; align-items: center; gap: 4px;">
            <!-- 展开/收起箭头 -->
            <svg v-if="folder.expanded" xmlns="http://www.w3.org/2000/svg" width="14" height="14" fill="none"
              stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" d="M6 9l6 6 6-6" />
            </svg>
            <svg v-else xmlns="http://www.w3.org/2000/svg" width="14" height="14" fill="none" stroke="currentColor"
              stroke-width="2" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" d="M9 6l6 6-6 6" />
            </svg>

            <!-- 文件夹图标 -->
            <svg t="1749520567168" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"
              p-id="11365" width="14" height="14" fill="currentColor">
              <path
                d="M870.4 217.6c14.1 0 25.6 11.5 25.6 25.6v537.6c0 14.1-11.5 25.6-25.6 25.6H153.6c-14.1 0-25.6-11.5-25.6-25.6v-448c0-14.1 11.5-25.6 25.6-25.6h384.9l18.7-18.7 70.9-70.9h242.3m0-64H620.2c-11.9 0-23.3 4.7-31.7 13.1L512 243.2H153.6c-49.5 0-89.6 40.1-89.6 89.6v448c0 49.5 40.1 89.6 89.6 89.6h716.8c49.5 0 89.6-40.1 89.6-89.6V243.2c0-49.5-40.1-89.6-89.6-89.6z"
                p-id="11366"></path>
              <path d="M768 741.4H256c-17.7 0-32-14.3-32-32s14.3-32 32-32h512c17.7 0 32 14.3 32 32s-14.3 32-32 32z"
                p-id="11367"></path>
            </svg>
          </span>

          <span class="folder-name">{{ folder.name }}</span>

          <!-- 文件夹操作按钮 -->
          <div class="folder-actions" v-if="true">
            <button class="file-action" @click.stop="createNewFile(folder)" title="新建文件">
              <!-- <span class="icon">📄</span> -->
              <svg t="1749537156322" class="icon" viewBox="0 0 1024 1024" version="1.1"
                xmlns="http://www.w3.org/2000/svg" p-id="13844" width="20" height="14" fill="currentColor">
                <path
                  d="M455.111111 967.111111H170.666667V56.888889h523.377777L853.333333 227.555556v341.333333h-56.888889V250.311111L671.288889 113.777778H227.555556v796.444444h227.555555z"
                  p-id="13845"></path>
                <path
                  d="M853.333333 341.333333h-284.444444V56.888889h56.888889v227.555555h227.555555zM455.111111 739.555556h398.222222v56.888888H455.111111z"
                  p-id="13846"></path>
                <path d="M625.777778 568.888889h56.888889v398.222222h-56.888889z" p-id="13847"></path>
              </svg>
            </button>
            <button class="file-action" @click.stop="createNewFolder(folder)" title="新建文件夹">
              <!-- <span class="icon">📁</span> -->
              <svg t="1749537248471" class="icon" viewBox="0 0 1024 1024" version="1.1"
                xmlns="http://www.w3.org/2000/svg" p-id="14863" width="20" height="14" fill="currentColor">
                <path
                  d="M703.8 547.8h-167v-167c0-13.8-11.2-25-25-25s-25 11.2-25 25v167h-167c-13.8 0-25 11.2-25 25s11.2 25 25 25h167v167c0 13.8 11.2 25 25 25s25-11.2 25-25v-167h167c13.8 0 25-11.2 25-25s-11.2-25-25-25z"
                  p-id="14864"></path>
                <path
                  d="M833.3 234.1H530.8l-29.6-58.5c-10.4-20.6-26.4-37.9-46.1-50.1-19.7-12.1-42.3-18.5-65.5-18.5H188.7c-68.9 0-125 56.1-125 125v513.5c0 96.5 78.5 175 175 175h544.7c96.5 0 175-78.5 175-175V359.1c-0.1-68.9-56.1-125-125.1-125z m75 511.5c0 68.9-56.1 125-125 125H238.7c-68.9 0-125-56.1-125-125V232c0-41.4 33.6-75 75-75h200.9c28.4 0 54.1 15.8 66.9 41.1l36.6 72.2c4.3 8.4 12.9 13.7 22.3 13.7h317.9c41.4 0 75 33.6 75 75v386.6z"
                  p-id="14865"></path>
              </svg>
            </button>
            <button class="file-action" @click.stop="deleteFolder(folder)" title="删除文件夹">
              <!-- <span class="icon">🗑️</span> -->
              <svg t="1749537357129" class="icon" viewBox="0 0 1024 1024" version="1.1"
                xmlns="http://www.w3.org/2000/svg" p-id="16943" width="20" height="14" fill="currentColor">
                <path
                  d="M96 320a32 32 0 1 1 0-64h832a32 32 0 0 1 0 64H96z m736 0h64v448a160 160 0 0 1-160 160H288a160 160 0 0 1-160-160V320h64v96H128v-96h64v448a96 96 0 0 0 96 96h448a96 96 0 0 0 96-96V320z m-512 112a32 32 0 0 1 64 0v320a32 32 0 0 1-64 0v-320z m320 0a32 32 0 0 1 64 0v320a32 32 0 0 1-64 0v-320zM288 256H224V192a96 96 0 0 1 96-96h384a96 96 0 0 1 96 96v64h-64V224h64v32h-64V192a32 32 0 0 0-32-32H320a32 32 0 0 0-32 32v64z"
                  p-id="16944"></path>
              </svg>
            </button>
          </div>
        </div>

        <div class="folder-content" v-if="folder.expanded">
          <div class="file-item" v-for="file in folder.files" :key="file.name"
            :class="{ active: activeFile === file.path }" @click="openFile(file)" @mouseenter="hoveredFile = file.path"
            @mouseleave="hoveredFile = ''">
            <span class="icon">{{ getFileIcon(file.name) }}</span>
            <span class="file-name">{{ file.name }}</span>

            <!-- 文件操作按钮 -->
            <div class="file-actions" v-if="hoveredFile === file.path">
              <button class="file-action" @click.stop="renameFile(file)" title="重命名">
                <!-- <span class="icon">✏️</span> -->
                <svg t="1749537449742" class="icon" viewBox="0 0 1024 1024" version="1.1"
                  xmlns="http://www.w3.org/2000/svg" p-id="18106" width="20" height="14" fill="currentColor">
                  <path
                    d="M391.467 736.322a31.99 31.99 0 0 1-13.783 8.126l-232.261 66.798c-12.088 3.477-23.276-7.711-19.799-19.799l66.798-232.261a32 32 0 0 1 8.126-13.782l472.869-472.87c12.496-12.496 32.758-12.496 45.254 0L864.335 218.2c12.497 12.496 12.497 32.758 0 45.255L391.467 736.322z m248.009-516.709l77.781 77.782 56.569-56.569-77.782-77.782-56.568 56.569z m-50.912 50.911L265.88 593.209l-31.401 109.182 109.182-31.401 322.685-322.684-77.782-77.782zM129.001 889h768v72h-768v-72z"
                    p-id="18107"></path>
                </svg>
              </button>
              <button class="file-action" @click.stop="deleteFile(file)" title="删除">
                <!-- <span class="icon">🗑️</span> -->
                <svg t="1749537357129" class="icon" viewBox="0 0 1024 1024" version="1.1"
                  xmlns="http://www.w3.org/2000/svg" p-id="16943" width="20" height="14" fill="currentColor">
                  <path
                    d="M96 320a32 32 0 1 1 0-64h832a32 32 0 0 1 0 64H96z m736 0h64v448a160 160 0 0 1-160 160H288a160 160 0 0 1-160-160V320h64v96H128v-96h64v448a96 96 0 0 0 96 96h448a96 96 0 0 0 96-96V320z m-512 112a32 32 0 0 1 64 0v320a32 32 0 0 1-64 0v-320z m320 0a32 32 0 0 1 64 0v320a32 32 0 0 1-64 0v-320zM288 256H224V192a96 96 0 0 1 96-96h384a96 96 0 0 1 96 96v64h-64V224h64v32h-64V192a32 32 0 0 0-32-32H320a32 32 0 0 0-32 32v64z"
                    p-id="16944"></path>
                </svg>
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 新建文件/文件夹对话框 -->
    <div class="modal" v-if="showDialog">
      <div class="modal-content">
        <h3>{{ dialogTitle }}</h3>
        <input v-model="dialogInput" ref="inputRef" class="modal-input" :placeholder="dialogPlaceholder"
          @keyup.enter="confirmDialog" />
        <div class="modal-actions">
          <button @click="cancelDialog">取消</button>
          <button @click="confirmDialog" class="primary">确认</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, nextTick } from 'vue';

interface FileItem {
  name: string;
  path: string;
  isDirectory?: boolean;
}

interface FolderItem {
  name: string;
  path: string;
  expanded: boolean;
  files: FileItem[];
}

const props = defineProps({
  initialFiles: {
    type: Array as () => FolderItem[],
    default: () => []
  }
});

const emit = defineEmits([
  'open-file',
  'create-file',
  'create-folder',
  'delete-file',
  'delete-folder',
  'rename-file'
]);

const fileList = ref<FolderItem[]>(props.initialFiles);
const activeFile = ref('');
const hoveredFolder = ref('');
const hoveredFile = ref('');

// 对话框相关状态
const showDialog = ref(false);
const dialogTitle = ref('');
const dialogInput = ref('');
const dialogPlaceholder = ref('');
const dialogAction = ref<'create-file' | 'create-folder' | 'rename-file' | null>(null);
const currentFolder = ref<FolderItem | null>(null);
const currentFile = ref<FileItem | null>(null);
const inputRef = ref<HTMLInputElement | null>(null);

// 文件图标映射
function getFileIcon(filename: string): string {
  const ext = filename.split('.').pop()?.toLowerCase();
  const icons: Record<string, string> = {
    'ipynb': '📓',
    'py': '🐍',
    'js': '📜',
    'ts': '📜',
    'json': '📋',
    'md': '📝',
    'txt': '📄',
    'csv': '📊',
    'jpg': '🖼️',
    'png': '🖼️',
    'gif': '🖼️',
  };
  return icons[ext || ''] || '📄';
}

// 切换文件夹展开状态
function toggleFolder(folder: FolderItem) {
  folder.expanded = !folder.expanded;
}

// 打开文件
function openFile(file: FileItem) {
  activeFile.value = file.path;
  emit('open-file', file);
}

// 新建文件
function createNewFile(folder: FolderItem) {
  currentFolder.value = folder;
  dialogTitle.value = `在 ${folder.name} 中新建文件`;
  dialogPlaceholder.value = '请输入文件名 (例如: script.py)';
  dialogAction.value = 'create-file';
  dialogInput.value = '';
  showDialog.value = true;
  nextTick(() => {
    inputRef.value?.focus();
  });
}

// 新建文件夹
function createNewFolder(folder: FolderItem) {
  currentFolder.value = folder;
  dialogTitle.value = `在 ${folder.name} 中新建文件夹`;
  dialogPlaceholder.value = '请输入文件夹名';
  dialogAction.value = 'create-folder';
  dialogInput.value = '';
  showDialog.value = true;
  nextTick(() => {
    inputRef.value?.focus();
  });
}

// 重命名文件
function renameFile(file: FileItem) {
  currentFile.value = file;
  dialogTitle.value = `重命名 ${file.name}`;
  dialogInput.value = file.name;
  dialogPlaceholder.value = '请输入新名称';
  dialogAction.value = 'rename-file';
  showDialog.value = true;
  nextTick(() => {
    inputRef.value?.focus();
    inputRef.value?.select();
  });
}

// 删除文件
function deleteFile(file: FileItem) {
  if (confirm(`确定要删除 ${file.name} 吗？`)) {
    emit('delete-file', file);
  }
}

// 删除文件夹
function deleteFolder(folder: FolderItem) {
  if (confirm(`确定要删除文件夹 ${folder.name} 及其所有内容吗？`)) {
    emit('delete-folder', folder);
  }
}

// 确认对话框操作
function confirmDialog() {
  if (!dialogInput.value.trim()) {
    alert('名称不能为空');
    return;
  }

  if (dialogAction.value === 'create-file' && currentFolder.value) {
    emit('create-file', {
      folder: currentFolder.value,
      name: dialogInput.value
    });
  } else if (dialogAction.value === 'create-folder' && currentFolder.value) {
    emit('create-folder', {
      parent: currentFolder.value,
      name: dialogInput.value
    });
  } else if (dialogAction.value === 'rename-file' && currentFile.value) {
    emit('rename-file', {
      file: currentFile.value,
      newName: dialogInput.value
    });
  }

  closeDialog();
}

// 取消对话框
function cancelDialog() {
  closeDialog();
}

// 关闭对话框
function closeDialog() {
  showDialog.value = false;
  dialogAction.value = null;
  currentFolder.value = null;
  currentFile.value = null;
  dialogInput.value = '';
}
</script>


<style scoped>
.explorer-content {
  padding: 8px 0;
  height: 100%;
  overflow-y: auto;
}

.file-tree {
  font-size: 13px;
  user-select: none;
}

.folder {
  margin-bottom: 4px;
}

.folder-header {
  display: flex;
  align-items: center;
  padding: 6px 8px;
  cursor: pointer;
  position: relative;
}

.folder-header:hover {
  background-color: var(--hover-bg);
  border-radius: 4px;
}

.folder-name {
  margin-left: 4px;
  flex: 1;
}

.folder-actions {
  display: none;
  position: absolute;
  right: 8px;
  /* background-color: var(--card-bg); */
  border-radius: 4px;
  padding: 2px;
  /* box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1); */
}

.folder-header:hover .folder-actions {
  display: flex;
  gap: 4px;
}

.folder-content {
  margin-left: 20px;
  padding-left: 8px;
  border-left: 1px dashed var(--border-color);
}

.file-item {
  display: flex;
  align-items: center;
  padding: 4px 8px;
  cursor: pointer;
  position: relative;
  border-radius: 4px;
}

.file-item:hover {
  background-color: var(--hover-bg);
}

.file-item.active {
  background-color: var(--highlight-color);
  font-weight: 500;
}

.file-name {
  margin-left: 4px;
  flex: 1;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.file-actions {
  display: none;
  position: absolute;
  right: 8px;
  /* background-color: var(--card-bg); */
  border-radius: 4px;
  padding: 2px;
  /* box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1); */

}

.file-item:hover .file-actions {
  display: flex;
  gap: 4px;
}

.file-action {
  background: none;
  border: none;
  cursor: pointer;
  color: var(--text-color);
  font-size: 12px;
  padding: 2px 0px;
  border-radius: 2px;
  display: flex;
  align-items: center;
  box-shadow: none;
}

.file-action:hover {
  background-color: var(--hover-bg);
}

/* 模态对话框样式 */
.modal {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal-content {
  background-color: var(--card-bg);
  border-radius: 8px;
  padding: 16px;
  min-width: 300px;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.2);
}

.modal-input {
  width: 96%;
  margin: 12px 0;
  padding: 8px 2%;
  border: 1px solid var(--border-color);
  border-radius: 4px;
  background-color: var(--input-bg);
  color: var(--text-color);
}

.modal-actions {
  display: flex;
  justify-content: flex-end;
  gap: 8px;
}

.modal-actions button {
  padding: 6px 12px;
  border-radius: 4px;
  border: none;
  cursor: pointer;
}

.modal-actions button.primary {
  background-color: var(--button-bg);
  color: var(--button-text-color);
}
</style>