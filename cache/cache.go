package cache

import (
	"container/list"
)

type Cache struct {
	Count int
	queue *list.List
	data  map[string]string
}

func NewCache(count int) *Cache {
	return &Cache{
		Count: count,
		queue: list.New(),
		data:  make(map[string]string),
	}
}

func (cache Cache) Add(key, value string) bool {
	if _, ok := cache.data[key]; ok {
		return false
	}

	if len(cache.data) >= cache.Count {
		cache.deleteLast()
	}

	cache.insert(key)
	cache.data[key] = value

	return true
}

func (cache Cache) Get(key string) (value string, ok bool) {
	value, ok = cache.data[key]
	if !ok {
		return value, ok
	}

	el := cache.findInQueue(key)
	cache.queue.MoveToFront(el)

	return value, ok
}

func (cache Cache) Remove(key string) (ok bool) {
	if _, ok = cache.data[key]; !ok {
		return ok
	}

	el := cache.findInQueue(key)
	cache.queue.Remove(el)

	delete(cache.data, key)
	return true
}

// deleteLast deletes oldest element
func (cache Cache) deleteLast() {
	lastEl := cache.queue.Back()
	delete(cache.data, lastEl.Value.(string))
	cache.queue.Remove(lastEl)
}

// insert addes element to queue head
func (cache Cache) insert(key string) {
	cache.queue.PushFront(key)
}

func (cache Cache) findInQueue(key string) *list.Element {
	el := cache.queue.Front()

	for {
		if el.Value.(string) == key {
			return el
		}
		el = el.Next()
	}
}
