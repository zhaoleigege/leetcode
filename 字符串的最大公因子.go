package leetcode

// strMod 最两个字符串进行求mod
// 这里已经假定len(b) <= len(a)
func strMod(a, b string) string {
	if len(b) <= 0 {
		return ""
	}

	for i := 0; i < len(a); {
		if len(a)-i < len(b) {
			return a[i:]
		}

		for j := 0; j < len(b); {
			if a[i] == b[j] {
				i++
				j++
			} else {
				return ""
			}
		}
	}

	return b
}

func StrGcd(a, b string) string {
	if len(a) < len(b) {
		c := a
		a = b
		b = c
	}

	mod := strMod(a, b)
	for mod != b {
		a = b
		b = mod
		mod = strMod(a, b)
	}

	return b
}
