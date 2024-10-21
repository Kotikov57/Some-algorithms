/* Для заданного числа n требуется найти число от 1 до n (включительно), которое имеет максимальное число положительных целых делителей. 
Если таких чисел несколько, требуется найти максимальное. */

package main

import (
	"fmt"
	"math"
)

func dividers(n int) (int, int) {
	if n == 1 {
		return 1, 1
	}
	number := 0
	max := 0
	for i := 1; i <= n; i++ {
		count := 0
		for j := 1; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				count += 2
			}
		}
		if count >= max {
			max = count
			number = i
		}
	}
	return number, max
}

func main() {
	var n int
	fmt.Scan(&n)
	a, b := dividers(n)
	fmt.Println(a)
	fmt.Println(b)
}
