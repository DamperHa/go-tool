package _defer

import "testing"

func sum(max int) int {
	total := 0
	for i := 0; i < max; i++ {
		total += i
	}

	return total
}

func fooWithDefer() {
	defer func() {
		sum(10)
	}()
}

func fooWithOutDefer() {
	sum(10)
}

func BenchmarkFooWithDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fooWithDefer()
	}
}

func BenchmarkFooWithOutDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fooWithOutDefer()
	}
}
