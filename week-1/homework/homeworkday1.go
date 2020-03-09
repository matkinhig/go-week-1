package homework

import "fmt"

// Max cmt
func Max(arr []int) (resul int, err string) {
	if len(arr) == 0 {
		return resul, "lenght not empty"
	}

	resul = arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > resul {
			resul = arr[i]
		}
	}

	return resul, ""
}

// Min cmt
func Min(arr []int) (resul int, err string) {
	if len(arr) == 0 {
		return resul, "lenght not empty"
	}

	resul = arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] < resul {
			return arr[i], ""
		}
	}

	return resul, ""
}

// Contains cmt, return index
func Contains(arr []int, value int) (result int, err string) {
	if len(arr) == 0 {
		return result, "lenght not empty"
	}

	for i := 0; i < len(arr); i++ {
		if value == arr[i] {
			return i, ""
		}
	}

	return result, fmt.Sprintf("%d not found", value)
}

// Difference cmt
func Difference(arr []int, others []int) (result []int, err string) {
	if len(arr) == 0 {
		return result, "arr not empty"
	}
	if len(others) == 0 {
		return result, "others not empty"
	}

	for i := 0; i < len(arr); i++ {
		found := false
		for e := 0; e < len(others); e++ {
			if arr[i] == others[e] {
				found = true
				break
			}
		}
		if !found {
			result = append(result, arr[i])
		}
	}

	if len(result) == 0 {
		return result, "no other value"
	}

	return result, ""
}

// IndexOf , return index
func IndexOf(arr []int, value int) (result int, err string) {
	if len(arr) == 0 {
		return result, "arr not empty"
	}

	result = -1
	for i := 0; i < len(arr); i++ {
		if arr[i] == value {
			result = i
			break
		}
	}

	if result == -1 {
		return result, "not found"
	}
	return result, ""

}
