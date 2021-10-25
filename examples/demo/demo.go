package main

import (
	"fmt"
	"time"

	"github.com/fufuok/timewheel"
)

var (
	TW *timewheel.TimeWheel
)

func init() {
	TW, _ = timewheel.NewTimeWheel(100*time.Millisecond, 600)
	TW.Start()
}

func TWStop() {
	TW.Stop()
}

func main() {
	defer TWStop()

	fmt.Println(time.Now())

	TW.Sleep(2 * time.Second)

	fmt.Println("sleep 2s:", time.Now())

	TW.AfterFunc(1*time.Second, func() {
		fmt.Println("after 1s:", time.Now())
	})

	TW.AddCron(1*time.Second, func() {
		fmt.Println("cron(1s):", time.Now())
	})

	ticker := TW.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for i := 0; i < 5; i++ {
		<-ticker.C
		fmt.Println("ticker(200ms):", i, time.Now())
	}

	fmt.Println()
	ticker.Reset(1 * time.Second)

	for i := 0; i < 5; i++ {
		<-ticker.C
		fmt.Println("ticker(1s):", i, time.Now())
	}

	// more timewheel
	TW50ms, _ := timewheel.NewTimeWheel(50*time.Millisecond, 10)
	TW50ms.Start()
	defer TW50ms.Stop()

	now := time.Now()
	TW50ms.AfterFunc(100*time.Millisecond, func() {
		fmt.Println("cost(tick:50ms+after:100ms):", time.Since(now))
	})

	TW1s, _ := timewheel.NewTimeWheel(1*time.Second, 10)
	TW1s.Start()
	defer TW1s.Stop()

	now = time.Now()
	TW1s.AfterFunc(1*time.Second, func() {
		fmt.Println("cost(tick:1s+after:1s):", time.Since(now))
	})

	TW1s.Sleep(2 * time.Second)
}
