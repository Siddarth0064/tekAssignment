package main

import (
	"fmt"
	"sync"
)

func sender(ch chan int, wg *sync.WaitGroup) { // this is sender func which accepts unbufferd channel refrence and   sends the values  to receiver
	go func() {
		defer wg.Done()
		ch <- 1
	}()
}
func receiver(ch chan int, wg *sync.WaitGroup) { // this is receiver func which accepts unbufferd channel refrence and collects the values from sender
	go func() {
		defer wg.Done()
		r := <-ch
		fmt.Println("frceived from sender  : ", r)
	}()
}
func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(2)
	sender(ch, wg)   //this line calling sender func with passing parameter of channel type and waitgroup
	receiver(ch, wg) //this line calling receiver func with passing parameter of channel type and waitgroup
	wg.Wait()
	close(ch) // closing the channel
	fmt.Println("main terminating normally")
}
