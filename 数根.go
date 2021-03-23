package leetcode

func AddDigitsRecursion(num int) int {
	if num < 10 {
		return num
	}

	root := 0
	for num > 0 {
		root = root + num%10
		num = num / 10
	}

	return AddDigitsRecursion(root)
}

func AddDigitsLoop(num int) int {
	root := num
	for root > 9 {
		num = root
		root = 0

		for num > 0 {
			root = root + num%10
			num = num / 10
		}
	}

	return root
}
