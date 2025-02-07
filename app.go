package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v3/pkg/application"
)

// App struct
type App struct {
	Filename string
}

// NewApp creates a new App application struct
func New() *App {
	var filename string
	if len(os.Args) >= 2 {
		// Get the first argument (os.Args[0] is the program path)
		filename = os.Args[1]
		log.Println("--> Filename in Args: ", filename)
	}
	return &App{Filename: filename}
}

// GetMdContent
func (a *App) GetMdContent() map[string]interface{} {
	if a.Filename == "" {
		// log.Println("### Error\nNo filename specified")
		return map[string]interface{}{"value": "", "filename": ""}
	}

	content, err := os.ReadFile(a.Filename)
	localDir = getFilePath(a.Filename)
	// log.Println("localDir:", localDir)

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
	dialog := application.OpenFileDialog().
		SetTitle("Open Markdown File").
		SetMessage("Select a markdown file to open").
		AddFilter("Markdown file(*.md)", "*.md")

	filename, _ := dialog.PromptForSingleSelection()
	if filename == "" {
		return map[string]interface{}{"value": "", "filename": ""}
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		return map[string]interface{}{"value": "", "filename": ""}
	}

	a.Filename = filename
	localDir = getFilePath(a.Filename)
	// log.Println("localDir:", localDir, " a.Filename:", a.Filename)

	return map[string]interface{}{"value": string(content), "filename": filename}
}

func getNewFileName() string {
	filename, _ := application.SaveFileDialog().
		CanCreateDirectories(true).
		SetMessage("Save to a markdow file").
		SetFilename("untitle.md").
		PromptForSingleSelection()

	// log.Println("getNewFileName filename: ", filename)
	return filename
}

// GoSaveFile
func (a *App) GoSaveFile(content string) string {
	if a.Filename == "" {
		a.Filename = getNewFileName()
	}

	if a.Filename != "" {
		localDir = getFilePath(a.Filename)
		// log.Println("localDir:", localDir)

		err := os.WriteFile(a.Filename, []byte(content), 0644)
		if err == nil {
			return "OK"
		}
	}
	return "Error"
}

// GoSaveAsFile resutl: {"value": "OK", "filename": "e:\\StXhCode\\go\\mdPad\\untitle.md"}
func (a *App) GoSaveAsFile(content string) map[string]interface{} {
	a.Filename = getNewFileName()

	if a.Filename != "" {
		localDir = getFilePath(a.Filename)
		// log.Println("localDir:", localDir)

		err := os.WriteFile(a.Filename, []byte(content), 0644)
		if err == nil {
			return map[string]interface{}{"value": "OK", "filename": a.Filename}
		}
	}
	return map[string]interface{}{"value": "Error", "filename": ""}
}

// get file path
func getFilePath(filename string) string {
	if filename != "" {
		return filepath.Dir(filename)
	} else {
		return ""
	}
}
