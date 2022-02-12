package main

import (
	"fmt"
	"os"
	"strconv"
)

func max(larger_value int, list []int, index_value int) int {

	if index_value > len(list)-1 {
		return larger_value
	}
	if list[index_value] > larger_value {
		larger_value = list[index_value]
	}

	return max(larger_value, list, index_value+1) //chamada recursiva
}

func main() {
	var list []int

	for _, a := range os.Args {
		a, err := strconv.Atoi(a)
		if err == nil {
			list = append(list, a)
		}
	}

	fmt.Println("O maior valor da lista é:", max(list[1], list, 2))
}
