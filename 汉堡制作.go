package leetcode

// NumOfChick 鸡兔同笼问题
// 一只鸡两条腿，一只兔子四条腿，问给定腿数动物数至多和至少是多少
func NumOfChick(num int) (int, int) {
	if num <= 0 {
		return 0, 0
	}

	// 最多，全是鸡的情况
	max := num / 2
	mod := num % 2
	if mod != 0 {
		return 0, 0
	}

	// 最少，全是兔子或者至多只有一只鸡
	min := num / 4
	mod = num % 4
	switch mod {
	case 0:
	case 1, 3:
		return 0, 0
	case 2:
		min = min + 1
	}

	return max, min
}

// NumOfBurger 不浪费原料的汉堡制作方案
func NumOfBurger(tomato, cheese int) []int {
	if tomato == 0 && cheese == 0 {
		return []int{0, 0}
	}

	if tomato <= 0 || cheese <= 0 {
		return []int{}
	}

	double := tomato - 2*cheese
	if double < 0 || double%2 != 0 {
		return []int{}
	}

	burger := double / 2
	ingredient := cheese - burger
	if ingredient < 0 {
		return []int{}
	}

	return []int{burger, ingredient}
}
