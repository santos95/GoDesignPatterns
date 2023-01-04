package ex1

type EvictionAlgo interface {
	evict(c *cache)
}
