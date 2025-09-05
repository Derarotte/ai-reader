package app

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
)

// Config 配置实现
type Config struct {
	configPath string
	data       map[string]interface{}
	mu         sync.RWMutex
}

// NewConfig 创建配置
func NewConfig(configPath string) *Config {
	return &Config{
		configPath: configPath,
		data:       make(map[string]interface{}),
	}
}

func (c *Config) Load() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	// 如果配置文件不存在，使用默认配置
	if _, err := os.Stat(c.configPath); os.IsNotExist(err) {
		c.setDefaults()
		return nil
	}
	
	data, err := os.ReadFile(c.configPath)
	if err != nil {
		c.setDefaults()
		return err
	}
	
	if err := json.Unmarshal(data, &c.data); err != nil {
		c.setDefaults()
		return err
	}
	
	return nil
}

func (c *Config) Save() error {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	// 确保配置目录存在
	if err := os.MkdirAll(filepath.Dir(c.configPath), 0755); err != nil {
		return err
	}
	
	data, err := json.MarshalIndent(c.data, "", "  ")
	if err != nil {
		return err
	}
	
	return os.WriteFile(c.configPath, data, 0644)
}

func (c *Config) Get(key string) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	return c.data[key]
}

func (c *Config) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	c.data[key] = value
}

func (c *Config) GetString(key string) string {
	if value := c.Get(key); value != nil {
		if str, ok := value.(string); ok {
			return str
		}
	}
	return ""
}

func (c *Config) GetInt(key string) int {
	if value := c.Get(key); value != nil {
		if num, ok := value.(float64); ok {
			return int(num)
		}
	}
	return 0
}

func (c *Config) GetBool(key string) bool {
	if value := c.Get(key); value != nil {
		if b, ok := value.(bool); ok {
			return b
		}
	}
	return false
}

func (c *Config) GetFloat(key string) float64 {
	if value := c.Get(key); value != nil {
		if f, ok := value.(float64); ok {
			return f
		}
	}
	return 0.0
}

// setDefaults 设置默认配置
func (c *Config) setDefaults() {
	c.data = map[string]interface{}{
		"window_width":      1200,
		"window_height":     800,
		"window_maximized":  false,
		"default_theme":     "classic",
		"font_size":         14,
		"auto_save":         true,
		"ai_provider":       "openai",
		"page_turn_animation": "flip",
		"enable_sounds":     true,
	}
}