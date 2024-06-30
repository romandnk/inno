package inmem

import "sync"

type Cache[T any] struct {
	mu sync.RWMutex
	m  map[string]T
}

func New[T any](capacity int) *Cache[T] {
	return &Cache[T]{
		m:  make(map[string]T, capacity),
		mu: sync.RWMutex{},
	}
}

func (c *Cache[T]) Set(key string, value T) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = value
}

func (c *Cache[T]) Get(key string) (value T, ok bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok = c.m[key]
	return value, ok
}
