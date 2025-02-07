package main

import (
	"bytes"
	"embed"
	_ "embed"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
)

//go:embed frontend/dist
var assets embed.FS

// localDir 设置为包级别变量
var localDir string = "./"

// gFileName 当前打开的文件名
var gFileName string = ""

func getMimeType(filename string) string {
	var mimeTypesByExt = map[string]string{
		".avif": "image/avif",
		".css":  "text/css; charset=utf-8",
		".gif":  "image/gif",
		".htm":  "text/html; charset=utf-8",
		".html": "text/html; charset=utf-8",
		".jpeg": "image/jpeg",
		".jpg":  "image/jpeg",
		".js":   "text/javascript; charset=utf-8",
		".json": "application/json",
		".mjs":  "text/javascript; charset=utf-8",
		".pdf":  "application/pdf",
		".png":  "image/png",
		".svg":  "image/svg+xml",
		".wasm": "application/wasm",
		".webp": "image/webp",
		".xml":  "text/xml; charset=utf-8",
	}
	return mimeTypesByExt[filepath.Ext(filename)]
}

func main() {

	app := application.New(application.Options{
		Name:        "mdPad",
		Description: "Markdown Editor Pad",
		Services: []application.Service{
			application.NewService(&App{}),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
			Middleware: func(next http.Handler) http.Handler {
				return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
					path := req.URL.Path
					// log.Println(" --> Middleware path: ", path)

					// 如果路径以 /dist 或 /wails 开头，或者以 http(s):// 开头，则直接跳过
					if strings.HasPrefix(path, "/dist") || strings.HasPrefix(path, "/wails") || strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
						next.ServeHTTP(rw, req)
						return
					}

					// 检查本地文件是否存在
					localFilePath := filepath.Join(localDir, path)
					// log.Println(" --> Middleware real path: ", localFilePath)

					if fileInfo, err := os.Stat(localFilePath); err == nil && !fileInfo.IsDir() {
						// 如果本地文件存在，则使用 assetserver.ServeFile 处理
						// log.Println(" --> Middleware find in local path: ", localFilePath)

						fileData, _ := os.ReadFile(localFilePath)

						header := rw.Header()
						header.Set("Content-Length", fmt.Sprintf("%d", len(fileData)))
						if mimeType := header.Get("Content-Type"); mimeType == "" {
							mimeType = getMimeType(localFilePath)
							header.Set("Content-Type", mimeType)
						}

						rw.WriteHeader(http.StatusOK)
						io.Copy(rw, bytes.NewReader(fileData))

						return
					}

					// 如果本地文件不存在，则调用 next.ServeHTTP
					next.ServeHTTP(rw, req)
				})
			},
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	// Listen for files being used to open the application
	app.OnApplicationEvent(events.Common.ApplicationStarted, func(event *application.ApplicationEvent) {
		if len(os.Args) >= 2 {
			if fileInfo, err := os.Stat(os.Args[1]); err == nil && !fileInfo.IsDir() {
				gFileName = os.Args[1]
			}
		}
	})

	app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title: "mdPad",
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		BackgroundColour: application.NewRGB(27, 38, 54),
		URL:              "/",
	})

	// Create a goroutine that emits an event containing the current time every second.
	// The frontend can listen to this event and update the UI accordingly.
	go func() {
		for {
			now := time.Now().Format(time.RFC1123)
			app.EmitEvent("time", now)
			time.Sleep(time.Second)
		}
	}()

	// Run the application. This blocks until the application has been exited.
	err := app.Run()

	// If an error occurred while running the application, log it and exit.
	if err != nil {
		log.Fatal(err)
	}
}
