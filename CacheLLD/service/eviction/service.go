package eviction

type IEvictionService interface {
	UpdateEviction(key string) error
	Evict() error
}
