/* Задана строка s, состоящая только из символов английского алфавита нижнего регистра ({a,…,z}), и множество символов английского алфавита нижнего регистра C
Подстрокой (i,j) (1≤i≤j≤∣s∣) назовем записанные подряд символы строки s с i-й по j-ю позиции
Назовем набором символов подстроки множество символов: T(i,j)={st​ ∣i≤t≤j}.
Вам необходимо найти подстроку (i,j) минимальной длины, для которой выполняется: T(i,j)=C. */

package main

import (
	"bufio"
	"fmt"
	"os"
)

func minWindow(s, c string) int {
	need := make(map[byte]bool)
	for i := 0; i < len(c); i++ {
		need[c[i]] = true
	}
	required := len(need)
	have := make(map[byte]int)
	formed := 0
	l, r := 0, 0
	minLen := len(s) + 1

	for r < len(s) {
		char := s[r]
		if _, ok := need[char]; ok {
			have[char]++
			if have[char] == 1 {
				formed++
			}
		}
		for l <= r && formed == required {
			if r-l+1 < minLen {
				minLen = r - l + 1
			}
			leftChar := s[l]
			if _, ok := need[leftChar]; ok {
				have[leftChar]--
				if have[leftChar] == 0 {
					formed--
				}
			}
			l++
		}
		r++
	}
	if minLen == len(s)+1 {
		return 0
	}
	return minLen
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	c, _ := reader.ReadString('\n')
	s = s[:len(s)-1]
	c = c[:len(c)-1]
	result := minWindow(s, c)
	fmt.Println(result)
}
