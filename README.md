# mdPad README

mdPad is a desktop Markdown note application developed using the Wails framework. It combines the powerful backend capabilities of Go with modern web frontend technologies to provide a smooth and efficient Markdown editing experience.

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

mdPad is built on the Wails framework, using the Vditor Markdown editor for the frontend and Go for the backend.

**Technology Stack:**

* **Frontend:** HTML, CSS, JavaScript, [vditor](https://github.com/Vanessa219/vditor)
* **Backend:** Go
* **Framework:** [Wails v2](https://wails.io)
* **Build Tools:** (Depending on the specific project configuration, Wails projects typically use Go's build tools)

**Development Environment Setup:**

1. **Install Go Environment:** Ensure you have the Go language environment installed.
2. **Install Bun:** Install Bun following the instructions on the official website: [https://bun.sh/](https://bun.sh/)
3. **Install Wails CLI:** `go install github.com/wailsapp/wails/v2/cmd/wails@latest`
4. **Clone the Repository:**

   ```bash
   git clone https://github.com/stxh/mdPad.git
   cd mdPad
   ```
5. **Initialize Wails Project (if needed):** If starting from scratch, you can use `wails init` to create a new project.
6. **Install Frontend Dependencies (using Bun):**

   ```bash
   cd frontend
   bun install
   ```
7. **Run the Development Server:**

   ```bash
   wails dev
   ```

**Code Structure:**

(The following is an example; describe based on the actual project structure)
