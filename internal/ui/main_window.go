package ui

import (
	"ai-reader/internal/events"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// MainWindow 主窗口结构
type MainWindow struct {
	app       fyne.App
	window    fyne.Window
	eventBus  *events.Bus
	
	// UI组件
	fileTree    *widget.Tree
	readerArea  *ReaderArea
	aiPanel     *AIPanel
	statusBar   *StatusBar
	menuBar     *fyne.MainMenu
	
	// 布局容器
	leftPanel   *container.Split
	rightPanel  *container.Split
	mainContent *container.Split
}

// NewMainWindow 创建主窗口
func NewMainWindow(eventBus *events.Bus) *MainWindow {
	fyneApp := app.New()
	fyneApp.SetIcon(nil) // TODO: 添加应用图标
	
	window := fyneApp.NewWindow("AI Reader")
	window.SetMaster()
	window.Resize(fyne.NewSize(1200, 800))
	window.CenterOnScreen()
	
	mw := &MainWindow{
		app:      fyneApp,
		window:   window,
		eventBus: eventBus,
	}
	
	mw.initializeComponents()
	mw.setupLayout()
	mw.setupEventHandlers()
	
	return mw
}

// initializeComponents 初始化UI组件
func (mw *MainWindow) initializeComponents() {
	// 文件树
	mw.fileTree = mw.createFileTree()
	
	// 阅读器区域
	mw.readerArea = NewReaderArea(mw.eventBus)
	
	// AI分析面板
	mw.aiPanel = NewAIPanel(mw.eventBus)
	
	// 状态栏
	mw.statusBar = NewStatusBar()
	
	// 菜单栏
	mw.menuBar = mw.createMenuBar()
}

// setupLayout 设置布局
func (mw *MainWindow) setupLayout() {
	// 左侧面板 - 文件树
	leftContainer := container.NewBorder(
		widget.NewLabel("文档浏览"), nil, nil, nil,
		container.NewScroll(mw.fileTree),
	)
	
	// 右侧面板 - AI分析
	rightContainer := container.NewBorder(
		widget.NewLabel("AI分析"), nil, nil, nil,
		mw.aiPanel.GetContainer(),
	)
	
	// 主要内容区域
	mw.mainContent = container.NewHSplit(
		leftContainer,
		container.NewHSplit(
			mw.readerArea.GetContainer(),
			rightContainer,
		),
	)
	
	// 设置分割比例
	mw.mainContent.SetOffset(0.2) // 左侧占20%
	mw.mainContent.Trailing.(*container.Split).SetOffset(0.75) // 中间占75%，右侧占25%
	
	// 主布局
	content := container.NewBorder(
		nil, // 顶部
		mw.statusBar.GetContainer(), // 底部
		nil, nil, // 左右
		mw.mainContent, // 中心
	)
	
	mw.window.SetContent(content)
	mw.window.SetMainMenu(mw.menuBar)
}

// createFileTree 创建文件树
func (mw *MainWindow) createFileTree() *widget.Tree {
	tree := widget.NewTree(
		func(uid string) []string {
			// TODO: 实现文件系统浏览
			if uid == "" {
				return []string{"Documents", "Recent"}
			}
			return []string{}
		},
		func(uid string) bool {
			return uid == ""
		},
		func(branch bool) fyne.CanvasObject {
			if branch {
				return widget.NewLabel("📁 Folder")
			}
			return widget.NewLabel("📄 Document")
		},
		func(uid string, branch bool, obj fyne.CanvasObject) {
			label := obj.(*widget.Label)
			if branch {
				label.SetText("📁 " + uid)
			} else {
				label.SetText("📄 " + uid)
			}
		},
	)
	
	tree.OnSelected = func(uid string) {
		// 处理文件选择
		mw.eventBus.Publish(events.Event{
			Type:    events.DocumentOpened,
			Payload: uid,
		})
	}
	
	return tree
}

// createMenuBar 创建菜单栏
func (mw *MainWindow) createMenuBar() *fyne.MainMenu {
	// 文件菜单
	fileMenu := fyne.NewMenu("文件",
		fyne.NewMenuItem("打开文档...", mw.handleOpenDocument),
		fyne.NewMenuItem("最近文档", nil),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("退出", mw.handleExit),
	)
	
	// 视图菜单
	viewMenu := fyne.NewMenu("视图",
		fyne.NewMenuItem("全屏模式", mw.handleToggleFullscreen),
		fyne.NewMenuItem("专注模式", mw.handleToggleFocusMode),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("缩放", nil),
	)
	
	// 主题菜单
	themeMenu := fyne.NewMenu("主题",
		fyne.NewMenuItem("经典主题", func() { mw.handleThemeChange("classic") }),
		fyne.NewMenuItem("夜间主题", func() { mw.handleThemeChange("dark") }),
		fyne.NewMenuItem("护眼主题", func() { mw.handleThemeChange("green") }),
		fyne.NewMenuItem("简约主题", func() { mw.handleThemeChange("minimal") }),
	)
	
	// AI菜单
	aiMenu := fyne.NewMenu("AI",
		fyne.NewMenuItem("分析设置", mw.handleAISettings),
		fyne.NewMenuItem("清除历史", mw.handleClearAIHistory),
	)
	
	// 帮助菜单
	helpMenu := fyne.NewMenu("帮助",
		fyne.NewMenuItem("使用说明", mw.handleShowHelp),
		fyne.NewMenuItem("关于", mw.handleShowAbout),
	)
	
	return fyne.NewMainMenu(fileMenu, viewMenu, themeMenu, aiMenu, helpMenu)
}

// setupEventHandlers 设置事件处理器
func (mw *MainWindow) setupEventHandlers() {
	// 监听主题变化事件
	mw.eventBus.Subscribe(events.ThemeChanged, func(event events.Event) {
		// TODO: 应用新主题
	})
	
	// 监听页面变化事件
	mw.eventBus.Subscribe(events.PageChanged, func(event events.Event) {
		mw.statusBar.UpdatePageInfo(event.Payload)
	})
}

// Show 显示窗口
func (mw *MainWindow) Show() {
	mw.window.ShowAndRun()
}

// 菜单事件处理器
func (mw *MainWindow) handleOpenDocument() {
	// TODO: 实现文档打开对话框
}

func (mw *MainWindow) handleExit() {
	mw.app.Quit()
}

func (mw *MainWindow) handleToggleFullscreen() {
	mw.window.SetFullScreen(!mw.window.FullScreen())
}

func (mw *MainWindow) handleToggleFocusMode() {
	// TODO: 实现专注模式
}

func (mw *MainWindow) handleThemeChange(themeName string) {
	mw.eventBus.Publish(events.Event{
		Type:    events.ThemeChanged,
		Payload: themeName,
	})
}

func (mw *MainWindow) handleAISettings() {
	// TODO: 显示AI设置对话框
}

func (mw *MainWindow) handleClearAIHistory() {
	// TODO: 清除AI历史
}

func (mw *MainWindow) handleShowHelp() {
	// TODO: 显示帮助对话框
}

func (mw *MainWindow) handleShowAbout() {
	// TODO: 显示关于对话框
}