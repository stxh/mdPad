package main

import (
	"embed"
	_ "embed"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed frontend/dist
var assets embed.FS

// localDir 设置为包级别变量
var localDir string = "./"

// localFileHandler 实现了 http.Handler 接口，用于优先服务嵌入资源，并在资源未找到时尝试服务本地文件
type localFileHandler struct {
	assetsHandler http.Handler // 用于服务嵌入资源的 Handler
	localDir      string       // 本地文件目录，用于查找 Markdown 文档相关的本地资源
}

func (h *localFileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("ServeHTTP -> ", r.URL)

	// 1. 优先尝试从嵌入的 assets 中查找文件
	rw := &responseRecorder{ResponseWriter: w}
	h.assetsHandler.ServeHTTP(rw, r)

	// 2. 如果 assets 中找不到文件 (状态码为 404)，则尝试从本地目录查找
	if rw.status == http.StatusNotFound {
		localFilePath := filepath.Join(h.localDir, r.URL.Path)
		log.Println("尝试加载本地文件:", localFilePath)

		_, err := os.Stat(localFilePath)
		if err == nil {
			// 本地文件存在，则使用 http.ServeFile 提供文件服务
			log.Println("本地文件找到，提供服务:", localFilePath)
			http.ServeFile(w, r, localFilePath)
			return // 成功服务本地文件后直接返回
		} else {
			log.Println("本地文件未找到:", localFilePath, err)
		}
	}

	// 3. 如果 assets 中找到文件 (无论本地文件是否存在)，或者本地文件不存在且 assets 中也找不到，则使用 assetsHandler 的处理结果
	if rw.status != 0 { // 只有当 assetsHandler 设置了状态码时才应用，避免覆盖 http.ServeFile 的状态码
		w.WriteHeader(rw.status)
	}
	if rw.header != nil {
		for k, v := range rw.header {
			for _, val := range v { // 复制所有 header 值
				w.Header().Add(k, val)
			}
		}
	}
	if rw.body != nil {
		w.Write(rw.body)
	}
}

// responseRecorder 用于记录 ServeHTTP 的状态码和响应体
type responseRecorder struct {
	http.ResponseWriter
	status int
	header http.Header
	body   []byte
}

func (rw *responseRecorder) WriteHeader(code int) {
	rw.status = code
}

func (rw *responseRecorder) Header() http.Header {
	if rw.header == nil {
		rw.header = make(http.Header)
	}
	return rw.header
}

func (rw *responseRecorder) Write(buf []byte) (int, error) {
	if rw.body == nil {
		rw.body = []byte{}
	}
	rw.body = append(rw.body, buf...)
	return rw.ResponseWriter.Write(buf)
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
			// &localFileHandler{
			// 	assets: application.AssetFileServerFS(assets),
			// },
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
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
