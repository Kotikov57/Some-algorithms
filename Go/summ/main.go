/* Вам дано целое положительное число N. Найдите все способы, которым это число можно получить, складывая целые положительные числа.
Для того, чтобы не выводить одинаковые способы много раз, суммы должны обладать следующими свойствами:
Слагаемые в рамках одного выражения идут в невозрастающем порядке.
Выражения перечисляются в лексикографическом порядке. Иначе говоря, для любой пары выражений верно, что раньше будет выписано то из них, в котором меньше число на первой несовпадающей позиции. 
Например, 2 + 1 + 1 должно быть выписано раньше, чем 2 + 2. */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func findPartitions(n int, max int, current []int, results *[][]int) {
	if n == 0 {
		partition := make([]int, len(current))
		copy(partition, current)
		*results = append(*results, partition)
		return
	}
	for i := 1; i <= max && i <= n; i++ {
		current = append(current, i)
		findPartitions(n-i, i, current, results)
		current = current[:len(current)-1]
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	var results [][]int
	findPartitions(n, n, []int{}, &results)
	var output []string
	for _, partition := range results {
		parts := make([]string, len(partition))
		for i, num := range partition {
			parts[i] = strconv.Itoa(num)
		}
		line := strings.Join(parts, " + ")
		output = append(output, line)
	}
	for _, line := range output {
		fmt.Println(line)
	}
}
