package main

import (
	"fmt"
)

func main() {
	var array [10]int
	var aSlice = array[5:6]
	fmt.Println("length of slice: ", len(aSlice))
	fmt.Println("capacity of slice: ", cap(aSlice))
	fmt.Println(&aSlice[0] == &array[5])
}
