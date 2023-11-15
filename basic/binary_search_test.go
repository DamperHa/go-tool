package basic

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestLeftBinary(t *testing.T) {
	nums := []int{1, 3, 4, 5, 6, 6, 6, 6, 7}
	testTarget := []int{1, 6, 7}
	expect := []int{0, 4, 8}

	for i, target := range testTarget {
		res := FindLeftBo(nums, target)

		assert.Equal(t, res, expect[i])
	}
}

func FindLeftBo(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			if mid == 0 || nums[mid-1] < target {
				return mid
			}

			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		}
	}

	return -1
}

func TestRightBinary(t *testing.T) {
	nums := []int{1, 3, 4, 5, 6, 6, 6, 6, 7}
	testTarget := []int{1, 6, 7}
	expect := []int{0, 7, 8}

	for i, target := range testTarget {
		res := FindRightBo(nums, target)

		assert.Equal(t, res, expect[i])
	}
}

func FindRightBo(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			if mid == len(nums)-1 || target < nums[mid+1] {
				return mid
			}

			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		}
	}

	return -1
}
