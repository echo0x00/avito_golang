package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mutex    sync.Mutex
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
		mutex:    sync.Mutex{},
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if elem, ok := c.items[key]; ok {
		c.queue.MoveToFront(elem)
		elem.Value.(*cacheItem).value = value
		return true
	}

	item := &cacheItem{key: key, value: value}
	elem := c.queue.PushFront(item)
	c.items[key] = elem
	if c.capacity < c.queue.Len() {
		backQueueElem := c.queue.Back()
		c.queue.Remove(backQueueElem)
		delete(c.items, backQueueElem.Value.(*cacheItem).key)
	}

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	elem, ok := c.items[key]
	if !ok {
		return nil, ok
	}

	c.queue.MoveToFront(elem)
	return elem.Value.(*cacheItem).value, ok
}

func (c *lruCache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}
