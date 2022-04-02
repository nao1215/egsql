package cache

import (
	"container/list"
	"sync"
)

// LRU is the thread-safe Least Recently Used cache.
type LRU struct {
	// capacity is cache capacity
	capacity int
	// evictList is eviction cache list. It is doubly linked list
	evictList *list.List
	// items is cache itself
	items map[interface{}]*list.Element
	// mutex is used by LRU operation
	mutex sync.RWMutex
}

// entry is a cache that is left in key-value format
type entry struct {
	key   interface{}
	value interface{}
}

// NewLRU return LRU pointer
func NewLRU(cap int) *LRU {
	return &LRU{
		capacity:  cap,
		evictList: list.New(),
		items:     make(map[interface{}]*list.Element),
	}
}

// Insert adds data to the LRU. If capacity is exceeded, return oldest data. otherwise, return nil.
func (l *LRU) Insert(key, value interface{}) interface{} {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	var victim interface{}
	entry := &entry{key, value}
	elment := l.evictList.PushFront(entry)
	l.items[key] = elment

	if l.needEvict() {
		victim = l.evictList.Back()
		l.removeOldest()
	}
	return victim
}

// Get returns the value corresponding to the key.
// If there is no corresponding value, nil is returned.
func (l *LRU) Get(key interface{}) interface{} {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	if element, ok := l.items[key]; ok {
		l.evictList.MoveToFront(element)
		return element.Value.(*entry).value
	}
	return nil
}

// Len return length of eviction list.
func (l *LRU) Len() int {
	return l.evictList.Len()
}

// needEvict returns whether eviction is necessary
func (l *LRU) needEvict() bool {
	return l.Len() > l.capacity
}

// removeOldest removes the oldest element in the eviction list.
func (l *LRU) removeOldest() {
	elm := l.evictList.Back()
	if elm != nil {
		l.evictList.Remove(elm)
		delete(l.items, elm.Value.(*entry).key)
	}
}
