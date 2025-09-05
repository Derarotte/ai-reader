package theme

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// Manager 主题管理器实现
type Manager struct {
	themes      map[string]Theme
	currentTheme Theme
	configPath  string
	mu          sync.RWMutex
}

// NewManager 创建主题管理器
func NewManager(configPath string) *Manager {
	manager := &Manager{
		themes:     make(map[string]Theme),
		configPath: configPath,
	}
	
	// 注册默认主题
	manager.RegisterTheme(NewClassicTheme())
	manager.RegisterTheme(NewDarkTheme())
	manager.RegisterTheme(NewGreenTheme())
	manager.RegisterTheme(NewMinimalTheme())
	
	// 设置默认主题
	manager.currentTheme = manager.themes["classic"]
	
	// 尝试加载配置
	manager.LoadThemeConfig()
	
	return manager
}

func (m *Manager) RegisterTheme(theme Theme) {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	m.themes[theme.GetName()] = theme
}

func (m *Manager) GetTheme(name string) Theme {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	if theme, exists := m.themes[name]; exists {
		return theme
	}
	return nil
}

func (m *Manager) GetAllThemes() []Theme {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	themes := make([]Theme, 0, len(m.themes))
	for _, theme := range m.themes {
		themes = append(themes, theme)
	}
	return themes
}

func (m *Manager) SetCurrentTheme(name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	theme, exists := m.themes[name]
	if !exists {
		return fmt.Errorf("theme not found: %s", name)
	}
	
	m.currentTheme = theme
	return nil
}

func (m *Manager) GetCurrentTheme() Theme {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	return m.currentTheme
}

// ThemeConfig 主题配置结构
type ThemeConfig struct {
	CurrentTheme string                    `json:"current_theme"`
	CustomColors map[string]interface{}    `json:"custom_colors,omitempty"`
	CustomFonts  map[string]interface{}    `json:"custom_fonts,omitempty"`
}

func (m *Manager) SaveThemeConfig() error {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	config := ThemeConfig{
		CurrentTheme: m.currentTheme.GetName(),
		CustomColors: make(map[string]interface{}),
		CustomFonts:  make(map[string]interface{}),
	}
	
	// 确保配置目录存在
	if err := os.MkdirAll(filepath.Dir(m.configPath), 0755); err != nil {
		return err
	}
	
	// 写入配置文件
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	
	return os.WriteFile(m.configPath, data, 0644)
}

func (m *Manager) LoadThemeConfig() error {
	// 如果配置文件不存在，使用默认配置
	if _, err := os.Stat(m.configPath); os.IsNotExist(err) {
		return nil
	}
	
	data, err := os.ReadFile(m.configPath)
	if err != nil {
		return err
	}
	
	var config ThemeConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return err
	}
	
	// 应用配置
	if config.CurrentTheme != "" {
		if theme := m.GetTheme(config.CurrentTheme); theme != nil {
			m.mu.Lock()
			m.currentTheme = theme
			m.mu.Unlock()
		}
	}
	
	return nil
}

// GetThemeNames 获取所有主题名称
func (m *Manager) GetThemeNames() []string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	names := make([]string, 0, len(m.themes))
	for name := range m.themes {
		names = append(names, name)
	}
	return names
}