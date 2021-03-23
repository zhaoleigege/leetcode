package leetcode

import "testing"

func arrCompare(a [][]int, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, ai := range a {
		if len(b) <= i {
			return false
		}
		bi := b[i]
		for j := range ai {
			if len(bi) <= j {
				return false
			}
			if bi[j] != ai[j] {
				return false
			}
		}
	}

	return true
}

func TestMerge(t *testing.T) {
	t.Run("test_0", func(t *testing.T) {
		expect := [][]int{
			{1, 6},
			{8, 10},
			{15, 18},
		}
		fact := Merge([][]int{
			{1, 3},
			{2, 6},
			{8, 10},
			{15, 18},
		})
		if !arrCompare(expect, fact) {
			t.Errorf("期望结果: %v, 实际: %v", expect, fact)
		}
	})

	t.Run("test_1", func(t *testing.T) {
		expect := [][]int{
			{1, 5},
		}
		fact := Merge([][]int{
			{1, 4},
			{4, 5},
		})
		if !arrCompare(expect, fact) {
			t.Errorf("期望结果: %v, 实际: %v", expect, fact)
		}
	})
}

func TestRotate(t *testing.T) {
	motrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}

	Rotate(motrix)
	t.Log(motrix)
}
