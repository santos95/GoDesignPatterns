package ex1

import "fmt"

type lru struct {
}

func (l *lru) evict(c *cache) {
	fmt.Println("Evicting by iru strategy")
}
