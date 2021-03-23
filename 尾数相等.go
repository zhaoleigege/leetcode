package leetcode

import (
	"math"
)

func TailNum(k int) (int, int) {
	baseNum := float64(1000)

	var n, m int
	m = int(math.Log(baseNum) / math.Log(float64(k)))
	num := math.Pow(float64(k), float64(m))
	if num < baseNum {
		m = m + 1
		baseNum = math.Pow(float64(k), float64(m))
	}

	appearArray := make([]int, 1000)
	appearArray[int(baseNum)%1000] = m

	for n = m + 1; ; n = n + 1 {
		lastThreeNum := numPowOfNLastThreeNum(k, n)
		if appearArray[lastThreeNum] != 0 {
			return appearArray[lastThreeNum], n
		} else {
			appearArray[lastThreeNum] = n
		}
	}
}

// numPowOfNLastThreeNum 计算num的n次幂的最后三位
func numPowOfNLastThreeNum(num, n int) int {
	result := 1
	num = num % 1000

	for i := 0; i < n; i++ {
		result = result * num
		if result/1000 >= 1 {
			result = result % 1000
		}
	}

	return result
}
