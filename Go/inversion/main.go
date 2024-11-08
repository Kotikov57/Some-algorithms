/* Пусть перестановка чисел от 1 до n. Будем говорить, что пара индексов (i,j) образует инверсию, если i<j и pi>pj
Задана некоторая перестановка, требуется определить среднее количество инверсий в перестановке, полученной из данной после одной перестановки пары элементов. 
При этом индексы переставляемых элементов выбираются равновероятно среди всех пар различных чисел от 1 до n. */

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

func updateFenwickTree(tree []int, index int, value int, n int) {
	for index <= n {
		tree[index] += value
		index += index & (-index)
	}
}
func queryFenwickTree(tree []int, index int) int {
	sum := 0
	for index > 0 {
		sum += tree[index]
		index -= index & (-index)
	}
	return sum
}
func countInversions(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	maxElement := 0
	for _, v := range arr {
		if v > maxElement {
			maxElement = v
		}
	}
	fenwickTree := make([]int, maxElement+1)
	inversionCount := 0
	for i := n - 1; i >= 0; i-- {
		inversionCount += queryFenwickTree(fenwickTree, arr[i]-1)
		updateFenwickTree(fenwickTree, arr[i], 1, maxElement)
	}

	return inversionCount
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	str, _ := reader.ReadString('\n')
	n, _ = strconv.Atoi(str[:(len(str))-2])
	array := make([]int, n)
	arrString, _ := reader.ReadString('\n')
	arrString = strings.TrimSpace(arrString)
	numbers := strings.Fields(arrString)
	for i, numStr := range numbers {
		num, _ := strconv.Atoi(numStr)
		array[i] = num
	}
	result := 0
	count := 0
	for i := 0; i < n - 1; i++ {
		for j := i + 1; j < n; j++ {
			count++
			array[i], array[j] = array[j], array[i]

			result += countInversions(array)
			array[i], array[j] = array[j], array[i]
		}
	}
	fmt.Println(result)
	fmt.Println(count)
	fmt.Printf("%d/%d", result / gcd(result, count), count / gcd(result,count))
}
