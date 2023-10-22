package basic

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"testing"
)

func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine"))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}

	return id
}

func TestGOID(t *testing.T) {
	fmt.Println("main.go", GoID())
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		j := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(j, GoID())
		}()
	}

	wg.Wait()
}
