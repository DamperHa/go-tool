package basic

import (
	"fmt"
	"testing"
)

// 关于slice有哪些需要注意的点呢？
// 1. 如何使用slice？创建slice，append，slice访问，slice的遍历；
// 2. slice的原理？makeSLice实现，append的实现原理，slice的扩容原理？

// 传指针能够改变，不传指针不能改变，为什么呢？
// 这里主要得看一下append的实现原理吧

// 切片的初始化，主要区分三个点，是否为切片对象分配内存，长度，容量；
func TestInitializeSlice(t *testing.T) {
	var s1 []int
	s2 := make([]int, 2)
	s3 := make([]int, 2, 3)
	s4 := []int{1, 2, 3, 4, 5}

	fmt.Println(s1 == nil)
	fmt.Printf("s1:[%v], len:[%d], cap:[%d]\n", s1, len(s1), cap(s1))
	fmt.Printf("s2:[%v], len:[%d], cap:[%d]\n", s2, len(s2), cap(s2))
	fmt.Printf("s3:[%v], len:[%d], cap:[%d]\n", s3, len(s3), cap(s3))
	fmt.Printf("s4:[%v], len:[%d], cap:[%d]\n", s4, len(s4), cap(s4))
}

func TestCrop(t *testing.T) {
	a1 := [2]int{1, 2}
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1[:]
	s2[1] = 100

	s3 := a1[:]
                    
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	s3[0] = 100
	fmt.Println(s3, a1)
}

func TestCopy(t *testing.T) {
	a1 := [4]int{}
	fmt.Println(a1)
}

func changeSLice(arr *[]int) {
	*arr = append(*arr, 5)
}

func TestSLice(t *testing.T) {
	s := []int{1}
	changeSLice(&s)
	fmt.Println(s)
}

func changeSliceByIndex(s []int) {
	if len(s) < 1 {
		return
	}

	s[0] = 10000
}

// 改变某个值是没有问题的
func TestChangeSliceByIndex(t *testing.T) {
	s := []int{1}
	changeSliceByIndex(s)
	fmt.Println(s)
}
