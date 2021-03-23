package leetcode

// 另一种方案：设置length + 1长度的数组标识每棵树的状态(默认为1)，下标对应数的位置
// 每当修建的铁路经过这个数组的某个下标时，就把值标记为0
// 最后统计剩下为1的元素的个数

// BuildTrain 路边的树
// 这里假设regions的起始位置<=终止位置
func BuildTrain(length int, regions [][2]int) int {
	freeRegions := make([][2]int, 0)
	freeRegions = append(freeRegions, [2]int{0, length})

	for _, region := range regions {
		start := region[0]
		end := region[1]

		// 找出这次需要移动的数在那几个空闲数组之间
		startIndex := 0
		endCount := 0
		for _, fr := range freeRegions {
			if fr[1] < start {
				startIndex++
				continue
			}
			if fr[1] < end {
				endCount++
			} else {
				if fr[0] > end {
					endCount--
				}
				break
			}
		}
		endIndex := startIndex + endCount

		// 把原来的空闲数组进行拆分，并且组装成新的空闲数组
		newFreeRegions := make([][2]int, 0)
		for i, fr := range freeRegions {
			if i < startIndex {
				newFreeRegions = append(newFreeRegions, fr)
				continue
			}

			if i > endIndex {
				newFreeRegions = append(newFreeRegions, fr)
				continue
			}

			frStart := fr[0]
			frEnd := fr[1]
			if i == startIndex {
				frEnd = start - 1
			}
			if i == endIndex {
				frStart = end + 1
			}

			if frEnd > fr[0] {
				newFreeRegions = append(newFreeRegions, [2]int{fr[0], frEnd})
			}

			if frStart < fr[1] {
				newFreeRegions = append(newFreeRegions, [2]int{frStart, fr[1]})
			}
		}

		freeRegions = newFreeRegions
	}

	remain := 0
	for _, fr := range freeRegions {
		remain = fr[1] - fr[0] + 1 + remain
	}
	return remain
}
