package grok

import (
	"fmt"
	"time"
)

func FromGrok() {
	// unbuffered channel example
	unbufferedCh := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		unbufferedCh <- "ping"
	}()
	// main goroutine waits to receive
	fmt.Println(<-unbufferedCh)

	// buffer channel example
	// bufferCh := make(chan string, 2)
	//
	// // send two message to bufferChannel
	// bufferCh <- "one"
	// bufferCh <- "two"
	//
	// go func() {
	// 	for msg := range bufferCh {
	// 		fmt.Println(msg)
	// 	}
	// }()
	//
	// bufferCh <- "three"
	// close(bufferCh)
	//
	// time.Sleep(time.Second * 1)

}
