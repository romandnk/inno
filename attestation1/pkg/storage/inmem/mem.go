package inmem

import (
	"context"
	"sync"
)

type Cache[K string, V any] struct {
	mu sync.RWMutex
	m  map[K][]V
}

func NewCache[K string, V any](capacity uint64) *Cache[K, V] {
	return &Cache[K, V]{
		m:  make(map[K][]V, capacity),
		mu: sync.RWMutex{},
	}
}

func (c *Cache[K, V]) Set(ctx context.Context, key K, value V) {
	select {
	case <-ctx.Done():
		return
	default:
		c.mu.Lock()
		defer c.mu.Unlock()
		c.m[key] = append(c.m[key], value)
	}
}

func (c *Cache[K, V]) All(ctx context.Context) []V {
	select {
	case <-ctx.Done():
		return nil
	default:
		c.mu.RLock()
		defer c.mu.RUnlock()
		data := make([]V, 0, len(c.m))
		for _, v := range c.m {
			data = append(data, v...)
			select {
			case <-ctx.Done():
				return data
			default:
			}
		}
		return data
	}
}

func (c *Cache[K, V]) Clear(ctx context.Context) {
	select {
	case <-ctx.Done():
		return
	default:
		c.mu.Lock()
		defer c.mu.Unlock()
		c.m = make(map[K][]V)
	}
}
