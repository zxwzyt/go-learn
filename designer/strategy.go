package designer

import (
	"fmt"
)

type EvictionAlgo interface {
	evict(c *Cache)
}

type Fifo struct{}

func (l *Fifo) evict(c *Cache) {
	fmt.Println("Eviction by fifo strtegy")
}

type Lru struct{}

func (l *Lru) evict(c *Cache) {
	fmt.Println("Evicting by lru strtegy")
}

type Lfu struct{}

func (l *Lfu) evict(c *Cache) {
	fmt.Println("Evictiong by lfu strtegy")
}

type Cache struct {
	storage      map[string]string
	evictionAlog EvictionAlgo
	capacity     int
	maxCapaticy  int
}

func initCache(e EvictionAlgo) *Cache {
	storage := make(map[string]string)
	return &Cache{
		storage:      storage,
		evictionAlog: e,
		capacity:     0,
		maxCapaticy:  2,
	}
}

func (c *Cache) setEvictionAlog(e EvictionAlgo) {
	c.evictionAlog = e
}

func (c *Cache) add(key, value string) {
	if c.capacity == c.maxCapaticy {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) get(key string) {
	delete(c.storage, key)
}

func (c *Cache) evict() {
	c.evictionAlog.evict(c)
	c.capacity--
}

func RunStrategy() {
	lfu := &Lfu{}
	cache := initCache(lfu)

	cache.add("a", "1")
	cache.add("b", "2")
	cache.add("c", "3")

	lru := &Lru{}
	cache.setEvictionAlog(lru)

	cache.add("d", "4")

	fifo := &Fifo{}
	cache.setEvictionAlog(fifo)

	cache.add("e", "5")
}
