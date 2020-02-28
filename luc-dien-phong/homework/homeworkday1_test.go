package homework

import "testing"

/// Max testing
func TestMax1(t *testing.T) {
	result, err := Max([]int{0, 1, 2, 3})
	if len(err) > 0 {
		t.Error("TestMax1 error", err)
	} else {
		t.Log("TestMax1 success. result: ", result)
	}
}
func TestMax2(t *testing.T) {
	result, err := Max([]int{})
	if len(err) > 0 {
		t.Error("TestMax2 error", err)
	} else {
		t.Log("TestMax2 success. result: ", result)
	}
}

/// Min testing
func TestMin1(t *testing.T) {
	result, err := Min([]int{0, 1, 2, 3})
	if len(err) > 0 {
		t.Error("TestMin1 error", err)
	} else {
		t.Log("TestMin1 success. result: ", result)
	}
}
func TestMin2(t *testing.T) {
	result, err := Min([]int{})
	if len(err) > 0 {
		t.Error("TestMin2 error", err)
	} else {
		t.Log("TestMin2 success. result: ", result)
	}
}
func TestMin3(t *testing.T) {
	result, err := Min([]int{0, -1, -2, 3, 4, 5})
	if len(err) > 0 {
		t.Error("TestMin2 error", err)
	} else {
		t.Log("TestMin2 success. result: ", result)
	}
}
