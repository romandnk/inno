package inmem

import (
	"errors"
	"sync"
)

var ErrNotFound = errors.New("not found")

type Cache struct {
	mu sync.RWMutex
	m  map[string]any
}

func New() *Cache {
	return &Cache{
		mu: sync.RWMutex{},
		m:  make(map[string]any),
	}
}

func (c *Cache) Set(key string, value any) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = value
	return nil
}

func (c *Cache) Get(key string) (any, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.m[key]
	if !ok {
		return nil, ErrNotFound
	}
	return v, nil
}
