package singleflight

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"testing"
	"time"

	"context"

	"golang.org/x/sync/singleflight"
)

type Handler struct {
	next http.Handler
}

func (handler *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler.next != nil {
		handler.next.ServeHTTP(w, r)

		return
	}

	_, _ = w.Write([]byte("body"))
}

type SingleFlightHandler struct {
	group *singleflight.Group
	next  http.Handler
}

func (sfHandler *SingleFlightHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sfHandler.group.DoChan("", func() (interface{}, error) {
		defer sfHandler.group.Forget("")

		sfHandler.next.ServeHTTP(w, r)

		return struct{}{}, nil
	})
}

type PanicHandler struct{}

func (pHandler *PanicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	panicErrAbortHandler()
}

func panicErrAbortHandler() {
	panic(http.ErrAbortHandler)
}

func TestHttpWithSingleFlight(t *testing.T) {
	sf := &singleflight.Group{}

	mux := &http.ServeMux{}
	mux.Handle("/", &Handler{
		next: &SingleFlightHandler{group: sf,
			next: &PanicHandler{},
		},
	})

	go func() {
		_ = http.ListenAndServe("localhost:8080", mux)
	}()

	// Wait for HTTP server to be up
	time.Sleep(time.Millisecond)

	client := &http.Client{}

	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Second)
	defer cancelCtx()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/", http.NoBody)

	_, _ = client.Do(req)
}

// 测试一下，panic后，fatal之后，defer会不会执行呢
func TestPanic2(t *testing.T) {
	var wg sync.WaitGroup
	type People struct {
		name string
	}

	var p *People

	wg.Add(1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("recover:", err)
			}
		}()

		// 比蹦
		// go panic("ok")
		// map呢
		//var mp map[int]int
		//mp[1] = 1

		// 越界是否能cover住
		sk := []int{1}
		print(sk[3])

		// 一定是panic
		print(p.name)
		// Golang panic之后，代码就不会执行了，即使recover了也不行
		panic("are you ok")

		wg.Done()
	}()

	go func() {
		wg.Wait()
		fmt.Println("wait for ok")
	}()

	time.Sleep(1 * time.Minute)
	fmt.Println("end main")
}

// 如果里面是go panic，那么就捕捉不到
func TestWithOutSingleFlight(t *testing.T) {
	mux := &http.ServeMux{}
	mux.Handle("/", &Handler{
		next: &PanicHandler{},
	})

	go func() {
		_ = http.ListenAndServe("localhost:8080", mux)
	}()

	// Wait for HTTP server to be up
	time.Sleep(time.Millisecond)

	client := &http.Client{}

	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Second)
	defer cancelCtx()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/", http.NoBody)

	_, _ = client.Do(req)
}

// 在一个协程中recover住panic，为什么还需要再调用一次panic
func TestPanicRecover(t *testing.T) {

	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("recover 2:", err)
			}
		}()

		go func() {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println("recover 1:", err)
					panic(err)
				}
			}()

			panic("you are ok")
		}()
	}()

	time.Sleep(10 * time.Second)
}

func First() {
	Second()
}

func Second() {
	Third()
}

func Third() {
	for c := 0; c < 5; c++ {
		fmt.Println(runtime.Caller(c))
	}
}

func stackExample() {
	stackSlice := make([]byte, 512)
	s := runtime.Stack(stackSlice, false)
	fmt.Printf("\n%s", stackSlice[0:s])
}

func TestCallerStacker(t *testing.T) {

	fmt.Println("######### STACK ################")
	stackExample()
	fmt.Println("\n\n######### CALLER ################")
	First()
}
