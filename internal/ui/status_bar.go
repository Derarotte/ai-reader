package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fmt"
)

// StatusBar 状态栏
type StatusBar struct {
	container   *fyne.Container
	
	// UI组件
	pageInfo    *widget.Label
	zoomInfo    *widget.Label
	docInfo     *widget.Label
	progressBar *widget.ProgressBar
	statusText  *widget.Label
}

// NewStatusBar 创建状态栏
func NewStatusBar() *StatusBar {
	sb := &StatusBar{}
	sb.initializeComponents()
	sb.setupLayout()
	return sb
}

// initializeComponents 初始化组件
func (sb *StatusBar) initializeComponents() {
	// 页面信息
	sb.pageInfo = widget.NewLabel("第 1 页，共 1 页")
	
	// 缩放信息
	sb.zoomInfo = widget.NewLabel("100%")
	
	// 文档信息
	sb.docInfo = widget.NewLabel("无文档")
	
	// 进度条
	sb.progressBar = widget.NewProgressBar()
	sb.progressBar.SetValue(0)
	sb.progressBar.Hide()
	
	// 状态文本
	sb.statusText = widget.NewLabel("就绪")
}

// setupLayout 设置布局
func (sb *StatusBar) setupLayout() {
	// 左侧信息
	leftInfo := container.NewHBox(
		sb.docInfo,
		widget.NewSeparator(),
		sb.pageInfo,
		widget.NewSeparator(),
		sb.zoomInfo,
	)
	
	// 右侧状态
	rightInfo := container.NewHBox(
		sb.progressBar,
		sb.statusText,
	)
	
	// 主容器
	sb.container = container.NewBorder(
		widget.NewSeparator(), // 顶部分隔线
		nil,                   // 底部
		leftInfo,              // 左侧
		rightInfo,             // 右侧
		nil,                   // 中心为空
	)
}

// UpdatePageInfo 更新页面信息
func (sb *StatusBar) UpdatePageInfo(pageData interface{}) {
	if data, ok := pageData.(map[string]interface{}); ok {
		current := data["current"].(int)
		total := data["total"].(int)
		sb.pageInfo.SetText(fmt.Sprintf("第 %d 页，共 %d 页", current, total))
	}
}

// UpdateZoomInfo 更新缩放信息
func (sb *StatusBar) UpdateZoomInfo(zoom float32) {
	sb.zoomInfo.SetText(fmt.Sprintf("%.0f%%", zoom*100))
}

// UpdateDocInfo 更新文档信息
func (sb *StatusBar) UpdateDocInfo(docName string) {
	sb.docInfo.SetText(docName)
}

// SetStatus 设置状态文本
func (sb *StatusBar) SetStatus(status string) {
	sb.statusText.SetText(status)
}

// ShowProgress 显示进度
func (sb *StatusBar) ShowProgress(message string) {
	sb.statusText.SetText(message)
	sb.progressBar.Show()
	sb.progressBar.SetValue(0)
	sb.container.Refresh()
}

// UpdateProgress 更新进度
func (sb *StatusBar) UpdateProgress(value float64) {
	sb.progressBar.SetValue(value)
}

// HideProgress 隐藏进度
func (sb *StatusBar) HideProgress() {
	sb.progressBar.Hide()
	sb.container.Refresh()
}

// GetContainer 获取容器
func (sb *StatusBar) GetContainer() *fyne.Container {
	return sb.container
}