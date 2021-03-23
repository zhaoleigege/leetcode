package leetcode

// TrailingZeros 计算num阶乘末尾0的个数
// 末尾0的个数其实就是这个数阶乘当中可以乘以的5的个数
// 25 -> 5 * 5，这里有两个5
func TrailingZeros(num int) int {
	count := 0
	for num != 0 {
		count = count + num/5
		num = num / 5
	}

	return count
}
