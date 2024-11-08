/* Рассмотрим три числа и упорядочим их по возрастанию.

Какое число будет стоять между двумя другими? */


package main

import (
	"fmt"
	"sort"
)

func main() {
	var array = make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Scan(&array[i])
	}
	sort.Slice(array, func(i, j int) bool {
		return array[i] < array[j]
	})
	fmt.Print(array[1])
}
