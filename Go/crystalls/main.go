/* https://coderun.yandex.ru/problem/crystals */

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type sequence struct {
	symbol byte
	count  int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	strA, _ := reader.ReadString('\n')
	strA = strings.TrimSpace(strA)
	strB, _ := reader.ReadString('\n')
	strB = strings.TrimSpace(strB)
	strC, _ := reader.ReadString('\n')
	strC = strings.TrimSpace(strC)
	arrA := []sequence{}
	arrB := []sequence{}
	arrC := []sequence{}
	countA := 1
	prevCharA := strA[0]
	for i := 1; i < len(strA); i++ {
		if strA[i] == prevCharA {
			countA++
		} else {
			arrA = append(arrA, sequence{prevCharA, countA})
			prevCharA = strA[i]
			countA = 1
		}
	}
	arrA = append(arrA, sequence{strA[len(strA)-1], countA})
	countB := 1
	prevCharB := strB[0]
	for i := 1; i < len(strB); i++ {
		if strB[i] == prevCharB {
			countB++
		} else {
			arrB = append(arrB, sequence{prevCharB, countB})
			prevCharB = strB[i]
			countB = 1
		}
	}
	arrB = append(arrB, sequence{strB[len(strB)-1], countB})
	countC := 1
	prevCharC := strC[0]
	for i := 1; i < len(strC); i++ {
		if strC[i] == prevCharC {
			countC++
		} else {
			arrC = append(arrC, sequence{prevCharC, countC})
			prevCharC = strC[i]
			countC = 1
		}
	}
	arrC = append(arrC, sequence{strC[len(strC)-1], countC})
	if len(arrA) != len(arrB) || len(arrA) != len(arrC) || len(arrB) != len(arrC) {
		fmt.Println("IMPOSSIBLE")
		return
	}
	for i := range arrA {
		if arrA[i].symbol != arrB[i].symbol || arrA[i].symbol != arrC[i].symbol || arrB[i].symbol != arrC[i].symbol {
			fmt.Println("IMPOSSIBLE")
			return
		}
	}
	arrMean := []int{}
	for i := 0; i < len(arrA); i++ {
		mean := arrA[i].count + arrB[i].count + arrC[i].count
		mean = int(math.Round(float64(mean) / float64(3)))
		arrMean = append(arrMean, mean)
	}
	out := ""
	for i := 0; i < len(arrA); i++ {
		out += strings.Repeat(string(arrA[i].symbol), arrMean[i])
	}

	fmt.Println(out)
}
