package linkedlist

// Do not edit the class below except for the insertKeyValuePair,
// getValueFromKey, and getMostRecentKey methods. Feel free
// to add new properties and methods to the class.
type LRUCache struct {
	index            map[string]*DoublyLinkedListNodeLRU
	maxSize          int
	currentSize      int
	listOfMostRecent *DoublyLinkedListLRU
}

func NewLRUCache(size int) *LRUCache {
	lru := &LRUCache{
		index:            map[string]*DoublyLinkedListNodeLRU{},
		maxSize:          size,
		currentSize:      0,
		listOfMostRecent: &DoublyLinkedListLRU{},
	}
	if lru.maxSize < 1 {
		lru.maxSize = 1
	}
	return lru
}

func (cache *LRUCache) InsertKeyValuePair(key string, value int) {
	if _, found := cache.index[key]; !found {
		if cache.currentSize == cache.maxSize {
			cache.evictLeastRecent()
		} else {
			cache.currentSize += 1
		}
		cache.index[key] = &DoublyLinkedListNodeLRU{
			key:   key,
			value: value,
		}
	} else {
		cache.replaceKey(key, value)
	}
	cache.updateMostRecent(cache.index[key])
}

// The second return value indicates whether or not the key was found
// in the cache.
func (cache *LRUCache) GetValueFromKey(key string) (int, bool) {
	if node, found := cache.index[key]; !found {
		return 0, false
	} else {
		cache.updateMostRecent(node)
		return node.value, true
	}
}

// The second return value is false if the cache is empty.
func (cache *LRUCache) GetMostRecentKey() (string, bool) {
	if cache.listOfMostRecent.head == nil {
		return "", false
	}
	return cache.listOfMostRecent.head.key, true
}

func (cache *LRUCache) evictLeastRecent() {
	key := cache.listOfMostRecent.tail.key
	cache.listOfMostRecent.removeTail()
	delete(cache.index, key)
}

func (cache *LRUCache) updateMostRecent(node *DoublyLinkedListNodeLRU) {
	cache.listOfMostRecent.setHeadTo(node)
}

func (cache *LRUCache) replaceKey(key string, value int) {
	if node, found := cache.index[key]; !found {
		panic("The provided key isn't in the cache!")
	} else {
		node.value = value
	}
}

type DoublyLinkedListLRU struct {
	head, tail *DoublyLinkedListNodeLRU
}

type DoublyLinkedListNodeLRU struct {
	key        string
	value      int
	prev, next *DoublyLinkedListNodeLRU
}

func (list *DoublyLinkedListLRU) setHeadTo(node *DoublyLinkedListNodeLRU) {
	if list.head == node {
		return
	}
	if list.head == nil {
		list.head, list.tail = node, node
		return
	}
	if list.head == list.tail {
		list.tail.prev = node
		list.head = node
		list.head.next = list.tail
		return
	}
	if list.tail == node {
		list.removeTail()
	}
	node.removeBindings()
	list.head.prev = node
	node.next = list.head
	list.head = node
}

func (list *DoublyLinkedListLRU) removeTail() {
	if list.tail == nil {
		return
	}
	if list.tail == list.head {
		list.head, list.tail = nil, nil
		return
	}
	list.tail = list.tail.prev
	list.tail.next = nil
}

func (node *DoublyLinkedListNodeLRU) removeBindings() {
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	node.prev, node.next = nil, nil
}
