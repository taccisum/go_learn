package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func randNum() chan int {
	c := make(chan int)
	go func() {
		n := rand.Intn(233)
		time.Sleep(1 * time.Second)
		c <- n
	}()
	return c
}

func add1(num int) chan int {
	c := make(chan int)
	go func() {
		// time.Sleep(1 * time.Second)
		c <- num + 1
	}()
	return c
}

func main() {
	fmt.Println("CPU Cores", runtime.NumCPU())

	for i := 0; i < 3; i++ {
		id := i

		go func() {
			fmt.Printf("[%d] goroutine start\n", id)
			num := <-randNum()
			fmt.Printf("[%d] rand num: %d\n", id, num)
			num = <-add1(num)
			fmt.Printf("[%d] goroutine ok. result: %d\n", id, num)
		}()
	}

	time.Sleep(3 * time.Second)
}
