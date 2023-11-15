package codetop

type LRUCache struct {
	linkedHashMap *LinkedHashMap
	capacity      int
}

type LinkedHashMap struct {
	head    *Node
	tail    *Node
	mapping map[int]*Node
}

type Node struct {
	key   int
	value int
	next  *Node
	prev  *Node
}

func NewLinkedHashMap() *LinkedHashMap {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return &LinkedHashMap{
		head:    head,
		tail:    tail,
		mapping: make(map[int]*Node),
	}
}

func (m LinkedHashMap) Size() int {
	return len(m.mapping)
}

func (m LinkedHashMap) Get(key int) int {
	if node, ok := m.mapping[key]; ok {
		return node.value
	} else {
		return -1
	}
}

func (m LinkedHashMap) Contains(key int) bool {
	_, ok := m.mapping[key]
	return ok
}

func (m LinkedHashMap) Put(key, value int) {
	m.Delete(key)
	node := &Node{
		key:   key,
		value: value,
		next:  m.head.next,
		prev:  m.head,
	}
	m.head.next.prev = node
	m.head.next = node
	m.mapping[key] = node
}

func (m LinkedHashMap) Delete(key int) {
	if node, ok := m.mapping[key]; ok {
		prev := node.prev
		next := node.next
		prev.next = next
		next.prev = prev
		delete(m.mapping, key)
	}
}

func (m LinkedHashMap) DeleteBack() {
	if m.Size() > 0 {
		m.Delete(m.tail.prev.key)
	}
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		linkedHashMap: NewLinkedHashMap(),
		capacity:      capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if this.linkedHashMap.Contains(key) {
		value := this.linkedHashMap.Get(key)
		this.Put(key, value)
		return value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	this.linkedHashMap.Put(key, value)
	if this.linkedHashMap.Size() > this.capacity {
		this.linkedHashMap.DeleteBack()
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
