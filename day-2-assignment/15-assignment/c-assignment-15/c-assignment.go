package main

import (
	"fmt"
	"time"
)

func sender(ch chan int) { // this is sender func which accepts unbufferd channel refrence and   sends the values  to receiver
	go func() { // this is a go rouiten of sender
		i := 1
		for i < 4 {
			ch <- i
			i++
		}
	}()
}
func receiver(ch chan int) { // this is receiver func which accepts unbufferd channel refrence and collects the values from sender
	go func() { //this is a go rouiten of recevier
		for r := range ch {
			fmt.Println("frceived from sender  : ", r)
		}
		fmt.Println()
	}()
}
func main() {

	ch := make(chan int) // this is unbufferd channel in int data type
	sendersend := 3      // this lines defindes 3 senders are present
	receiverreciv := 2   //this lines defindes 2 receiver are presents
	for i := 1; i <= sendersend; i++ {
		sender(ch) //calling sender func with passing channel type
	}
	for j := 1; j <= receiverreciv; j++ {
		receiver(ch) //calling receiver func with passing channel type
	}
	time.Sleep(4 * time.Second)
	fmt.Println("main terminating normally")
}
