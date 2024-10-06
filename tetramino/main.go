/* На шахматном поле 8×8 некоторые клетки пустые, а некоторые заняты фигурами.
Определите количество способов разместить тетрамино на этом поле, чтобы фигура занимала целиком четыре свободные клетки.
В задаче мы рассматриваем тетрамино только одного типа. 
***
 *
 */

package main

import (
	"bufio"
	"fmt"
	"os"
)

func left(array [][]byte, i, j int) bool {
	if array[i][j] == '.' && array[i+1][j] == '.' && array[i+2][j] == '.' && array[i+1][j-1] == '.' {
		return true
	}
	return false
}

func right(array [][]byte, i, j int) bool {
	if array[i][j] == '.' && array[i+1][j] == '.' && array[i+2][j] == '.' && array[i+1][j+1] == '.' {
		return true
	}
	return false
}

func up(array [][]byte, i, j int) bool {
	if array[i][j] == '.' && array[i][j+1] == '.' && array[i][j+2] == '.' && array[i-1][j+1] == '.' {
		return true
	}
	return false
}

func down(array [][]byte, i, j int) bool {
	if array[i][j] == '.' && array[i][j+1] == '.' && array[i][j+2] == '.' && array[i+1][j+1] == '.' {
		return true
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	array := make([][]byte, 8)
	for i := 0; i < 8; i++ {
		array[i], _ = reader.ReadBytes('\n')
	}
	count := 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if j < 6 {
				if i != 0 {
					if up(array, i, j) {
						count++
					}
				}
				if i != 7 {
					if down(array, i, j) {
						count++
					}
				}
			}
			if i < 6 {
				if j != 0 {
					if left(array, i, j) {
						count++
					}
				}
				if j != 7 {
					if right(array, i, j) {
						count++
					}
				}
			}
		}
	}
	fmt.Print(count)
}
