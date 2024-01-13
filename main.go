package main

import "fmt"

func main() {
	nums := []int{39, 4323, 2, 569}
	for _, num := range nums {
		fmt.Println(num, "plus 10 = ", addTen(num))
	}
}
