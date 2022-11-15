/**
 * @author trung
 * @created 15/11/2022
 * @project go-play-around
 */
package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
	var (
		ch   = make(chan int)
		done = make(chan bool)
	)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		done <- true
	}()
	for {
		select {
		case v := <-ch:
			fmt.Printf("channel value: %d\n", v)
		case <-done:
			close(ch)
			fmt.Printf("channel value: %d\n", <-ch)
			fmt.Printf("Done!\n")
			return
		}
	}

}
