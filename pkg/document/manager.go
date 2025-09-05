package document

import (
	"fmt"
	"path/filepath"
	"strings"
	"sync"
)

// Manager 文档管理器实现
type Manager struct {
	loaders []DocumentLoader
	mu      sync.RWMutex
}

// NewManager 创建新的文档管理器
func NewManager() *Manager {
	manager := &Manager{
		loaders: make([]DocumentLoader, 0),
	}
	
	// 注册默认加载器
	manager.RegisterLoader(NewTxtLoader())
	// TODO: 添加其他格式的加载器
	// manager.RegisterLoader(NewMarkdownLoader())
	// manager.RegisterLoader(NewPDFLoader())
	
	return manager
}

func (m *Manager) RegisterLoader(loader DocumentLoader) {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	m.loaders = append(m.loaders, loader)
}

func (m *Manager) LoadDocument(filename string) (Document, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	// 寻找合适的加载器
	for _, loader := range m.loaders {
		if loader.CanHandle(filename) {
			return loader.LoadFromFile(filename)
		}
	}
	
	return nil, fmt.Errorf("%w: %s", ErrUnsupportedFormat, filepath.Ext(filename))
}

func (m *Manager) GetSupportedFormats() []string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	formats := make(map[string]bool)
	
	// 这里简化实现，实际应该让每个loader提供支持的格式列表
	testFiles := []string{
		"test.txt",
		"test.md",
		"test.pdf",
		"test.docx",
		"test.html",
	}
	
	for _, testFile := range testFiles {
		for _, loader := range m.loaders {
			if loader.CanHandle(testFile) {
				ext := strings.ToLower(filepath.Ext(testFile))
				formats[ext] = true
			}
		}
	}
	
	result := make([]string, 0, len(formats))
	for format := range formats {
		result = append(result, format)
	}
	
	return result
}