/* https://coderun.yandex.ru/selections/backend/problems/sorting-reverse-order */

package main

import(
	"fmt"
	"encoding/json"
	"sort"
	"net/http"
	"io/ioutil"
)

func main() {
	var host string
	fmt.Scanln(&host)
	var port, a, b int
	fmt.Scanln(&port)
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	url := fmt.Sprintf("%s:%d?a=%d&b=%d", host, port, a, b)
	res, _ := http.Get(url)
	data, _ := ioutil.ReadAll(res.Body)
	var values []int
	json.Unmarshal(data, &values)
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	for _, n := range(values) {
		if n > 0 {
			fmt.Println(n)
		}
	}
}