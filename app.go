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
}

// GetMdContent
func (a *App) GetMdContent() map[string]interface{} {
	if gFileName == "" {
		// log.Println("### Error\nNo filename specified")
		return map[string]interface{}{"value": "", "filename": ""}
	}

	content, err := os.ReadFile(gFileName)
	localDir = getFilePath(gFileName)
	// log.Println("localDir:", localDir)

	if err != nil {
		errMsg := fmt.Sprintf("### Error\nError reading file: %v", err)
		log.Println(errMsg)
		return map[string]interface{}{"value": "", "filename": ""}
	}

	// log.Println("-->Content: ", string(content))
	return map[string]interface{}{"value": string(content), "filename": gFileName}
}

// GoNewFile
func (a *App) GoNewFile() string {
	gFileName = ""
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

	gFileName = filename
	localDir = getFilePath(gFileName)
	// log.Println("localDir:", localDir, " gFileName:", gFileName)

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
	if gFileName == "" {
		gFileName = getNewFileName()
	}

	if gFileName != "" {
		localDir = getFilePath(gFileName)
		// log.Println("localDir:", localDir)

		err := os.WriteFile(gFileName, []byte(content), 0644)
		if err == nil {
			return "OK"
		}
	}
	return "Error"
}

// GoSaveAsFile resutl: {"value": "OK", "filename": "e:\\StXhCode\\go\\mdPad\\untitle.md"}
func (a *App) GoSaveAsFile(content string) map[string]interface{} {
	gFileName = getNewFileName()

	if gFileName != "" {
		localDir = getFilePath(gFileName)
		// log.Println("localDir:", localDir)

		err := os.WriteFile(gFileName, []byte(content), 0644)
		if err == nil {
			return map[string]interface{}{"value": "OK", "filename": gFileName}
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
