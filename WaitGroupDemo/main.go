package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func hello(i int, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		fmt.Println("i=", i, " end")
	}()
	wg.Add(1)
	second := rand.Intn(5)
	fmt.Println("i=", i, " run", "sleep=", second)
	time.Sleep(time.Duration(second) * time.Second)

}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		go hello(i, &wg)
	}

	wg.Wait()
	fmt.Println("run end")
}
