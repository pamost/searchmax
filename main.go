package main

import (
	"fmt"
	"reflect"
)

func main() {

	notSlice := "notSlice"
	fmt.Println(searchMax(notSlice, func(i, j int) bool { return notSlice[i] > notSlice[j] }))

	sliceEmpty := []int{}
	fmt.Println(searchMax(sliceEmpty, func(i, j int) bool { return sliceEmpty[i] > sliceEmpty[j] }))

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

func searchMax(slice interface{}, compare func(i, j int) bool) interface{} {

	var result interface{}

	//рефлексия среза для доступа к элементам
	data := reflect.ValueOf(slice)

	if data.Kind() != reflect.Slice {
		result = "not a slice"
	} else if data.Len() == 0 {
		result = "empty slice"
	} else {
		// поиск максимального элемента в срезе
		var j int // индекс первого элемента среза
		for i := 1; i < data.Len(); i++ {
			if compare(i, j) {
				j = i // презаписываем значение j если i > j
			}
		}
		result = data.Index(j) // выводим значение j-го элемента среза
	}
	return result
}
