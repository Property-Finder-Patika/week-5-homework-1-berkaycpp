package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu1   sync.Mutex
	mu2   sync.Mutex
	count int
}

func newCounter() *Counter {
	var counter = Counter{count: 0}
	counter.mu2.Lock()
	return &counter
}

func (c Counter) addCounter(i int) {
	c.mu1.Lock()
	if i != 0 {
		c.mu2.Unlock()
	}
	fmt.Println("Before count up", c.count)
	defer fmt.Println("After count up: ", c.count)

	c.count = c.count + 1

	c.mu1.Unlock()
	c.mu2.Lock()
}

func SequentialProgram() {
	var counter = Counter{}
	for i := 0; i < 100; i++ {
		counter.addCounter(i)
	}
}

func AsyncProgram() {
	var counter = Counter{}
	for i := 0; i < 100; i++ {
		go counter.addCounter(i)
	}
}

func main() {
	SequentialProgram()
	AsyncProgram()
}
