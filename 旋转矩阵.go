package leetcode

// Rotate 原地对二维数组进行旋转
func Rotate(matrix [][]int) {
	sIndex := 0
	eIndex := len(matrix) - 1

	for sIndex < eIndex {
		start := sIndex
		end := eIndex

		for start < eIndex {
			swapArr := [][]int{
				{sIndex, start},
				{start, eIndex},
				{eIndex, end},
				{end, sIndex},
			}

			swap := swapArr[3]
			temp := matrix[swap[0]][swap[1]]
			for i := 3; i > 0; i-- {
				matrix[swapArr[i][0]][swapArr[i][1]] = matrix[swapArr[i-1][0]][swapArr[i-1][1]]
			}
			swap = swapArr[0]
			matrix[swap[0]][swap[1]] = temp

			start = start + 1
			end = end - 1
		}

		sIndex = sIndex + 1
		eIndex = eIndex - 1
	}
}
