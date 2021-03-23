package leetcode

// Gcd 最大公约数
func Gcd(a, b int) int {
	if a < b {
		c := a
		a = b
		b = c
	}

	if b <= 1 {
		return 1
	}

	mod := a % b
	if mod == 0 {
		return b
	}

	return Gcd(mod, b)
}

// LoopGcd 使用循环的方式计算最大公约数
func LoopGcd(a, b int) int {
	if a < b {
		c := a
		a = b
		b = c
	}

	mod := a % b
	for mod != 0 {
		a = b
		b = mod
		mod = a % b
	}

	return b
}
