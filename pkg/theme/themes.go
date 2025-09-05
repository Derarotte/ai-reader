package theme

import (
	"image/color"
)

// ClassicTheme 经典主题
type ClassicTheme struct{}

func NewClassicTheme() *ClassicTheme {
	return &ClassicTheme{}
}

func (t *ClassicTheme) GetName() string {
	return "classic"
}

func (t *ClassicTheme) GetColors() ColorScheme {
	return ColorScheme{
		Background:    color.RGBA{245, 245, 220, 255}, // 米黄色
		Text:          color.RGBA{47, 27, 20, 255},    // 深棕色
		Selection:     color.RGBA{255, 215, 0, 100},   // 金色高亮
		Border:        color.RGBA{139, 69, 19, 255},   // 棕色边框
		Highlight:     color.RGBA{255, 165, 0, 255},   // 橙色
		ButtonPrimary: color.RGBA{139, 69, 19, 255},   // 棕色按钮
		ButtonHover:   color.RGBA{160, 82, 45, 255},   // 悬停棕色
	}
}

func (t *ClassicTheme) GetFonts() FontScheme {
	return FontScheme{
		Primary: FontConfig{
			Family: "serif",
			Size:   14,
			Weight: "normal",
		},
		Secondary: FontConfig{
			Family: "serif",
			Size:   12,
			Weight: "normal",
		},
		Monospace: FontConfig{
			Family: "monospace",
			Size:   12,
			Weight: "normal",
		},
	}
}

func (t *ClassicTheme) GetTexture() TextureConfig {
	return TextureConfig{
		Type:    "paper",
		Opacity: 0.3,
		Scale:   1.0,
		BlendMode: "overlay",
		Parameters: map[string]interface{}{
			"roughness": 0.2,
			"aging":     0.1,
		},
	}
}

func (t *ClassicTheme) GetAnimation() AnimationConfig {
	return AnimationConfig{
		PageTurnType:     "flip",
		PageTurnDuration: 500,
		Easing:          "ease-out",
		EnableSounds:    true,
	}
}

// DarkTheme 夜间主题
type DarkTheme struct{}

func NewDarkTheme() *DarkTheme {
	return &DarkTheme{}
}

func (t *DarkTheme) GetName() string {
	return "dark"
}

func (t *DarkTheme) GetColors() ColorScheme {
	return ColorScheme{
		Background:    color.RGBA{26, 26, 26, 255},    // 深灰
		Text:          color.RGBA{224, 224, 224, 255}, // 浅灰文字
		Selection:     color.RGBA{64, 128, 255, 100},  // 蓝色高亮
		Border:        color.RGBA{64, 64, 64, 255},    // 灰色边框
		Highlight:     color.RGBA{255, 193, 7, 255},   // 黄色高亮
		ButtonPrimary: color.RGBA{52, 58, 64, 255},    // 深灰按钮
		ButtonHover:   color.RGBA{73, 80, 87, 255},    // 悬停灰色
	}
}

func (t *DarkTheme) GetFonts() FontScheme {
	return FontScheme{
		Primary: FontConfig{
			Family: "sans-serif",
			Size:   14,
			Weight: "normal",
		},
		Secondary: FontConfig{
			Family: "sans-serif",
			Size:   12,
			Weight: "light",
		},
		Monospace: FontConfig{
			Family: "monospace",
			Size:   12,
			Weight: "normal",
		},
	}
}

func (t *DarkTheme) GetTexture() TextureConfig {
	return TextureConfig{
		Type:    "solid",
		Opacity: 1.0,
		Scale:   1.0,
		BlendMode: "normal",
		Parameters: map[string]interface{}{
			"smoothness": 1.0,
		},
	}
}

func (t *DarkTheme) GetAnimation() AnimationConfig {
	return AnimationConfig{
		PageTurnType:     "fade",
		PageTurnDuration: 300,
		Easing:          "ease-in-out",
		EnableSounds:    false,
	}
}

// GreenTheme 护眼绿色主题
type GreenTheme struct{}

func NewGreenTheme() *GreenTheme {
	return &GreenTheme{}
}

func (t *GreenTheme) GetName() string {
	return "green"
}

func (t *GreenTheme) GetColors() ColorScheme {
	return ColorScheme{
		Background:    color.RGBA{240, 248, 240, 255}, // 淡绿色
		Text:          color.RGBA{34, 87, 34, 255},    // 深绿色
		Selection:     color.RGBA{144, 238, 144, 100}, // 浅绿高亮
		Border:        color.RGBA{107, 142, 35, 255},  // 橄榄绿边框
		Highlight:     color.RGBA{50, 205, 50, 255},   // 酸橙绿
		ButtonPrimary: color.RGBA{34, 139, 34, 255},   // 森林绿按钮
		ButtonHover:   color.RGBA{60, 179, 113, 255},  // 悬停绿色
	}
}

func (t *GreenTheme) GetFonts() FontScheme {
	return FontScheme{
		Primary: FontConfig{
			Family: "sans-serif",
			Size:   14,
			Weight: "normal",
		},
		Secondary: FontConfig{
			Family: "sans-serif",
			Size:   12,
			Weight: "normal",
		},
		Monospace: FontConfig{
			Family: "monospace",
			Size:   12,
			Weight: "normal",
		},
	}
}

func (t *GreenTheme) GetTexture() TextureConfig {
	return TextureConfig{
		Type:    "linen",
		Opacity: 0.2,
		Scale:   1.0,
		BlendMode: "multiply",
		Parameters: map[string]interface{}{
			"naturalness": 0.8,
		},
	}
}

func (t *GreenTheme) GetAnimation() AnimationConfig {
	return AnimationConfig{
		PageTurnType:     "slide",
		PageTurnDuration: 400,
		Easing:          "ease-out",
		EnableSounds:    false,
	}
}

// MinimalTheme 简约主题
type MinimalTheme struct{}

func NewMinimalTheme() *MinimalTheme {
	return &MinimalTheme{}
}

func (t *MinimalTheme) GetName() string {
	return "minimal"
}

func (t *MinimalTheme) GetColors() ColorScheme {
	return ColorScheme{
		Background:    color.RGBA{255, 255, 255, 255}, // 纯白
		Text:          color.RGBA{33, 37, 41, 255},    // 深灰文字
		Selection:     color.RGBA{0, 123, 255, 100},   // 蓝色高亮
		Border:        color.RGBA{206, 212, 218, 255}, // 浅灰边框
		Highlight:     color.RGBA{255, 193, 7, 255},   // 黄色高亮
		ButtonPrimary: color.RGBA{0, 123, 255, 255},   // 蓝色按钮
		ButtonHover:   color.RGBA{0, 86, 179, 255},    // 悬停蓝色
	}
}

func (t *MinimalTheme) GetFonts() FontScheme {
	return FontScheme{
		Primary: FontConfig{
			Family: "sans-serif",
			Size:   14,
			Weight: "normal",
		},
		Secondary: FontConfig{
			Family: "sans-serif",
			Size:   12,
			Weight: "light",
		},
		Monospace: FontConfig{
			Family: "monospace",
			Size:   12,
			Weight: "normal",
		},
	}
}

func (t *MinimalTheme) GetTexture() TextureConfig {
	return TextureConfig{
		Type:    "solid",
		Opacity: 1.0,
		Scale:   1.0,
		BlendMode: "normal",
		Parameters: map[string]interface{}{
			"flatness": 1.0,
		},
	}
}

func (t *MinimalTheme) GetAnimation() AnimationConfig {
	return AnimationConfig{
		PageTurnType:     "fade",
		PageTurnDuration: 200,
		Easing:          "linear",
		EnableSounds:    false,
	}
}