package main

import "container/list"

type entry[K comparable, V any] struct {
	key   K
	value V
}
type LRUCache[K comparable, V any] struct {
	capacity int
	ll       list.List
	items    map[K]*list.Element
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		capacity: capacity,
		items:    make(map[K]*list.Element),
	}
}
func (c *LRUCache[K, V]) Get(key K) (value V, ok bool) {
	el, ok := c.items[key]

	if !ok {
		var zero V
		return zero, false
	}
	c.ll.MoveToFront(el)
	return el.Value.(*entry[K, V]).value, true
}
func (c *LRUCache[K, V]) Set(key K, value V) {
	if c.capacity = 0 {
		return
	}

	if el, ok := c.items[key]; ok {
		el.Value.(*entry[K, V]).value = value
	} else {
		if c.ll.Len() == c.capacity {
			back := c.ll.Back()
			e := back.Value.(*entry[K, V])

			c.ll.Remove(back)
			delete(c.items, e.key)
		}
		el := c.ll.PushFront(&entry[K, V]{
			key:   key,
			value: value,
		})
		c.items[key] = el
	}
}

type LRU[K comparable, V any] interface {
	Get(K) (V, bool)
	Set(K, V)
}
