/* Задан массив a размера n. 
Необходимо посчитать количество уникальных элементов в данном массиве.
Элемент называется уникальным, если встречается в массиве ровно один раз. */

package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	array := make([]int, n)
	a := make(map[int]int)
	uniqueCount := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
		if a[array[i]] == 0 {
			uniqueCount++
			a[array[i]] = 1
		} else if a[array[i]] == 1 {
			uniqueCount--
			a[array[i]] = 2
		}
	}
	fmt.Println(uniqueCount)
}
