package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- "A"
	}()
	go func() {
		time.Sleep(10 * time.Second)
		ch2 <- "B"
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("OK:", msg1)
		case msg2 := <-ch2:
			fmt.Println("OK:", msg2)
		}
	}
}
