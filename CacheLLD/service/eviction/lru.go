package eviction

import (
	"CacheLLD/service/algorithms"
	"CacheLLD/service/storage"
	"errors"
)

type LRUEviction struct {
	dll     *algorithms.DoublyLinkedList
	mapping map[string]*algorithms.Node
	storage storage.IStorageService
}

func NewLRUEviction(storage storage.IStorageService) *LRUEviction {
	return &LRUEviction{
		dll:     algorithms.NewDoublyLinkedList(),
		mapping: make(map[string]*algorithms.Node),
		storage: storage,
	}
}

func (lru *LRUEviction) UpdateEviction(key string) error {
	if lru.mapping[key] == nil {
		node := lru.dll.Add(key)
		lru.mapping[key] = node
	} else {
		curNode := lru.mapping[key]
		if err := lru.dll.Remove(curNode); err != nil {
			return errors.New("error in updating eviction")
		}
		lru.dll.Add(curNode.Val())
	}
	return nil
}

func (lru *LRUEviction) Evict() error {
	nodeToBeRemoved := lru.dll.End().Prev()

	err := lru.dll.Remove(nodeToBeRemoved)
	if err != nil {
		return errors.New("error in eviction ")
	}
	err = lru.storage.Remove(nodeToBeRemoved.Val().(string))
	if err != nil {
		return err
	}
	return nil
}
