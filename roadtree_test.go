package leetcode

import "testing"

func TestBuildTrain(t *testing.T) {
	t.Run("test_0", func(t *testing.T) {
		result := BuildTrain(500, [][2]int{
			{150, 300},
			{100, 200},
			{470, 471},
		})
		expect := 298
		if result != expect {
			t.Errorf("不等于期望值%d, result: %d", expect, result)
		}
	})
}
