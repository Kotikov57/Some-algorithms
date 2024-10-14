/* Студенты, разбившись на две команды, решали контест по программированию. 
Каждый студент решил не менее 11 и не более NN задач.
Известно, что первая команда суммарно решила A задач, а вторая — B задач.
Определите, могло ли быть в первой команде больше студентов, чем во второй. */

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
	lineA, _ := reader.ReadString('\n')
	a, _ := strconv.Atoi(strings.Fields(lineA)[0])
	lineB, _ := reader.ReadString('\n')
	b, _ := strconv.Atoi(strings.Fields(lineB)[0])
	lineN, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.Fields(lineN)[0])
	/* 	if a > b{
		fmt.Println("Yes")
		return
	} */
	if int(math.Ceil(float64(b)/float64(n))) >= a {
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")
}
