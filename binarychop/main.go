package main

import (
	"fmt"
)

func main() {
	var array  = []int {1,2,23,2}
	slice := array[0:3];
	mutate(slice)

	fmt.Println(array)
}

func mutate(slice []int) {
	slice[0] = 10
}

// func main() {
// 	var array  = GetArray();
// 	var result = Chop(20, array);
// 	fmt.Println(result);
// }


// func Chop(value int, array []int) int {
// 	var result = -1;
// 	for idx, v := range array {
// 		if v == value {
// 			result = idx;
// 		}
// 	}
	
// 	return result;
// }

// func GetArray() []int {
// 	array :=  make([]int, 200000);
// 	for i := range  array {
// 		array[i] = i;
// 	}

// 	return array;
// }