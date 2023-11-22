package main

import "fmt"

func SliceDelete[T any](vals []T, idx int) []T {
	if idx < 0 || idx > len(vals) {
		return vals
	}
	res := append(vals[:idx], vals[idx+1:]...)
	return res

}

func main() {
	s1 := make([]int, 0)
	for i := 0; i < 10; i++ {
		s1 = append(s1, i)
	}
	res := SliceDelete(s1, 2)
	fmt.Printf("res:%v \n", res)

}
