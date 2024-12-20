package storage

type IStorageService interface {
	Put(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Remove(key string) error
}
