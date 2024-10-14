package main

import (
	"fmt"
	"strconv"
)

func partition(n int){
	if n == 1{
		return
	}
	var out string
	for i := 1; i <= n / 2; i++{
		out += strconv.Itoa(n - i) + " + " + strconv.Itoa(i)
		fmt.Println(out)
		partition(n - i)
	}

}

func main() {
/* 	var n int
	fmt.Scan(&n) */
	n := 4
	partition(n)
}