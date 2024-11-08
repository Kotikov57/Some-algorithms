package main

import (
	"fmt"
)

func main() {
	var first, second, third, fourth int
	fmt.Scan(&first, &second, &third, &fourth)
	summary := first + second + third + fourth
	count := 0
	for summary > 4 && count != 4 {
		count++
		summary -= 4
	}
	if summary < 2 {
		count = 1
	}
	fmt.Println(count)
}
