package main

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

type contextKey string

func (c contextKey) String() string {
	return string(c)
}

func TestCOntextWithvalue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, contextKey("a"), "B")
	contextC := context.WithValue(contextA, contextKey("c"), "C")

	contextD := context.WithValue(contextB, contextKey("d"), "D")
	contextE := context.WithValue(contextB, contextKey("e"), "E")

	contextF := context.WithValue(contextC, contextKey("f"), "F")
	contextG := context.WithValue(contextF, contextKey("g"), "G")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	fmt.Println(contextF.Value(contextKey("f")))
	fmt.Println(contextF.Value(contextKey("c")))

}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return destination
}

func TestContextWithCancel(t *testing.T) {
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	destination := CreateCounter(ctx)

	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}
	cancel()

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}
func TestContextWithTimeOut(t *testing.T) {
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	destination := CreateCounter(ctx)

	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}

func TestContextWithDeadLine(t *testing.T) {
	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	destination := CreateCounter(ctx)

	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}
