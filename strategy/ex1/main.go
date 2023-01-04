package main

import main2 "desingPatterns/strategy/ex1"

func main() {
	lfu := &main2.lfu{}
	cache := main2.initCache(lfu)
	cache.add("a", "1")
	cache.add("b", "2")
	cache.add("c", "3")
	lru := &main2.lru{}
	cache.setEvictionAlgo(lru)
	cache.add("d", "4")
	fifo := &main2.fifo{}
	cache.setEvictionAlgo(fifo)
	cache.add("e", "5")
}
