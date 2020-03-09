package homework

import "testing"

/// Max testing
func TestMax1(t *testing.T) {
	result, err := Max([]int{0, 1, 2, 3})
	if len(err) > 0 {
		t.Error("TestMax1 error: ", err)
	} else {
		t.Log("TestMax1 success. result: ", result)
	}
}
func TestMax2(t *testing.T) {
	result, err := Max([]int{})
	if len(err) > 0 {
		t.Error("TestMax2 error: ", err)
	} else {
		t.Log("TestMax2 success. result: ", result)
	}
}

/// Min testing
func TestMin1(t *testing.T) {
	result, err := Min([]int{0, 1, 2, 3})
	if len(err) > 0 {
		t.Error("TestMin1 error: ", err)
	} else {
		t.Log("TestMin1 success. result: ", result)
	}
}
func TestMin2(t *testing.T) {
	result, err := Min([]int{})
	if len(err) > 0 {
		t.Error("TestMin2 error: ", err)
	} else {
		t.Log("TestMin2 success. result: ", result)
	}
}
func TestMin3(t *testing.T) {
	result, err := Min([]int{0, -1, -2, 3, 4, 5})
	if len(err) > 0 {
		t.Error("TestMin3 error: ", err)
	} else {
		t.Log("TestMin3 success. result: ", result)
	}
}

func TestContain1(t *testing.T) {
	result, err := Contains([]int{0, -1, -2, 3, 4, 5}, -2)
	if len(err) > 0 {
		t.Error("TestContain1 error: ", err)
	} else {
		t.Log("TestContain1 success. result: ", result)
	}
}
func TestContain2(t *testing.T) {
	result, err := Contains([]int{}, -2)
	if len(err) > 0 {
		t.Error("TestContain2 error: ", err)
	} else {
		t.Log("TestContain2 success. result: ", result)
	}
}
func TestContain3(t *testing.T) {
	result, err := Contains([]int{0, -1, -2, 3, 4, 5}, 100)
	if len(err) > 0 {
		t.Error("TestContain3 error: ", err)
	} else {
		t.Log("TestContain3 success. result: ", result)
	}
}

// Difference testing
func TestDifference1(t *testing.T) {
	result, err := Difference([]int{0, -1, -2, 3, 4, 5}, []int{0, -2, 4})
	if len(err) > 0 {
		t.Error("TestDifference1 error: ", err)
	} else {
		t.Log("TestContTestDifference1ain23 success. result: ", result)
	}
}
func TestDifference2(t *testing.T) {
	result, err := Difference([]int{0, -1, -2, 3, 4, 5}, []int{6, 7, 9})
	if len(err) > 0 {
		t.Error("TestDifference2 error: ", err)
	} else {
		t.Log("TestDifference2 success. result: ", result)
	}
}
func TestDifference3(t *testing.T) {
	result, err := Difference([]int{0, -1, -2, 3, 4, 5}, []int{0, -1, -2, 3, 4, 5})
	if len(err) > 0 {
		t.Error("TestDifference3 error: ", err)
	} else {
		t.Log("TestDifference3 success. result: ", result)
	}
}
func TestDifference4(t *testing.T) {
	result, err := Difference([]int{1, 2, 3}, []int{})
	if len(err) > 0 {
		t.Error("TestDifference4 error: ", err)
	} else {
		t.Log("TestDifference4 success. result: ", result)
	}
}
func TestDifference5(t *testing.T) {
	result, err := Difference([]int{}, []int{1, 2, 3})
	if len(err) > 0 {
		t.Error("TestDifference5 error: ", err)
	} else {
		t.Log("TestDifference5 success. result: ", result)
	}
}

// IndexOf testing
func TestIndexOf1(t *testing.T) {
	result, err := IndexOf([]int{1, 2, 3}, 1)
	if len(err) > 0 {
		t.Error("TestIndexOf1 error: ", err)
	} else {
		t.Log("TestIndexOf1 success. result: ", result)
	}
}
func TestIndexOf2(t *testing.T) {
	result, err := IndexOf([]int{1, 2, 3}, 0)
	if len(err) > 0 {
		t.Error("TestIndexOf2 error: ", err)
	} else {
		t.Log("TestIndexOf2 success. result: ", result)
	}
}
func TestIndexOf3(t *testing.T) {
	result, err := IndexOf([]int{}, 0)
	if len(err) > 0 {
		t.Error("TestIndexOf3 error: ", err)
	} else {
		t.Log("TestIndexOf3 success. result: ", result)
	}
}
