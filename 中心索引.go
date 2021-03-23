package leetcode

// PivotIndex 返回中心索引
func PivotIndex(nums []int) int {
	if len(nums) <= 1 {
		return -1
	}

	leftSum := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		leftSum[i] = leftSum[i-1] + nums[i-1]
	}

	rightSum := make([]int, len(nums))
	for i := len(nums) - 2; i >= 0; i-- {
		rightSum[i] = rightSum[i+1] + nums[i+1]
	}

	for i := 0; i < len(leftSum); i++ {
		if leftSum[i] == rightSum[i] {
			return i
		}
	}

	return -1
}
