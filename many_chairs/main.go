package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.Fields(line)[0])
	m, _ := strconv.Atoi(strings.Fields(line)[1])
	demand := make([]int,n)
	supply := make([]int,m)
	linea, _ := reader.ReadString('\n')
	demands := strings.Fields(linea)
	for i := 0; i < n; i++{
		demand[i], _ = strconv.Atoi(demands[i])
	}
	lineb, _ := reader.ReadString('\n')
	supplys := strings.Fields(lineb)
	for i := 0; i < m; i++{
		supply[i], _ = strconv.Atoi(supplys[i])
	}
}