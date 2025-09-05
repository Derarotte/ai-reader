package document

import (
	"io"
)

// Document 文档接口 - 统一不同格式文档的处理
type Document interface {
	// GetContent 获取文档内容
	GetContent() (string, error)
	
	// GetTitle 获取文档标题
	GetTitle() string
	
	// GetPages 获取总页数（如果支持分页）
	GetPages() int
	
	// GetPage 获取指定页的内容
	GetPage(pageNum int) (string, error)
	
	// GetMetadata 获取文档元数据
	GetMetadata() Metadata
	
	// Search 在文档中搜索
	Search(query string) ([]SearchResult, error)
	
	// Close 关闭文档资源
	Close() error
}

// Metadata 文档元数据
type Metadata struct {
	Title       string
	Author      string
	Subject     string
	Creator     string
	CreatedAt   string
	ModifiedAt  string
	PageCount   int
	WordCount   int
	FileSize    int64
	Format      string
}

// SearchResult 搜索结果
type SearchResult struct {
	PageNumber int
	Context    string
	Position   int
}

// DocumentLoader 文档加载器接口
type DocumentLoader interface {
	// CanHandle 检查是否能处理指定格式
	CanHandle(filename string) bool
	
	// LoadFromFile 从文件加载文档
	LoadFromFile(filename string) (Document, error)
	
	// LoadFromReader 从Reader加载文档
	LoadFromReader(reader io.Reader, filename string) (Document, error)
}

// DocumentManager 文档管理器接口
type DocumentManager interface {
	// RegisterLoader 注册文档加载器
	RegisterLoader(loader DocumentLoader)
	
	// LoadDocument 加载文档
	LoadDocument(filename string) (Document, error)
	
	// GetSupportedFormats 获取支持的文档格式
	GetSupportedFormats() []string
}