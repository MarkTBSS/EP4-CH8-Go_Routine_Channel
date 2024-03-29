package main

import (
	"fmt"
	"time"
)

func slow(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(s, ":", i)
	}
}
func main() {
	done := make(chan bool)
	go func() {
		slow("A")
		done <- true
	}()
	slow("B")
	<-done
	fmt.Println("all task done.")
}
