package homework

import "fmt"

//PrintName comment
func PrintName() {
	fmt.Println("Test")
}

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
