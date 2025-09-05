package app

import (
	"ai-reader/internal/ai"
	"ai-reader/internal/events"
	"ai-reader/internal/reader"
	"ai-reader/internal/ui"
	"ai-reader/pkg/document"
	"ai-reader/pkg/theme"
	"os"
	"path/filepath"
)

// App 应用程序实现
type App struct {
	eventBus         *events.Bus
	documentManager  document.DocumentManager
	themeManager     theme.ThemeManager
	readerController reader.ReaderController
	aiService        ai.AIService
	mainWindow       *ui.MainWindow
	serviceContainer ServiceContainer
	config           Configuration
}

// NewApp 创建新应用程序
func NewApp() *App {
	app := &App{
		eventBus:         events.NewBus(),
		serviceContainer: NewServiceContainer(),
	}
	
	app.initializeServices()
	app.setupEventHandlers()
	
	return app
}

// initializeServices 初始化服务
func (a *App) initializeServices() {
	// 获取配置目录
	configDir := a.getConfigDir()
	
	// 初始化配置
	a.config = NewConfig(filepath.Join(configDir, "config.json"))
	
	// 初始化文档管理器
	a.documentManager = document.NewManager()
	
	// 初始化主题管理器
	a.themeManager = theme.NewManager(filepath.Join(configDir, "theme.json"))
	
	// TODO: 初始化AI服务
	// a.aiService = ai.NewService()
	
	// TODO: 初始化阅读器控制器
	// a.readerController = reader.NewController(a.eventBus)
	
	// 注册服务到容器
	a.serviceContainer.Register("eventBus", a.eventBus)
	a.serviceContainer.Register("documentManager", a.documentManager)
	a.serviceContainer.Register("themeManager", a.themeManager)
	a.serviceContainer.Register("config", a.config)
}

// setupEventHandlers 设置全局事件处理器
func (a *App) setupEventHandlers() {
	// 监听主题变化事件
	a.eventBus.Subscribe(events.ThemeChanged, func(event events.Event) {
		themeName := event.Payload.(string)
		if err := a.themeManager.SetCurrentTheme(themeName); err == nil {
			// 保存主题配置
			a.themeManager.SaveThemeConfig()
		}
	})
	
	// 监听AI分析请求
	a.eventBus.Subscribe(events.AIAnalysisRequest, func(event events.Event) {
		// TODO: 处理AI分析请求
		selectedText := event.Payload.(string)
		
		// 模拟AI分析结果
		result := "这是对文本 \"" + selectedText + "\" 的分析结果。\n\n" +
				 "这段文本包含了重要的概念和背景信息。AI分析功能正在开发中，" +
				 "将来会提供更深入的语义分析、概念解释和相关背景知识。"
		
		// 发布分析结果
		a.eventBus.Publish(events.Event{
			Type:    events.AIAnalysisResult,
			Payload: result,
		})
	})
}

// getConfigDir 获取配置目录
func (a *App) getConfigDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return ".ai-reader"
	}
	return filepath.Join(homeDir, ".ai-reader")
}

// Initialize 初始化应用程序
func (a *App) Initialize() error {
	// 加载配置
	if err := a.config.Load(); err != nil {
		// 配置加载失败不是致命错误，使用默认配置
	}
	
	// 创建主窗口
	a.mainWindow = ui.NewMainWindow(a.eventBus)
	
	return nil
}

// Run 运行应用程序
func (a *App) Run() error {
	if a.mainWindow == nil {
		if err := a.Initialize(); err != nil {
			return err
		}
	}
	
	a.mainWindow.Show()
	return nil
}

// Shutdown 关闭应用程序
func (a *App) Shutdown() error {
	// 保存配置
	a.config.Save()
	a.themeManager.SaveThemeConfig()
	
	return nil
}

// 接口实现
func (a *App) GetEventBus() *events.Bus {
	return a.eventBus
}

func (a *App) GetDocumentManager() document.DocumentManager {
	return a.documentManager
}

func (a *App) GetThemeManager() theme.ThemeManager {
	return a.themeManager
}

func (a *App) GetReaderController() reader.ReaderController {
	return a.readerController
}

func (a *App) GetAIService() ai.AIService {
	return a.aiService
}