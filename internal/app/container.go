package app

import "sync"

// Container 服务容器实现
type Container struct {
	services map[string]interface{}
	mu       sync.RWMutex
}

// NewServiceContainer 创建新的服务容器
func NewServiceContainer() *Container {
	return &Container{
		services: make(map[string]interface{}),
	}
}

func (c *Container) Register(name string, service interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	c.services[name] = service
}

func (c *Container) Get(name string) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	return c.services[name]
}

func (c *Container) Has(name string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	_, exists := c.services[name]
	return exists
}

func (c *Container) Remove(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	delete(c.services, name)
}