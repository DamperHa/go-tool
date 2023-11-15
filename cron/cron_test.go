package cron

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/robfig/cron"
)

func TestCron(t *testing.T) {
	c := cron.New()
	//
	//c.AddFunc("* * * * *", func() {
	//	fmt.Println(time.Now())
	//	time.Sleep(time.Second * 12)
	//	fmt.Println("Hello world!")
	//})

	// Start the Cron job scheduler
	c.Start()
	sh := cron.Every(10 * time.Second)
	c.Schedule(sh, cron.FuncJob(func() {
		fmt.Println(time.Now())
		time.Sleep(time.Second * 12)
		fmt.Println("you are ok")
	}))

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
