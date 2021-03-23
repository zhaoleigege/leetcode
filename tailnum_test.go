package leetcode

import "testing"

func TestTailNum(t *testing.T) {
	t.Run("test_10", func(t *testing.T) {
		min, max := TailNum(10)
		eMin, eMax := 3, 4
		if min != eMin || max != eMax {
			t.Errorf("结果: min: %d, max: %d不等于预期: min: %d, max: %d", min, max, eMin, eMax)
		}
	})

	t.Run("test_20", func(t *testing.T) {
		min, max := TailNum(20)
		eMin, eMax := 3, 4
		if min != eMin || max != eMax {
			t.Errorf("结果: min: %d, max: %d不等于预期: min: %d, max: %d", min, max, eMin, eMax)
		}
	})

	t.Run("test_125", func(t *testing.T) {
		min, max := TailNum(125)
		eMin, eMax := 2, 4
		if min != eMin || max != eMax {
			t.Errorf("结果: min: %d, max: %d不等于预期: min: %d, max: %d", min, max, eMin, eMax)
		}
	})

	t.Run("test_1000", func(t *testing.T) {
		min, max := TailNum(10)
		eMin, eMax := 3, 4
		if min != eMin || max != eMax {
			t.Errorf("结果: min: %d, max: %d不等于预期: min: %d, max: %d", min, max, eMin, eMax)
		}
	})
}
