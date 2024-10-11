package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	nStr, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nStr))
	products := make(map[int]int)
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		parts := strings.Fields(line)
		productId, _ := strconv.Atoi(parts[0])
		categoryId, _ := strconv.Atoi(parts[1])
		products[productId] = categoryId
	}
	permStr, _ := reader.ReadString('\n')
	permParts := strings.Fields(permStr)
	perm := make([]int, n)
	for i := 0; i < n; i++ {
		perm[i], _ = strconv.Atoi(permParts[i])
	}
	categoryPositions := make(map[int][]int)
	for i := 0; i < n; i++ {
		productId := perm[i] 
		category := products[productId]
		categoryPositions[category] = append(categoryPositions[category], i)
	}
	minDiversity := math.MaxInt
	for _, positions := range categoryPositions {
		if len(positions) > 1 {
			for i := 1; i < len(positions); i++ {
				diff := positions[i] - positions[i-1]
				if diff < minDiversity {
					minDiversity = diff
				}
			}
		}
	}
	if minDiversity == math.MaxInt {
		minDiversity = n
	}

	fmt.Println(minDiversity)
}
