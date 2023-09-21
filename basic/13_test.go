package basic

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

// 当我们把slice看做是数组的窗口，一切都水到渠成了
// 比如，slice的底层数组指针，通过arr[low:high]创建的切片

func TestSliceNil(t *testing.T) {
	var slice []int
	fmt.Println(len(slice))
	fmt.Printf("slice pointer：%p", &slice)
}

func TestDynamicScale(t *testing.T) {
	s1 := make([]int, 1, 3)
	s1[0] = 0

	s2 := s1
	fmt.Printf("s1:%v, cap：%v s2:%v， cap：%v \n", s1, cap(s1), s2, cap(s2))

	s2[0] = 100
	fmt.Printf("s1:%v, cap：%v s2:%v， cap：%v \n", s1, cap(s1), s2, cap(s2))

	s3 := append(s1, 2)
	fmt.Printf("s1:%v, cap：%v s2:%v， cap：%v s3:%v， cap：%v\n", s1, cap(s1), s2, cap(s2), s3, cap(s3))

}

func printSliceArray(slice []int) {
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	// 这里由于数组里面只接受常量，所以这里默认4
	data := *(*[4]int)(unsafe.Pointer(hdr.Data))
	fmt.Println(cap(slice), len(slice), data)
}

// append的时候，如果容量足够。就在原有的底层数组上append，如果容量不足够，那么就会重新
// 分配底层数组
// https://stackoverflow.com/questions/36706843/how-to-get-the-underlying-array-of-a-slice-in-go
func TestPrintArr(t *testing.T) {
	s1 := make([]int, 1, 3)
	s1[0] = 100

	s2 := s1
	printSliceArray(s1)
	printSliceArray(s2)

	s2 = append(s2, 1)
	printSliceArray(s1)
	printSliceArray(s2)

	s2 = append(s2, 1)
	printSliceArray(s1)
	printSliceArray(s2)

	fmt.Println(s1)

	// 当超过底层容量是，会重新给s2分配数组
	s2 = append(s2, 1)
	printSliceArray(s1)
	printSliceArray(s2)
}

const sliceSize = 1000

// 对于benchmark的操作
// go test -bench=.
// go test -bench=BenchmarkWithCap
func BenchmarkWithoutCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sl := make([]int, 0)
		for i := 0; i < sliceSize; i++ {
			sl = append(sl, i)
		}
	}
}

func BenchmarkWithCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sl := make([]int, sliceSize)
		for i := 0; i < b.N; i++ {
			sl = append(sl, i)
		}
	}
}
