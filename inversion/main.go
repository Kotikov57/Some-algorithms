package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

func mergeSortAndCount(arr []int) ([]int, int) {
	if len(arr) < 2 {
		return arr, 0
	}

	// Разделяем массив на две части
	mid := len(arr) / 2
	left, leftInv := mergeSortAndCount(arr[:mid])
	right, rightInv := mergeSortAndCount(arr[mid:])

	// Сливаем две половины и считаем инверсии
	merged, splitInv := mergeAndCount(left, right)
	
	// Общее количество инверсий = инверсии в левой части + в правой части + при слиянии
	totalInv := leftInv + rightInv + splitInv

	return merged, totalInv
}

// Функция для слияния двух массивов и подсчёта инверсий
func mergeAndCount(left, right []int) ([]int, int) {
	merged := make([]int, 0, len(left)+len(right))
	i, j, invCount := 0, 0, 0

	// Сливаем два отсортированных подмассива
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
			// Все оставшиеся элементы в left больше, чем right[j], значит есть инверсии
			invCount += len(left) - i
		}
	}

	// Добавляем оставшиеся элементы
	merged = append(merged, left[i:]...)
	merged = append(merged, right[j:]...)

	return merged, invCount
}

/* func isInversion(array []int, i, j int) bool {
	if i < j && array[i] > array[j] {
		return true
	}
	return false
} */

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
	var arrayCopy []int
	for i := 0; i < n - 1; i++ {
		for j := i + 1; j < n; j++ {
			count++
			arrayCopy[i], arrayCopy[j] = arrayCopy[j], arrayCopy[i]
			_, result = mergeSortAndCount(arrayCopy)
		}
	}
	fmt.Println(result)
	fmt.Println(count)
	fmt.Printf("%d/%d", result / gcd(result, count), count / gcd(result,count))
}
