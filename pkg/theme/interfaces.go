package theme

import (
	"image/color"
)

// Theme 主题接口
type Theme interface {
	// GetName 获取主题名称
	GetName() string
	
	// GetColors 获取颜色配置
	GetColors() ColorScheme
	
	// GetFonts 获取字体配置
	GetFonts() FontScheme
	
	// GetTexture 获取纹理配置
	GetTexture() TextureConfig
	
	// GetAnimation 获取动画配置
	GetAnimation() AnimationConfig
}

// ColorScheme 颜色方案
type ColorScheme struct {
	Background    color.Color
	Text          color.Color
	Selection     color.Color
	Border        color.Color
	Highlight     color.Color
	ButtonPrimary color.Color
	ButtonHover   color.Color
}

// FontScheme 字体方案
type FontScheme struct {
	Primary   FontConfig
	Secondary FontConfig
	Monospace FontConfig
}

// FontConfig 字体配置
type FontConfig struct {
	Family string
	Size   float32
	Weight string // "normal", "bold", "light"
}

// TextureConfig 纹理配置
type TextureConfig struct {
	Type       string  // "solid", "paper", "parchment", etc.
	Opacity    float32 // 0.0 - 1.0
	Scale      float32
	BlendMode  string
	Parameters map[string]interface{} // 额外参数
}

// AnimationConfig 动画配置
type AnimationConfig struct {
	PageTurnType     string        // "fade", "slide", "flip", "wave"
	PageTurnDuration int          // 毫秒
	Easing          string        // "linear", "ease-in", "ease-out"
	EnableSounds    bool
}

// ThemeManager 主题管理器接口
type ThemeManager interface {
	// RegisterTheme 注册主题
	RegisterTheme(theme Theme)
	
	// GetTheme 获取主题
	GetTheme(name string) Theme
	
	// GetAllThemes 获取所有主题
	GetAllThemes() []Theme
	
	// SetCurrentTheme 设置当前主题
	SetCurrentTheme(name string) error
	
	// GetCurrentTheme 获取当前主题
	GetCurrentTheme() Theme
	
	// SaveThemeConfig 保存主题配置
	SaveThemeConfig() error
	
	// LoadThemeConfig 加载主题配置
	LoadThemeConfig() error
}