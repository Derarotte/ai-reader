package app

import (
	"ai-reader/internal/ai"
	"ai-reader/internal/events"
	"ai-reader/internal/reader"
	"ai-reader/pkg/document"
	"ai-reader/pkg/theme"
)

// Application 应用程序主接口
type Application interface {
	// Initialize 初始化应用程序
	Initialize() error
	
	// Run 运行应用程序
	Run() error
	
	// Shutdown 关闭应用程序
	Shutdown() error
	
	// GetEventBus 获取事件总线
	GetEventBus() *events.Bus
	
	// GetDocumentManager 获取文档管理器
	GetDocumentManager() document.DocumentManager
	
	// GetThemeManager 获取主题管理器
	GetThemeManager() theme.ThemeManager
	
	// GetReaderController 获取阅读器控制器
	GetReaderController() reader.ReaderController
	
	// GetAIService 获取AI服务
	GetAIService() ai.AIService
}

// ServiceContainer 服务容器接口 - 依赖注入
type ServiceContainer interface {
	// Register 注册服务
	Register(name string, service interface{})
	
	// Get 获取服务
	Get(name string) interface{}
	
	// Has 检查服务是否存在
	Has(name string) bool
	
	// Remove 移除服务
	Remove(name string)
}

// Configuration 配置管理接口
type Configuration interface {
	// Load 加载配置
	Load() error
	
	// Save 保存配置
	Save() error
	
	// Get 获取配置值
	Get(key string) interface{}
	
	// Set 设置配置值
	Set(key string, value interface{})
	
	// GetString 获取字符串配置
	GetString(key string) string
	
	// GetInt 获取整数配置
	GetInt(key string) int
	
	// GetBool 获取布尔配置
	GetBool(key string) bool
	
	// GetFloat 获取浮点数配置
	GetFloat(key string) float64
}

// Plugin 插件接口
type Plugin interface {
	// GetName 获取插件名称
	GetName() string
	
	// GetVersion 获取插件版本
	GetVersion() string
	
	// Initialize 初始化插件
	Initialize(app Application) error
	
	// Activate 激活插件
	Activate() error
	
	// Deactivate 停用插件
	Deactivate() error
	
	// GetDependencies 获取依赖项
	GetDependencies() []string
}

// PluginManager 插件管理器接口
type PluginManager interface {
	// LoadPlugin 加载插件
	LoadPlugin(path string) error
	
	// UnloadPlugin 卸载插件
	UnloadPlugin(name string) error
	
	// GetPlugin 获取插件
	GetPlugin(name string) Plugin
	
	// GetLoadedPlugins 获取已加载的插件
	GetLoadedPlugins() []Plugin
	
	// ActivatePlugin 激活插件
	ActivatePlugin(name string) error
	
	// DeactivatePlugin 停用插件
	DeactivatePlugin(name string) error
}