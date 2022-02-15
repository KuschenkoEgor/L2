package pkg

type EvictionAlgo interface {
	Evict(c *Cache)
}
