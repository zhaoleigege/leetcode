package leetcode

import "testing"

func TestSearchInsert(t *testing.T) {
	t.Run("test_0", func(t *testing.T) {
		expect := 2
		fact := SearchInsert([]int{1,3,5,6}, 5)
		if expect != fact {
			t.Errorf("期望: %d, 实际: %d", expect, fact)
		}
	})
	t.Run("test_1", func(t *testing.T) {
		expect := 1
		fact := SearchInsert([]int{1,3,5,6}, 2)
		if expect != fact {
			t.Errorf("期望: %d, 实际: %d", expect, fact)
		}
	})
	t.Run("test_1", func(t *testing.T) {
		expect := 4
		fact := SearchInsert([]int{1,3,5,6}, 7)
		if expect != fact {
			t.Errorf("期望: %d, 实际: %d", expect, fact)
		}
	})
	t.Run("test_0", func(t *testing.T) {
		expect := 0
		fact := SearchInsert([]int{1,3,5,6}, 0)
		if expect != fact {
			t.Errorf("期望: %d, 实际: %d", expect, fact)
		}
	})
}
