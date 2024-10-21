/* https://coderun.yandex.ru/problem/tic-tac-toe/description */

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func vert(n, m int, arr [][]byte) bool {
	for i := 0; i < m; i++ {
		prev := arr[0][i]
		if prev == '.' {
			continue
		}
		
		count := 1
		for j := 1; j < n; j++ {
			if count < 5-(n-j) {
				continue
			}
			cur := arr[j][i]
			if cur == prev && cur != '.' {
				count++
			} else {
				prev = cur
				count = 1
			}
			if count == 5 {
				return true
			}
		}
	}
	return false
}

func hor(n, m int, arr [][]byte) bool {
	count := 1
	for i := 0; i < n; i++ {
		prev := arr[i][0]
		if prev == '.' {
			continue
		}
		count = 1
		for j := 1; j < m; j++ {
			if count < 5-(m-j) {
				continue
			}
			cur := arr[i][j]
			if cur == prev && cur != '.' {
				count++
			} else {
				prev = cur
				count = 1
			}
			if count == 5 {
				return true
			}
		}
	}
	return false
}

func diag(n, m int, arr [][]byte) bool {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if arr[i][j] == '.' {
				continue
			}
			if diagLeft(n, m, i, j, arr) || diagRight(n, i, j, arr) {
				return true
			}
		}
	}
	return false
}

func diagLeft(n, m, i, j int, arr [][]byte) bool {
	count := 1
	prev := arr[i][j]
	k, l := i+1, j+1
	if count < 5-int(math.Min(float64(n-i), float64(m-j))) {
		return false
	}
	for k < n && l < m {
		cur := arr[k][l]
		if cur == prev && cur != '.' {
			count++
		} else {
			prev = cur
			count = 1
		}
		if count == 5 {
			return true
		}
		k++
		l++
	}
	return false
}

func diagRight(n, i, j int, arr [][]byte) bool {
	count := 1
	if count < 5-int(math.Min(float64(n-i), float64(j))) {
		return false
	}
	prev := arr[i][j]
	k, l := i+1, j-1
	for k < n && l >= 0 {
		cur := arr[k][l]
		if cur == prev && cur != '.' {
			count++
		} else {
			prev = cur
			count = 1
		}
		if count == 5 {
			return true
		}
		k++
		l--
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.Fields(str)[0])
	m, _ := strconv.Atoi(strings.Fields(str)[1])
	array := make([][]byte, n)
	for i := 0; i < n; i++ {
		array[i] = make([]byte, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			bit, _ := reader.ReadByte()
			if bit == '\n' {
				j--
				continue
			}
			array[i][j] = bit
		}
	}
	if vert(n, m, array) || hor(n, m, array) || diag(n, m, array) {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}
