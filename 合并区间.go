package leetcode

func sortIntervals(intervals [][]int) [][]int {
	if len(intervals) <= 0 {
		return [][]int{}
	}

	result := make([][]int, len(intervals))
	for i, interval := range intervals {
		index := i - 1
		for index >= 0 && result[index][0] > interval[0] {
			result[index+1] = []int{result[index][0], result[index][1]}
			index = index - 1
		}

		result[index+1] = []int{interval[0], interval[1]}
	}

	return result
}

func Merge(intervals [][]int) [][]int {
	// 先对intervals的数组进行排序，按照从小到大的顺序先进行排序
	intervals = sortIntervals(intervals)

	result := make([][]int, 0, len(intervals))
	for _, interval := range intervals {
		isAppend := true
		for i, ri := range result {
			smallArr := interval
			bigArr := ri
			if smallArr[0] > bigArr[0] {
				smallArr = ri
				bigArr = interval
			}

			if smallArr[1] >= bigArr[0] {
				end := bigArr[1]
				if smallArr[1] > bigArr[1] {
					end = smallArr[1]
				}
				result[i] = []int{smallArr[0], end}
				isAppend = false
				break
			}
		}

		if isAppend {
			result = append(result, []int{interval[0], interval[1]})
		}
	}

	return result
}
