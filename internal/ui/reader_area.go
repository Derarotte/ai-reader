package ui

import (
	"ai-reader/internal/events"
	"ai-reader/pkg/document"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

// ReaderArea 阅读器区域
type ReaderArea struct {
	eventBus    *events.Bus
	container   *fyne.Container
	
	// UI组件
	contentArea *SelectableText
	toolbar     *fyne.Container
	pageInfo    *widget.Label
	prevBtn     *widget.Button
	nextBtn     *widget.Button
	
	// 状态
	currentDoc  document.Document
	currentPage int
	totalPages  int
	zoom        float32
}

// NewReaderArea 创建阅读器区域
func NewReaderArea(eventBus *events.Bus) *ReaderArea {
	ra := &ReaderArea{
		eventBus:    eventBus,
		currentPage: 1,
		totalPages:  1,
		zoom:        1.0,
	}
	
	ra.initializeComponents()
	ra.setupLayout()
	ra.setupEventHandlers()
	
	return ra
}

// initializeComponents 初始化组件
func (ra *ReaderArea) initializeComponents() {
	// 内容显示区域 - 使用可选择文本组件
	ra.contentArea = NewSelectableText("欢迎使用AI阅读器！\n\n请打开文档开始阅读。\n\n您可以选择文本进行AI分析。")
	
	// 页面信息
	ra.pageInfo = widget.NewLabel("第 1 页，共 1 页")
	
	// 导航按钮
	ra.prevBtn = widget.NewButton("◀ 上一页", ra.handlePreviousPage)
	ra.nextBtn = widget.NewButton("下一页 ▶", ra.handleNextPage)
	
	// 工具栏
	zoomInBtn := widget.NewButton("🔍+", ra.handleZoomIn)
	zoomOutBtn := widget.NewButton("🔍-", ra.handleZoomOut)
	zoomResetBtn := widget.NewButton("100%", ra.handleZoomReset)
	
	toolbar := container.NewHBox(
		ra.prevBtn,
		ra.pageInfo,
		ra.nextBtn,
		widget.NewSeparator(),
		zoomOutBtn,
		zoomResetBtn,
		zoomInBtn,
	)
	
	ra.toolbar = toolbar
}

// setupLayout 设置布局
func (ra *ReaderArea) setupLayout() {
	// 内容滚动区域
	scrollContent := container.NewScroll(ra.contentArea)
	scrollContent.SetMinSize(fyne.NewSize(400, 300))
	
	// 主容器
	ra.container = container.NewBorder(
		nil,          // 顶部
		ra.toolbar,   // 底部工具栏
		nil, nil,     // 左右
		scrollContent, // 中心内容
	)
}

// setupEventHandlers 设置事件处理器
func (ra *ReaderArea) setupEventHandlers() {
	// 监听文档打开事件
	ra.eventBus.Subscribe(events.DocumentOpened, func(event events.Event) {
		filename := event.Payload.(string)
		ra.loadDocument(filename)
	})
	
	// 添加文本选择处理 - 使用鼠标事件实现
	ra.setupTextSelection()
}

// loadDocument 加载文档
func (ra *ReaderArea) loadDocument(filename string) {
	// TODO: 通过文档管理器加载文档
	// 这里先用占位文本
	content := "# " + filename + "\n\n这是一个示例文档内容。\n\n## 主要功能\n\n- 智能文本分析\n- 多种主题切换\n- 丰富的翻页动画\n- 文档格式支持\n\n您可以选择文本进行AI分析，系统会为您提供背景信息和概念解释。"
	ra.contentArea.SetContent(content)
	ra.currentPage = 1
	ra.totalPages = 1
	ra.updatePageInfo()
}

// setupTextSelection 设置文本选择功能
func (ra *ReaderArea) setupTextSelection() {
	// 设置文本选择回调
	ra.contentArea.OnSelectionChanged = func(selectedText string) {
		if selectedText != "" {
			ra.eventBus.Publish(events.Event{
				Type:    events.TextSelected,
				Payload: selectedText,
			})
		}
	}
}

// getSelectedText 获取当前选中的文本
func (ra *ReaderArea) getSelectedText() string {
	return ra.contentArea.GetSelectedText()
}

// handlePreviousPage 处理上一页
func (ra *ReaderArea) handlePreviousPage() {
	if ra.currentPage > 1 {
		ra.currentPage--
		ra.updatePage()
		ra.publishPageChanged()
	}
}

// handleNextPage 处理下一页
func (ra *ReaderArea) handleNextPage() {
	if ra.currentPage < ra.totalPages {
		ra.currentPage++
		ra.updatePage()
		ra.publishPageChanged()
	}
}

// handleZoomIn 放大
func (ra *ReaderArea) handleZoomIn() {
	if ra.zoom < 3.0 {
		ra.zoom += 0.1
		ra.applyZoom()
	}
}

// handleZoomOut 缩小
func (ra *ReaderArea) handleZoomOut() {
	if ra.zoom > 0.5 {
		ra.zoom -= 0.1
		ra.applyZoom()
	}
}

// handleZoomReset 重置缩放
func (ra *ReaderArea) handleZoomReset() {
	ra.zoom = 1.0
	ra.applyZoom()
}

// applyZoom 应用缩放
func (ra *ReaderArea) applyZoom() {
	// TODO: 实现文本缩放
	// 可能需要调整字体大小或容器缩放
}

// updatePage 更新页面内容
func (ra *ReaderArea) updatePage() {
	if ra.currentDoc != nil {
		content, err := ra.currentDoc.GetPage(ra.currentPage)
		if err == nil {
			ra.contentArea.SetContent(content)
		}
	}
	ra.updatePageInfo()
}

// updatePageInfo 更新页面信息
func (ra *ReaderArea) updatePageInfo() {
	ra.pageInfo.SetText("第 " + strconv.Itoa(ra.currentPage) + " 页，共 " + strconv.Itoa(ra.totalPages) + " 页")
	
	// 更新按钮状态
	ra.prevBtn.Enable()
	ra.nextBtn.Enable()
	
	if ra.currentPage <= 1 {
		ra.prevBtn.Disable()
	}
	if ra.currentPage >= ra.totalPages {
		ra.nextBtn.Disable()
	}
}

// publishPageChanged 发布页面变化事件
func (ra *ReaderArea) publishPageChanged() {
	ra.eventBus.Publish(events.Event{
		Type: events.PageChanged,
		Payload: map[string]interface{}{
			"current": ra.currentPage,
			"total":   ra.totalPages,
		},
	})
}

// GetContainer 获取容器
func (ra *ReaderArea) GetContainer() *fyne.Container {
	return ra.container
}