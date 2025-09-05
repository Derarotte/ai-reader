package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/canvas"
	"image/color"
	"strings"
)

// SelectableText 可选择的文本组件
type SelectableText struct {
	widget.BaseWidget
	
	content      string
	richText     *widget.RichText
	overlay      *canvas.Rectangle
	
	// 选择状态
	isSelecting  bool
	selectionStart fyne.Position
	selectionEnd   fyne.Position
	selectedText   string
	
	// 回调函数
	OnSelectionChanged func(selectedText string)
}

// NewSelectableText 创建新的可选择文本组件
func NewSelectableText(content string) *SelectableText {
	st := &SelectableText{
		content:  content,
		richText: widget.NewRichText(),
		overlay:  canvas.NewRectangle(color.RGBA{0, 123, 255, 50}), // 半透明蓝色选择框
	}
	
	st.richText.ParseMarkdown(content)
	st.richText.Wrapping = fyne.TextWrapWord
	st.overlay.Hide()
	
	st.ExtendBaseWidget(st)
	return st
}

// SetContent 设置文本内容
func (st *SelectableText) SetContent(content string) {
	st.content = content
	st.richText.ParseMarkdown(content)
	st.clearSelection()
}

// GetContent 获取文本内容
func (st *SelectableText) GetContent() string {
	return st.content
}

// GetSelectedText 获取选中的文本
func (st *SelectableText) GetSelectedText() string {
	return st.selectedText
}

// clearSelection 清除选择
func (st *SelectableText) clearSelection() {
	st.selectedText = ""
	st.overlay.Hide()
	st.isSelecting = false
	if st.OnSelectionChanged != nil {
		st.OnSelectionChanged("")
	}
}

// CreateRenderer 创建渲染器
func (st *SelectableText) CreateRenderer() fyne.WidgetRenderer {
	return &selectableTextRenderer{
		selectableText: st,
		objects:       []fyne.CanvasObject{st.richText, st.overlay},
	}
}

// Tapped 处理单击事件
func (st *SelectableText) Tapped(evt *fyne.PointEvent) {
	st.clearSelection()
}

// TappedSecondary 处理右键点击
func (st *SelectableText) TappedSecondary(evt *fyne.PointEvent) {
	// 可以在这里添加右键菜单
}

// Dragged 处理拖拽事件（用于文本选择）
func (st *SelectableText) Dragged(evt *fyne.DragEvent) {
	if !st.isSelecting {
		st.isSelecting = true
		st.selectionStart = evt.Position
		st.overlay.Show()
	}
	
	st.selectionEnd = evt.Position
	st.updateSelection()
}

// DragEnd 拖拽结束
func (st *SelectableText) DragEnd() {
	if st.isSelecting {
		st.finalizeSelection()
	}
}

// updateSelection 更新选择区域
func (st *SelectableText) updateSelection() {
	if !st.isSelecting {
		return
	}
	
	// 计算选择矩形
	x1, y1 := st.selectionStart.X, st.selectionStart.Y
	x2, y2 := st.selectionEnd.X, st.selectionEnd.Y
	
	// 确保矩形坐标正确
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	if y1 > y2 {
		y1, y2 = y2, y1
	}
	
	st.overlay.Move(fyne.NewPos(x1, y1))
	st.overlay.Resize(fyne.NewSize(x2-x1, y2-y1))
	st.overlay.Refresh()
}

// finalizeSelection 完成选择
func (st *SelectableText) finalizeSelection() {
	st.isSelecting = false
	
	// 这里应该根据选择区域提取实际的文本
	// 为了简化，我们提取内容的一部分作为示例
	st.selectedText = st.extractTextFromSelection()
	
	if st.OnSelectionChanged != nil && st.selectedText != "" {
		st.OnSelectionChanged(st.selectedText)
	}
}

// extractTextFromSelection 从选择区域提取文本
func (st *SelectableText) extractTextFromSelection() string {
	if st.content == "" {
		return ""
	}
	
	// 简化实现：基于选择区域的相对位置提取文本
	// 实际实现中需要更复杂的文本位置计算
	lines := strings.Split(st.content, "\n")
	if len(lines) == 0 {
		return ""
	}
	
	// 模拟选择了第一段内容
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "#") {
			// 返回第一个非标题行
			if len(line) > 50 {
				return line[:50] + "..."
			}
			return line
		}
	}
	
	return "选中的文本内容"
}

// selectableTextRenderer 渲染器实现
type selectableTextRenderer struct {
	selectableText *SelectableText
	objects        []fyne.CanvasObject
}

func (r *selectableTextRenderer) Layout(size fyne.Size) {
	r.selectableText.richText.Resize(size)
}

func (r *selectableTextRenderer) MinSize() fyne.Size {
	return r.selectableText.richText.MinSize()
}

func (r *selectableTextRenderer) Refresh() {
	r.selectableText.richText.Refresh()
}

func (r *selectableTextRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *selectableTextRenderer) Destroy() {
	// 清理资源
}