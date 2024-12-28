package main

import (
	"context"
	"fmt"
	"os"
	"log"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
	Filename string
}


// NewApp creates a new App application struct
func NewApp() *App {
    var filename string
    if len(os.Args) >= 2 {
		// 获取第一个参数（os.Args[0] 是程序路径）
		filename = os.Args[1]
		log.Println("--> Filename in Args: ", filename)
    }
    return &App{Filename: filename}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Greet(ctl string, content string) string {
    switch ctl {
    case "init":
        return a.getMdContent()
    case "open":
        return a.goOpenFile()
    case "save":
        return a.goSaveFile(content)
    case "saveas":
        return a.goSaveAsFile(content)
	case "new":
		return a.goNewFile()
    default:
        return "Command Error"
    }
}

// Greet returns a greeting for the given name
func (a *App) getMdContent() string {
    if a.Filename == "" {
        log.Println("### Error\n没有指定文件名")
        return "### Error\n没有指定文件名"
    }

    log.Println("要处理的文件是：", a.Filename)
   
    content, err := os.ReadFile(a.Filename)
    if err != nil {
        errMsg := fmt.Sprintf("### Error\n读取文件出错：%v", err)
        log.Println(errMsg)
        return errMsg
    }

	// log.Println("-->Content: ", string(content))
    return string(content)
}

// goOpenFile
func (a *App) goOpenFile() string {
	odo := runtime.OpenDialogOptions{
		DefaultDirectory: "",           // string
		DefaultFilename:  "",           // string
		Title:            "Open File",  // string
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Markdown Files (*.md)",
				Pattern:     "*.md",
			},
		}, // []FileFilter
	}

	filename, err := runtime.OpenFileDialog(a.ctx, odo)
	if err != nil {
		return "### Test Message\n测试消息"
	}

	a.Filename = filename
	content, err := os.ReadFile(filename)
	if err != nil {
		return "### Test Message\n测试消息"
	}

	return string(content)
}

func (a *App) goSaveFile(content string) string {
	if a.Filename == "" {
		sdo := runtime.SaveDialogOptions {
			DefaultDirectory: "",           // string
			DefaultFilename:  "nutitile.md",           // string
			Title:            "Save File",  // string
			Filters: []runtime.FileFilter{
				{
					DisplayName: "Markdown Files (*.md)",
					Pattern:     "*.md",
				},
			}, // []FileFilter
		}
		filename, _ := runtime.SaveFileDialog(a.ctx, sdo)
		a.Filename = filename
	}

	if a.Filename != "" {
		err := os.WriteFile(a.Filename, []byte(content), 0644)
		if err == nil {
			return "OK"
		}
	}
	return "Error"
}

func (a *App) goSaveAsFile(content string) string {
	a.Filename = ""

	sdo := runtime.SaveDialogOptions {
		DefaultDirectory: "",           // string
		DefaultFilename:  "nutitile.md",           // string
		Title:            "Save File",  // string
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Markdown Files (*.md)",
				Pattern:     "*.md",
			},
		}, // []FileFilter
	}
	
	filename, _ := runtime.SaveFileDialog(a.ctx, sdo)

	if filename != "" {
		a.Filename = filename
		err := os.WriteFile(a.Filename, []byte(content), 0644)
		if err == nil {
			return "OK"
		}
	}
	return "Error"
}

func (a *App) goNewFile() string {
	a.Filename = ""
	return "OK"
}
