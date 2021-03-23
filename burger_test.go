package leetcode

import "testing"

func TestNumOfChick(t *testing.T) {
	t.Run("test_0", func(t *testing.T) {
		max, min := NumOfChick(3)
		if max != 0 && min != 0 {
			t.Errorf("最大: %d 最小: %d不为0", max, min)
		}
	})

	t.Run("test_exist1", func(t *testing.T) {
		max, min := NumOfChick(20)
		if max != 10 && min != 5 {
			t.Errorf("最大: %d不为10，最小: %d不为5", max, min)
		}
	})

	t.Run("test_exist2", func(t *testing.T) {
		max, min := NumOfChick(100)
		if max != 50 && min != 25 {
			t.Errorf("最大: %d不为50，最小: %d不为25", max, min)
		}
	})

	t.Run("test_exist3", func(t *testing.T) {
		max, min := NumOfChick(58)
		if max != 29 && min != 15 {
			t.Errorf("最大: %d不为19，最小: %d不为15", max, min)
		}
	})
}

func TestNumOfBurger(t *testing.T) {
	t.Run("test_1_6", func(t *testing.T) {
		result := NumOfBurger(16, 7)
		if len(result) != 2 {
			t.Errorf("结果不对result: %v", result)
		}
		burger := 1
		ingredient := 6
		if result[0] != burger || result[1] != ingredient {
			t.Errorf("结果不符合预期burger: %d != %d, ingredient: %d != %d", result[0], burger,
				result[1], ingredient)
		}
	})
	t.Run("test_0_0_1", func(t *testing.T) {
		result := NumOfBurger(17, 4)
		if len(result) != 0 {
			t.Errorf("结果不对result: %v", result)
		}
	})
	t.Run("test_0_0_2", func(t *testing.T) {
		result := NumOfBurger(4, 17)
		if len(result) != 0 {
			t.Errorf("结果不对result: %v", result)
		}
	})
	t.Run("test_0_0_3", func(t *testing.T) {
		result := NumOfBurger(0, 0)
		if len(result) != 2 {
			t.Errorf("结果不对result: %v", result)
		}
		burger := 0
		ingredient := 0
		if result[0] != burger || result[1] != ingredient {
			t.Errorf("结果不符合预期burger: %d != %d, ingredient: %d != %d", result[0], burger,
				result[1], ingredient)
		}
	})
	t.Run("test_0_1", func(t *testing.T) {
		result := NumOfBurger(2, 1)
		if len(result) != 2 {
			t.Errorf("结果不对result: %v", result)
		}
		burger := 0
		ingredient := 1
		if result[0] != burger || result[1] != ingredient {
			t.Errorf("结果不符合预期burger: %d != %d, ingredient: %d != %d", result[0], burger,
				result[1], ingredient)
		}
	})
	t.Run("test_1_6", func(t *testing.T) {
		result := NumOfBurger(16, 7)
		if len(result) != 2 {
			t.Errorf("结果不对result: %v", result)
		}
		burger := 1
		ingredient := 6
		if result[0] != burger || result[1] != ingredient {
			t.Errorf("结果不符合预期burger: %d != %d, ingredient: %d != %d", result[0], burger,
				result[1], ingredient)
		}
	})
	t.Run("test_negative", func(t *testing.T) {
		result := NumOfBurger(2385088, 164763)
		if len(result) != 0 {
			t.Errorf("结果不对result: %v", result)
		}
	})
}
