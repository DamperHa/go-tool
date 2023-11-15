package timer

import (
	"fmt"
	"testing"
	"time"
)

func TestShangHaiAndNewYork(t *testing.T) {
	shanghaiTimer, _ := time.LoadLocation("Asia/Shanghai")
	newYorkLoc, _ := time.LoadLocation("America/New_York")

	fmt.Printf("shanghai: [%v]\n", time.Now().In(shanghaiTimer))
	fmt.Printf("newYork:[%v]\n", time.Now().In(newYorkLoc))
}
