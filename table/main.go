package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	nn, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.Fields(nn)[0])
	names := make([]string, n)
	prevScore1 := 0
	prevScore2 := 0
	for i := 0; i < n; i++ {
		name, _ := reader.ReadString('\n')
		names[i] = strings.TrimSpace(name)
	}
	players := make(map[string]int)
	mm, _ := reader.ReadString('\n')
	m, _ := strconv.Atoi(strings.Fields(mm)[0])
	max := 0
	for i := 0; i < m; i++ {
		tmp, _ := reader.ReadString('\n')
		game := strings.Fields(tmp)
		score := game[0]
		name := game[1]
		scoreTmp := strings.Split(score, ":")
		score1, _ := strconv.Atoi(scoreTmp[0])
		score2, _ := strconv.Atoi(scoreTmp[1])
		players[name] += score1 - prevScore1
		players[name] += score2 - prevScore2
		if players[name] > max {
			max = players[name]
		}
		prevScore1 = score1
		prevScore2 = score2
	}
	for key, val := range players {
		if val == max {
			fmt.Println(key, val)
		}
	}

}
