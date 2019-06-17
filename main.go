package main

import (
	"fmt"
	"reflect"
)

func main() {

	sliceInt := []int{3, 1, 2, 5, 7, 4}
	fmt.Println(searchMax(sliceInt, func(i, j int) bool { return sliceInt[i] > sliceInt[j] }))

	sliceStr := []string{"ab", "bc", "ba", "da", "ad"}
	fmt.Println(searchMax(sliceStr, func(i, j int) bool { return sliceStr[i] > sliceStr[j] }))

	sliceStruct := []struct {
		Name string
		Age  int
	}{
		{"Pavel", 30},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}

	fmt.Println(searchMax(sliceStruct, func(i, j int) bool { return sliceStruct[i].Name > sliceStruct[j].Name }))
	fmt.Println(searchMax(sliceStruct, func(i, j int) bool { return sliceStruct[i].Age > sliceStruct[j].Age }))
}

func searchMax(slice interface{}, c func(i, j int) bool) interface{} {

	s := reflect.ValueOf(slice)

	var j int

	for i := 1; i < s.Len(); i++ {
		if c(i, j) {
			j = i
		}
	}
	return s.Index(j)
}
