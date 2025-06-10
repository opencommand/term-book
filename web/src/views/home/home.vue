<template>
  <div :class="['vscode-container', theme]" v-if="!loading">
    <!-- 活动栏 -->
    <div class="activity-bar">
      <div class="activity-bar-items">
        <button class="activity-item active" title="文件资源管理器" @click="toggleFileExplorer">
          <!-- <span class="icon">🗂</span> 黑白文件夹图标 -->

          <svg t="1749517893653" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"
            p-id="3369" width="20" height="20">
            <path
              d="M773.82 338.84L625.06 194.28v144.56z m-220.1 69.33V165.51H268.47q-15 0-25.34 10.06a33 33 0 0 0-10.31 24.61v624a32.84 32.84 0 0 0 10.31 24.61q10.29 10.05 25.34 10.05h499.2q15.08 0 25.37-10.05a33 33 0 0 0 10.27-24.61v-416z m-285.23-312h356.57l249.58 242.67v485.33c0 28.65-10.45 53.27-31.33 73.49s-46.09 30.5-75.65 30.5H268.47q-44.25 0-75.62-30.5c-20.9-20.22-31.34-44.84-31.34-73.49v-624q0-43 31.34-73.49t75.62-30.51z m0 0"
              fill="#949DA6" p-id="3370"></path>
          </svg>


        </button>
        <button class="activity-item" title="搜索">
          <!-- <span class="icon">🔍</span> 黑白放大镜 -->
          <svg t="1749519568254" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"
            p-id="4379" width="20" height="20">
            <path
              d="M446.112323 177.545051c137.567677 0.219798 252.612525 104.59798 266.162424 241.493333 13.562828 136.895354-78.778182 261.818182-213.617777 289.008485-134.852525 27.203232-268.386263-52.156768-308.945455-183.608889s25.018182-272.252121 151.738182-325.779394A267.235556 267.235556 0 0 1 446.112323 177.545051m0-62.060607c-182.794343 0-330.989899 148.195556-330.989899 330.989899s148.195556 330.989899 330.989899 330.989899 330.989899-148.195556 330.989899-330.989899-148.195556-330.989899-330.989899-330.989899z m431.321212 793.341415a30.849293 30.849293 0 0 1-21.94101-9.102223l-157.220202-157.220202c-11.752727-12.179394-11.584646-31.534545 0.37495-43.50707 11.972525-11.972525 31.327677-12.140606 43.494141-0.37495l157.220202 157.220202a31.036768 31.036768 0 0 1 6.723232 33.810101 31.004444 31.004444 0 0 1-28.651313 19.174142z m0 0"
              p-id="4380" fill="#868686"></path>
          </svg>
        </button>
        <button class="activity-item" title="源代码管理">
          <!-- <span class="icon">↻</span> 更简洁的同步/刷新 -->
          <svg t="1749519674845" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"
            p-id="7104" width="20" height="20">
            <path
              d="M246.72 327.296c-12.032-12.032-31.424-12.032-43.456 0L9.024 521.344c-12.032 11.968-12.032 31.36 0 43.392l194.304 194.048c12.032 11.968 31.424 11.968 43.456 0 12.032-12.032 12.032-31.36 0-43.392L74.112 543.04l172.608-172.416C258.752 358.656 258.752 339.264 246.72 327.296zM954.368 521.344l-194.304-194.048c-12.032-12.032-31.488-12.032-43.456 0-12.032 11.968-12.032 31.36 0 43.328l172.608 172.416-172.608 172.352c-12.032 12.032-12.032 31.36 0 43.392 11.968 11.968 31.424 11.968 43.456 0l194.304-194.048C966.4 552.768 966.4 533.312 954.368 521.344zM689.664 132.608c-15.296-8.96-34.88-3.648-43.712 11.84l-384 766.848c-8.896 15.424-3.648 35.2 11.712 44.032 15.296 8.96 34.88 3.648 43.776-11.84l383.936-766.848C710.272 161.28 705.024 141.504 689.664 132.608z"
              p-id="7105" fill="#798686"></path>
          </svg>
        </button>
        <button class="activity-item" title="运行和调试">
          <!-- <span class="icon">▶</span> 黑色播放符号 -->
          <svg t="1749519783494" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"
            p-id="8324" width="20" height="20">
            <path d="M64 128v768h896V128H64zM0 64h1024v896H0V64z" fill="#798686" p-id="8325"></path>
            <path
              d="M64 256h896v64H64zM448 704H768V768H448zM290.2784 772.548267l162.269867-162.269867L290.2784 448l-45.2608 45.252267 117.0176 117.026133L245.0176 727.296z"
              fill="#798686" p-id="8326"></path>
          </svg>
        </button>
        <button class="activity-item" title="扩展">
          <!-- <span class="icon">＋</span> 黑白加号 -->
          <svg t="1749519839280" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"
            p-id="10297" width="20" height="20">
            <path
              d="M120.123733 60.629333h296.751787c3.636907 0 7.227733 0.730453 10.763947 2.193067h0.093866c3.495253 1.455787 6.570667 3.495253 9.186987 6.111573l31.634773 31.634774c2.6112 2.6112 4.652373 5.684907 6.111574 9.186986v0.093867c1.4592 3.531093 2.193067 7.12192 2.193066 10.763947V417.365333c0 3.636907-0.730453 7.227733-2.193066 10.763947v0.093867c-1.457493 3.49696-3.49696 6.572373-6.111574 9.186986l-31.634773 31.634774c-2.614613 2.614613-5.690027 4.65408-9.186987 6.111573h-0.093866c-3.536213 1.462613-7.12704 2.193067-10.763947 2.193067H120.123733c-3.642027 0-7.232853-0.733867-10.763946-2.193067h-0.093867c-3.50208-1.4592-6.575787-3.500373-9.186987-6.111573l-31.634773-31.634774c-2.61632-2.61632-4.655787-5.691733-6.111573-9.186986v-0.093867c-1.462613-3.536213-2.193067-7.12704-2.193067-10.763947V120.613547c0-3.642027 0.733867-7.232853 2.193067-10.763947v-0.093867c1.4592-3.498667 3.49696-6.572373 6.111573-9.186986L100.078933 68.933973c2.614613-2.614613 5.68832-4.652373 9.186987-6.111573h0.093867c3.531093-1.4592 7.12192-2.193067 10.763946-2.193067z m288.54784 48.165547H128.32768L108.305067 128.817493v280.343894l20.022613 20.022613h280.343893l20.022614-20.022613V128.817493l-20.022614-20.022613zM606.28992 60.629333h296.751787c3.642027 0 7.232853 0.733867 10.763946 2.193067h0.093867c3.498667 1.4592 6.572373 3.49696 9.186987 6.111573l31.634773 31.634774c2.614613 2.614613 4.652373 5.68832 6.111573 9.186986v0.093867c1.4592 3.531093 2.193067 7.12192 2.193067 10.763947V417.365333c0 3.636907-0.730453 7.227733-2.193067 10.763947v0.093867c-1.455787 3.495253-3.495253 6.570667-6.111573 9.186986l-31.634773 31.634774c-2.6112 2.6112-5.684907 4.652373-9.186987 6.111573h-0.093867c-3.531093 1.4592-7.12192 2.193067-10.763946 2.193067H606.28992c-3.636907 0-7.227733-0.730453-10.763947-2.193067h-0.093866c-3.49696-1.457493-6.572373-3.49696-9.186987-6.111573l-31.634773-31.634774c-2.614613-2.614613-4.65408-5.690027-6.111574-9.186986v-0.093867c-1.462613-3.536213-2.193067-7.12704-2.193066-10.763947V120.613547c0-3.642027 0.733867-7.232853 2.193066-10.763947v-0.093867c1.4592-3.50208 3.500373-6.575787 6.111574-9.186986l31.634773-31.634774c2.61632-2.61632 5.691733-4.655787 9.186987-6.111573h0.093866c3.536213-1.462613 7.12704-2.193067 10.763947-2.193067z m288.54784 48.165547H614.493867l-20.022614 20.022613v280.343894l20.022614 20.022613h280.343893l20.022613-20.022613V128.817493l-20.022613-20.022613zM120.123733 546.404693h296.751787c3.636907 0 7.227733 0.730453 10.763947 2.193067h0.093866c3.49696 1.457493 6.572373 3.49696 9.186987 6.111573l31.634773 31.634774c2.614613 2.614613 4.65408 5.690027 6.111574 9.186986v0.093867c1.462613 3.536213 2.193067 7.12704 2.193066 10.763947v296.751786c0 3.642027-0.733867 7.232853-2.193066 10.763947v0.093867c-1.4592 3.50208-3.500373 6.575787-6.111574 9.186986l-31.634773 31.634774c-2.61632 2.61632-5.691733 4.655787-9.186987 6.111573h-0.093866c-3.536213 1.462613-7.12704 2.193067-10.763947 2.193067H120.123733c-3.642027 0-7.232853-0.733867-10.763946-2.193067h-0.093867c-3.498667-1.4592-6.572373-3.49696-9.186987-6.111573l-31.634773-31.634774c-2.614613-2.614613-4.652373-5.68832-6.111573-9.186986v-0.093867c-1.4592-3.531093-2.193067-7.12192-2.193067-10.763947V606.388907c0-3.636907 0.730453-7.227733 2.193067-10.763947v-0.093867c1.455787-3.495253 3.495253-6.570667 6.111573-9.186986l31.634773-31.634774c2.6112-2.6112 5.684907-4.652373 9.186987-6.111573h0.093867c3.531093-1.4592 7.12192-2.193067 10.763946-2.193067z m288.54784 48.165547H128.32768L108.305067 614.592853v280.343894l20.022613 20.022613h280.343893l20.022614-20.022613V614.592853l-20.022614-20.022613zM606.28992 546.404693h296.751787c3.642027 0 7.232853 0.733867 10.763946 2.193067h0.093867c3.50208 1.4592 6.575787 3.500373 9.186987 6.111573l31.634773 31.634774c2.61632 2.61632 4.655787 5.691733 6.111573 9.186986v0.093867c1.462613 3.536213 2.193067 7.12704 2.193067 10.763947v296.751786c0 3.642027-0.733867 7.232853-2.193067 10.763947v0.093867c-1.4592 3.498667-3.49696 6.572373-6.111573 9.186986l-31.634773 31.634774c-2.614613 2.614613-5.68832 4.652373-9.186987 6.111573h-0.093867c-3.531093 1.4592-7.12192 2.193067-10.763946 2.193067H606.28992c-3.636907 0-7.227733-0.730453-10.763947-2.193067h-0.093866c-3.495253-1.455787-6.570667-3.495253-9.186987-6.111573l-31.634773-31.634774c-2.6112-2.6112-4.652373-5.684907-6.111574-9.186986v-0.093867c-1.4592-3.531093-2.193067-7.12192-2.193066-10.763947V606.388907c0-3.636907 0.730453-7.227733 2.193066-10.763947v-0.093867c1.457493-3.49696 3.49696-6.572373 6.111574-9.186986l31.634773-31.634774c2.614613-2.614613 5.690027-4.65408 9.186987-6.111573h0.093866c3.536213-1.462613 7.12704-2.193067 10.763947-2.193067z m288.54784 48.165547H614.493867l-20.022614 20.022613v280.343894l20.022614 20.022613h280.343893l20.022613-20.022613V614.592853l-20.022613-20.022613z"
              p-id="10298" fill="#798686"></path>
          </svg>
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
          <!-- <span class="icon">↻</span> -->
          <svg t="1749521061940" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"
            p-id="12515" width="200" height="200" fill="currentColor">
            <path
              d="M55.935033 264.48948c0 0 85.897017-132.548409 221.81443-203.673173 135.916406-71.121743 303.368504-50.646859 413.187968 18.319527 109.819465 68.970415 146.791894 127.160016 146.791894 127.160016l94.59499-53.879895c0 0 19.576483-9.697092 19.576483 12.932142l0 338.379961c0 0 0 30.17399-22.837719 19.395191-19.210878-9.062571-226.959086-127.198289-292.424528-164.466828-35.950145-16.035251-4.365101-29.062068-4.365101-29.062068l91.284402-52.173738c0 0-52.068992-65.209619-128.278989-99.744682-81.576231-42.501826-157.948384-47.541735-251.497925-12.224097-61.002644 23.025054-132.823368 81.988166-184.553949 169.082716L55.935033 264.48948 55.935033 264.48948 55.935033 264.48948zM904.056909 711.697844c0 0-85.897017 132.550423-221.816444 203.671159-135.917413 71.12275-303.366489 50.651895-413.186961-18.315498-109.825508-68.972429-146.790886-127.165052-146.790886-127.165052L27.662591 823.768348c0 0-19.572454 9.703135-19.572454-12.932142L8.090137 472.459267c0 0 0-30.170968 22.831676-19.397205 19.211885 9.067607 226.965129 127.198289 292.430571 164.470856 35.950145 16.035251 4.366109 29.058039 4.366109 29.058039l-91.285409 52.175753c0 0 52.071006 65.206598 128.279996 99.744682 81.57321 42.498804 157.942341 47.540728 251.496918 12.222082 60.998616-23.026061 132.820346-81.983131 184.546898-169.082716L904.056909 711.697844 904.056909 711.697844 904.056909 711.697844z"
              p-id="12516"></path>
          </svg>

        </button>
      </div>
      <!-- <div class="explorer-content">
        <div class="file-tree">
          <div class="folder" v-for="folder in fileList" :key="folder.name">
            <div class="folder-header" @click="toggleFolder(folder)">
              <span class="icon" style="display: flex; align-items: center; gap: 4px;">
                
                <svg v-if="folder.expanded" xmlns="http://www.w3.org/2000/svg" width="14" height="14" fill="none"
                  stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M6 9l6 6 6-6" />
                </svg>

                
                <svg v-else xmlns="http://www.w3.org/2000/svg" width="14" height="14" fill="none" stroke="currentColor"
                  stroke-width="2" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M9 6l6 6-6 6" />
                </svg>

             
                <svg t="1749520567168" class="icon" viewBox="0 0 1024 1024" version="1.1"
                  xmlns="http://www.w3.org/2000/svg" p-id="11365" width="14" height="14" fill="currentColor">
                  <path
                    d="M870.4 217.6c14.1 0 25.6 11.5 25.6 25.6v537.6c0 14.1-11.5 25.6-25.6 25.6H153.6c-14.1 0-25.6-11.5-25.6-25.6v-448c0-14.1 11.5-25.6 25.6-25.6h384.9l18.7-18.7 70.9-70.9h242.3m0-64H620.2c-11.9 0-23.3 4.7-31.7 13.1L512 243.2H153.6c-49.5 0-89.6 40.1-89.6 89.6v448c0 49.5 40.1 89.6 89.6 89.6h716.8c49.5 0 89.6-40.1 89.6-89.6V243.2c0-49.5-40.1-89.6-89.6-89.6z"
                    p-id="11366"></path>
                  <path d="M768 741.4H256c-17.7 0-32-14.3-32-32s14.3-32 32-32h512c17.7 0 32 14.3 32 32s-14.3 32-32 32z"
                    p-id="11367"></path>
                </svg>

              </span>

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
      </div> -->
      <FileExplorer :initial-files="fileList" @open-file="handleOpenFile" @create-file="handleCreateFile"
        @create-folder="handleCreateFolder" @delete-file="handleDeleteFile" @delete-folder="handleDeleteFolder"
        @rename-file="handleRenameFile" />
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
                  <span class="icon">
                    <svg v-if="cell.isRunning" xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="none"
                      stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
                      <!-- 时钟图标 (类似 ⏳) -->
                      <circle cx="12" cy="12" r="10" />
                      <path d="M12 6v6l4 2" />
                    </svg>

                    <svg v-else xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="currentColor"
                      viewBox="0 0 24 24">
                      <!-- 播放图标 (类似 ▶) -->
                      <path d="M8 5v14l11-7z" />
                    </svg>
                  </span>
                </button>

                <!-- 添加单元格 -->
                <button @click="addCell(index)" class="icon-button" title="添加单元格" :disabled="cell.isRunning">
                  <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="none" stroke="currentColor"
                    stroke-width="2" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M12 5v14m7-7H5" />
                  </svg>
                </button>

                <!-- 删除单元格 -->
                <button @click="removeCell(index)" class="icon-button" title="删除单元格" :disabled="cell.isRunning">
                  <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="none" stroke="currentColor"
                    stroke-width="2" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
                  </svg>
                </button>

                <!-- 保存单元格 -->
                <button @click="saveCell(index)" class="icon-button" title="保存单元格" :disabled="cell.isRunning">
                  <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="none" stroke="currentColor"
                    stroke-width="2" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round"
                      d="M17 21H7a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h6l6 6v10a2 2 0 0 1-2 2z" />
                    <path stroke-linecap="round" stroke-linejoin="round" d="M12 17a2 2 0 1 0 0-4 2 2 0 0 0 0 4z" />
                    <path stroke-linecap="round" stroke-linejoin="round" d="M13 3v6h6" />
                  </svg>
                </button>

              </div>

            </div>
            <div class="code-editor">
              <!-- <textarea v-model="cell.content" class="code-input"
                :placeholder="index === 0 ? '输入代码并按Shift+Enter运行' : '输入代码...'" :disabled="cell.isRunning"></textarea> -->
              <CodeEditor v-model="cell.content" :language="language" :readOnly="cell.isRunning" :theme="monacoTheme"
                @run="executeCell" @save="saveCell" ref="codeEditorRef" />
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
  <div v-else class="advanced-loading">
    <div class="loading-content">
      <div class="loading-logo">
        <svg viewBox="0 0 100 100" class="loading-svg">
          <path d="M50 10 L90 50 L50 90 L10 50 Z" fill="none" stroke="var(--accent-color)" stroke-width="8" />
        </svg>
      </div>
      <div class="loading-dots">
        <div class="dot"></div>
        <div class="dot"></div>
        <div class="dot"></div>
      </div>
      <div class="loading-status">
        <div class="status-text">正在加载资源</div>
        <div class="status-progress">0%</div>
      </div>
      <div class="loading-hint">请稍候，这可能需要几秒钟...</div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { inject, ref, watch, nextTick, onMounted, onBeforeUnmount, toRaw, computed } from 'vue'
import { ThemeSymbol, Theme } from '../../theme-context'
import { getFileListApi, openFileApi, runCellApi, saveFileApi } from '../../api/Document.ts'
import FileExplorer from '../../components/FileExplorer.vue'
import { ElMessage } from 'element-plus'
const themeContext = inject(ThemeSymbol)
if (!themeContext) throw new Error('Theme context not provided')
const currentFile = ref<FileItem | null>(null)
const { theme, setTheme } = themeContext
const selected = ref<Theme>(theme.value)
import CodeEditor from '../../components/CodeEditor.vue'


const props = defineProps({
  cell: {
    type: Object as () => Cell,
    required: true
  },
  language: {
    type: String,
    default: 'python'
  }
})

const emit = defineEmits(['update:cell', 'execute', 'add', 'remove'])

const codeEditorRef = ref<InstanceType<typeof CodeEditor> | null>(null)

// 计算Monaco主题名称
// 根据 theme 动态返回 monaco 主题名
const monacoTheme = computed(() => {
  switch (theme.value) {
    case 'light':
      return 'vs'
    case 'dark':
      return 'vs-dark'
    case 'solarized':
      return 'solarized'
    case 'dracula':
      return 'dracula'
    default:
      return 'vs'
  }
})

watch(theme, (val) => {
  selected.value = val
})

const loading = ref(false)

const activeFile = ref<string | null>(null)
const showFileExplorer = ref(true)

async function handleOpenFile(file: any) {
  console.log('打开文件:', file);
  // 实现打开文件逻辑
  activeFile.value = file.name
  currentFile.value = file
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

function handleCreateFile({ folder, name }: any) {
  console.log('新建文件:', folder, name);
  // 实现新建文件逻辑
}


// 创建新文件夹
function handleCreateFolder({ parent, name }: { parent: FolderItem, name: string }) {
  if (!name.trim()) return;


}

// 删除文件
function handleDeleteFile(file: FileItem) {

}

// 删除文件夹
function handleDeleteFolder(folder: FolderItem) {

}

// 重命名文件
function handleRenameFile({ file, newName }: { file: FileItem, newName: string }) {
  if (!newName.trim()) return;


}



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
  currentFile.value = file
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



async function saveCell(index: number) {
  const rawCell = toRaw(cells.value[index]);

  // 手动映射字段
  const formattedCell = {
    input: rawCell.content || "",
    output: rawCell.output || "",
    input_time: rawCell.input_time || new Date().toISOString(),
    output_time: rawCell.output_time || new Date().toISOString(),
    exec_time: rawCell.exec_time || "0ms"
  };

  const documentToSave = {
    author: "defineAuthor",
    filename: currentFile.value?.name || `cell-${index + 1}.json`,
    cells: [formattedCell]
  };

  try {
    const result = await saveFileApi(documentToSave);
    ElMessage.success(result.msg || `单元格 ${index + 1} 保存成功`)
    // alert(result.msg || `单元格 ${index + 1} 保存成功`);
  } catch (error) {
    ElMessage.error(`保存失败：${(error as Error).message}`)
    // alert(`保存失败：${(error as Error).message}`);
  }
}


async function executeCell(index: number) {
  const cell = cells.value[index];

  if (cell.isRunning) return;

  cell.isRunning = true;
  cell.output = "正在执行...";
  cell.progress = 0;
  cell.currentExecutionTime = 0;

  const startTime = Date.now();

  // 实时更新执行时间和模拟进度条
  const timer = setInterval(() => {
    cell.currentExecutionTime = Date.now() - startTime;

    // 模拟进度条
    const elapsed = cell.currentExecutionTime;
    cell.progress = Math.min(90, Math.floor((elapsed / 1000) * 90));
  }, 100); // 每 100ms 更新一次

  try {
    await new Promise(resolve => setTimeout(resolve, 1000));

    console.log(cell.content);

    const response = await runCellApi(cell.content);
    const result = response.data;

    cell.output = result.output || "执行完成";
    cell.executionTime = result.exec_time || (Date.now() - startTime);
    cell.executionCount = (cell.executionCount || 0) + 1;
  } catch (error) {
    cell.output = `执行出错：${(error as Error).message}`;
  } finally {
    clearInterval(timer); // 清除定时器
    cell.isRunning = false;
    cell.progress = 100;
    cell.currentExecutionTime = Date.now() - startTime;
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
  padding: 0;
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
  padding: 0px 4px;
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









.advanced-loading {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100vh;
  background-color: var(--bg-color);
}

.loading-content {
  text-align: center;
  max-width: 300px;
}

.loading-logo {
  margin-bottom: 30px;
}

.loading-svg {
  width: 80px;
  height: 80px;
  animation: rotate 2s linear infinite;
}

.loading-dots {
  display: flex;
  justify-content: center;
  gap: 10px;
  margin-bottom: 20px;
}

.dot {
  width: 12px;
  height: 12px;
  border-radius: 50%;
  background-color: var(--accent-color);
  opacity: 0.6;
  animation: pulse 1.5s infinite ease-in-out;
}

.dot:nth-child(2) {
  animation-delay: 0.2s;
}

.dot:nth-child(3) {
  animation-delay: 0.4s;
}

.loading-status {
  display: flex;
  justify-content: space-between;
  margin-bottom: 10px;
  font-size: 14px;
  color: var(--text-color);
}

.status-text {
  opacity: 0.8;
}

.status-progress {
  font-weight: bold;
  color: var(--accent-color);
}

.loading-hint {
  font-size: 12px;
  color: var(--text-color);
  opacity: 0.6;
}

@keyframes rotate {
  from {
    transform: rotate(0deg);
  }

  to {
    transform: rotate(360deg);
  }
}

@keyframes pulse {

  0%,
  100% {
    transform: scale(0.8);
    opacity: 0.6;
  }

  50% {
    transform: scale(1.2);
    opacity: 1;
  }
}
</style>