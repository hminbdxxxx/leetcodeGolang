package leetcode_go

import (
	"container/list"
)

type kvNode struct {
	key, val int
}

type LRUCache struct {
	ls  *list.List
	m   map[int]*list.Element
	cap int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{list.New(), make(map[int]*list.Element), capacity}
}

func (this *LRUCache) Get(key int) int {
	elem, ok := this.m[key]
	if !ok {
		return -1
	}
	this.ls.MoveToFront(elem)
	return elem.Value.(*kvNode).val
}

func (this *LRUCache) Put(key int, value int) {
	if elem, ok := this.m[key]; ok {
		kv := elem.Value.(*kvNode)
		kv.val = value
		this.ls.MoveToFront(elem)
		return
	}

	if this.ls.Len() < this.cap {
		kv := kvNode{key, value}
		elem := this.ls.PushFront(&kv)
		this.m[key] = elem
	} else {
		tail := this.ls.Back()
		this.ls.Remove(tail)
		delete(this.m, tail.Value.(*kvNode).key)
		kv := kvNode{key, value}
		elem := this.ls.PushFront(&kv)
		this.m[key] = elem
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
