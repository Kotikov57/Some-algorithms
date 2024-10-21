/* Имеется текст, состоящий из заглавных латинских символов и пробелов, длиной не более 255 символов. Словом в тексте называется последовательность букв, ограниченная пробелами, началом или концом строки. 
Пара соседних слов разделена хотя бы одним пробелом.
Найдите такую пару латинских символов, которая чаще всего встречается в тексте в качестве подстроки. Например, «LL» является подстрокой в тексте «HELLO», а «HO» — нет.
Гарантируется, что в тексте есть хотя бы одно слово из 2 и более букв. */

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	primStr, _ := reader.ReadString('\n')
	divStr := strings.Fields(primStr)
	mapa := make(map[string]int)
	max := 0
	for _, val := range divStr {
		for j := 0; j < len(val)-1; j++ {
			mapa[string(val[j])+string(val[j+1])]++
			count := mapa[string(val[j])+string(val[j+1])]
			if count > max {
				max = count
			}
		}
	}
	out := []string{}
	for key, value := range mapa {
		if value == max {
			out = append(out, key)
		}
	}
	sort.Slice(out, func(i, j int) bool {
		return out[i] > out[j]
	})
	fmt.Println(out[0])
}
