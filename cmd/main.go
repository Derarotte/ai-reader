package main

import (
	"ai-reader/internal/app"
	"log"
)

func main() {
	// 创建应用程序实例
	application := app.NewApp()
	
	// 初始化应用程序
	if err := application.Initialize(); err != nil {
		log.Fatal("Failed to initialize application:", err)
	}
	
	// 运行应用程序
	if err := application.Run(); err != nil {
		log.Fatal("Failed to run application:", err)
	}
	
	// 关闭应用程序
	if err := application.Shutdown(); err != nil {
		log.Printf("Error during shutdown: %v", err)
	}
}