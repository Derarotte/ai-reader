package ui

import (
	"ai-reader/internal/events"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// AIPanel AI分析面板
type AIPanel struct {
	eventBus    *events.Bus
	container   *fyne.Container
	
	// UI组件
	analysisText  *widget.RichText
	statusLabel   *widget.Label
	analyzeBtn    *widget.Button
	clearBtn      *widget.Button
	historyList   *widget.List
	
	// 状态
	isAnalyzing   bool
	analysisHistory []string
	selectedText  string
}

// NewAIPanel 创建AI分析面板
func NewAIPanel(eventBus *events.Bus) *AIPanel {
	ap := &AIPanel{
		eventBus:        eventBus,
		analysisHistory: make([]string, 0),
	}
	
	ap.initializeComponents()
	ap.setupLayout()
	ap.setupEventHandlers()
	
	return ap
}

// initializeComponents 初始化组件
func (ap *AIPanel) initializeComponents() {
	// 分析结果显示区域
	ap.analysisText = widget.NewRichText()
	ap.analysisText.Wrapping = fyne.TextWrapWord
	ap.analysisText.ParseMarkdown("*选择文本进行AI分析*")
	
	// 状态标签
	ap.statusLabel = widget.NewLabel("就绪")
	
	// 分析按钮
	ap.analyzeBtn = widget.NewButton("📝 分析选中文本", ap.handleAnalyze)
	ap.analyzeBtn.Disable() // 初始禁用
	
	// 清除按钮
	ap.clearBtn = widget.NewButton("🗑️ 清除", ap.handleClear)
	
	// 历史记录列表
	ap.historyList = widget.NewList(
		func() int {
			return len(ap.analysisHistory)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id int, obj fyne.CanvasObject) {
			if id < len(ap.analysisHistory) {
				obj.(*widget.Label).SetText(ap.analysisHistory[id])
			}
		},
	)
	
	ap.historyList.OnSelected = func(id int) {
		if id < len(ap.analysisHistory) {
			// 显示历史分析结果
			ap.analysisText.ParseMarkdown(ap.analysisHistory[id])
		}
	}
}

// setupLayout 设置布局
func (ap *AIPanel) setupLayout() {
	// 按钮栏
	buttonBar := container.NewHBox(
		ap.analyzeBtn,
		ap.clearBtn,
	)
	
	// 分析结果区域（可滚动）
	analysisScroll := container.NewScroll(ap.analysisText)
	analysisScroll.SetMinSize(fyne.NewSize(250, 200))
	
	// 历史记录区域
	historyContainer := container.NewBorder(
		widget.NewCard("", "分析历史", nil),
		nil, nil, nil,
		container.NewScroll(ap.historyList),
	)
	historyContainer.Resize(fyne.NewSize(250, 150))
	
	// 主容器 - 垂直分割
	mainContent := container.NewVSplit(
		container.NewBorder(
			widget.NewCard("", "分析结果", nil),
			buttonBar,
			nil, nil,
			analysisScroll,
		),
		historyContainer,
	)
	mainContent.SetOffset(0.7) // 分析结果占70%
	
	// 底部状态
	ap.container = container.NewBorder(
		nil,             // 顶部
		ap.statusLabel,  // 底部状态
		nil, nil,        // 左右
		mainContent,     // 中心
	)
}

// setupEventHandlers 设置事件处理器
func (ap *AIPanel) setupEventHandlers() {
	// 监听文本选择事件
	ap.eventBus.Subscribe(events.TextSelected, func(event events.Event) {
		ap.selectedText = event.Payload.(string)
		ap.analyzeBtn.Enable()
		ap.statusLabel.SetText("已选择文本，可进行分析")
	})
	
	// 监听AI分析结果事件
	ap.eventBus.Subscribe(events.AIAnalysisResult, func(event events.Event) {
		result := event.Payload.(string)
		ap.displayAnalysisResult(result)
		ap.addToHistory(result)
		ap.isAnalyzing = false
		ap.updateUIState()
	})
}

// handleAnalyze 处理分析请求
func (ap *AIPanel) handleAnalyze() {
	if ap.selectedText == "" {
		return
	}
	
	ap.isAnalyzing = true
	ap.updateUIState()
	
	// 发布AI分析请求事件
	ap.eventBus.Publish(events.Event{
		Type:    events.AIAnalysisRequest,
		Payload: ap.selectedText,
	})
}

// handleClear 处理清除操作
func (ap *AIPanel) handleClear() {
	ap.analysisText.ParseMarkdown("*选择文本进行AI分析*")
	ap.selectedText = ""
	ap.analyzeBtn.Disable()
	ap.statusLabel.SetText("就绪")
}

// displayAnalysisResult 显示分析结果
func (ap *AIPanel) displayAnalysisResult(result string) {
	ap.analysisText.ParseMarkdown("## 分析结果\n\n" + result)
	ap.statusLabel.SetText("分析完成")
}

// addToHistory 添加到历史记录
func (ap *AIPanel) addToHistory(result string) {
	// 限制历史记录数量
	if len(ap.analysisHistory) >= 10 {
		ap.analysisHistory = ap.analysisHistory[1:]
	}
	
	// 添加新记录（只保存前50个字符作为标题）
	title := result
	if len(title) > 50 {
		title = title[:50] + "..."
	}
	
	ap.analysisHistory = append(ap.analysisHistory, result)
	ap.historyList.Refresh()
}

// updateUIState 更新UI状态
func (ap *AIPanel) updateUIState() {
	if ap.isAnalyzing {
		ap.analyzeBtn.SetText("🔄 分析中...")
		ap.analyzeBtn.Disable()
		ap.statusLabel.SetText("正在分析...")
	} else {
		ap.analyzeBtn.SetText("📝 分析选中文本")
		if ap.selectedText != "" {
			ap.analyzeBtn.Enable()
		}
	}
}

// GetContainer 获取容器
func (ap *AIPanel) GetContainer() *fyne.Container {
	return ap.container
}