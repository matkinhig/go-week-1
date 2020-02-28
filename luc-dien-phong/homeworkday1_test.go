package homework

import "testing"

func Test1(t *testing.T) {
	result, err := MinInt([]int{0, 1, 2, 3})
	if err == true {
		t.Log("Test1 success", "Test1", result)
	} else {
		t.Error("Function error", "Test1", result)
	}
}
