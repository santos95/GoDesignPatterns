package ex1

import "fmt"

type fifo struct {
}

func (f *fifo) evict(c *cache) {
	fmt.Println("Evicting by fifo strategy")
}
