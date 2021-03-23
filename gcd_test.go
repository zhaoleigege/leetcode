package leetcode

import "testing"

func TestGcd(t *testing.T) {
	t.Run("test_0", func(t *testing.T) {
		if Gcd(5, 6) != 1 {
			t.Errorf("不等于1, 结果: %d", Gcd(5, 6))
		}
	})

	t.Run("test_8", func(t *testing.T) {
		if Gcd(24, 16) != 8 {
			t.Errorf("不等于8, 结果: %d", Gcd(5, 6))
		}
	})
}

func TestLoopGcd(t *testing.T) {
	t.Run("test_0", func(t *testing.T) {
		if LoopGcd(5, 6) != 1 {
			t.Errorf("不等于1, 结果: %d", Gcd(5, 6))
		}
	})

	t.Run("test_8", func(t *testing.T) {
		if LoopGcd(24, 16) != 8 {
			t.Errorf("不等于8, 结果: %d", Gcd(5, 6))
		}
	})
}
