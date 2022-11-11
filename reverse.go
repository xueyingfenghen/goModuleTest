package goModuleTest

import "fmt"

func ReverseRunes(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	fmt.Println("=====success=====")

	fmt.Println("======commit1======")

	return string(r)
}

func testAdd(a *int) int {
	*a += 1
	return *a
}
