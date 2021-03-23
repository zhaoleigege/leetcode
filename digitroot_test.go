package leetcode

import "testing"

func TestAddDigitsRecursion(t *testing.T) {
	t.Run("test_0", func(t *testing.T) {
		if AddDigitsRecursion(0) != 0 {
			t.Errorf("0的数根不等于0, 为: %d", AddDigitsRecursion(0))
		}
	})

	t.Run("test_2", func(t *testing.T) {
		if AddDigitsRecursion(38) != 2 {
			t.Errorf("38的数根不等于2，为: %d", AddDigitsRecursion(38))
		}
	})
}

func TestAddDigitsLoop(t *testing.T) {
	t.Run("test_0", func(t *testing.T) {
		if AddDigitsLoop(0) != 0 {
			t.Errorf("0的数根不等于0, 为: %d", AddDigitsLoop(0))
		}
	})

	t.Run("test_2", func(t *testing.T) {
		if AddDigitsLoop(38) != 2 {
			t.Errorf("38的数根不等于2，为: %d", AddDigitsLoop(38))
		}
	})
}
