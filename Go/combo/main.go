package main

import (
	"fmt"
//	"bufio"
//	"os"
)

func main() {
	var n, x, k int
	fmt.Scan(&n)
	prices := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&prices[i])
	}
	fmt.Scan(&x)
	combo := make([]int, 4)
	for i := 0; i < 4; i++ {
		fmt.Scan(&combo[i])
	}
	fmt.Scan(&k)
	productList := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Scan(&productList[i])
	}

}