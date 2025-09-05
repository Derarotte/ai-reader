package reader

import (
	"ai-reader/pkg/document"
	"ai-reader/pkg/theme"
)

// ReaderView 阅读器视图接口
type ReaderView interface {
	// DisplayDocument 显示文档
	DisplayDocument(doc document.Document) error
	
	// SetPage 设置当前页
	SetPage(pageNum int) error
	
	// GetCurrentPage 获取当前页
	GetCurrentPage() int
	
	// SetZoom 设置缩放级别
	SetZoom(level float32)
	
	// GetZoom 获取当前缩放级别
	GetZoom() float32
	
	// ApplyTheme 应用主题
	ApplyTheme(theme theme.Theme)
	
	// GetSelectedText 获取选中的文本
	GetSelectedText() string
	
	// ClearSelection 清除选择
	ClearSelection()
	
	// ScrollToPosition 滚动到指定位置
	ScrollToPosition(position float32)
	
	// Refresh 刷新显示
	Refresh()
}

// TextSelector 文本选择器接口
type TextSelector interface {
	// StartSelection 开始选择
	StartSelection(x, y float32)
	
	// UpdateSelection 更新选择
	UpdateSelection(x, y float32)
	
	// EndSelection 结束选择
	EndSelection() string
	
	// GetSelectionBounds 获取选择范围
	GetSelectionBounds() (start, end int)
	
	// SelectWord 选择单词
	SelectWord(x, y float32) string
	
	// SelectSentence 选择句子
	SelectSentence(x, y float32) string
	
	// SelectParagraph 选择段落
	SelectParagraph(x, y float32) string
}

// PageTurner 翻页器接口
type PageTurner interface {
	// NextPage 下一页
	NextPage() error
	
	// PreviousPage 上一页
	PreviousPage() error
	
	// GoToPage 跳转到指定页
	GoToPage(pageNum int) error
	
	// SetTransition 设置翻页动画
	SetTransition(transitionType string)
	
	// GetTransitionTypes 获取支持的翻页动画类型
	GetTransitionTypes() []string
}

// ReaderController 阅读器控制器接口
type ReaderController interface {
	// Initialize 初始化
	Initialize() error
	
	// LoadDocument 加载文档
	LoadDocument(filename string) error
	
	// CloseDocument 关闭文档
	CloseDocument() error
	
	// GetView 获取视图
	GetView() ReaderView
	
	// GetSelector 获取选择器
	GetSelector() TextSelector
	
	// GetPageTurner 获取翻页器
	GetPageTurner() PageTurner
	
	// HandleUserInput 处理用户输入
	HandleUserInput(inputType string, data interface{})
	
	// Cleanup 清理资源
	Cleanup()
}