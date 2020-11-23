package main

import (
	"log"
	"time"
)

func main() {
	log.Println("started")
	c := make(chan int)
	go func(ch chan int) {
		time.Sleep(time.Second * 5)
		<-ch
		log.Println("read")
	}(c)

	log.Println("before writing")
	c <- 1 // it will wait reading
	log.Println("after writing")
	close(c)
}
