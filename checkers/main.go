/* Чтобы избежать ошибок, разработчики решили написать программу, которая будет по текущей позиции определять, можно ли сходить так, чтобы срубить шашку противника.
Но прямо сейчас у них много других важных проектов, поэтому запрограммировать анализатор позиции попросили вас.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func captureableWhite(i, j int, array [][]int) bool {
	if array[i-1][j-1] == 2 && array[i-2][j-2] == 0 {
		return true
	}
	if array[i-1][j+1] == 2 && array[i-2][j+2] == 0 {
		return true
	}
	if array[i+1][j-1] == 2 && array[i+2][j-2] == 0 {
		return true
	}
	if array[i+1][j+1] == 2 && array[i+2][j+2] == 0 {
		return true
	}
	return false
}

func captureableBlack(i, j int, array [][]int) bool {
	if array[i-1][j-1] == 1 && array[i-2][j-2] == 0 {
		return true
	}
	if array[i-1][j+1] == 1 && array[i-2][j+2] == 0 {
		return true
	}
	if array[i+1][j-1] == 1 && array[i+2][j-2] == 0 {
		return true
	}
	if array[i+1][j+1] == 1 && array[i+2][j+2] == 0 {
		return true
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	size := strings.TrimSpace(line)
	sizes := strings.Split(size, " ")
	n, _ := strconv.Atoi(sizes[0])
	m, _ := strconv.Atoi(sizes[1])
	field := make([][]int, n+4)
	for i := 0; i < n+4; i++ {
		field[i] = make([]int, m+4)
	}
	for i := 0; i < n+4; i++ {
		for j := 0; j < m+4; j++ {
			if i < 2 || j < 2 || i > n+1 || j > m+1 {
				field[i][j] = 9
			}
		}
	}
	whites, _ := reader.ReadString('\n')
	white := strings.Fields(whites)
	whiteCount, _ := strconv.Atoi(white[0])
	whiteArray := make([][2]int, whiteCount)
	for i := 0; i < whiteCount; i++ {
		line, _ := reader.ReadString('\n')
		lines := strings.Fields(line)
		whiteArray[i][0], _ = strconv.Atoi(lines[0])
		whiteArray[i][1], _ = strconv.Atoi(lines[1])
	}
	blacks, _ := reader.ReadString('\n')
	black := strings.Fields(blacks)
	blackCount, _ := strconv.Atoi(black[0])
	blackArray := make([][2]int, blackCount)
	for i := 0; i < blackCount; i++ {
		line, _ := reader.ReadString('\n')
		lines := strings.Fields(line)
		blackArray[i][0], _ = strconv.Atoi(lines[0])
		blackArray[i][1], _ = strconv.Atoi(lines[1])
	}
	for i := range whiteArray {
		x := whiteArray[i][0] + 1
		y := whiteArray[i][1] + 1
		field[x][y] = 1
	}
	for i := range blackArray {
		x := blackArray[i][0] + 1
		y := blackArray[i][1] + 1
		field[x][y] = 2
	}
	turn, _ := reader.ReadString('\n')
	turn = strings.TrimSpace(turn)
	if turn == "white" {
		for i := range whiteArray {
			x := whiteArray[i][0] + 1
			y := whiteArray[i][1] + 1
			if captureableWhite(x, y, field) {
				fmt.Println("Yes")
				return
			}
		}
	}
	if turn == "black" {
		for i := range blackArray {
			x := blackArray[i][0] + 1
			y := blackArray[i][1] + 1
			if captureableBlack(x, y, field) {
				fmt.Println("Yes")
				return
			}
		}
	}
	fmt.Println("No")
}
