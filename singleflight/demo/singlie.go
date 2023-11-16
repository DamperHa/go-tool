package main

import (
	"context"
	"net/http"
	"time"

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
	sfHandler.group.Do("", func() (interface{}, error) {
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

func main() {
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

	time.Sleep(10 * time.Second)

	print("ok")
}
