package ai

// AIProvider AI服务提供者接口
type AIProvider interface {
	// GetName 获取AI服务名称
	GetName() string
	
	// AnalyzeText 分析文本
	AnalyzeText(request AnalysisRequest) (*AnalysisResult, error)
	
	// IsAvailable 检查服务是否可用
	IsAvailable() bool
	
	// GetSupportedAnalysisTypes 获取支持的分析类型
	GetSupportedAnalysisTypes() []string
}

// AnalysisRequest 分析请求
type AnalysisRequest struct {
	Text        string                 `json:"text"`
	Context     string                 `json:"context,omitempty"`     // 上下文信息
	AnalysisType string                `json:"analysis_type"`         // 分析类型
	Language    string                 `json:"language,omitempty"`    // 文本语言
	Parameters  map[string]interface{} `json:"parameters,omitempty"`  // 额外参数
}

// AnalysisResult 分析结果
type AnalysisResult struct {
	ID          string                 `json:"id"`
	Type        string                 `json:"type"`
	Content     string                 `json:"content"`
	Summary     string                 `json:"summary,omitempty"`
	Keywords    []string               `json:"keywords,omitempty"`
	Concepts    []Concept              `json:"concepts,omitempty"`
	References  []Reference            `json:"references,omitempty"`
	Confidence  float32                `json:"confidence"`
	ProcessTime int64                  `json:"process_time"` // 处理时间（毫秒）
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

// Concept 概念信息
type Concept struct {
	Name        string `json:"name"`
	Definition  string `json:"definition"`
	Category    string `json:"category,omitempty"`
	Importance  float32 `json:"importance"` // 重要性 0-1
}

// Reference 参考信息
type Reference struct {
	Title       string `json:"title"`
	URL         string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type"` // "web", "book", "paper", etc.
}

// AIService AI服务管理器接口
type AIService interface {
	// RegisterProvider 注册AI提供者
	RegisterProvider(provider AIProvider)
	
	// GetProvider 获取AI提供者
	GetProvider(name string) AIProvider
	
	// GetAvailableProviders 获取可用的AI提供者
	GetAvailableProviders() []AIProvider
	
	// SetDefaultProvider 设置默认提供者
	SetDefaultProvider(name string) error
	
	// AnalyzeText 分析文本（使用默认提供者）
	AnalyzeText(request AnalysisRequest) (*AnalysisResult, error)
	
	// AnalyzeTextWithProvider 使用指定提供者分析文本
	AnalyzeTextWithProvider(providerName string, request AnalysisRequest) (*AnalysisResult, error)
}

// AnalysisCache 分析缓存接口
type AnalysisCache interface {
	// Get 获取缓存的分析结果
	Get(key string) (*AnalysisResult, bool)
	
	// Set 设置缓存
	Set(key string, result *AnalysisResult)
	
	// Delete 删除缓存
	Delete(key string)
	
	// Clear 清空缓存
	Clear()
	
	// GetCacheKey 生成缓存键
	GetCacheKey(request AnalysisRequest) string
}