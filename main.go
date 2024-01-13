package main

import "fmt"

func main() {
	nums := []int{8, 10, 209}
	for _, num := range nums {
		fmt.Println(num, "plus 10 = ", addTen(num))
	}
}
