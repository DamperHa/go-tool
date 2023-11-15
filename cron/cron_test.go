package cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"math"
	"testing"
	"time"
)

func TestCron(t *testing.T) {
	c := cron.New(cron.WithSeconds())

	c.AddFunc("1 * * * * *", func() {
		fmt.Println("Hello world!")
	})

	sh := cron.Every(10 * time.Second)
	c.Schedule(sh, cron.FuncJob(func() {
		fmt.Println("you are ok")
	}))

	go func() {
		ticker := time.NewTicker(time.Second * 4)
		for {
			select {
			case <-ticker.C:
				fmt.Println("length: ", len(c.Entries()))
			}
		}
	}()

	// c.Start()
	c.Start()

	// Wait for the Cron job to run
	time.Sleep(5 * time.Minute)

	// Stop the Cron job scheduler
	c.Stop()
}

func getBits(min, max, step uint) uint64 {
	var bits uint64

	// If step is 1, use shifts.
	if step == 1 {
		return ^(math.MaxUint64 << (max + 1)) & (math.MaxUint64 << min)
	}

	// Else, use a simple loop.
	for i := min; i <= max; i += step {
		bits |= 1 << i
	}
	return bits
}

func TestGetBits(t *testing.T) {
	res := getBits(1, 3, 1)

	fmt.Printf("%d 的二进制表示是 %b\n", res, res)
}

func TestSlice(t *testing.T) {
	a := []int{1, 2, 3}

	fmt.Println(a)

	fmt.Println(a[:len(a)-1])
}
