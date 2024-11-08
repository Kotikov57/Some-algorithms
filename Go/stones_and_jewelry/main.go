/* Даны две строки строчных латинских символов: строка J и строка S. Символы, входящие в строку J, — «драгоценности», входящие в строку S — «камни». 
Нужно определить, какое количество символов из S одновременно являются «драгоценностями». Проще говоря, нужно проверить, какое количество символов из S входит в J. */

package main

import "fmt"

func main() {
	var j, s string
	fmt.Scanln(&j)
	fmt.Scanln(&s)
	count := 0
	jewelrys := make(map[rune]bool)
	for _, valj := range j {
		jewelrys[valj] = true
	}
	for _, vals := range s {
		if jewelrys[vals] == true {
			count++
		}
	}
	fmt.Println(count)
}
