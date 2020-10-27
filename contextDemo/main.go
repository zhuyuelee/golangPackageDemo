package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())

	go work(ctx, ch)
	time.Sleep(5 * time.Second)
	cancel()
	<-ch

	fmt.Println("回家")

}

func work(ctx context.Context, ch chan<- bool) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("下班")
			ch <- true
			return
		default:
			fmt.Println("上班中....")
			time.Sleep(2 * time.Second)
		}
	}
}
