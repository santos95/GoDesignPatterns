package ex1

type cache struct {
	storage      map[string]string
	evictionAlgo EvictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache(e EvictionAlgo) *cache {
	storage := make(map[string]string)
	return &cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}

}

func (c *cache) setEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}

func (c *cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}
