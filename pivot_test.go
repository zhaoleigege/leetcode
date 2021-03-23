package leetcode

import "testing"

func TestPivotIndex(t *testing.T) {
	t.Run("test_-1", func(t *testing.T) {
		expect := -1
		fact := PivotIndex([]int{1, 2, 3})
		if fact != expect {
			t.Errorf("中心索引期望为: %d, 实际为: %d", expect, fact)
		}
	})

	t.Run("test_3", func(t *testing.T) {
		expect := 3
		fact := PivotIndex([]int{1, 7, 3, 6, 5, 6})
		if fact != expect {
			t.Errorf("中心索引期望为: %d, 实际为: %d", expect, fact)
		}
	})

	t.Run("test_0", func(t *testing.T) {
		expect := 0
		fact := PivotIndex([]int{-1,-1,-1,0,1,1})
		if fact != expect {
			t.Errorf("中心索引期望为: %d, 实际为: %d", expect, fact)
		}
	})
}
