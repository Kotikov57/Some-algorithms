/* https://coderun.yandex.ru/problem/nearest-bus-stop/description */

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
	k, _ := strconv.Atoi(strings.Fields(line)[1])
	stations := make([]int, n)
	requests := make([]int, k)
	lineN, _ := reader.ReadString('\n')
	for i, val := range strings.Fields(lineN) {
		stations[i], _ = strconv.Atoi(val)
	}
	lineK, _ := reader.ReadString('\n')
	for i, val := range strings.Fields(lineK) {
		requests[i], _ = strconv.Atoi(val)
	}
    for _, x := range requests {
        idx := sort.Search(n, func(i int) bool {
            return stations[i] >= x
        })
        if idx < n && stations[idx] == x {
            left := idx
            for left > 0 && stations[left-1] == x {
                left--
            }
            fmt.Println(left+1)
        } else {
            if idx == 0 {
                fmt.Println(1)
            } else if idx == n {
                fmt.Println(n)
            } else {
                leftIdx := idx - 1
                rightIdx := idx
                if x-stations[leftIdx] <= stations[rightIdx]-x {
                    fmt.Println(leftIdx+1)
                } else {
                    fmt.Println(rightIdx+1)
                }
            }
        }
    }
}
