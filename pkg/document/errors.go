package document

import "errors"

var (
	// ErrUnsupportedFormat 不支持的文档格式
	ErrUnsupportedFormat = errors.New("unsupported document format")
	
	// ErrInvalidPage 无效的页码
	ErrInvalidPage = errors.New("invalid page number")
	
	// ErrInvalidEncoding 无效的编码
	ErrInvalidEncoding = errors.New("invalid text encoding")
	
	// ErrDocumentClosed 文档已关闭
	ErrDocumentClosed = errors.New("document is closed")
	
	// ErrFileNotFound 文件未找到
	ErrFileNotFound = errors.New("file not found")
	
	// ErrReadPermission 读取权限错误
	ErrReadPermission = errors.New("insufficient read permissions")
)