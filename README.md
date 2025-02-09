# mdPad README

[中文](README.cn.md) | English

mdPad is a desktop Markdown note application developed using the Wails3 framework. It combines the powerful backend capabilities of Go with modern web frontend technologies to provide a smooth and efficient Markdown editing experience.

**Key Features:**

* **Instant Rendering:** Uses the Vditor Markdown editor for instant rendering as you type, providing a WYSIWYG (What You See Is What You Get) experience without needing to manually switch to a preview.
* **Rich Markdown Syntax Support:** Vditor supports standard Markdown syntax and provides various extended syntaxes, such as:
  * **GFM** (GitHub Flavored Markdown)
  * **Mathematical Formulas:** Supports both KaTeX and MathJax rendering engines for mathematical formulas.
  * **Flowcharts, Sequence Diagrams, Gantt Charts:** Uses mermaid syntax for drawing these diagrams.
  * **Mind Maps:** Uses markmap syntax for drawing mind maps.
  * **Charts:** Supports the ECharts charting library.
  * **Musical Notation:** Supports ABC.js for rendering musical notation.
* **Desktop Application Experience:** Built with Wails, providing the performance and experience of a native desktop application, with a small footprint and fast startup speed.
* **Cross-Platform Support:** Wails supports cross-platform building, so mdPad can run on operating systems such as Windows, macOS, and Linux.

**Usage:**

1. **Download and Install:** Download the installation package for your operating system from the release page and install it.
2. **Launch mdPad:** Run the application.
3. **Edit Markdown:** Enter Markdown text in the editor area, and the rendered result will be displayed instantly.
4. **Save/Open:** Use the provided save and open functions to manage your notes.

# mdPad Development Instructions

mdPad is built on the Wails3 framework, using the Vditor Markdown editor for the frontend and Go for the backend.

**Technology Stack:**

* **Frontend:** HTML, CSS, JavaScript, [vditor](https://github.com/Vanessa219/vditor)
* **Backend:** Go
* **Framework:** [Wails v3](https://wails.io)
* **Build Tools:** (Depending on the specific project configuration, Wails projects typically use Go's build tools)

**Development Environment Setup:**

1. **Install Go Environment:** Ensure you have the Go language environment installed.
2. **Install Bun:** Install Bun following the instructions on the official website: [https://bun.sh/](https://bun.sh/)
3. **Install Wails CLI:** `go install github.com/wailsapp/wails/v3/cmd/wails3@latest`
4. **Clone the Repository:**

```bash
git clone https://github.com/stxh/mdPad.git
cd mdPad
```

5. **Initialize Wails Project (if needed):** If starting from scratch, you can use `wails init` to create a new project.
6. **Install Frontend Dependencies (using Bun):**

```bash
wails3 task install:frontend:deps
```

7. **Run the Development Server:**

```bash
wails3 task dev
```

**Code Structure:**

（The description should be based on the actual project; the following is an example）

```text
mdPad/
├── Taskfile.yml         // Wails configuration file
├── frontend/          // Frontend code
│   ├── index.html
│   ├── src/           // JavaScript/TypeScript source code
│   │   ├── main.js      // or main.ts
│   │   └── components/  // Components
│   │       └── ...
│   ├── package.json
│   └── ...
├── app.go             // Go backend code
├── main.go            // Program entry point
└── ...
```

**Key Code Snippets:**

* **Frontend Vditor Integration:**

  [vidtor](https://github.com/Vanessa219/vditor)

  **JavaScript**

  ```javascript
  import Vditor from 'vditor'

  const vditor = new Vditor('vditor', {
      // Vditor configuration options, e.g.:
      mode: 'ir', // Instant rendering mode
      height: '100%',
      // ... other configurations
  })
  ```
* **Go Backend and Frontend Interaction:** (Using the binding mechanism provided by Wails)
  **Go**

  ```go
  // Go code
  type App struct {
          runtime *wails.Runtime
  }

  func (a *App) Greet(name string) string {
          return "Hello " + name + ", from Go!"
  }

  // Binding in main.go
  app := NewApp()
  err := wails.CreateApp(&wails.AppConfig{
          // ...
          Bind: []interface{}{
                  app,
          },
          // ...
  })

  // Frontend JavaScript calling Go function
  window.backend.App.Greet("World").then(result => {
          console.log(result) // Outputs "Hello World, from Go!"
  })
  ```

**Contributing Code, Building, Other Instructions** are similar to previous versions; please refer to previous responses.

**Important Changes and Notes:**

* All `npm install` commands have been replaced with `bun install`.
* To add new packages, use `bun add <package-name>` instead of `npm install <package-name>`. For devDependencies, use `bun add -d <package-name>`.
* When running scripts, use `bun run <script>` instead of `npm run <script>`.

With these modifications, you can now use Bun to manage the frontend dependencies of the mdPad project, leveraging the speed benefits that Bun offers. Please ensure that Bun is correctly installed in your development environment.
