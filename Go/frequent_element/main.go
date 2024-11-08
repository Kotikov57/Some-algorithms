/* Дан массив a из n целых чисел. Напишите программу, которая найдет число, которое встречается в массиве наибольшее число раз.
Формат ввода
В первой строке входных данных записано число (1≤n≤300000). Во второй строке записаны n целых чисел (0≤a≤1000000000).
Формат вывода
Выведите единственное число x, наибольшее из чисел, которое встречается в a максимальное количество раз. */

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.Fields(line)[0])
	arr, _ := reader.ReadString('\n')
	array := strings.Fields(arr)
	numbers := make([]int, n)
	mapa := make(map[int]int)
	for i := range array {
		numbers[i], _ = strconv.Atoi(array[i])
		mapa[numbers[i]]++
	}
	max := -1
	for _, value := range mapa {
		if value > max {
			max = value
		}
	}
	out := []int{}
	for key, value := range mapa {
		if value == max {
			out = append(out, key)
		}
	}
	sort.Slice(out, func(i, j int) bool {
		return out[i] > out[j]
	})
	fmt.Println(out[0])
}
