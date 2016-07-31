package main

import (
	"fmt"
	"runtime"
	"time"
)

//by default, there will only be one logical processor available to the go runtime
func main() {
	// basic()
	// basicWithLoops()
	// yielding()
	// exitBeforeFinishing()
	parallelExecutions()
}

func basic() {
	go func() {
		fmt.Println("Hello")
	}()
	go func() {
		fmt.Println("Go")
	}()

	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
}

//goroutines are scheduled in order
func basicWithLoops() {
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Hello")
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Go")
		}
	}()

	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
}

//time.Sleep() is yielding to the other goroutine
func yielding() {
	godur, _ := time.ParseDuration("10ms")
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Hello")
			time.Sleep(godur)
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Go")
			time.Sleep(godur)
		}
	}()

	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
}

//goroutines are killed in the middle of executions when the main goroutine exits
func exitBeforeFinishing() {
	godur, _ := time.ParseDuration("10ms")
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Hello")
			time.Sleep(godur)
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Go")
			time.Sleep(godur)
		}
	}()

	dur, _ := time.ParseDuration("20ms")
	time.Sleep(dur)
}

func parallelExecutions() {
	godur, _ := time.ParseDuration("10ms")
	runtime.GOMAXPROCS(2)

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Hello")
			time.Sleep(godur)
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Go")
			time.Sleep(godur)
		}
	}()

	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
}
