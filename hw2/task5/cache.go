package main

import "sync"

type Cache[T comparable, M any] struct {
	mu sync.RWMutex
	m  map[T]M
}

func NewCache[T comparable, M any](size int) *Cache[T, M] {
	return &Cache[T, M]{
		mu: sync.RWMutex{},
		m:  make(map[T]M, size),
	}
}

func (c *Cache[T, M]) Get(key T) M {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.m[key]
}

func (c *Cache[T, M]) Set(key T, val M) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = val
}
