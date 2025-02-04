package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/wailsapp/wails/v3/pkg/application"
)

// App struct
type App struct {
	ctx      context.Context
	Filename string
}

// NewApp creates a new App application struct
func NewApp() *App {
	var filename string
	if len(os.Args) >= 2 {
		// Get the first argument (os.Args[0] is the program path)
		filename = os.Args[1]
		log.Println("--> Filename in Args: ", filename)
	}
	return &App{Filename: filename}
}

// startup is called when the app starts. The context is saved
// so we can call the application methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// GetMdContent
func (a *App) GetMdContent() map[string]interface{} {
	if a.Filename == "" {
		// log.Println("### Error\nNo filename specified")
		return map[string]interface{}{"value": "", "filename": ""}
	}

	// log.Println("File to process:", a.Filename)

	content, err := os.ReadFile(a.Filename)
	if err != nil {
		errMsg := fmt.Sprintf("### Error\nError reading file: %v", err)
		log.Println(errMsg)
		return map[string]interface{}{"value": "", "filename": ""}
	}

	// log.Println("-->Content: ", string(content))
	return map[string]interface{}{"value": string(content), "filename": a.Filename}
}

// GoNewFile
func (a *App) GoNewFile() string {
	a.Filename = ""
	return "OK"
}

// goOpenFile
func (a *App) GoOpenFile() map[string]interface{} {
	dialog := application.OpenFileDialog()
	// dialog.SetTitle("Select Markdown file")
	// dialog.SetFilters([]*application.FileFilter{
	// 	{
	// 		DisplayName: "Markdown (*.md)",
	// 		Pattern:     "*.md",
	// 	},
	// })

	// Single file selection
	filename, err := dialog.PromptForSingleSelection()
	if err != nil {
		return map[string]interface{}{"value": "", "filename": ""}
	}

	// odo := application.OpenDialogOptions{
	// 	DefaultDirectory: "",          // string
	// 	DefaultFilename:  "",          // string
	// 	Title:            "Open File", // string
	// 	Filters: []application.FileFilter{
	// 		{
	// 			DisplayName: "Markdown Files (*.md)",
	// 			Pattern:     "*.md",
	// 		},
	// 	}, // []FileFilter
	// }

	// filename, err := application.OpenFileDialog(a.ctx, odo)
	// if err != nil {
	// 	return map[string]interface{}{"value": "", "filename": ""}
	// }

	a.Filename = filename
	content, err := os.ReadFile(filename)
	if err != nil {
		return map[string]interface{}{"value": "", "filename": ""}
	}

	return map[string]interface{}{"value": string(content), "filename": filename}
}

// GoSaveFile
func (a *App) GoSaveFile(content string) string {
	// if a.Filename == "" {
	// 	sdo := application.SaveDialogOptions{
	// 		DefaultDirectory: "",            // string
	// 		DefaultFilename:  "untitile.md", // string
	// 		Title:            "Save File",   // string
	// 		Filters: []application.FileFilter{
	// 			{
	// 				DisplayName: "Markdown Files (*.md)",
	// 				Pattern:     "*.md",
	// 			},
	// 		}, // []FileFilter
	// 	}
	// 	filename, _ := application.SaveFileDialog(a.ctx, sdo)
	// 	a.Filename = filename
	// }

	// if a.Filename != "" {
	// 	err := os.WriteFile(a.Filename, []byte(content), 0644)
	// 	if err == nil {
	// 		return "OK"
	// 	}
	// }
	return "Error"
}

// GoSaveAsFile resutl: {"value": "OK", "filename": "e:\\StXhCode\\go\\mdPad\\untitle.md"}
func (a *App) GoSaveAsFile(content string) map[string]interface{} {
	// a.Filename = ""

	// sdo := application.SaveDialogOptions{
	// 	DefaultDirectory: "",           // string
	// 	DefaultFilename:  "untitle.md", // string
	// 	Title:            "Save File",  // string
	// 	Filters: []application.FileFilter{
	// 		{
	// 			DisplayName: "Markdown Files (*.md)",
	// 			Pattern:     "*.md",
	// 		},
	// 	}, // []FileFilter
	// }

	// filename, _ := application.SaveFileDialog(a.ctx, sdo)

	// if filename != "" {
	// 	a.Filename = filename
	// 	err := os.WriteFile(a.Filename, []byte(content), 0644)
	// 	if err == nil {
	// 		return map[string]interface{}{"value": "OK", "filename": a.Filename}
	// 	}
	// }
	return map[string]interface{}{"value": "Error", "filename": ""}
}
