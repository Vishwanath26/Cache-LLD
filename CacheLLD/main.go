package main

import (
	"CacheLLD/service/eviction"
	"CacheLLD/service/storage"
	"fmt"
)

var mapStorage = storage.NewMapStorage(4)
var lruEviction = eviction.NewLRUEviction(mapStorage)

func Put(key string, val interface{}) {
	err := mapStorage.Put(key, val)
	if err != nil && err.Error() == "max capacity reached" {
		err := lruEviction.Evict()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	err = lruEviction.UpdateEviction(key)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func Get(key string) {
	val, err := mapStorage.Get(key)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = lruEviction.UpdateEviction(key)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print("key " + key + " has value ")
	fmt.Println(val)
}

func main() {
	Put("first", 1)
	Put("second", 2)
	Get("first")

	Put("third", 3)
	Put("fourth", 4)
	Put("fifth", 5)
	Get("second")
}
