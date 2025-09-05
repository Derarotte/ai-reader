package events

import (
	"sync"
)

// EventType 事件类型
type EventType string

const (
	DocumentOpened    EventType = "document_opened"
	DocumentClosed    EventType = "document_closed"
	TextSelected      EventType = "text_selected"
	PageChanged       EventType = "page_changed"
	ThemeChanged      EventType = "theme_changed"
	AIAnalysisRequest EventType = "ai_analysis_request"
	AIAnalysisResult  EventType = "ai_analysis_result"
)

// Event 事件数据结构
type Event struct {
	Type    EventType
	Payload interface{}
}

// Handler 事件处理器函数类型
type Handler func(event Event)

// Bus 事件总线 - 实现组件间松耦合通信
type Bus struct {
	mu       sync.RWMutex
	handlers map[EventType][]Handler
}

// NewBus 创建新的事件总线
func NewBus() *Bus {
	return &Bus{
		handlers: make(map[EventType][]Handler),
	}
}

// Subscribe 订阅事件
func (b *Bus) Subscribe(eventType EventType, handler Handler) {
	b.mu.Lock()
	defer b.mu.Unlock()
	
	b.handlers[eventType] = append(b.handlers[eventType], handler)
}

// Publish 发布事件
func (b *Bus) Publish(event Event) {
	b.mu.RLock()
	handlers := b.handlers[event.Type]
	b.mu.RUnlock()
	
	// 异步执行所有处理器
	for _, handler := range handlers {
		go handler(event)
	}
}

// Unsubscribe 取消订阅（可选功能）
func (b *Bus) Unsubscribe(eventType EventType, handler Handler) {
	b.mu.Lock()
	defer b.mu.Unlock()
	
	handlers := b.handlers[eventType]
	for i, h := range handlers {
		// 比较函数指针（简化实现）
		if &h == &handler {
			b.handlers[eventType] = append(handlers[:i], handlers[i+1:]...)
			break
		}
	}
}