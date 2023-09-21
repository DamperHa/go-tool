package basic

import (
	"fmt"
	"strings"
	"testing"
)

var sl = []string{
	"Rob Pike",
	"Robert Griesemer",
	"Ken Thompson",
}

func concatStringByOperator(sl []string) string {
	var s string
	for _, v := range sl {
		s += v
	}

	return s
}

func concatStringBySprintf(sl []string) string {
	var s string
	for _, v := range sl {
		s = fmt.Sprintf("%s%s", s, v)
	}

	return s
}

func concatStringByJoin(sl []string) string {
	return strings.Join(sl, "")
}

func BenchmarkConcatStringByOperator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatStringByOperator(sl)
	}
}

func BenchmarkConcatStringBySprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatStringBySprintf(sl)
	}
}

func BenchmarkStringByJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatStringByJoin(sl)
	}
}

//➜  basic git:(master) ✗ go test -bench=ByJoin ./benchmark_test.go
//goos: darwin
//goarch: amd64
//cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
//BenchmarkStringByJoin-12        24681794                46.58 ns/op
//PASS
//ok      command-line-arguments  1.308s
// 每次for循环的平均执行时间
//
//var (
//	m     map[int64]struct{} = make(map[int64]struct{}, 10)
//	mu    sync.Mutex
//	round int64 = 1
//)
//
//func BenchmarkSequential(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		//		fmt.Printf("\ngoroutinue [%d] enter BenchMarkSequentiaal: round:[%d], b.N[%d]\n", tls.ID(), atomic.LoadInt64(&round), b.N)
//	}
//}
