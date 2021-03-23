package leetcode

import "testing"

func TestTrailingZeros(t *testing.T) {
	t.Run("test_0", func(t *testing.T) {
		result := TrailingZeros(4)
		expect := 0
		if result != expect {
			t.Errorf("结果: %d不等于期望值: %d", result, expect)
		}
	})

	t.Run("test_1", func(t *testing.T) {
		result := TrailingZeros(5)
		expect := 1
		if result != expect {
			t.Errorf("结果: %d不等于期望值: %d", result, expect)
		}
	})

	t.Run("test_6", func(t *testing.T) {
		result := TrailingZeros(25)
		expect := 6
		if result != expect {
			t.Errorf("结果: %d不等于期望值: %d", result, expect)
		}
	})
}
