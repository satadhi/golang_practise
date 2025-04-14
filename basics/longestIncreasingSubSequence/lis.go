package main

import (
	"fmt"
	"math"
)

func lengthOfLIS(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}
	
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[j]+1)))
			}
		}
	}

	maxLength := 0
	for _, length := range dp {
		maxLength = int(math.Max(float64(maxLength), float64(length)))
	}

	return maxLength
}

func main() {
	nums1 := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums1)) // Output: 4

	nums2 := []int{0, 1, 0, 3, 2, 3}
	fmt.Println(lengthOfLIS(nums2)) // Output: 4

	nums3 := []int{7, 7, 7, 7, 7, 7, 7}
	fmt.Println(lengthOfLIS(nums3)) // Output: 1

	nums4 := []int{}
	fmt.Println(lengthOfLIS(nums4)) // Output: 0
}
