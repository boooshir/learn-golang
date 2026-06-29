package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)

	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)
}

func TestAfter(t *testing.T) {
	channel := time.After(5 * time.Second)

	fmt.Println(time.Now())

	time := <-channel
	fmt.Println(time)
}

func TestAfterFunction(t *testing.T) {
	group := sync.WaitGroup{}

	group.Add(1)

	time.AfterFunc(5*time.Second, func() {
		fmt.Println("Execute after 1 second")
		fmt.Println(time.Now())
		group.Done()
	})
	fmt.Println(time.Now())
	group.Wait()
}

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for time := range ticker.C {
		fmt.Println(time)
	}

}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for time := range channel {
		fmt.Println(time)
	}

}
