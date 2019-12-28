package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan bool)

	go func() {
		for {
			done <- true
			time.Sleep((time.Second * 1) / 100)
		}
	}()

	var c int
	for {
		select {
		case ok := <-done:
			if ok {
				c++
			}
		default:
			loading(c)
		}
	}
}

func loading(s int) {
	chars := "|/-\\"
	for _, c := range chars {
		fmt.Printf("\r%s %ds", string(c), s)
		time.Sleep((time.Second * 1) / 100)
	}
}
