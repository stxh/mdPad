# **mdPad Readme**

中文 | [English](README.md)

mdPad 是一款基于 Wails3 框架开发的桌面 Markdown 笔记应用。它结合了 Go 语言的强大后端能力和现代 Web 前端技术，提供流畅、高效的 Markdown 编辑体验。

**主要功能：**

* **即时渲染：** 采用 Vditor Markdown 编辑器，实现输入即时渲染，所见即所得，无需手动切换预览。
* **丰富的 Markdown 语法支持：** Vditor 支持标准 Markdown 语法，并提供多种扩展语法，例如：

  * **GFM** (GitHub Flavored Markdown)
  * **数学公式：** 支持 KaTeX 和 MathJax 两种数学公式渲染引擎。
  * **流程图、时序图、甘特图：** 使用 mermaid 语法绘制流程图。
  * **脑图：** 使用 markmap 语法绘制脑图。
  * **图表：** 支持 ECharts 图表库。
  * **五线谱：** 支持 ABC.js 渲染五线谱。
* **桌面应用体验：** 基于 Wails3 构建，提供原生桌面应用的性能和体验，体积小巧，启动速度快。
* **跨平台支持：** Wails3 支持跨平台构建，mdPad 可以在 Windows、macOS 和 Linux 等操作系统上运行。

**使用方法：**

1. **下载安装：** 从发布页面下载对应操作系统的安装包并安装。
2. **启动 mdPad：** 运行应用程序。
3. **编辑 Markdown：** 在编辑器区域输入 Markdown 文本，会即时显示渲染后的效果。
4. **保存/打开：** 使用提供的保存和打开功能管理笔记。
5. **关联Markdown**：关联md文件后可双击打开md文件

**mdPad 开发说明**

mdPad 基于 Wails3 框架构建，前端使用 Vditor Markdown 编辑器，后端使用 Go 语言。

**技术栈：**

* **前端：** HTML、CSS、JavaScript、[vditor](https://github.com/Vanessa219/vditor)
* **后端：** Go 语言
* **框架：** [Wails v3](https://v3alpha.wails.io/getting-started/installation/)
* **构建工具：** (根据项目实际情况，通常 Wails3 项目使用 Go 的构建工具)

**开发环境搭建：**

1. **安装 Go 环境：** 确保已安装 Go 语言环境。
2. **安装 Bun：** 按照官方网站（[https://bun.sh/](https://www.google.com/url?sa=E&source=gmail&q=https://bun.sh/)）的指引安装 Bun。
3. **安装 Wails CLI：**`go install -v github.com/wailsapp/wails/v3/cmd/wails3@latest`
4. **克隆代码库：**
   **Bash**

```bash
git clone https://github.com/stxh/mdPad.git
cd mdPad
```

5. **初始化 Wails 项目（如果需要）：** 如果是从零开始，可以使用 `wails3 init` 创建项目。
6. **安装前端依赖（使用 Bun）：**
   **Bash**

```bash
wails3 task dev
```

7. **运行开发服务器：**
   **Bash**

```
wails dev
```

**代码结构：**

(根据项目实际情况进行描述，以下是一个示例)

```
mdPad/
├── Taskfile.yml         // Wails 配置文件
├── frontend/          // 前端代码
│   ├── index.html
│   ├── src/           // JavaScript/TypeScript 源代码
│   │   ├── main.js      // 或 main.ts
│   │   └── components/  // 组件
│   │       └── ...
│   ├── package.json
│   └── ...
├── app.go             // Go 后端代码
├── main.go            // 程序入口
└── ...
```

**关键代码片段：**

* **前端 Vditor 集成：**

  [vidtor](https://github.com/Vanessa219/vditor)

  **JavaScript**

  ```
  import Vditor from 'vditor'

  const vditor = new Vditor('vditor', {
      // Vditor 配置项，例如：
      mode: 'ir', // 即时渲染模式
      height: '100%',
      // ...其他配置
  })
  ```
* **Go 后端与前端交互：** (使用 Wails 提供的绑定机制)
  **Go**

  ```
  // Go 代码
  type App struct {
          runtime *wails.Runtime
  }

  func (a *App) Greet(name string) string {
          return "Hello " + name + ", from Go!"
  }

  // 在 main.go 中绑定
  app := NewApp()
  err := wails.CreateApp(&wails.AppConfig{
          // ...
          Bind: []interface{}{
                  app,
          },
          // ...
  })

  // 前端 JavaScript 调用 Go 函数
  window.backend.App.Greet("World").then(result => {
          console.log(result) // 输出 "Hello World, from Go!"
  })
  ```

**贡献代码、构建、其他说明** 与之前的版本类似，请参考之前的回复。

**重要变更和说明：**

* 所有 `npm install` 命令都替换为 `bun install`。
* 如果需要添加新的包，使用 `bun add <package-name>` 代替 `npm install <package-name>`。devDependencies 使用 `bun add -d <package-name>`。
* 运行脚本时，使用 `bun run <script>` 代替 `npm run <script>`。

通过以上修改，您现在可以使用 Bun 来管理 mdPad 项目的前端依赖，享受 Bun 带来的速度优势。请确保您的开发环境中已正确安装 Bun。
