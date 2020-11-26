package main

import (
	"fmt"
	"time"
)

func main() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err=", err)
		}
	}()
	go panicFunc()

	time.Sleep(time.Second * 10)
	fmt.Println("end")
}

func panicFunc() {
	panic("panicFunc panic")

}
