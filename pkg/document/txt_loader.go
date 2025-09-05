package document

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

// TextDocument TXT文档实现
type TextDocument struct {
	title    string
	content  string
	pages    []string
	metadata Metadata
}

// NewTextDocument 创建新的文本文档
func NewTextDocument(title, content string) *TextDocument {
	doc := &TextDocument{
		title:   title,
		content: content,
	}
	doc.generatePages()
	doc.generateMetadata()
	return doc
}

func (d *TextDocument) GetContent() (string, error) {
	return d.content, nil
}

func (d *TextDocument) GetTitle() string {
	return d.title
}

func (d *TextDocument) GetPages() int {
	return len(d.pages)
}

func (d *TextDocument) GetPage(pageNum int) (string, error) {
	if pageNum < 1 || pageNum > len(d.pages) {
		return "", ErrInvalidPage
	}
	return d.pages[pageNum-1], nil
}

func (d *TextDocument) GetMetadata() Metadata {
	return d.metadata
}

func (d *TextDocument) Search(query string) ([]SearchResult, error) {
	var results []SearchResult
	query = strings.ToLower(query)
	
	for pageNum, pageContent := range d.pages {
		lowerContent := strings.ToLower(pageContent)
		index := 0
		
		for {
			pos := strings.Index(lowerContent[index:], query)
			if pos == -1 {
				break
			}
			
			actualPos := index + pos
			// 获取上下文（前后50个字符）
			start := max(0, actualPos-50)
			end := min(len(pageContent), actualPos+len(query)+50)
			context := pageContent[start:end]
			
			results = append(results, SearchResult{
				PageNumber: pageNum + 1,
				Context:    context,
				Position:   actualPos,
			})
			
			index = actualPos + len(query)
		}
	}
	
	return results, nil
}

func (d *TextDocument) Close() error {
	// 文本文档无需特殊关闭操作
	return nil
}

// generatePages 将内容分页（每页约1000字符）
func (d *TextDocument) generatePages() {
	const pageSize = 1000
	content := d.content
	
	if len(content) <= pageSize {
		d.pages = []string{content}
		return
	}
	
	var pages []string
	for len(content) > pageSize {
		// 寻找合适的分页点（段落或句子结束）
		breakPoint := pageSize
		for i := pageSize; i > pageSize-200 && i > 0; i-- {
			if content[i] == '\n' || content[i] == '.' {
				breakPoint = i + 1
				break
			}
		}
		
		pages = append(pages, content[:breakPoint])
		content = content[breakPoint:]
	}
	
	if len(content) > 0 {
		pages = append(pages, content)
	}
	
	d.pages = pages
}

// generateMetadata 生成文档元数据
func (d *TextDocument) generateMetadata() {
	wordCount := len(strings.Fields(d.content))
	d.metadata = Metadata{
		Title:     d.title,
		Author:    "",
		Subject:   "",
		Creator:   "AI Reader",
		PageCount: len(d.pages),
		WordCount: wordCount,
		FileSize:  int64(len(d.content)),
		Format:    "text/plain",
	}
}

// TxtLoader TXT文件加载器
type TxtLoader struct{}

func NewTxtLoader() *TxtLoader {
	return &TxtLoader{}
}

func (l *TxtLoader) CanHandle(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".txt"
}

func (l *TxtLoader) LoadFromFile(filename string) (Document, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	return l.LoadFromReader(file, filename)
}

func (l *TxtLoader) LoadFromReader(reader io.Reader, filename string) (Document, error) {
	var content strings.Builder
	scanner := bufio.NewScanner(reader)
	
	for scanner.Scan() {
		content.WriteString(scanner.Text())
		content.WriteString("\n")
	}
	
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	
	// 验证UTF-8编码
	contentStr := content.String()
	if !utf8.ValidString(contentStr) {
		return nil, ErrInvalidEncoding
	}
	
	title := filepath.Base(filename)
	if ext := filepath.Ext(title); ext != "" {
		title = title[:len(title)-len(ext)]
	}
	
	return NewTextDocument(title, contentStr), nil
}

// 工具函数
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}