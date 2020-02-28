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
			resul = arr[i]
		}
	}

	return resul, ""
}
