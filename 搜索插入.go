package leetcode

func SearchInsert(nums []int, target int) int {
	for i, num := range nums {
		if target <= num {
			return i
		}
	}

	return len(nums)
}
