package storage

import (
	"errors"
	"fmt"
)

type MapStorage struct {
	capacity int
	store    map[string]interface{}
}

func NewMapStorage(size int) IStorageService {
	return &MapStorage{
		store:    make(map[string]interface{}),
		capacity: size,
	}
}

func (ms *MapStorage) Put(key string, value interface{}) error {
	if len(ms.store) == ms.capacity {
		fmt.Println("max capacity reached, will try to evict least recently used")
		return errors.New("max capacity reached")
	}
	ms.store[key] = value
	return nil
}

func (ms *MapStorage) Get(key string) (interface{}, error) {
	if ms.store[key] == nil {
		return nil, errors.New("key " + key + " not found")
	}
	return ms.store[key], nil
}

func (ms *MapStorage) Remove(key string) error {
	if ms.store[key] == nil {
		return errors.New("key " + key + " not found")
	}
	delete(ms.store, key)
	return nil
}
