package leetcode

import "testing"

func TestStrGcd(t *testing.T) {
	t.Run("test_ab", func(t *testing.T) {
		mod := StrGcd("ABABAB", "ABAB")
		if mod != "AB" {
			t.Errorf("结果不等于AB, result: %s", mod)
		}
	})
	t.Run("test_abc", func(t *testing.T) {
		mod := StrGcd("ABCABC", "ABC")
		if mod != "ABC" {
			t.Errorf("结果不等于ABC, result: %s", mod)
		}
	})
	t.Run("test_", func(t *testing.T) {
		mod := StrGcd("LEET", "CODE")
		if mod != "" {
			t.Errorf("结果不等于空, result: %s", mod)
		}
	})
	t.Run("test_nono", func(t *testing.T) {
		mod := StrGcd("", "")
		if mod != "" {
			t.Errorf("结果不等于空, result: %s", mod)
		}
	})
}
