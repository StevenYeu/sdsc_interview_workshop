package main

import (
	"fmt"
	"presentation/lib"
)

func main() {
    numbers := []int{1,3,5,6,7}
    index := lib.BinarySearch(numbers,1)
    fmt.Printf("Searching 1 expectecing index 0, got:%d", index)
}
