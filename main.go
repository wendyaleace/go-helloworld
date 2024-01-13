package main

import "fmt"

func main() {
	nums := []int{9, 19, 209, 569}
	for _, num := range nums {
		fmt.Println(num, "plus 10 = ", addTen(num))
	}
}
