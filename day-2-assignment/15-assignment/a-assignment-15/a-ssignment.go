package main

import (
	"fmt"
	"time"
)

func sender(ch chan int) { // this is sender func which accepts unbufferd channel refrence and   sends the values  to receiver
	i := 1
	for i <= 5 {
		ch <- i
		i++

	}
}
func receiver(ch chan int) { // this is receiver func which accepts unbufferd channel refrence and collects the values from sender
	for r := range ch {
		fmt.Printf("receiver : %d.\n", r)
	}
}
func main() {
	ch := make(chan int) // this is unbufferd channel in int data type
	go sender(ch)        // this is a go rouiten of sender
	go receiver(ch)      //this is a go rouiten of recevier
	time.Sleep(1 * time.Second)
	fmt.Println("main programm ending normally")
}
