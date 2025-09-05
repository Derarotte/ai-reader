# AI Reader

一个基于Go和Fyne开发的智能桌面阅读器，支持内容选取和AI背景分析。

## ✨ 特性

- 📖 **多格式文档支持**: TXT、Markdown、PDF等
- 🎨 **多种主题风格**: 经典、夜间、护眼、简约四种主题
- 📑 **丰富的翻页动画**: 经典翻页、淡入淡出、滑动等7种动画效果
- 🤖 **AI智能分析**: 选择文本即可获得背景分析和概念解释
- 🏗️ **松耦合架构**: 高度模块化，易于扩展和维护

## 🚀 快速开始

### 环境要求

- Go 1.21+
- 支持OpenGL的系统
- Linux: 需要安装X11开发库

### 安装依赖 (Linux)

```bash
sudo apt-get install -y pkg-config libgl1-mesa-dev libx11-dev libxrandr-dev libxinerama-dev libxcursor-dev libxi-dev
```

### 编译运行

```bash
git clone https://github.com/your-username/ai-reader.git
cd ai-reader
go mod tidy
go build -o ai-reader ./cmd/
./ai-reader
```

## 🏗️ 架构设计

项目采用高度松耦合的架构设计，主要模块包括：

### 核心组件

- **事件总线**: 组件间异步通信，实现完全解耦
- **文档管理器**: 支持多格式文档加载，易于扩展新格式
- **主题管理器**: 可插拔的主题系统，支持自定义主题
- **UI框架**: 基于Fyne的现代化界面
- **AI服务**: 可切换的AI提供者，支持多种AI服务

### 目录结构

```
ai-reader/
├── cmd/                    # 应用入口
├── internal/
│   ├── app/               # 应用核心逻辑和服务容器
│   ├── ui/                # 用户界面组件
│   ├── reader/            # 阅读器核心功能
│   ├── ai/               # AI服务集成
│   ├── events/           # 事件总线系统
│   └── config/           # 配置管理
├── pkg/
│   ├── document/         # 文档处理接口和实现
│   ├── theme/            # 主题系统
│   └── animation/        # 动画效果
└── assets/               # 静态资源
```

## 🎨 主题系统

### 内置主题

1. **📖 经典主题**: 温暖的米黄色，适合长时间阅读
2. **🌙 夜间主题**: 深色模式，保护眼睛
3. **🌿 护眼主题**: 淡绿色背景，减少眼部疲劳  
4. **✨ 简约主题**: 纯净简洁的现代设计

### 自定义主题

支持通过JSON配置文件自定义主题：

```json
{
  "current_theme": "classic",
  "custom_colors": {
    "background": "#F5F5DC",
    "text": "#2F1B14"
  }
}
```

## 📑 翻页动画

支持多种翻页动画效果：

- **📄 经典翻页**: 3D书本翻页效果
- **💫 淡入淡出**: 柔和的页面切换
- **⚡ 滑动翻页**: 类似手机的滑动效果
- **🌊 波浪翻页**: 独特的波浪形翻转
- **🎭 卷轴展开**: 古典卷轴样式
- **🔄 立体翻转**: 3D立体旋转效果
- **✨ 粒子过渡**: 科幻风格的粒子效果

## 🤖 AI功能

### 支持的分析类型

- **背景分析**: 提供文本的历史背景和语境
- **概念解释**: 解释专业术语和概念
- **关联推荐**: 推荐相关知识和资料
- **深度扩展**: 深入分析和延伸阅读

### AI服务接入

框架设计支持多种AI服务：
- OpenAI GPT系列
- Claude API  
- 本地大语言模型
- 自定义AI服务

## 🔧 开发

### 添加新的文档格式

1. 实现 `document.DocumentLoader` 接口
2. 在文档管理器中注册新的加载器

```go
type MyDocumentLoader struct{}

func (l *MyDocumentLoader) CanHandle(filename string) bool {
    return strings.HasSuffix(filename, ".mydoc")
}

func (l *MyDocumentLoader) LoadFromFile(filename string) (Document, error) {
    // 实现文档加载逻辑
}
```

### 添加新主题

1. 实现 `theme.Theme` 接口
2. 在主题管理器中注册

```go
type MyTheme struct{}

func (t *MyTheme) GetName() string { return "my_theme" }
func (t *MyTheme) GetColors() ColorScheme { /* 返回颜色配置 */ }
```

## 📝 TODO

- [ ] 完善PDF文档支持
- [ ] 实现Markdown渲染
- [ ] 添加书签和笔记功能
- [ ] 支持EPUB格式
- [ ] 实现全文搜索
- [ ] 添加阅读统计
- [ ] 支持文档导出

## 🤝 贡献

欢迎提交Issue和Pull Request！

1. Fork项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启Pull Request

## 📄 许可证

本项目采用MIT许可证 - 查看 [LICENSE](LICENSE) 文件了解详情

## 🙏 致谢

- [Fyne](https://fyne.io/) - 优秀的Go GUI框架
- 所有为开源社区贡献的开发者们